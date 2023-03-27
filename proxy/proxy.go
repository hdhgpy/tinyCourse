package proxy

import (
	"balancer/balancer"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

var (
	XRealIP       = http.CanonicalHeaderKey("X-Real-IP")
	XProxy        = http.CanonicalHeaderKey("X-Proxy")
	XForwardedFor = http.CanonicalHeaderKey("X-Forwarded-For")
)

var (
	ReverseProxy = "Balancer-Reverse-Proxy"
)

// HTTPProxy means reverse proxy in the balancer
type HTTPProxy struct {
	hostMap map[string]*httputil.ReverseProxy
	lb      balancer.Balancer

	sync.RWMutex // protect alive
	alive        map[string]bool
}

// NewHTTPProxy create a new reverse proxy with url and balancer algorithm
func NewHTTPProxy(targetHosts []string, algorithm string) (*HTTPProxy, error) {
	hosts := make([]string, 0)
	hostMap := make(map[string]*httputil.ReverseProxy)
	alive := make(map[string]bool)

	for _, targetHost := range targetHosts {
		targetUrl, err := url.Parse(targetHost) // 解析Url
		if err != nil {
			return nil, err
		}
		proxy := httputil.NewSingleHostReverseProxy(targetUrl)

		originDirector := proxy.Director

		proxy.Director = func(req *http.Request) {
			originDirector(req)
			req.Header.Set(XProxy, ReverseProxy)
			req.Header.Set(XRealIP, GetIP(req))
		}

		host := GetHost(targetUrl)
		alive[host] = true
		hostMap[host] = proxy
		hosts = append(hosts, host)
	}

	lb, err := balancer.Build(algo, hosts)
	if err != nil {
		return nil, err
	}

	return &HTTPProxy{
		hostMap: hostMap,
		lb:      lb,
		alive:   alive,
	}, nil
}
