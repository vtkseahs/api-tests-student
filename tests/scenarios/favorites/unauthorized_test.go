package favorites

import (
	"net/http"
	"testing"

	"api-tests-template/internal/managers/favorites"
	"api-tests-template/internal/managers/response"
	"api-tests-template/tests"
	"api-tests-template/tests/fixtures"
)

const advertisementID = "00000000-0000-4000-8000-000000000001"

func TestFavoritesWithoutAuthorization(t *testing.T) {
	tests.SetupSuite()
	t.Log("Подготовка: запросы без токена и с неверным токеном")

	for _, tc := range []struct{ name, token, message string }{
		{"missing_token", "", fixtures.MissingAuthMessage},
		{"invalid_token", fixtures.InvalidToken, fixtures.InvalidAuthMessage},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Run("GET", func(t *testing.T) {
				body := favorites.List(t, tc.token, "", http.StatusUnauthorized)
				response.CheckError(t, body, tc.message)
				t.Log("Список избранного недоступен")
			})
			t.Run("POST", func(t *testing.T) {
				body := favorites.Add(t, tc.token, advertisementID, http.StatusUnauthorized)
				response.CheckError(t, body, tc.message)
				t.Log("Добавление недоступно")
			})
			t.Run("DELETE", func(t *testing.T) {
				body := favorites.Remove(t, tc.token, advertisementID, http.StatusUnauthorized)
				response.CheckError(t, body, tc.message)
				t.Log("Удаление недоступно")
			})
		})
	}
}
