package interfaces

import "net/http"

type HttpBinder interface {
	BindToHttpServer(server *http.Server)
}
