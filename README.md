# LogLintGo

LogLintGo - это кастомный линтер для Go совместимый с [golangci-lint](https://golangci-lint.run/), который проверяет логи. 
Поддерживаемые логгеры:`log/slog`, `go.uber.org/zap`.

## Rules

| Rule | ❌ Bad | ✅ Good |
|------|--------|---------|
| Message must start with a lowercase letter | `slog.Info("Starting server on port 8080")` | `slog.Info("starting server on port 8080")` |
| Message must be in English | `slog.Error("ошибка подключения к базе данных")` | `slog.Error("failed to connect to database")` |
| Message must not contain special characters (`!`, `?`, `#`, etc.) | `slog.Warn("something went wrong!🚀")` | `slog.Warn("something went wrong")` |
| Message must not contain sensitive keywords (`token`, `password`, `secret`, etc.) | `slog.Info("user token: " + token)` | `slog.Info("user authenticated successfully")` |

## Сборка бинарного файла

### Склонировать репозиторий
```bash
git clone https://github.com/kirillveshnyakov/LogLintGo.git
```

### Сборка loglinter
```bash
go build -o loglinter ./cmd/loglinter
```

### Сборка custom-gcl для golangci-lint
```bash
golangci-lint custom -v
```

После в корне репозитория появятся 2 бинарных файла: `loglinter` и `custom-gcl`

## Использование

### Requirements

- Go version 1.22+
- golangci-lint v2.0.0+ (в случае использования custom-gcl)

### Вариант 1:
```bash
# Скопировать бинарник в PATH
cp ./loglinter /usr/local/bin/loglinter

# Запустить в проекте
loglinter ./...

# Или через go vet
go vet -vettool=$(which loglinter) ./...
```


### Вариант 2:

### Установка golangci-lint

```bash
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
  | sh -s -- -b $(go env GOPATH)/bin latest
```

#### Скопировать необходимые файлы в свой проект
```bash
cp ./custom-gcl /path/to/other-project/
cp ./.golangci.yml /path/to/other-project/
```

#### Если файл `.golangci.yml` уже был в проекте, то нужно добавить в уже существующую конфигурацию
```bash
version: "2"

linters:
  default: none
  enable:
    - govet        # уже был
    - staticcheck  # уже был
    - loglinter    # добавить

  settings:
    custom:
      loglinter:
        type: "module"
        description: "checks log messages style and content"
```

#### Запуск
```bash
./custom-gcl run ./<path_to_file_or_dir>
```

## Running Tests
```bash
 go test -v ./tests/...
```

## License

MIT
