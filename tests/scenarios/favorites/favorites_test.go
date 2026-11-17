package favorites

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"

	"api-tests-template/internal/managers/auth"
	"api-tests-template/internal/managers/favorites"
	"api-tests-template/tests/fixtures"
)

// создаёт объявление от второго аккаунта
func foreignAdvertisement(t *testing.T) (token, id, title string, remove func()) {
	t.Helper()
	token = fixtures.Login(t)
	ownerToken := secondToken(t)
	fields, photo := fixtures.Advertisement()
	id, remove = fixtures.CreateAdvertisement(t, ownerToken, fields, photo)
	return token, id, fields["title"], remove
}

// получает токен второго аккаунта
func secondToken(t *testing.T) string {
	t.Helper()
	ownerLogin := os.Getenv("TEST_SECOND_LOGIN")
	ownerPassword := os.Getenv("TEST_SECOND_PASSWORD")
	require.NotEmpty(t, ownerLogin, "задайте TEST_SECOND_LOGIN в .env.override")
	require.NotEmpty(t, ownerPassword, "задайте TEST_SECOND_PASSWORD в .env.override")
	require.NotEqual(t, os.Getenv("TEST_LOGIN"), ownerLogin, "нужен второй аккаунт TestBoard")
	return auth.Login(t, ownerLogin, ownerPassword).Token
}

// проверяет наличие в избранном
func checkFavorite(t *testing.T, token, id, title string, expected bool) {
	t.Helper()
	body := favorites.List(t, token, title, http.StatusOK)
	require.True(t, gjson.Get(body, "items").Exists(), "нет поля items: %s", body)
	found := gjson.Get(body, "items.#(advertisement.id=="+id+").advertisement.id").String()
	require.Equal(t, expected, found == id, "id %s, список избранного: %s", id, body)
	t.Logf("Объявление в избранном: %t", expected)
}

func TestFavorites(t *testing.T) {
	token, id, title, deleteAd := foreignAdvertisement(t)

	t.Log("Добавляем объявление в избранное")
	favorites.Add(t, token, id, http.StatusCreated)
	checkFavorite(t, token, id, title, true)

	t.Log("Удаляем объявление из избранного")
	favorites.Remove(t, token, id, http.StatusNoContent)
	checkFavorite(t, token, id, title, false)

	t.Log("Снова добавляем объявление и удаляем его от имени владельца")
	favorites.Add(t, token, id, http.StatusCreated)
	checkFavorite(t, token, id, title, true)
	deleteAd()
	checkFavorite(t, token, id, title, false)
}
