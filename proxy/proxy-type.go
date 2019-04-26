package proxy

type (
	Type int
)

const (
	Socks5Type Type = iota + 1
	HttpType
	NullType

	Socks5TypeString = "socks5"
	HttpTypeString   = "http"
	NullTypeString   = "null"
)

func (t Type) String() string {
	switch t {
	case Socks5Type:
		return Socks5TypeString
	case HttpType:
		return HttpTypeString
	case NullType:
		return NullTypeString
	}

	return ""
}
