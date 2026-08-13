package main_test

import (
	"fmt"
	"math/rand"
	randv2 "math/rand/v2"
	"testing"
)

const MAX = 1e9

func BenchmarkRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Intn(MAX)
	}
}

func BenchmarkRandV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randv2.IntN(MAX)
	}
}

// 本质还是下面的函数的逻辑，
// 用了randv2这个包下的intN这个函数，rander变成全局的，每次都不一样，和时间戳有关
func TestRand(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d  ", randv2.IntN(100))
	}
}

func TestRandSeed(t *testing.T) {
	// 初始种子都是123456，所以这个结果肯定是固定的
	source := randv2.NewPCG(123, 456)
	for i := 0; i < 5; i++ {
		// source.Seed(123, 456)
		// 如果把上面注释打开，那每次随机数都一样了
		rander := randv2.New(source)
		fmt.Printf("%d  ", rander.IntN(100))
		//每次生成随机数的时候会有一个新的种子，这个种子会作用到新的source里面去，
		//所以每次生成随机数都是一个新的种子在这个循环里
	}
}

// go test -v ./basic -run=^TestRand$ -count=1
// go test -v ./basic -run=^TestRandSeed$ -count=1
// go test ./basic -bench=Rand -run=^$ -count=1

// PS E:\Project\go_project\go_basic\basic> go test . -bench=Rand -run=^$ -count=1
// goos: windows
// goarch: amd64
// pkg: shy/go_basic/basic
// cpu: AMD Ryzen 7 4800H with Radeon Graphics
// BenchmarkRand-16        84267295                15.11 ns/op
// BenchmarkRandV2-16      147941905                7.794 ns/op
// PASS
// ok      shy/go_basic/basic      4.024s
// PS E:\Project\go_project\go_basic\basic>
