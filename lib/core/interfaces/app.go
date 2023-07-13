package interfaces

type App interface {
	Initializer
	HttpBinder
	Stopper
}
