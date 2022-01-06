package components

import (
	"reflect"
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestNewHttpAppInstance(t *testing.T) {
	instance := NewHttpAppInstance()
	typeOfInstance := reflect.TypeOf(instance).String()
	assert.Assert(t, is.Equal(typeOfInstance, "*gin.Engine"))
}

func TestContextType(t *testing.T) {
	instance := &Context{}
	typeOfInstance := reflect.TypeOf(instance).String()
	assert.Assert(t, is.Equal(typeOfInstance, "*gin.Context"))
}

func TestEngineType(t *testing.T) {
	instance := &Engine{}
	typeOfInstance := reflect.TypeOf(instance).String()
	assert.Assert(t, is.Equal(typeOfInstance, "*gin.Engine"))
}
