package errors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFuncName(t *testing.T) {
	assert.Equal(t, "TestFuncName", FuncName())
}