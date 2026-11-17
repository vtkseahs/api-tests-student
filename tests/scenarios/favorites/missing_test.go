package favorites

import (
	"net/http"
	"testing"

	"api-tests-template/internal/managers/favorites"
	"api-tests-template/internal/managers/response"
	"api-tests-template/tests/fixtures"
)

func TestMissingAdvertisement(t *testing.T) {
	token := fixtures.Login(t)
	fields, photo := fixtures.Advertisement()
	id, deleteAd := fixtures.CreateAdvertisement(t, token, fields, photo)
	t.Log("Подготовка: удаляем объявление, чтобы получить несуществующий id")
	deleteAd()

	t.Run("add", func(t *testing.T) {
		t.Log("Добавляем удалённое объявление")
		body := favorites.Add(t, token, id, http.StatusNotFound)
		response.CheckError(t, body, "")
		t.Log("Получили 404 и сообщение об ошибке")
	})
	t.Run("remove", func(t *testing.T) {
		t.Log("Удаляем из избранного несуществующее объявление")
		body := favorites.Remove(t, token, id, http.StatusNotFound)
		response.CheckError(t, body, "")
		t.Log("Получили 404 и сообщение об ошибке")
	})
}
