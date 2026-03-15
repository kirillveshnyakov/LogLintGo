# LogLintGo

A Go linter for checking log messages style and content. Compatible with `golangci-lint`.

## Rules

| Rule | Description | Example ❌ | Example ✅ |
|------|-------------|-----------|-----------|
| **lowercase** | Log messages must start with a lowercase letter | `"Starting server"` | `"starting server"` |
| **english** | Log messages must be in English only | `"запуск сервера"` | `"starting server"` |
| **no-specials** | Log messages must not contain special characters or emoji | `"server started! 🚀"` | `"server started"` |
| **no-secrets** | Log messages must not contain sensitive data | `"user password: " + pwd` | `"user authenticated"` |

## Supported Loggers

- `log/slog`
- `go.uber.org/zap` (`Logger`, `SugaredLogger`)

## Installation

```bash
go install github.com/kirillveshnyakov/LogLintGo/cmd/loglinter@latest
