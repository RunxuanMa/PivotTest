package mapq

import (
	"errors"
	"regexp"
	"strings"
)

const (
	TYPE_PLUS      = iota // "+"
	TYPE_SUB              // "-"
	TYPE_MUL              // "*"
	TYPE_DIV              // "/"
	TYPE_LP               // "("
	TYPE_RP               // ")"
	TYPE_VAR              // "([a-z]|[A-Z])([a-z]|[A-Z]|[0-9])*"
	TYPE_RES_TRUE         // "true"
	TYPE_RES_FALSE        // "false"
	TYPE_AND              // "&&"
	TYPE_OR               // "||"
	TYPE_EQ               // "=="
	TYPE_LG               // ">"
	TYPE_SM               // "<"
	TYPE_LEQ              // ">="
	TYPE_SEQ              // "<="
	TYPE_NEQ              // "!="
	TYPE_STR              // a quoted string(单引号)
	TYPE_INT              // an integer
	TYPE_FLOAT            // 小数，x.y这种
	TYPE_UNKNOWN          // 未知的类型
	TYPE_NOT              // "!"
	TYPE_DOT              // "."
	TYPE_RES_NULL         // "null"
)

// Lexer 词法分析器
type Lexer struct {
	input string
	pos   int
	runes []rune
}

// SetInput 设置输入
func (l *Lexer) SetInput(s string) {
	l.input=s
	l.runes=[]rune(s)
}

// Peek 看下一个字符
func (l *Lexer) Peek() (ch rune, end bool) {

	if l.pos> len(l.runes)-1 {
		return 0,false
	}
	l.pos++
	return l.runes[l.pos-1],true

}

