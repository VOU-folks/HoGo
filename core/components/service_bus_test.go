package components

import (
    "reflect"
    "testing"

    "gotest.tools/assert"
    is "gotest.tools/assert/cmp"
)

func TestGetServiceBus(t *testing.T) {
    instance := GetServiceBus()
    typeOfInstance := reflect.TypeOf(instance).String()
    assert.Assert(t, is.Equal(typeOfInstance, "*components.ServiceBus"))
}
