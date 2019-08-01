package round

import (
	"sync"
	"time"
)

type Round struct {
	ServerList		[]*Server
	KeepAlive		bool
	HeartBeatTime	time.Duration
	Lock            sync.Mutex
	totalWeight		int
	CheckAlive		func(string,int)bool
}

func NewRound() *Round{
	r := new(Round)
	return r
}

func (r *Round) WithKeepAlive(b bool){
	r.KeepAlive = b
}

func (r *Round) WithHeartBeatTime(t time.Duration){
	r.HeartBeatTime = t
}

func (r *Round) Start(){
	if r.KeepAlive {
		if r.CheckAlive == nil {
			return
		}
		go func(){
			for{
				time.Sleep(r.HeartBeatTime)
				for _,v := range r.ServerList {
					if r.CheckAlive(v.IP,v.Port) {
						if !v.Alive {
							r.Lock.Lock()
							v.Alive = true
							r.totalWeight += v.Weight
							r.reSetCurrentWeight()
							r.Lock.Unlock()
						}
					}else{
						if v.Alive {
							r.Lock.Lock()
							v.Alive = false
							r.totalWeight -= v.Weight
							r.reSetCurrentWeight()
							r.Lock.Unlock()
						}
					}
				}
			}
		}()
	}
}

func (r *Round) AddServer(s *Server)error{
    if s.Weight<=0{
        return ErrInvalidWeight
    }
    s.CurrentWeight = s.Weight
	r.ServerList = append(r.ServerList, s)
	r.addTotalWeight(s.Weight)
	return nil
}

func (r *Round) GetServer()(*Server, error){
	index := -1
	maxWeight := 0
	r.Lock.Lock()
	defer r.Lock.Unlock()
	for k,v := range r.ServerList {
		if v.Alive {
			if v.CurrentWeight > maxWeight{
				index = k
				maxWeight = v.CurrentWeight
			}
			v.CurrentWeight += v.Weight
		}
	}
	if index<0{
		return &Server{}, ErrNoServerAlive
	}
	r.ServerList[index].CurrentWeight -= r.totalWeight
	return r.ServerList[index], nil
}

func (r *Round) setTotalWeight(w int){
	r.Lock.Lock()
	defer r.Lock.Unlock()
	r.totalWeight = w
}

func (r *Round) addTotalWeight(w int){
	r.Lock.Lock()
	defer r.Lock.Unlock()
	r.totalWeight += w
}

func (r *Round) reduceTotalWeight(w int){
	r.Lock.Lock()
	defer r.Lock.Unlock()
	r.totalWeight -= w
}

func (r *Round) reSetCurrentWeight(){
	for _,v := range r.ServerList {
		v.CurrentWeight = v.Weight
	}
}

func (r *Round) WithCheckAlive(fn func(string,int)bool){
	r.CheckAlive = fn
}