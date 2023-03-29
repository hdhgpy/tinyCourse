package balancer

import (
	"math/rand"
	"sync"
	"time"
)

type Random struct {
	sync.RWMutex
	hosts []string
	rnd   *rand.Rand
}

func init() {
	factories[RandomBalancer] = NewRandom
}

func NewRandom(hosts []string) Balancer {
	return &Random{
		hosts: hosts,
		rnd:   rand.New(rand.NewSource(time.Now().UnixNano()))}
}

func (r *Random) Add(string) {

}

func (r *Random) Del(string) {

}

func (r *Random) Balancer(string) (string, error) {
	return "", nil
}

func (r *Random) Inc(string) {

}
func (r *Random) Done(string) {

}
