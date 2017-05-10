package btioc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIoc_Register(t *testing.T) {
	ResetGlobalState()

	obj := "foobar"

	Register("foo", obj)

	assert.NotNil(t, Container()["foo"])
}

func TestIoc_IsRegistered(t *testing.T) {
	ResetGlobalState()

	obj := "foobar"

	Register("foo", obj)

	assert.True(t, IsRegistered("foo"))
}

func TestIoc_Retrieve(t *testing.T) {
	ResetGlobalState()

	obj := "foobar"

	Register("foo", obj)
	res, err := Retrieve("foo")

	assert.Nil(t, err)
	assert.Equal(t, "foobar", res)

	// Retrieve non-existent
	nonexist, err := Retrieve("nonexist")
	assert.EqualError(t, err, "'nonexist' is not registered in IoC container.")
	assert.Nil(t, nonexist)
}

func TestIoc_RetrieveIf(t *testing.T) {
	ResetGlobalState()

	obj := "foobar"

	// Register once first
	Register("foo", obj)

	// Cannot register again
	err := RegisterIf("foo", obj)

	assert.NotNil(t, err)
}

func TestIoc_RegisterIf(t *testing.T) {
	ResetGlobalState()

	obj := "foobar"

	// Register once first
	Register("foo", obj)

	var err error

	// Cannot register again
	err = RegisterIf("foo", obj)
	assert.NotNil(t, err)

	// Register with another name
	err = RegisterIf("foobaz", obj)
	assert.NotNil(t, Container()["foobaz"])
}

func TestIoc_Unregister(t *testing.T) {
	ResetGlobalState()

	obj := "foobar"

	// Register
	Register("foo", obj)

	// Unregister that
	Unregister("foo")

	assert.Nil(t, Container()["foo"])
}
