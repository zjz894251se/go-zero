package clickhouse

import (
	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/zjz894251se/go-zero/core/stores/sqlx"
)

const clickHouseDriverName = "clickhouse"

func New(datasource string, opts ...sqlx.SqlOption) sqlx.SqlConn {
	return sqlx.NewSqlConn(clickHouseDriverName, datasource, opts...)
}
