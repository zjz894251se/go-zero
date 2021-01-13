package postgres

import (
	_ "github.com/lib/pq"
	"github.com/zjz894251se/go-zero/core/stores/sqlx"
)

const postgresDriverName = "postgres"

func NewPostgres(datasource string, opts ...sqlx.SqlOption) sqlx.SqlConn {
	return sqlx.NewSqlConn(postgresDriverName, datasource, opts...)
}
