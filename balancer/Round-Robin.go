package balancer

import "sync"

type RoundRobin struct {
	sync.RWMutex
	i     uint64
	hosts []string
}

func init() {
	factories[R2Balancer] = NewRoundRobin
}

func NewRoundRobin(hosts []string) Balancer {
	// todo
	/*...*/
	return &RoundRobin{hosts: hosts, i: 0}
}

func (r *RoundRobin) Balance(_ string) (string, error) {
	r.RLock()
	defer r.RUnlock()
	if len(r.hosts) == 0 {
		return "", NoHostError
	}
	host := r.hosts[r.i%uint64(len(r.hosts))]
	r.i++
	return host, nil
}

func (r *RoundRobin) Add(string) {
	
}

func (r *RoundRobin) Del(string) {

}

func (r *RoundRobin) Balancer(string) (string, error) {
	return "", nil
}

func (r *RoundRobin) Inc(string) {

}
func (r *RoundRobin) Done(string) {

}
