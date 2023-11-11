package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

type DataSource struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func GetDefaultDataSource() string {
	return fmt.Sprintf("postgres://postgres:%v@%v:%v/%v?sslmode=%v", os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_SSLMODE"))
}

func NewPgConn(ctx context.Context, source string) (*pgx.Conn, error) {

	conn, err := pgx.Connect(ctx, source)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		return nil, err
	}

	return conn, nil
}
