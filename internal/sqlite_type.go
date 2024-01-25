package python

import (
	"log"
	"strings"

	"github.com/sqlc-dev/plugin-sdk-go/plugin"
	"github.com/sqlc-dev/plugin-sdk-go/sdk"
)

func sqliteType(req *plugin.GenerateRequest, col *plugin.Column) string {
	dt := strings.ToLower(sdk.DataType(col.Type))
	// notNull := col.NotNull || col.IsArray

	switch dt {
	case "int", "integer", "tinyint", "smallint", "mediumint", "bigint", "unsignedbigint", "int2", "int8":
		return "int"

	case "blob":
		return "memoryview"

	case "real", "double", "doubleprecision", "float":
		return "float"

	case "boolean", "bool":
		return "bool"

	case "date":
		return "datetime.date"

	case "datetime", "timestamp":
		return "datetime.datetime"

	case "any":
		return "Any"
	}

	switch {
	case strings.HasPrefix(dt, "character"),
		strings.HasPrefix(dt, "varchar"),
		strings.HasPrefix(dt, "varyingcharacter"),
		strings.HasPrefix(dt, "nchar"),
		strings.HasPrefix(dt, "nativecharacter"),
		strings.HasPrefix(dt, "nvarchar"),
		dt == "text",
		dt == "clob":
		return "str"

	case strings.HasPrefix(dt, "decimal"), dt == "numeric":
		return "float"

	default:
		log.Printf("unknown SQLite type: %s\n", dt)
		return "Any"
	}
}
