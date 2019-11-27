package cmd

import (
	"strings"
	"testing"

	"github.com/spf13/viper"

	"github.com/stretchr/testify/assert"
)

var (
	expSettings = `  .apps:  ["go","git","git-lfs","git-crypt","zsh","keychain","yarn"]
`
	expCfg = "\nconfig:\n" + expSettings
)

func init() {
	viper.AddConfigPath(".")
	initConfig()
}

func TestConfigPrinting(t *testing.T) {
	sb := new(strings.Builder)
	writeSettings(sb, viper.AllSettings())

	assert.Equal(t, expSettings, sb.String(), "Expected printed configuration to match expectation")
}

func TestSprintConfig(t *testing.T) {
	cfg := sprintConfig()

	assert.Equal(t, expCfg, cfg, "Expected configuration to match expectation")
}

func TestConfigCommand(t *testing.T) {
	assert := assert.New(t)
	output, err := executeCommand(rootCmd, "config")
	assert.NoError(err)
	assert.Equal(output, expCfg, "Expected configation to be outputted in a different format")
}
