package favorites

import (
	"net/http"
	"testing"

	"api-tests-template/internal/client/http/favorites"
	"api-tests-template/internal/managers/response"
)

// получает список избранного
func List(t *testing.T, token, title string, status int) string {
	t.Helper()
	return response.Read(t, favorites.List(t, token, title), status)
}

// добавляет объявление в избранное
func Add(t *testing.T, token, id string, status int) string {
	t.Helper()
	return response.Read(t, favorites.Add(t, token, id), status)
}

// удаляет объявление из избранного
func Remove(t *testing.T, token, id string, status int) string {
	t.Helper()
	return response.Read(t, favorites.Remove(t, token, id), status)
}

// добавляет несколько объявлений в избранное
func AddBulk(t *testing.T, token string, ids []string) string {
	t.Helper()
	return response.Read(t, favorites.AddBulk(t, token, ids), http.StatusOK)
}
