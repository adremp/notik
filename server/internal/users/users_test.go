package users_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	migrations "notik"
	"notik/internal/users"
	"notik/internal/users/http"
	"notik/internal/users/usecase"
	"notik/internal/users/users_repo"
	"notik/pkg/postgres"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func MigrateDb(t *testing.T, dbURI string) error {
	psql, err := goose.OpenDBWithDriver("postgres", dbURI)
	require.NoError(t, err)
	goose.SetBaseFS(migrations.Migrations)

	require.NoError(t, goose.SetDialect("postgres"))
	require.NoError(t, goose.Up(psql, "migrations"))

	return nil
}

func TestContainer(t *testing.T) {

	ctx := context.Background()

	pgEnv := map[string]string{
		"POSTGRES_PASSWORD": "postgres",
		"POSTGRES_USER":     "postgres",
		"POSTGRES_DB":       "pgtest",
	}

	creq := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres:latest",
			ExposedPorts: []string{"5432/tcp"},
			Env:          pgEnv,
			WaitingFor:   wait.ForListeningPort("5432/tcp"),
		},
		Started: true,
	}

	pgCont, err := testcontainers.GenericContainer(ctx, creq)
	if err != nil {
		t.Fatal(err)
	}

	pgPort, _ := pgCont.MappedPort(ctx, "5432")
	pgHost, _ := pgCont.Host(ctx)

	defer pgCont.Terminate(ctx)

	testDataSource := fmt.Sprintf("postgres://postgres:postgres@%s:%s/pgtest?sslmode=disable", pgHost, pgPort.Port())
	pgConn, err := postgres.NewPgConn(ctx, testDataSource)
	if err != nil {
		t.Fatal(err)
	}

	usersR := users_repo.New(pgConn)
	usersUc := usecase.New(usersR)
	usersH := http.New(usersUc)

	handler := usersH.Create()

	input := users.CreateInput{
		Username: "Andre",
		Email:    "arer@rt.rt",
		Password: "12345678",
	}

	inputBytes, err := json.Marshal(input)
	require.NoError(t, err)

	req := httptest.NewRequest("POST", "/", bytes.NewReader(inputBytes))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	ectx := echo.New().NewContext(req, res)

	MigrateDb(t, testDataSource)

	err = handler(ectx)
	require.Equal(t, res.Result().StatusCode, 201)
}
