package advertisements

import (
	"encoding/json"
	"io"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/client/http/advertisements"
	"api-tests-template/internal/managers/advertisements/models"
)

func GetAdvertisementsList(t *testing.T, token, text string) models.AdvertisementsList {
	response := advertisements.HttpGetAdvertisements(t, token, text)
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	var list models.AdvertisementsList
	require.NoError(t, json.Unmarshal(body, &list))
	return list
}
