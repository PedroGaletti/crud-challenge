package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	t.Run("GetConfig - Test runnning the application without a .env file", func(t *testing.T) {
		config := GetConfig()

		assert.Equal(t, Config{
			GinMode:     "",
			LogLevel:    "",
			SqlDb:       "",
			SqlHost:     "",
			SqlPassword: "",
			SqlPort:     "",
			SqlUser:     "",
		}, config)
	})
}
