package createAdvertisement

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"

	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/advertisements/models"
	"api-tests-template/tests/fixtures"
)

func TestCreateAdvertisement(t *testing.T) {
	token := fixtures.Login(t)
	fields, photo := fixtures.Advertisement()
	id, _ := fixtures.CreateAdvertisement(t, token, fields, photo)

	t.Log("Проверяем поля созданного объявления")
	body := advertisements.Get(t, token, id)
	for field, value := range fields {
		require.Equal(t, value, gjson.Get(body, field).String(), "поле %s", field)
	}
	require.Equal(t, id, gjson.Get(body, "id").String())

	t.Log("Ищем объявление по полному названию")
	list := advertisements.GetAdvertisementsList(t, token, fields["title"])
	require.Contains(t, list.Items, models.Advertisement{ID: id, Title: fields["title"]}, "объявление не найдено")

	t.Log("Проверяем загруженную фотографию")
	photos := gjson.Parse(advertisements.Photos(t, token, id)).Array()
	require.Len(t, photos, 1, "ожидалась одна фотография")
	require.NotEmpty(t, photos[0].Get("id").String(), "нет id фотографии")
	require.NotEmpty(t, photos[0].Get("url").String(), "нет URL фотографии")
	t.Log("Поля, поиск и фотография проверены")
}
