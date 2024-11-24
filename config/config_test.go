package config_test

import (
	"os"
	"testing"

	"github.com/KazikovAP/log_analyzer/config"
	"github.com/stretchr/testify/assert"
)

func TestInitConfig_MissingPathFlag(t *testing.T) {
	os.Args = []string{
		"cmd",
	}

	_, err := config.Init()
	if err == nil {
		t.Fatal("Ожидалась ошибка, но она не была возвращена")
	}

	assert.EqualError(t, err, "обязательный флаг -path не указан")
}
