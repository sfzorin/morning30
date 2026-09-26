// Package mcp is the agent interface: stdio JSON-RPC (Model Context Protocol)
// plus optional Bearer HTTP /api. Agents write custom exercises and upload a
// 30-day cycle into one account's personal JSON. Progress is never reset.
package mcp

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"morning30/internal/app"
	"morning30/internal/catalog"
	"morning30/internal/content"
	"morning30/internal/db"
	"morning30/internal/i18n"
)

const protocolVersion = "2024-11-05"

type rpcReq struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

type rpcResp struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id,omitempty"`
	Result  any             `json:"result,omitempty"`
	Error   *rpcErr         `json:"error,omitempty"`
}

type rpcErr struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type toolCall struct {
	Name      string          `json:"name"`
	Arguments json.RawMessage `json:"arguments"`
}

type accountArgs struct {
	Account string `json:"account"`
}

type programArgs struct {
	Account string          `json:"account"`
	Program json.RawMessage `json:"program"`
}

type exerciseArgs struct {
	Account  string          `json:"account"`
	ID       string          `json:"id"`
	Exercise json.RawMessage `json:"exercise"`
}

type programView struct {
	Account  db.Account       `json:"account"`
	Personal bool             `json:"personal"`
	Program  content.Resolved `json:"program"`
}

type toolResult struct {
	Content []toolContent `json:"content"`
	IsError bool          `json:"isError,omitempty"`
}

type toolContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func textResult(v any, err error) toolResult {
	if err != nil {
		return toolResult{Content: []toolContent{{Type: "text", Text: err.Error()}}, IsError: true}
	}
	b, _ := json.MarshalIndent(v, "", "  ")
	return toolResult{Content: []toolContent{{Type: "text", Text: string(b)}}}
}

func tools() []map[string]any {
	account := map[string]any{"type": "string", "description": "User id, email, or unique display name"}
	return []map[string]any{
		{
			"name":        "list_accounts",
			"description": "List morning30 accounts (id, email, name, whether they have a personal 30-day plan).",
			"inputSchema": map[string]any{"type": "object", "properties": map[string]any{}},
		},
		{
			"name":        "get_program",
			"description": "Get an account's 30-day cycle JSON (personal file, or the live demo if they have none).",
			"inputSchema": map[string]any{"type": "object", "properties": map[string]any{"account": account}, "required": []string{"account"}},
		},
		{
			"name":        "set_program",
			"description": "Upload a 30-day cycle JSON into the account's personal file. Does not reset day position or streaks.",
			"inputSchema": map[string]any{"type": "object", "properties": map[string]any{"account": account, "program": map[string]any{"description": "Resolved program object or JSON string"}}, "required": []string{"account", "program"}},
		},
		{
			"name":        "restore_demo",
			"description": "Clear the personal file so the account runs the public demo. Progress is kept.",
			"inputSchema": map[string]any{"type": "object", "properties": map[string]any{"account": account}, "required": []string{"account"}},
		},
		{
			"name":        "list_custom_exercises",
			"description": "List this account's custom exercises (overlay on the built-in library).",
			"inputSchema": map[string]any{"type": "object", "properties": map[string]any{"account": account}, "required": []string{"account"}},
		},
		{
			"name":        "list_exercises",
			"description": "List every exercise this account can use (built-in plus custom overlay).",
			"inputSchema": map[string]any{"type": "object", "properties": map[string]any{"account": account}, "required": []string{"account"}},
		},
		{
			"name":        "get_exercise",
			"description": "Get one exercise document (custom overlay wins, else built-in).",
			"inputSchema": map[string]any{"type": "object", "properties": map[string]any{"account": account, "id": map[string]any{"type": "string"}}, "required": []string{"account", "id"}},
		},
		{
			"name":        "put_exercise",
			"description": "Write a custom exercise into the account's library. svg must be a built-in figure id (no image upload).",
			"inputSchema": map[string]any{"type": "object", "properties": map[string]any{"account": account, "exercise": map[string]any{"type": "object"}}, "required": []string{"account", "exercise"}},
		},
		{
			"name":        "list_figures",
			"description": "Built-in illustration ids a custom exercise may reuse as svg.",
			"inputSchema": map[string]any{"type": "object", "properties": map[string]any{}},
		},
	}
}

