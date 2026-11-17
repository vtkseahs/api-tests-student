package favorites

import (
	"net/http"
	"testing"

	"api-tests-template/internal/managers/favorites"
	"api-tests-template/internal/managers/response"
)

func TestAddFavoriteTwice(t *testing.T) {
	token, id, title, _ := foreignAdvertisement(t)

	t.Log("Добавляем объявление в избранное")
	favorites.Add(t, token, id, http.StatusCreated)

	t.Log("Повторяем запрос и проверяем конфликт")
	body := favorites.Add(t, token, id, http.StatusConflict)
	response.CheckError(t, body, "")
	checkFavorite(t, token, id, title, true)
}
