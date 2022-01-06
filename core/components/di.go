package components

type DI struct {
	ServiceBus *ServiceBus
	App        *Engine
}

var di *DI

func GetDI() *DI {
	if di == nil {
		di = &DI{}
	}
	return di
}
