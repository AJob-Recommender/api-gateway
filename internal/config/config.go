package config

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type (
	Config struct {
		API      API      `mapstructure:"api"`
		Services Services `mapstructure:"services"`
		Logger   Logger   `mapstructure:"logger"`
	}

	Logger struct {
		Level string `mapstructure:"level"`
	}

	API struct {
		Port string `mapstructure:"port"`
	}

	Services struct {
		Seer Seer `mapstructure:"seer"`
	}

	Seer struct {
		Timeout            time.Duration `mapstructure:"timeout"`
		MaxRetry           int           `mapstructure:"max_retry"`
		MaxIdleConsPerHost int           `mapstructure:"max_idle_cons_per_host"`
		MaxIdleCons        int           `mapstructure:"max_idle_cons"`
		MaxConsPerHost     int           `mapstructure:"max_cons_per_host"`
		URL                string        `mapstructure:"url"`
		AuthUser           string        `mapstructure:"auth_user"`
		AuthPass           string        `mapstructure:"auth_pass"`
		Enable             bool          `mapstructure:"enable"`
	}
)

func Load() (cfg *Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return cfg, fmt.Errorf("failed to load config: %w", err)
	}

	if err = viper.Unmarshal(&cfg); err != nil {
		return nil, errors.WithStack(err)
	}

	return cfg, nil
}
