package errors

import (
	"github.com/VividCortex/mysqlerr"
	"github.com/go-sql-driver/mysql"
)

// NOTE: add other entries when necessary
var mySQLErrors = map[int]Class{
	mysqlerr.ER_DUP_ENTRY:           Conflict,
	mysqlerr.ER_INVALID_JSON_TEXT:   Invalid,
	mysqlerr.ER_DATA_TOO_LONG:       Invalid,
	mysqlerr.ER_ACCESS_DENIED_ERROR: PermDenied,
}

func WrapMySQLError(err error) error {
	mysqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		return err
	}

	if class, exists := mySQLErrors[int(mysqlErr.Number)]; exists {
		return E(err, class)
	}

	return E(err, Internal)
}
