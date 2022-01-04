package helpers

import "os"

type SocketAddr struct {
  HOST    string
  PORT    string
  ADDRESS string
}

func GetSocketAddr() SocketAddr {
  HOST := os.Getenv("HOST")
  if HOST == "" {
    HOST = "0.0.0.0"
  }

  PORT := os.Getenv("PORT")
  if PORT == "" {
    PORT = "10001"
  }

  ADDRESS := HOST + ":" + PORT

  return SocketAddr{HOST, PORT, ADDRESS}
}