func callTool(name string, args json.RawMessage) toolResult {
	app.SnapshotAll()
	switch name {
	case "list_accounts":
		accs, err := app.DB.ListAccounts()
		if accs == nil {
			accs = []db.Account{}
		}
		return textResult(accs, err)
	case "get_program":
		var a accountArgs
		if err := json.Unmarshal(args, &a); err != nil {
			return textResult(nil, err)
		}
		acc, err := app.DB.FindAccount(a.Account)
		if err != nil {
			return textResult(nil, err)
		}
		return textResult(programView{
			Account:  acc,
			Personal: app.ActiveProgramKey(acc.ID) == "custom",
			Program:  app.UserProgram(acc.ID),
		}, nil)
	case "set_program":
		var a programArgs
		if err := json.Unmarshal(args, &a); err != nil {
			return textResult(nil, err)
		}
		acc, err := app.DB.FindAccount(a.Account)
		if err != nil {
			return textResult(nil, err)
		}
		raw, err := programJSON(a.Program)
		if err != nil {
			return textResult(nil, err)
		}
		if err := app.SetUserProgram(acc.ID, raw); err != nil {
			return textResult(nil, err)
		}
		acc, _ = app.DB.AccountByID(acc.ID)
		return textResult(programView{Account: acc, Personal: true, Program: app.UserProgram(acc.ID)}, nil)
	case "restore_demo":
		var a accountArgs
		if err := json.Unmarshal(args, &a); err != nil {
			return textResult(nil, err)
		}
		acc, err := app.DB.FindAccount(a.Account)
		if err != nil {
			return textResult(nil, err)
		}
		if err := app.RestoreDemo(acc.ID); err != nil {
			return textResult(nil, err)
		}
		acc, _ = app.DB.AccountByID(acc.ID)
		return textResult(programView{Account: acc, Personal: false, Program: app.UserProgram(acc.ID)}, nil)
	case "list_custom_exercises":
		var a accountArgs
		if err := json.Unmarshal(args, &a); err != nil {
			return textResult(nil, err)
		}
		acc, err := app.DB.FindAccount(a.Account)
		if err != nil {
			return textResult(nil, err)
		}
		j, err := app.DB.GetCustomExercises(acc.ID)
		if err != nil {
			return textResult(nil, err)
		}
		c := catalog.New(j)
		return textResult(c.Custom(), nil)
	case "list_exercises":
		var a accountArgs
		if err := json.Unmarshal(args, &a); err != nil {
			return textResult(nil, err)
		}
		acc, err := app.DB.FindAccount(a.Account)
		if err != nil {
			return textResult(nil, err)
		}
		return textResult(app.UserCatalog(acc.ID).List(i18n.EN), nil)
	case "get_exercise":
		var a exerciseArgs
		if err := json.Unmarshal(args, &a); err != nil {
			return textResult(nil, err)
		}
		acc, err := app.DB.FindAccount(a.Account)
		if err != nil {
			return textResult(nil, err)
		}
		d, ok := app.UserCatalog(acc.ID).Doc(strings.TrimSpace(a.ID))
		if !ok {
			return textResult(nil, fmt.Errorf("exercise %q not found", a.ID))
		}
		return textResult(d, nil)
	case "put_exercise":
		var a exerciseArgs
		if err := json.Unmarshal(args, &a); err != nil {
			return textResult(nil, err)
		}
		acc, err := app.DB.FindAccount(a.Account)
		if err != nil {
			return textResult(nil, err)
		}
		if len(bytes.TrimSpace(a.Exercise)) == 0 {
			return textResult(nil, errors.New("exercise is required"))
		}
		var d catalog.Doc
		if err := json.Unmarshal(a.Exercise, &d); err != nil {
			return textResult(nil, err)
		}
		if err := app.SaveCustomExercise(acc.ID, d); err != nil {
			return textResult(nil, err)
		}
		got, _ := app.UserCatalog(acc.ID).Doc(d.ID)
		return textResult(got, nil)
	case "list_figures":
		return textResult(content.FigureIDs(), nil)
	default:
		return textResult(nil, fmt.Errorf("unknown tool %q", name))
	}
}

