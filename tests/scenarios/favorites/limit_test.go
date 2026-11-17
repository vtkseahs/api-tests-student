package favorites

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"

	"api-tests-template/internal/managers/favorites"
	"api-tests-template/internal/managers/response"
	"api-tests-template/tests/fixtures"
)

const (
	favoritesLimit  = 50
	fixtureInterval = 300 * time.Millisecond
	cleanupPause    = 2 * time.Second
)

func TestFavoritesLimit(t *testing.T) {
	token := fixtures.Login(t)
	owner := secondToken(t)

	t.Log("Подготовка: определяем, сколько мест осталось в избранном")
	list := favorites.List(t, token, "", http.StatusOK)
	total := gjson.Get(list, "total")
	require.Equal(t, gjson.Number, total.Type, "нет количества избранных: %s", list)
	remaining := favoritesLimit - int(total.Int())
	require.GreaterOrEqual(t, remaining, 0, "в избранном уже больше %d объявлений", favoritesLimit)
	require.LessOrEqual(t, remaining, favoritesLimit, "некорректное количество избранных")

	t.Logf("Создаём %d временных объявлений", remaining+1)
	t.Cleanup(func() { time.Sleep(cleanupPause) })
	ids := make([]string, 0, remaining)
	for range remaining {
		fields, photo := fixtures.Advertisement()
		id, _ := fixtures.CreateAdvertisement(t, owner, fields, photo)
		ids = append(ids, id)
		time.Sleep(fixtureInterval)
	}

	if len(ids) > 0 {
		t.Logf("Заполняем оставшиеся места в избранном: %d", len(ids))
		result := favorites.AddBulk(t, token, ids)
		require.Len(t, gjson.Get(result, "added").Array(), remaining, "результат добавления: %s", result)
	}

	fields, photo := fixtures.Advertisement()
	id, _ := fixtures.CreateAdvertisement(t, owner, fields, photo)

	t.Log("Пробуем добавить объявление сверх лимита")
	body := favorites.Add(t, token, id, http.StatusBadRequest)
	response.CheckError(t, body, "")
	checkFavorite(t, token, id, fields["title"], false)
	t.Log("Получили 400, дополнительное объявление не добавилось")
}
