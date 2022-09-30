package main

import (
	"fmt"
	"github.com/Pivot-Studio/mapq"
)

func main() {
	a:=mapq.Lexer{}
	a.SetInput("(a+b)*b<5")
	scan := a.Scan()
	for true {
		fmt.Println(scan.Value)
		if scan.SideNode != nil {
			scan=*scan.SideNode
		} else {
			break
		}
	}
}

