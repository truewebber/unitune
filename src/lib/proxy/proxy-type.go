package proxy

type (
	Type int
)

const (
	Socks5Type Type = iota + 1
	HttpType

	Socks5TypeString = "socks5"
	HttpTypeString   = "http"
)

func (t Type) String() string {
	switch t {
	case Socks5Type:
		return Socks5TypeString
	case HttpType:
		return HttpTypeString
	}

	return ""
}
