FROM golang:1.25-bookworm AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o /out/morning .

FROM debian:bookworm-slim

WORKDIR /app

RUN useradd --system --create-home --home-dir /app --shell /usr/sbin/nologin app

COPY --from=build /out/morning /app/morning
# Root certificates for outgoing HTTPS (the AI coach calls the model API).
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENV MORNING_DB=/data/morning.db

RUN mkdir -p /data && chown -R app:app /app /data

USER app

EXPOSE 8080

CMD ["/app/morning"]
