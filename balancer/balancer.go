package balancer

import "errors"

const (
	IPHashBalancer         = "ip-hash"
	ConsistentHashBalancer = "consistent-hash"
	P2CBalancer            = "p2c"
	RandomBalancer         = "random"
	R2Balancer             = "round-robin"
	LeastLoadBalancer      = "least-load"
	BoundedBalancer        = "bounded"
)

//type HTTPProxy struct {
//	hostMap map[string]*httputil.ReverseProxy
//	lb      balancer.Balancer
//
//	sync.RWMutex // protect alive
//	alive        map[string]bool
//}

var (
	NoHostError                = errors.New("no host")
	AlgorithmNotSupportedError = errors.New("algorithm not supported")
)

type Balancer interface {
	Add(string)
	Del(string)
	Balancer(string) (string, error)
	Inc(string)
	Done(string)
}

type Factory func([]string) Balancer

var factories = make(map[string]Factory)

// Build generates the corresponding Balancer according to the algorithm
func Build(algorithm string, hosts []string) (Balancer, error) {
	factory, ok := factories[algorithm]
	if !ok {
		return nil, AlgorithmNotSupportedError
	}
	return factory(hosts), nil
}
