package config

import (
	"fmt"
	"github.com/halm4d/kubeforward/util"
	"github.com/spf13/viper"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

type AppConfig struct {
	KubeConfig string            `yaml:"kubeConfig"`
	Namespaces []NameSpaceConfig `yaml:"namespaces"`
}

func New() *AppConfig {
	return &AppConfig{
		KubeConfig: filepath.Join(homedir.HomeDir(), ".kube", "config"),
	}
}

func (c *AppConfig) FindNamespace(s string) *NameSpaceConfig {
	for _, namespace := range c.Namespaces {
		if namespace.Name == s {
			return &namespace
		}
	}
	panic(fmt.Sprintf("%s namespace not found in config file.", s))
}

func (c *AppConfig) Load() {
	viper.AddConfigPath(filepath.Join(util.GetEnvOrDefault("KUBEFORWARD_CONFIG_PATH", util.GetExecutablePath())))
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(c); err != nil {
		panic(fmt.Sprintf("unable to decode into config struct, %v", err))
	}
}
