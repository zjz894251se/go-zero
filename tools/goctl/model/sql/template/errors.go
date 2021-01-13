package template

var Error = `package {{.pkg}}

import "github.com/zjz894251se/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound
`
