package controllers

import (
	"github.com/labring/sealos/controllers/pkg/config"
)

type Config struct {
	config.Global  `yaml:",inline"`
	TerminalConfig TerminalConfig `yaml:"terminalController"`
}

type TerminalConfig struct {
	IngressTlsSecretName string `yaml:"ingressTlsSecretName"`
}
