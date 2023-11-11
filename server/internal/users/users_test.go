package users_test

import (
	"net/http/httptest"
	"notik/internal/users"
	"notik/internal/users/http"
	"notik/internal/users/usecase"
	"notik/internal/users/users_repo"
	"notik/pkg/httpErrors"
	"notik/pkg/tests"
	"notik/pkg/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContainer(t *testing.T) {
	t.Parallel()

	type Case struct {
		Name    string
		Body    interface{}
		CheckFn func(res *httptest.ResponseRecorder)
	}

	cases := []Case{
		{
			Name: "Success",
			Body: users.CreateInput{
				Username: "Andre1",
				Email:    "arer@rt.rt",
				Password: "12345678",
			}, CheckFn: func(res *httptest.ResponseRecorder) {
				tokenCookie := tests.GetRespCookie(t, res, "token")
				tokenEl, err := utils.ParseToken(tokenCookie.Value)
				require.NoError(t, err)
				require.Equal(t, tokenEl.Email, "arer@rt.rt")
				require.Equal(t, res.Result().StatusCode, 201)
			},
		},
		{
			Name: "Email exist",
			Body: users.CreateInput{
				Username: "Andre1",
				Email:    "arer@rt.rt",
				Password: "12345678",
			}, CheckFn: func(res *httptest.ResponseRecorder) {
				tests.EqBodyMessageError(t, res, httpErrors.ErrEmailExist)
				require.Equal(t, res.Result().StatusCode, 400)
			},
		},
		{
			Name: "Invalid email",
			Body: users.CreateInput{
				Username: "Andrew2",
				Email:    "adrer",
				Password: "123143",
			},
			CheckFn: func(res *httptest.ResponseRecorder) {
				require.Equal(t, res.Result().StatusCode, 400)
			},
		},
		{
			Name: "Empty email",
			Body: users.CreateInput{
				Username: "Andrew2",
				Email:    "",
				Password: "123143",
			},
			CheckFn: func(res *httptest.ResponseRecorder) {
				require.Equal(t, res.Result().StatusCode, 400)
			},
		},
		{
			Name: "Empty password",
			Body: users.CreateInput{
				Username: "Andrew2",
				Email:    "qwe@qwe.fr",
				Password: "",
			},
			CheckFn: func(res *httptest.ResponseRecorder) {
				require.Equal(t, res.Result().StatusCode, 400)
			},
		},
	}

	pgConn, Close, testSource := tests.SetupTestPostgres(t)
	defer Close()
	tests.MigrateDb(t, testSource)

	usersR := users_repo.New(pgConn)
	usersUc := usecase.New(usersR)
	usersH := http.New(usersUc)
	handler := usersH.Create()

	for _, caseEl := range cases {
		t.Run(caseEl.Name, func(t *testing.T) {
			res, _, ectx := tests.SetupRequest(t, tests.SetupConfig{
				Path:   "/",
				Method: "POST",
				Body:   caseEl.Body,
			})
			handler(*ectx)
			caseEl.CheckFn(res)
		})
	}
}
