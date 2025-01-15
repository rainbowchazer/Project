package main

import (
	"idel/config"

	"go.uber.org/zap"
)

func main() {

	// Инициализация zap-логгера
	logger, err := zap.NewProduction()
	if err != nil {
		panic("Ошибка инициализации логгера: " + err.Error())
	}
	defer logger.Sync() // Flush буфер перед завершением программы

	sugar := logger.Sugar()

	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		sugar.Fatalf("Ошибка при загрузке конфигурации: %v", err)
	}

	// Используем конфигурацию
	sugar.Infof("Server Host: %s", cfg.Server.Host)
	sugar.Infof("Database URL: %s", cfg.Database.URL)

}
