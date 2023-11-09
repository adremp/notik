package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	migrations "notik"
	"notik/pkg/httpErrors"
	"notik/pkg/postgres"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type SetupConfig struct {
	Path, Method string
	Body         interface{}
}

func EqBodyMessageError(t *testing.T, res *httptest.ResponseRecorder, err error) {
	var bodyErr httpErrors.Error
	if err := json.NewDecoder(res.Body).Decode(&bodyErr); err != nil {
		t.Fatal(err)
	}
	require.Equal(t, bodyErr.ErrMessage, err.Error())
}

func SetupRequest(t *testing.T, config SetupConfig) (*httptest.ResponseRecorder, *echo.Echo, *echo.Context) {
	inputBytes, err := json.Marshal(config.Body)
	require.NoError(t, err)

	req := httptest.NewRequest(config.Method, config.Path, bytes.NewReader(inputBytes))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	e := echo.New()
	ectx := e.NewContext(req, res)
	return res, e, &ectx
}

func MigrateDb(t *testing.T, dbURI string) error {
	psql, err := goose.OpenDBWithDriver("postgres", dbURI)
	require.NoError(t, err)
	goose.SetBaseFS(migrations.Migrations)

	require.NoError(t, goose.SetDialect("postgres"))
	require.NoError(t, goose.Up(psql, "migrations"))

	return nil
}

func SetupTestPostgres(t *testing.T) (*pgx.Conn, func(), string) {
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

	testDataSource := fmt.Sprintf("postgres://postgres:postgres@%s:%s/pgtest?sslmode=disable", pgHost, pgPort.Port())
	pgConn, err := postgres.NewPgConn(ctx, testDataSource)
	if err != nil {
		t.Fatal(err)
	}

	return pgConn, func() {
		pgCont.Terminate(ctx)
	}, testDataSource
}

func GetRespCookie(t *testing.T, res *httptest.ResponseRecorder, key string) *http.Cookie {
	var cookie *http.Cookie
	for _, c := range res.Result().Cookies() {
		if c.Name == key {
			cookie = c
			break
		}
	}
	require.NotNil(t, cookie)
	return cookie
}
