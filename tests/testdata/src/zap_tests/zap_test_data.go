package zap_tests

import "go.uber.org/zap"

func test() {
	logger := &zap.Logger{}
	sugar := &zap.SugaredLogger{}

	// correct logs
	logger.Info("starting server")                     // ok
	logger.Error("failed to connect to database")      // ok
	logger.Warn("warning something went wrong")        // ok
	logger.Debug("debug logging something went wrong") // ok

	// lowercase error
	logger.Info("Starting server")    // want `log message should start with a lowercase letter`
	logger.Error("Failed to connect") // want `log message should start with a lowercase letter`

	// english error
	logger.Info("запуск сервера")                    // want `log message must be in English`
	logger.Error("ошибка подключения к базе данных") // want `log message must be in English`
	logger.Warn("предупреждение")                    // want `log message must be in English`

	// special symbols error
	logger.Info("server started! 🚀")                // want `log message must not contain special symbols`
	logger.Error("connection failed!!!((")          // want `log message must not contain special symbols`
	logger.Warn("warning: something went wrong...") // want `log message must not contain special symbols`

	// sensitive words
	logger.Info("user password here")   // want `log message may contain sensitive data`
	logger.Debug("api key received")    // want `log message may contain sensitive data`
	logger.Info("auth token received")  // want `log message may contain sensitive data`
	logger.Warn("session expired info") // want `log message may contain sensitive data`

	password := "123"
	token := "xyz"
	apiKey := "abc"

	logger.Info("password is " + password) // want `log message may contain sensitive data`
	logger.Info("token is " + token)       // want `log message may contain sensitive data`
	logger.Info("apiKey is " + apiKey)     // want `log message may contain sensitive data`

	// combined error
	logger.Info("Starting server! 🚀")                        // want `log message should start with a lowercase letter` `log message must not contain special symbols`
	logger.Error("ошибка: нет подключения")                  // want `log message must be in English` `log message must not contain special symbols`
	logger.Info("Auth token received")                       // want `log message should start with a lowercase letter` `log message may contain sensitive data`
	logger.Warn("предупреждения api key не инициализирован") // want `log message must be in English` `log message may contain sensitive data`

	sugar.Infow("starting server", "key", "value")     // ok
	sugar.Infow("Starting server", "key", "value")     // want `log message should start with a lowercase letter`
	sugar.Errorw("ошибка подключения", "key", "value") // want `log message must be in English`
	sugar.Warnw("server started!", "key", "value")     // want `log message must not contain special symbols`
	sugar.Infow("user password here", "key", "value")  // want `log message may contain sensitive data`
	sugar.Infow("Auth token received", "key", "value") // want `log message should start with a lowercase letter` `log message may contain sensitive data`
}
