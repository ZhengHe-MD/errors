package errors

import (
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWrapMySQLError(t *testing.T) {
	tests := map[string]struct {
		givenError     error
		wantErrorClass Class
	}{
		"3140": {
			givenError: &mysql.MySQLError{
				Number:  3140,
				Message: "Invalid JSON text",
			},
			wantErrorClass: Invalid,
		},
		"1062": {
			givenError: &mysql.MySQLError{
				Number:  1062,
				Message: "Duplicate entry '1' for key 'PRIMARY'",
			},
			wantErrorClass: Conflict,
		},
		"1045": {
			givenError: &mysql.MySQLError{
				Number:  1045,
				Message: "Access Denied for user 'root'@'localhost'",
			},
			wantErrorClass: PermDenied,
		},
		"1406": {
			givenError: &mysql.MySQLError{
				Number:  1406,
				Message: "Data too long for column 'content' at row 1",
			},
			wantErrorClass: Invalid,
		},
		"unknown": {
			givenError: New("test"),
		},
		"default": {
			givenError: &mysql.MySQLError{
				Number:  0,
				Message: "",
			},
			wantErrorClass: Internal,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			e := WrapMySQLError(tc.givenError)
			if name == "unknown" {
				assert.False(t, Is(e, tc.wantErrorClass))
			} else {
				assert.True(t, Is(e, tc.wantErrorClass))
			}
		})
	}
}
