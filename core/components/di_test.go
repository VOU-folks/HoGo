package components

import (
    "reflect"

    "testing"

    "gotest.tools/assert"
    is "gotest.tools/assert/cmp"
)

func TestGetDI(t *testing.T) {
    instance := GetDI()
    typeOfInstance := reflect.TypeOf(instance).String()
    assert.Assert(t, is.Equal(typeOfInstance, "*components.DI"))
}

func TestDIAppEngine(t *testing.T) {
    instance := GetDI().App
    typeOfInstance := reflect.TypeOf(instance).String()
    assert.Assert(t, is.Equal(typeOfInstance, "*gin.Engine"))
}

func TestDIServiceBus(t *testing.T) {
    instance := GetDI().ServiceBus
    typeOfInstance := reflect.TypeOf(instance).String()
    assert.Assert(t, is.Equal(typeOfInstance, "*components.ServiceBus"))
}
