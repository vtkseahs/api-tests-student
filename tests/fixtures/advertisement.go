package fixtures

import (
	"crypto/rand"
	_ "embed"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"

	"api-tests-template/internal/managers/advertisements"
)

const InvalidToken = "incorrect_token"
const InvalidAuthMessage = "Invalid or expired token"

//go:embed photo.jpg
var advertisementPhoto string

// готовит данные и фото
func Advertisement() (map[string]string, []byte) {
	return map[string]string{
		"title":       "Тестовое объявление " + rand.Text(),
		"description": "Описание тестового объявления",
		"price":       "125.5",
		"quantity":    "3",
	}, []byte(advertisementPhoto)
}

// создаёт временное объявление и регистрирует очистку
func CreateAdvertisement(t *testing.T, token string, fields map[string]string, photo []byte) (string, func()) {
	t.Helper()
	body := advertisements.Create(t, token, fields, photo, http.StatusCreated)
	id := gjson.Get(body, "id").String()
	require.NotEmpty(t, id, "нет id объявления: %s", body)
	deleted := false
	remove := func() {
		t.Helper()
		if !deleted {
			advertisements.Delete(t, token, id)
			deleted = true
		}
	}
	t.Cleanup(func() {
		if !deleted {
			advertisements.CleanUp(t, token, id)
		}
	})
	return id, remove
}
