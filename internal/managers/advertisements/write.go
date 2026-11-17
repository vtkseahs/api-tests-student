package advertisements

import (
	"net/http"
	"testing"
	"time"

	"api-tests-template/internal/client/http/advertisements"
	"api-tests-template/internal/managers/response"
)

// создаёт объявление и проверяет статус
func Create(t *testing.T, token string, fields map[string]string, photo []byte, status int) string {
	t.Helper()
	return response.Read(t, advertisements.Post(t, token, fields, photo), status)
}

// получает объявление по id
func Get(t *testing.T, token, id string) string {
	t.Helper()
	return response.Read(t, advertisements.Get(t, token, id), http.StatusOK)
}

// получает фотографии
func Photos(t *testing.T, token, id string) string {
	t.Helper()
	return response.Read(t, advertisements.Photos(t, token, id), http.StatusOK)
}

// удаляет объявление
func Delete(t *testing.T, token, id string) {
	t.Helper()
	response.Read(t, advertisements.Delete(t, token, id), http.StatusNoContent)
}

// удаляет временное объявление с повторами при 503
func CleanUp(t *testing.T, token, id string) {
	t.Helper()
	const attempts = 3
	for attempt := 1; attempt <= attempts; attempt++ {
		resp := advertisements.Delete(t, token, id)
		if resp.StatusCode != http.StatusServiceUnavailable || attempt == attempts {
			response.Read(t, resp, http.StatusNoContent, http.StatusNotFound)
			return
		}
		_ = resp.Body.Close()
		t.Logf("Очистка объявления %s: получили 503, повторяем запрос", id)
		time.Sleep(time.Duration(attempt) * time.Second)
	}
}
