// Package coach is the AI coach: a chat on the program page that edits the
// user's current program (and only it) through a small set of tools, using any
// OpenAI-compatible chat API (OpenRouter by default).
package coach

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// The models used when AI_MODEL / AI_FALLBACK are not set, as "model@provider"
// (chosen with cmd/coachbench). "~…-latest" names are OpenRouter aliases that
// move to the maker's newest model of the family by themselves, so a new
// version needs no code change. The provider pins requests to the maker's own
// API, never a reseller. The fallback answers when the main model fails.
const (
	DefaultModel    = "~deepseek/deepseek-flash-latest@deepseek"
	DefaultFallback = "qwen/qwen3.8-flash@alibaba"
)

// Config says where the model lives. The key comes only from the environment.
type Config struct {
	BaseURL string // e.g. https://openrouter.ai/api/v1
	Key     string
	Model   string
	// Providers pins OpenRouter to the model maker's own servers, e.g.
	// "deepseek" or "alibaba" — never a reseller.
	Providers []string
	// Fallback is tried when this model fails (nil = none).
	Fallback *Config
}

// parseModel reads "model@provider1,provider2".
func parseModel(spec string) (model string, providers []string) {
	model, prov, _ := strings.Cut(strings.TrimSpace(spec), "@")
	for _, x := range strings.Split(prov, ",") {
		if x = strings.TrimSpace(x); x != "" {
			providers = append(providers, x)
		}
	}
	return strings.TrimSpace(model), providers
}

// FromEnv reads AI_BASE_URL, AI_API_KEY, AI_MODEL and AI_FALLBACK (both as
// "model@provider"; AI_FALLBACK=none turns the fallback off).
func FromEnv() Config {
	c := Config{
		BaseURL: strings.TrimRight(strings.TrimSpace(os.Getenv("AI_BASE_URL")), "/"),
		Key:     strings.TrimSpace(os.Getenv("AI_API_KEY")),
	}
	if c.BaseURL == "" {
		c.BaseURL = "https://openrouter.ai/api/v1"
	}
	spec := strings.TrimSpace(os.Getenv("AI_MODEL"))
	if spec == "" {
		spec = DefaultModel
	}
	c.Model, c.Providers = parseModel(spec)
	fb := strings.TrimSpace(os.Getenv("AI_FALLBACK"))
	if fb == "" {
		fb = DefaultFallback
	}
	if fb != "none" {
		f := Config{BaseURL: c.BaseURL, Key: c.Key}
		f.Model, f.Providers = parseModel(fb)
		if f.Model != c.Model {
			c.Fallback = &f
		}
	}
	return c
}

// WithModel returns the config for one "model@provider" without a fallback.
func (c Config) WithModel(spec string) Config {
	c.Model, c.Providers = parseModel(spec)
	c.Fallback = nil
	return c
}

// Enabled reports whether a key is configured.
func (c Config) Enabled() bool { return c.Key != "" }

// Message is one chat message in the OpenAI format.
type Message struct {
	Role       string     `json:"role"`
	Content    string     `json:"content"`
	ToolCalls  []ToolCall `json:"tool_calls,omitempty"`
	ToolCallID string     `json:"tool_call_id,omitempty"`
}

// ToolCall is a function call requested by the model.
type ToolCall struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Function struct {
		Name      string `json:"name"`
		Arguments string `json:"arguments"`
	} `json:"function"`
}

// Usage is what a call cost.
type Usage struct {
	PromptTokens     int     `json:"prompt_tokens"`
	CompletionTokens int     `json:"completion_tokens"`
	Cost             float64 `json:"cost"` // USD, reported by OpenRouter
}

var httpClient = &http.Client{Timeout: 90 * time.Second}

// chat sends one request and returns the model's message.
func (c Config) chat(ctx context.Context, msgs []Message, tools []map[string]any) (Message, Usage, error) {
	body := map[string]any{
		"model":       c.Model,
		"messages":    msgs,
		"temperature": 0.2,
		"max_tokens":  4000,
	}
	if len(tools) > 0 {
		body["tools"] = tools
		body["tool_choice"] = "auto"
	}
	if strings.Contains(c.BaseURL, "openrouter.ai") {
		body["usage"] = map[string]any{"include": true}
		body["reasoning"] = map[string]any{"effort": "low", "exclude": true}
		if len(c.Providers) > 0 {
			body["provider"] = map[string]any{"only": c.Providers, "allow_fallbacks": false}
		}
	}
	raw, err := json.Marshal(body)
	if err != nil {
		return Message{}, Usage{}, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.BaseURL+"/chat/completions", bytes.NewReader(raw))
	if err != nil {
		return Message{}, Usage{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Key)
	req.Header.Set("HTTP-Referer", "https://morning30.com")
	req.Header.Set("X-Title", "morning30")
	resp, err := httpClient.Do(req)
	if err != nil {
		return Message{}, Usage{}, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(io.LimitReader(resp.Body, 4<<20))
	if err != nil {
		return Message{}, Usage{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return Message{}, Usage{}, fmt.Errorf("model api: %s: %s", resp.Status, clip(string(data), 300))
	}
	var out struct {
		Choices []struct {
			Message      Message `json:"message"`
			FinishReason string  `json:"finish_reason"`
		} `json:"choices"`
		Usage Usage `json:"usage"`
		Error *struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.Unmarshal(data, &out); err != nil {
		return Message{}, Usage{}, err
	}
	if out.Error != nil {
		return Message{}, out.Usage, errors.New("model api: " + out.Error.Message)
	}
	if len(out.Choices) == 0 {
		return Message{}, out.Usage, errors.New("model api: no answer")
	}
	if os.Getenv("COACH_DEBUG") != "" {
		fmt.Fprintf(os.Stderr, "[coach] %s finish=%s content=%d tools=%d tokens=%d/%d\n", c.Model, out.Choices[0].FinishReason,
			len(out.Choices[0].Message.Content), len(out.Choices[0].Message.ToolCalls), out.Usage.PromptTokens, out.Usage.CompletionTokens)
	}
	return out.Choices[0].Message, out.Usage, nil
}

func clip(s string, n int) string {
	if r := []rune(s); len(r) > n {
		return string(r[:n]) + "…"
	}
	return s
}
