# round

Weighted Round-Robin written in Go （平滑权重轮询算法）

#### 1.get
```bash
go get github.com/chi-chu/round
```

#### 2.example
```bash
import github/chi-chu/round
func main(){
	r := round.NewRound()
    r.WithHeartBeatTime(10*time.Second)  //set the server heatbeat check time
    r.WithKeepAlive(true)				//set need to check
    //set the function use ip and port to check heartbeat
    r.WithCheckAlive(func(IP string,Port int)bool{
        //this function use to simulate The Server was Down
        return Port%7 != 0
    })
    r.AddServer(&round.Server{"1", 5, 0,"111:111:111",7, true})
    r.AddServer(&round.Server{"2", 1, 0,"222:222:222",2,true})
    r.AddServer(&round.Server{"3", 1, 0,"333:333:333",3,true})
    r.AddServer(&round.Server{"4", 10, 0,"444:444:444",123,true})
    r.AddServer(&round.Server{"5", 10, 0,"555:555:555",123,true})
    r.Start()
    for i:=0;i<20;i++{
        time.Sleep(time.Second)
        s,err := r.GetServer()
        if err != nil {
            panic(err)
        }
        fmt.Printf("choose : %s\n",s.Name)
    }
}
```

It will returns 
(server 1 lostNode 10s later)
```bash
choose : 1
choose : 4
choose : 5
choose : 1
choose : 2
choose : 3
choose : 1
choose : 4
choose : 5
choose : 5
choose : 2
choose : 5
choose : 3
choose : 5
choose : 5
choose : 2
choose : 5
choose : 3
choose : 5
choose : 5
```
