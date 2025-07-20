package belajar_golang_viper

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViper(t *testing.T) {
	var config *viper.Viper = viper.New()
	assert.NotNil(t, config)
}

func TestJSON(t *testing.T) {
	config := viper.New()
	// config.SetConfigName("salah") // error: viper.ConfigFileNotFoundError
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	// read file
	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "belajar-golang-viper", config.GetString("app.name"))
	assert.Equal(t, "1.0.0", config.GetString("app.version"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
	assert.Equal(t, true, config.GetBool("database.show_sql"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
}

func TestYAML(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("config.yaml")
	config.AddConfigPath(".")

	// read file
	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "belajar-golang-viper", config.GetString("app.name"))
	assert.Equal(t, "1.0.0", config.GetString("app.version"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
	assert.Equal(t, true, config.GetBool("database.show_sql"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
}
