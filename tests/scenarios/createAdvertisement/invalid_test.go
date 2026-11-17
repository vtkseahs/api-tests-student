package createAdvertisement

import (
	"net/http"
	"strings"
	"testing"

	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/response"
	"api-tests-template/tests/fixtures"
)

const titleLimit = 255

func TestInvalidAdvertisement(t *testing.T) {
	token := fixtures.Login(t)
	cases := []struct {
		name, field, value string
		omit               bool
	}{
		{name: "missing_title", field: "title", omit: true},
		{name: "missing_description", field: "description", omit: true},
		{name: "missing_photo", field: "photos", omit: true},
		{name: "empty_title", field: "title"},
		{name: "negative_price", field: "price", value: "-1"},
		{name: "title_too_long", field: "title", value: strings.Repeat("a", titleLimit+1)},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log("Подготовка: меняем одно поле объявления")
			fields, photo := fixtures.Advertisement()
			switch {
			case tc.field == "photos":
				photo = nil
			case tc.omit:
				delete(fields, tc.field)
			default:
				fields[tc.field] = tc.value
			}

			t.Log("Отправляем объявление и проверяем ошибку")
			body := advertisements.Create(t, token, fields, photo, http.StatusBadRequest)
			response.CheckError(t, body, "")
			t.Log("Получили 400 и сообщение об ошибке")
		})
	}
}
