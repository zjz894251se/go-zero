package template

var (
	Imports = `import (
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	"github.com/zjz894251se/go-zero/core/stores/cache"
	"github.com/zjz894251se/go-zero/core/stores/sqlc"
	"github.com/zjz894251se/go-zero/core/stores/sqlx"
	"github.com/zjz894251se/go-zero/core/stringx"
	"github.com/zjz894251se/go-zero/tools/goctl/model/sql/builderx"
)
`
	ImportsNoCache = `import (
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	"github.com/zjz894251se/go-zero/core/stores/sqlc"
	"github.com/zjz894251se/go-zero/core/stores/sqlx"
	"github.com/zjz894251se/go-zero/core/stringx"
	"github.com/zjz894251se/go-zero/tools/goctl/model/sql/builderx"
)
`
)
