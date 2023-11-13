package http_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"notik/config"
	"notik/internal/users"
	"notik/internal/users/http"
	mock_users "notik/internal/users/mocks"
	"notik/internal/users/usecase"
	"notik/internal/users/users_repo"
	"notik/pkg/logger"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {

	input := users.CreateInput{
		Username: "Andre",
		Password: "12345678",
	}

	appLogger := logger.NewApiLogger(&config.Config{})

	inputBytes, err := json.Marshal(input)
	if err != nil {
		t.Fatal(err)
	}

	ctl := gomock.NewController(t)
	userRepo := mock_users.NewMockRepo(ctl)
	userUc := usecase.New(userRepo, appLogger)
	userH := http.New(userUc, appLogger)
	handler := userH.Create()

	req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(inputBytes))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	res := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, res)

	userRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(users_repo.User{Username: "Andre"}, nil)

	err = handler(c)
	require.NoError(t, err)
	require.Equal(t, 201, res.Code)
}