// some finction maybe useful for your implementation
//是字母吗
func isLetter(ch rune) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z')
}
//字母or下划线
func isLetterOrUnderscore(ch rune) bool {
	return isLetter(ch) || ch == '_'
}
//数字
func isNum(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

// Checkpoint 检查点
type Checkpoint struct {
	pos int
}

// SetCheckpoint 设置检查点
func (l *Lexer) SetCheckpoint() Checkpoint {
	return Checkpoint{pos: l.pos}
}

// GobackTo 回到一个检查点
func (l *Lexer) GobackTo(c Checkpoint) {
	l.pos=c.pos
}

// ScanType 扫描一个特定Token，下一个token不是这个类型则自动回退，返回err
func (l *Lexer) ScanType(code int) (token string, err error) {

	defer func() {
		i := recover()
		if i!=nil {
			err=errors.New("not implemented")
		}
	}()

	str:=l.input
	var flag =false

	switch code{
	case 0 :
		flag = strings.Contains(str, "+")
		token="+"
		break
	case 1:

	case 2:
	case 3:
	case 4:
	case 5:
	case 6:
		flag,token=Regexp(str,"([a-z]|[A-Z]|_)([a-z]|[A-Z]|[0-9]|_)*")
		break
	case 7:
		flag = strings.Contains(str, "true")
		token="true"
		break
	case 8:
	case 9:
	case 10:
	case 11:
	case 12:
		flag = strings.Contains(str, ">")
		token=">"
		break
	case 13:
	case 14:
		flag = strings.Contains(str, ">=")
		token=">="
		break
	case 15:
	case 16:
	case 17:
		copy:=str
		token= strings.ReplaceAll(str,"'","")
		flag= len(copy)-len(token)==2
		break
	case 18:
		flag,token=Regexp(str,"[0-9]+")
		break
	case 19:
		flag,token=Regexp(str,"[0-9]+(.[0-9]+)")
		break
	case 20:
	case 21:
	case 22:
	case 23:
	default:
		break
	}

	if !flag {
		token=""
		panic("not implemented")
	}

	return token, err
}

// Scan scan a token
func (l *Lexer) Scan() (node Node) {
	root:=Node{}

	startNum:=0

	l.scan(&root,startNum,0)

	return root
}
func (l *Lexer) scan(node *Node,startNum int,id int) () {
	if startNum >= len(l.runes) {
		return
	}
	str:=string(l.runes[startNum])


	node.SideNode=new(Node)


	if flag, _ := Regexp(str, "([a-zA-Z])"); flag {
		node.Value=str
		node.Type=TYPE_VAR
		node.Id=id
		id++

		if startNum+1 >= len(l.runes) {
			return
		}
		//true null false特殊情况
		if str == "f"&&string(l.runes[startNum+1])=="a" {
			node.Value="FALSE"
			node.Type=TYPE_RES_FALSE
			startNum+=5
			node.Id=id
			id++
			l.scan(node.SideNode,startNum,id)
		}else if str == "t"&&string(l.runes[startNum+1])=="r" {
			node.Value="TRUE"
			node.Type=TYPE_RES_TRUE
			startNum+=4
			node.Id=id
			id++
			l.scan(node.SideNode,startNum,id)
		}else if str == "n"&&string(l.runes[startNum+1])=="u" {
			node.Value="NULL"
			node.Type=TYPE_RES_NULL
			startNum+=4
			node.Id=id
			id++
			l.scan(node.SideNode,startNum,id)
		}else {

			for true {
				startNum++
				if startNum == len(l.runes) {
					return
				}
				str = string(l.runes[startNum])
				flag2 := str == "_"
				if flag, _ := Regexp(str, "[a-zA-Z]"); flag || flag2 {
					node.Value += str
				} else {
					break
				}

			}
			l.scan(node.SideNode, startNum, id)
		}
	}else if str == "|" {
		node.Value="_OR"
		node.Type=TYPE_OR
		startNum+=2
		node.Id=id
		id++
		l.scan(node.SideNode,startNum,id)
	}else if str == "&" {
		node.Value="_AND"
		node.Type=TYPE_AND
		startNum+=2
		node.Id=id
		id++
		l.scan(node.SideNode,startNum,id)
	}else if str=="!" {
		if string(l.runes[startNum+1])=="=" {
			node.Value="_NEQ"
			node.Type=TYPE_NEQ
			startNum+=2
			node.Id=id
			id++
			l.scan(node.SideNode,startNum,id)
		}else {
			node.Value="_NOT"
			node.Type=TYPE_NOT
			startNum+=1
			node.Id=id
			id++
			l.scan(node.SideNode,startNum,id)
		}

	}else if str=="=" {
		node.Value="_EQ"
		node.Type=TYPE_EQ
		startNum+=2
		node.Id=id
		id++
		l.scan(node.SideNode,startNum,id)
	}else if str == ">"&&string(l.runes[startNum+1])=="=" {
		node.Value="_LRE"
		node.Type=TYPE_LEQ
		startNum+=2
		node.Id=id
		id++
		l.scan(node.SideNode,startNum,id)
	}else if str == ">"&&string(l.runes[startNum+1])!="=" {
		node.Value="_LR"
		node.Type=TYPE_LG
		startNum+=1
		node.Id=id
		id++
		l.scan(node.SideNode,startNum,id)
	}else if str == "<"&&string(l.runes[startNum+1])=="=" {
		node.Value="_SME"
		node.Type=TYPE_SEQ
		startNum+=2
		node.Id=id
		id++
		l.scan(node.SideNode,startNum,id)
	}else if str == "<"&&string(l.runes[startNum+1])!="=" {
		node.Value="_SM"
		node.Type=TYPE_SM
		startNum+=1
		node.Id=id
		id++
		l.scan(node.SideNode,startNum,id)
	}else if flag, _ := Regexp(str, "[0-9]"); flag {
		node.Value=str
		node.Type=TYPE_INT

		for true {
			startNum++
			if startNum == len(l.runes) {
				return
			}
			str=string(l.runes[startNum])
			if flag, _ := Regexp(str, "[0-9]"); flag {
				node.Value+=str
			}else if str=="." {
				node.Value+=str
				node.Type=TYPE_FLOAT
			}else {
				break
			}

		}
		
	//、、startNum+=1
		node.Id=id
		id++
		l.scan(node.SideNode,startNum,id)
	}else if str=="(" {
		node.Value="("
		node.Type=TYPE_LP
		startNum+=1
		node.Id=id
		id++
		l.scan(node.SideNode,startNum,id)
	}else if str==")" {
		node.Value=")"
		node.Type=TYPE_RP
		startNum+=1
		node.Id=id
		id++
		l.scan(node.SideNode,startNum,id)
	}else if str=="+" {
		node.Value="+"
		node.Type=TYPE_PLUS
		startNum+=1
		node.Id=id
		id++
		l.scan(node.SideNode,startNum,id)
	}else if str=="-" {
		node.Value="-"
		node.Type=TYPE_SUB
		startNum+=1
		node.Id=id
		id++
		l.scan(node.SideNode,startNum,id)
	}else if str=="*" {
		node.Value="*"
		node.Type=TYPE_MUL
		startNum+=1
		node.Id=id
		id++
		l.scan(node.SideNode,startNum,id)
	}else if str=="/" {
		node.Value="/"
		node.Type=TYPE_DIV
		startNum+=1
		node.Id=id
		id++
		l.scan(node.SideNode,startNum,id)
	}else if str=="." {
		node.Value="."
		node.Type=TYPE_DOT
		startNum+=1
		node.Id=id
		id++
		l.scan(node.SideNode,startNum,id)
	}else if str=="'" {
		for true {
			str=string(l.runes[startNum+1])
			node.SideNode=new(Node)
			node.Value=str
			node.Type=TYPE_STR
			startNum+=1
			node.Id=id
			id++
			if string(l.runes[startNum+1])=="'" {
				startNum+=2
				break
			}
			node=node.SideNode
		}
		l.scan(node.SideNode,startNum,id)
	}




}
//正则判断
func Regexp(str string,reg string)(bool,string)  {
	matchString, _ := regexp.MatchString(reg, str)


	compile:= regexp.MustCompile(reg)

	return matchString,compile.FindString(str)
}
