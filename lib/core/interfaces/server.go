package interfaces

type Server interface {
	Initializer
	Stopper
	Listener
}
