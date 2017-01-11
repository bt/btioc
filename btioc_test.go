package btioc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIocRegister(t *testing.T) {
	ioc := New()
	obj := "foobar"

	ioc.Register("foo", obj)

	assert.True(t, ioc.IsRegistered("foo"))
}

func TestIocRetrieve(t *testing.T) {
	ioc := New()
	obj := "foobar"

	ioc.Register("foo", obj)
	err, res := ioc.Retrieve("foo")

	assert.Nil(t, err)
	assert.Equal(t, "foobar", res)
}