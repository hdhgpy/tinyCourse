package balancer

//type HTTPProxy struct {
//	hostMap map[string]*httputil.ReverseProxy
//	lb      balancer.Balancer
//
//	sync.RWMutex // protect alive
//	alive        map[string]bool
//}

type Balancer interface {
	Add(string)
	Del(string)
	Balancer(string) (string, error)
	Inc(string)
	Done(string)
}
