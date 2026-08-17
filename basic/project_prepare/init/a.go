package main

import (
	"fmt"
	_ "net/http/pprof" //在线pprof
	projectprepare "shy/go_basic/basic/project_prepare"

	_ "github.com/go-sql-driver/mysql" //注册mysql驱动
)

func InitLogger() {
	fmt.Println("init logger")
	fmt.Println("main是否匹配正则表达式", projectprepare.Reg.Match([]byte("hello123")))
}

func main() {
	projectprepare.CheckReg()
	InitLogger()
	InitDatabase()

	fmt.Println("server start")
}

func InitDatabase() {
	fmt.Println("init database")
}
