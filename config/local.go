package config

import (
    "errors"
    log "github.com/siruspen/logrus"
    "github.com/spf13/viper"
    "os"
    "strings"
)

type LocalRepository struct {
    viper *viper.Viper
}

func readLocalConfig() (*viper.Viper, error) {
    log.Info("Reading environment variables.")
    v := viper.New()
    v.AutomaticEnv()
    v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
    v.AddConfigPath(".")
    v.SetConfigName("config")
    env := os.Getenv("ENV")
    if env == "local" {
        log.Infof("Environment: %s", env)
        v.SetConfigName("config.local")

    }
    if env == "test" {
        log.Infof("Environment: %s", env)
        v.AddConfigPath("../../../../")
        v.SetConfigName("config.test")
    }
    err := v.ReadInConfig()
    if err != nil {
        log.Fatalf("error reading config file or env variable '%s' ", err.Error())
        return nil, err
    }
    return v, nil
}

func NewLocalRepository() *LocalRepository {
    v, err := readLocalConfig()
    if err != nil {
        log.Fatalf("failed to read local config: %s", err)
        return nil
    }
    return &LocalRepository{viper: v}
}
func (lr *LocalRepository) GetStringSlice(key string) []string {
    return lr.viper.GetStringSlice(key)
}

func (lr *LocalRepository) GetBool(key string) bool {
    return lr.viper.GetBool(key)
}

func (lr *LocalRepository) Get(key string) (interface{}, error) {
    configMap := lr.viper.Get(key)
    if configMap == nil {
        return nil, errors.New("config isn't initialised")
    }

    return configMap, nil

}

// GetString get a secret as a string
func (lr *LocalRepository) GetString(key string) string {
    return lr.viper.GetString(key)
}

func (lr *LocalRepository) GetInt(key string) int64 {
    return lr.viper.GetInt64(key)
}

func (lr *LocalRepository) GetFloat(key string) float64 {
    return lr.viper.GetFloat64(key)
}
