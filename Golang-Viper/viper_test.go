package Golang_Viper

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

var config = viper.New()

func TestViper(t *testing.T) {
	assert.NotNil(t, config)

}

func TestViperJson(t *testing.T) {
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "golang-viper", config.GetString("app.name"))
	assert.Equal(t, "hanif", config.GetString("app.author"))
	assert.Equal(t, "1.0.0", config.GetString("app.version"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
	assert.Equal(t, true, config.GetBool("database.show_sql"))
}

func TestYaml(t *testing.T) {
	config.SetConfigFile("config.yaml")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "golang-viper", config.GetString("app.name"))
	assert.Equal(t, "hanif", config.GetString("app.author"))
	assert.Equal(t, "1.0.0", config.GetString("app.version"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
	assert.Equal(t, true, config.GetBool("database.show_sql"))
}

func TestEnv(t *testing.T) {
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "golang-viper", config.GetString("APP_NAME"))
	assert.Equal(t, "hanif", config.GetString("APP_AUTHOR"))
	assert.Equal(t, "1.0.0", config.GetString("APP_VERSION"))
	assert.Equal(t, "localhost", config.GetString("DATABASE_HOST"))
	assert.Equal(t, 3306, config.GetInt("DATABASE_PORT"))
	assert.Equal(t, true, config.GetBool("DATABASE_SHOW_SQL"))
}

func TestEnvVariable(t *testing.T) {
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "golang-viper", config.GetString("APP_NAME"))
	assert.Equal(t, "hanif", config.GetString("APP_AUTHOR"))
	assert.Equal(t, "1.0.0", config.GetString("APP_VERSION"))
	assert.Equal(t, "localhost", config.GetString("DATABASE_HOST"))
	assert.Equal(t, 3306, config.GetInt("DATABASE_PORT"))
	assert.Equal(t, true, config.GetBool("DATABASE_SHOW_SQL"))

	assert.Equal(t, "Hello", config.GetString("FROM_ENV"))
}
