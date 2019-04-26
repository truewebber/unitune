package proxy

func GenerateProxyList(proxies []*Proxy) []HttpProxyClient {
	out := make([]HttpProxyClient, 0)

	for _, prx := range proxies {
		var p HttpProxyClient

		switch prx.Type {
		case Socks5Type:
			p = NewSocks5(prx)
		case HttpType:
			p = NewHttp(prx)
		default:
			continue
		}

		out = append(out, p)
	}

	return out
}
