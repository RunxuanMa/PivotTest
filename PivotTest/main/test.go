package main

import (
	"fmt"
	"github.com/Pivot-Studio/mapq"
)

func main() {
	a:=mapq.Lexer{}
	a.SetInput("a==1&&!(b==2||b==3)")
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

