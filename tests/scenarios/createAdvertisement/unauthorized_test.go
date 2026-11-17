package createAdvertisement

import (
	"net/http"
	"testing"

	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/response"
	"api-tests-template/tests"
	"api-tests-template/tests/fixtures"
)

func TestCreateWithoutAuthorization(t *testing.T) {
	tests.SetupSuite()
	for _, tc := range []struct{ name, token, message string }{
		{"missing_token", "", fixtures.MissingAuthMessage},
		{"invalid_token", fixtures.InvalidToken, fixtures.InvalidAuthMessage},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Log("Подготовка: собираем корректное объявление")
			fields, photo := fixtures.Advertisement()

			t.Log("Отправляем объявление и проверяем ошибку авторизации")
			body := advertisements.Create(t, tc.token, fields, photo, http.StatusUnauthorized)
			response.CheckError(t, body, tc.message)
			t.Log("Получили 401 и сообщение об ошибке")
		})
	}
}
