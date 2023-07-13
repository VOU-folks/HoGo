package interfaces

type Stopper interface {
	Stop()
	IsStopped() bool
}
