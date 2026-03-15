package slog_tests

import "log/slog"

func test() {
	// correct logs
	slog.Info("starting server")                     // ok
	slog.Error("failed to connect to database")      // ok
	slog.Warn("warning something went wrong")        // ok
	slog.Debug("debug logging something went wrong") // ok

	// lowercase error
	slog.Info("Starting server")    // want `log message should start with a lowercase letter`
	slog.Error("Failed to connect") // want `log message should start with a lowercase letter`

	// english error
	slog.Info("запуск сервера")                    // want `log message must be in English`
	slog.Error("ошибка подключения к базе данных") // want `log message must be in English`
	slog.Warn("предупреждение")                    // want `log message must be in English`

	// special symbols error
	slog.Info("server started! 🚀")                // want `log message must not contain special symbols`
	slog.Error("connection failed!!!((")          // want `log message must not contain special symbols`
	slog.Warn("warning: something went wrong...") // want `log message must not contain special symbols`

	// sensitive_words
	slog.Info("user password here")   // want `log message may contain sensitive data`
	slog.Debug("api key received")    // want `log message may contain sensitive data`
	slog.Info("auth token received")  // want `log message may contain sensitive data`
	slog.Warn("session expired info") // want `log message may contain sensitive data`

	password := "123"
	token := "xyz"
	apiKey := "abc"

	slog.Info("password is " + password) // want `log message may contain sensitive data`
	slog.Info("token is " + token)       // want `log message may contain sensitive data`
	slog.Info("apiKey is " + apiKey)     // want `log message may contain sensitive data`

	// combined error
	slog.Info("Starting server! 🚀")                        // want `log message should start with a lowercase letter` `log message must not contain special symbols`
	slog.Error("ошибка: нет подключения")                  // want `log message must be in English` `log message must not contain special symbols`
	slog.Info("Auth token received")                       // want `log message should start with a lowercase letter` `log message may contain sensitive data`
	slog.Warn("предупреждения api key не инициализирован") // want `log message must be in English` `log message may contain sensitive data`
}
