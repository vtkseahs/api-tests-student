package fixtures

import (
	"os"
	"testing"

	"api-tests-template/internal/managers/auth"
	"api-tests-template/tests"
)

// получает токен основного аккаунта
func Login(t *testing.T) string {
	t.Helper()
	tests.SetupSuite()
	t.Log("Подготовка: авторизуемся")
	return auth.Login(t, os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD")).Token
}

const MissingAuthMessage = "Authorization header required"
