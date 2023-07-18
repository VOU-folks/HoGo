package interfaces

type DBConnection interface {
	Opener
	Closer
}
