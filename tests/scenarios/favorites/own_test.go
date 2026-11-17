package favorites

import (
	"net/http"
	"testing"

	"api-tests-template/internal/managers/favorites"
	"api-tests-template/internal/managers/response"
	"api-tests-template/tests/fixtures"
)

func TestAddOwnAdvertisement(t *testing.T) {
	token := fixtures.Login(t)
	fields, photo := fixtures.Advertisement()
	id, _ := fixtures.CreateAdvertisement(t, token, fields, photo)

	t.Log("Пробуем добавить в избранное своё объявление")
	body := favorites.Add(t, token, id, http.StatusForbidden)
	response.CheckError(t, body, "")
	checkFavorite(t, token, id, fields["title"], false)
}
