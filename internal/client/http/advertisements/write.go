package advertisements

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
)

// отправляет объявление с фото
func Post(t *testing.T, token string, fields map[string]string, photo []byte) *http.Response {
	t.Helper()
	var body bytes.Buffer
	form := multipart.NewWriter(&body)
	for name, value := range fields {
		require.NoError(t, form.WriteField(name, value))
	}
	if len(photo) > 0 {
		part, err := form.CreateFormFile("photos", "photo.jpg")
		require.NoError(t, err)
		_, err = part.Write(photo)
		require.NoError(t, err)
	}
	require.NoError(t, form.Close())

	return apiRunner.GetRunner().Auth(token).Create().Post(path.AdvertisementPath).
		ContentType(form.FormDataContentType()).Body(body.String()).Expect(t).End().Response
}

// получает объявление по id
func Get(t *testing.T, token, id string) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().Get(path.AdvertisementPath).Query("id", id).Expect(t).End().Response
}

// получает фотографии
func Photos(t *testing.T, token, id string) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().Get(path.AdvertisementsPath + "/" + id + "/photos").Expect(t).End().Response
}

// удаляет объявление
func Delete(t *testing.T, token, id string) *http.Response {
	return apiRunner.GetRunner().Auth(token).Create().Delete(path.AdvertisementPath).
		Query("id", id).Expect(t).End().Response
}