func programJSON(raw json.RawMessage) (string, error) {
	raw = bytes.TrimSpace(raw)
	if len(raw) == 0 {
		return "", app.ErrInvalidProgram
	}
	if raw[0] == '"' {
		var s string
		if err := json.Unmarshal(raw, &s); err != nil {
			return "", err
		}
		return strings.TrimSpace(s), nil
	}
	return string(raw), nil
}

func handleRPC(body []byte) []byte {
	var req rpcReq
	if err := json.Unmarshal(body, &req); err != nil {
		b, _ := json.Marshal(rpcResp{JSONRPC: "2.0", Error: &rpcErr{Code: -32700, Message: "parse error"}})
		return b
	}
	if req.JSONRPC == "" {
		req.JSONRPC = "2.0"
	}
	if req.ID == nil && strings.HasPrefix(req.Method, "notifications/") {
		return nil
	}
	resp := rpcResp{JSONRPC: "2.0", ID: req.ID}
	switch req.Method {
	case "initialize":
		resp.Result = map[string]any{
			"protocolVersion": protocolVersion,
			"capabilities":    map[string]any{"tools": map[string]any{}},
			"serverInfo":      map[string]any{"name": "morning30", "version": "1.0.0"},
		}
	case "ping", "notifications/initialized":
		resp.Result = map[string]any{}
	case "tools/list":
		resp.Result = map[string]any{"tools": tools()}
	case "tools/call":
		var call toolCall
		if err := json.Unmarshal(req.Params, &call); err != nil {
			resp.Error = &rpcErr{Code: -32602, Message: err.Error()}
			break
		}
		resp.Result = callTool(call.Name, call.Arguments)
	case "resources/list":
		resp.Result = map[string]any{"resources": []any{}}
	case "prompts/list":
		resp.Result = map[string]any{"prompts": []any{}}
	default:
		resp.Error = &rpcErr{Code: -32601, Message: "method not found: " + req.Method}
	}
	b, _ := json.Marshal(resp)
	return b
}

// ServeStdio runs the MCP JSON-RPC server on stdin/stdout until EOF.
func ServeStdio() error {
	app.SnapshotAll()
	in := bufio.NewReader(os.Stdin)
	for {
		body, err := readMessage(in)
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return err
		}
		if len(bytes.TrimSpace(body)) == 0 {
			continue
		}
		resp := handleRPC(body)
		if resp == nil {
			continue
		}
		if err := writeMessage(os.Stdout, resp); err != nil {
			return err
		}
	}
}

func readMessage(r *bufio.Reader) ([]byte, error) {
	b, err := r.Peek(1)
	if err != nil {
		return nil, err
	}
	if b[0] == '{' {
		line, err := r.ReadBytes('\n')
		if err != nil && !errors.Is(err, io.EOF) {
			return nil, err
		}
		if errors.Is(err, io.EOF) && len(bytes.TrimSpace(line)) == 0 {
			return nil, io.EOF
		}
		return bytes.TrimSpace(line), nil
	}
	n := 0
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			return nil, err
		}
		line = strings.TrimRight(line, "\r\n")
		if line == "" {
			break
		}
		if i := strings.IndexByte(line, ':'); i >= 0 && strings.EqualFold(strings.TrimSpace(line[:i]), "content-length") {
			n, _ = strconv.Atoi(strings.TrimSpace(line[i+1:]))
		}
	}
	if n <= 0 {
		return nil, errors.New("missing Content-Length")
	}
	body := make([]byte, n)
	_, err = io.ReadFull(r, body)
	return body, err
}

func writeMessage(w io.Writer, body []byte) error {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "Content-Length: %d\r\n\r\n", len(body))
	buf.Write(body)
	_, err := w.Write(buf.Bytes())
	return err
}
