package utils

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type envModel struct {
	Port      string `env:"PORT" envDefault:"8084"`
	DebugMode bool   `env:"DEBUG_MODE" envDefault:"false"`

	EvolutionIntegration string `env:"EVOLUTION_INTEGRATION" envDefault:"https://evolution-integration-592799294413.us-central1.run.app"`
}

var Env envModel

func LoadEnvs() error {
	if err := godotenv.Load(".env"); err != nil {
		zap.L().Warn("error loading .env file", zap.Error(err))
	}

	return env.Parse(&Env)
}
