package response

import (
	"io"
	"net/http"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

// проверяет статус и читает ответ
func Read(t *testing.T, resp *http.Response, statuses ...int) string {
	t.Helper()
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.True(t, slices.Contains(statuses, resp.StatusCode), "ожидали HTTP %v, получили %d: %s", statuses, resp.StatusCode, body)
	return string(body)
}

// проверяет поля ошибки
func CheckError(t *testing.T, body, message string) {
	t.Helper()
	require.NotEmpty(t, gjson.Get(body, "error").String(), "нет кода ошибки: %s", body)
	actual := gjson.Get(body, "message").String()
	require.NotEmpty(t, actual, "нет сообщения об ошибке: %s", body)
	if message != "" {
		require.Equal(t, message, actual)
	}
}
