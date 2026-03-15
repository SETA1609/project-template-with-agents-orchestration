FROM golang:1.23-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o /server ./cmd/server

FROM gcr.io/distroless/static:nonroot

ENV MCP_TRANSPORT=http
ENV HOST=0.0.0.0
ENV PORT=3000
ENV LOG_LEVEL=info

COPY --from=builder /server /server

EXPOSE 3000

ENTRYPOINT ["/server"]
