package config

var proxySearch string

func SetProxySearch(ps string) {
	proxySearch = ps
	if proxySearch == "" {
		proxySearch = "LOCAL"
	}
}

func GetProxySearch() string {
	return proxySearch
}
