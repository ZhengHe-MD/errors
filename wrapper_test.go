package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEConflict(t *testing.T) {
	msg := "username already exists"
	err := EConflict(msg)
	errV := err.(*Error)
	assert.Equal(t, Conflict, errV.Class)
	assert.Equal(t, Op("TestEConflict"), errV.Op)
	assert.Equal(t, msg, errV.Msg)
}

func TestEInternal(t *testing.T) {
	msg := "connection reset by peer"
	err := EInternal(msg)
	errV := err.(*Error)
	assert.Equal(t, Internal, errV.Class)
	assert.Equal(t, Op("TestEInternal"), errV.Op)
	assert.Equal(t, msg, errV.Msg)
}

func TestEInvalid(t *testing.T) {
	msg := "value should be larger than 1"
	err := EInvalid(msg)
	errV := err.(*Error)
	assert.Equal(t, Invalid, errV.Class)
	assert.Equal(t, Op("TestEInvalid"), errV.Op)
	assert.Equal(t, msg, errV.Msg)
}

func TestENotFound(t *testing.T) {
	msg := "user not found"
	err := ENotFound(msg)
	errV := err.(*Error)
	assert.Equal(t, NotFound, errV.Class)
	assert.Equal(t, Op("TestENotFound"), errV.Op)
	assert.Equal(t, msg, errV.Msg)
}

func TestEPermDenied(t *testing.T) {
	msg := "only admin is allowed to change settings"
	err := EPermDenied(msg)
	errV := err.(*Error)
	assert.Equal(t, PermDenied, errV.Class)
	assert.Equal(t, Op("TestEPermDenied"), errV.Op)
	assert.Equal(t, msg, errV.Msg)
}
