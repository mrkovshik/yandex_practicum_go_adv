package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserViewHandler(t *testing.T) {
	type (
		want struct {
			code     int
			response string
		}
	)
	tests := []struct {
		name  string
		want  want
		users map[string]User
		url   string
	}{
		{
			name: "positive test #1",
			users: map[string]User{
				"u1": {
					ID:        "u1",
					FirstName: "Misha",
					LastName:  "Popov",
				},
				"u2": {
					ID:        "u2",
					FirstName: "Sasha",
					LastName:  "Popov",
				},
			},
			url: "/users/?user_id=u1",
			want: want{
				code:     200,
				response: `{"ID":"u1","FirstName":"Misha","LastName":"Popov"}`,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			request := httptest.NewRequest(http.MethodGet, test.url, nil)
			// создаём новый Recorder
			w := httptest.NewRecorder()
			UserViewHandler(test.users)(w, request)

			res := w.Result()
			// проверяем код ответа
			assert.Equal(t, test.want.code, res.StatusCode)
			// получаем и проверяем тело запроса
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)

			require.NoError(t, err)
			assert.JSONEq(t, test.want.response, string(resBody))
			assert.Equal(t, "application/json", res.Header.Get("Content-Type"))
		})
	}
}
