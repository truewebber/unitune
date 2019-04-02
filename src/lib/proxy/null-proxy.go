package proxy

import (
	"net/http"
)

type (
	Null struct{}
)

func NewNull() *Null {
	return &Null{}
}

func (n *Null) HttpClient() *http.Client {
	return &http.Client{}
}

func (n *Null) Type() Type {
	return 0
}
