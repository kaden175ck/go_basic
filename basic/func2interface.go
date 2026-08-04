package main

type Selector interface {
	Select([]string) int
}

type ConnectionPool struct {
	Servers      []string
	LoadBalancer Selector //成员变量是接口类型，名字叫做LoadBalancer
}

func f1([]string) int {
	return 0
}

func f2([]string) int {
	return 1
}

type RoundRobin struct{}

//roundrobin这个结构体实现了这个函数，就意味着实现了selector这个接口

func (RoundRobin) Select(s []string) int { return f1(s) }

type Interleave struct{}

func (Interleave) Select(s []string) int { return f2(s) }

type ConnectionPool2 struct {
	Servers      []string
	LoadBalancer func([]string) int //成员变量是函数类型
}

func main29() {
	cp := ConnectionPool{
		Servers:      []string{"127.0.0.1:1234", "127.0.0.1:5678"},
		LoadBalancer: RoundRobin{},
		// LoadBalancer: Interleave{},
	}
	// _ = cp
	cp.LoadBalancer.Select(cp.Servers)

	cp2 := ConnectionPool2{
		Servers: []string{"127.0.0.1:1234", "127.0.0.1:5678"},
		// LoadBalancer: f1,
		LoadBalancer: f2,
	}
	// _ = cp2
	cp2.LoadBalancer(cp.Servers)
}
