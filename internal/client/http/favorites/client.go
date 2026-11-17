package favorites

import (
	"net/http"
	"testing"

	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
)

// получает список избранного
func List(t *testing.T, token, title string) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().Get(path.FavoritesPath).
		Query("search", title).Expect(t).End().Response
}

// добавляет объявление в избранное
func Add(t *testing.T, token, id string) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().Post(path.FavoritesPath).
		JSON(map[string]string{"advertisement_id": id}).Expect(t).End().Response
}

// удаляет объявление из избранного
func Remove(t *testing.T, token, id string) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().Delete(path.FavoritesPath + "/" + id).
		Expect(t).End().Response
}

// добавляет несколько объявлений в избранное
func AddBulk(t *testing.T, token string, ids []string) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().Post(path.FavoritesPath + "/bulk").
		JSON(map[string][]string{"advertisement_ids": ids}).Expect(t).End().Response
}
