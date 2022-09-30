package mapq

import (
	"reflect"
	"strconv"
)

// Node 节点
type Node struct {

	//Eval(data map[string]interface{}) interface{}
	SideNode *Node
	Value string
	Type int
	Num int
	Bool bool
	Id int

}


// 这些函数提供给你，也许可以帮上忙。。。
func toF64(i interface{}) float64 {
	switch v := i.(type) {
	case int:
		return float64(v)
	case float64:
		return v
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case uint32:
		return float64(v)
	case uint64:
		return float64(v)
	case float32:
		return float64(v)
	}
	return 0
}
func trytoF64(i interface{}) interface{} {
	switch v := i.(type) {
	case int, float64, int32, int64, uint32, uint64, float32:
		return toF64(v)

	}
	return i
}

func equal(left, right interface{}) bool {
	return reflect.DeepEqual(trytoF64(left), trytoF64(right))
}

//// BinNode 双目表达式节点
//type BinNode struct {
//	Left, Right Node
//	Op          int
//}

// Eval 查询
// 对map 里的数据和 node链表作匹配  并得到一个 由 || && ! () true false 组成的切片
func (n *Node) Eval(data map[string]interface{}) interface{} {

	var identifier []interface{}

	for true {
		if n.SideNode != nil {

			Type:=n.Type
			switch Type {
			case TYPE_VAR:
				va := data[n.Value]

				_, b :=va.(int)
				if !b&&!n.Bool { //不是int类型
					_,bool1:=va.(map[string]interface{})
					if bool1 {//map类型
						aMap:=va.(map[string]interface{})
						num:= aMap[n.SideNode.SideNode.Value]

						switch n.SideNode.SideNode.SideNode.Type {
						case TYPE_EQ:
							if i, _ := strconv.Atoi(n.SideNode.SideNode.SideNode.SideNode.Value); i == num {
								n=n.SideNode.SideNode.SideNode
								identifier = append(identifier, true)
							} else {
								n=n.SideNode.SideNode.SideNode
								identifier = append(identifier, false)
							}
							break
						}
					}
					if _, bool2 := va.(string); bool2 {
						str:=va.(string)
						var strFromLinkList string
						relationBetweenTwoStrs:=n.SideNode.Type
						n=n.SideNode.SideNode
						for true {
							if n.Type == TYPE_STR {
								strFromLinkList+=n.Value
								n=n.SideNode
							}else {
								break
							}
						}


						switch relationBetweenTwoStrs {
						case TYPE_EQ:
							if str==strFromLinkList {
								identifier = append(identifier, true)
							}else {
								identifier = append(identifier, false)
							}
							break
						case TYPE_NEQ:
							if str==strFromLinkList {
								identifier = append(identifier, false)
							}else {
								identifier = append(identifier, true)
							}
						}



					}
					if _, bool3 := va.(float64); bool3{
						num:=va.(float64)
						relationBetweenTwoNums:=n.SideNode.Type
						n=n.SideNode.SideNode
						num2, _ := strconv.ParseFloat(n.Value, 64)

						switch relationBetweenTwoNums {
						case TYPE_EQ:
							if num == num2 {
								identifier = append(identifier, true)
							}else {
								identifier = append(identifier, false)
							}
							break
						case TYPE_NEQ:
							if num != num2 {
								identifier = append(identifier, true)
							}else {
								identifier = append(identifier, false)
							}
							break
						}

					}
				} else if va==nil&&!n.Bool { //null特殊情况
					Rel := n.SideNode.Type
					switch Rel {
					case TYPE_EQ:
						if n.SideNode.SideNode.Type == TYPE_RES_NULL {
							identifier = append(identifier, true)
						}else {
							identifier = append(identifier, false)
						}
						break
					case TYPE_NEQ:
						if n.SideNode.SideNode.Type == TYPE_RES_NULL {
							identifier = append(identifier, false)
						}else {
							identifier = append(identifier, true)
						}
						break
					}
				} else {
					var i int
					if n.Bool {
						i,_=strconv.Atoi(n.Value)
					}else {	i=va.(int)
					}
					Rel := n.SideNode.Type

					switch Rel {
					case TYPE_EQ:
						if num, _ := strconv.Atoi(n.SideNode.SideNode.Value); i == num {
							identifier = append(identifier, true)
						} else {
							identifier = append(identifier, false)
						}
						break
					case TYPE_LG:
						if num, _ := strconv.Atoi(n.SideNode.SideNode.Value); i > num {
							identifier = append(identifier, true)
						} else {
							identifier = append(identifier, false)
						}
						break
					case TYPE_NEQ:
						if num, _ := strconv.Atoi(n.SideNode.SideNode.Value); i != num {
							identifier = append(identifier, true)
						} else {
							identifier = append(identifier, false)
						}
						break
					case TYPE_SM:
						if num, _ := strconv.Atoi(n.SideNode.SideNode.Value); i < num {
							identifier = append(identifier, true)
						} else {
							identifier = append(identifier, false)
						}
						break
					case TYPE_LEQ:
						if num, _ := strconv.Atoi(n.SideNode.SideNode.Value); i >= num {
							identifier = append(identifier, true)
						} else {
							identifier = append(identifier, false)
						}
						break
					case TYPE_SEQ:
						if num, _ := strconv.Atoi(n.SideNode.SideNode.Value); i <= num {
							identifier = append(identifier, true)
						} else {
							identifier = append(identifier, false)
						}
						break

					case TYPE_PLUS:
						sideChar := n.SideNode.SideNode.Value
						selfNum := data[sideChar].(int)
						selfNum=selfNum+i
						data[sideChar]=selfNum
						n=n.SideNode
						break
					case TYPE_MUL:
						sideChar := n.SideNode.SideNode.Value
						selfNum := data[sideChar].(int)
						selfNum=selfNum*i
						data[sideChar]=selfNum
						n=n.SideNode
						break

					case TYPE_DIV:
						sideChar := n.SideNode.SideNode.Value
						selfNum := toF64(data[sideChar])
						a:=float64(selfNum)
						b:=float64(i)
						selfNum=(b/a)
						data[sideChar]=selfNum
						n=n.SideNode
						break
					case TYPE_SUB:
						sideChar := n.SideNode.SideNode.Value
						selfNum := data[sideChar].(int)
						selfNum=i-selfNum
						data[sideChar]=selfNum
						n=n.SideNode
						break
					}
				}
					break
			case TYPE_AND:
				identifier= append(identifier, "&&")
				break
			case TYPE_OR:
				identifier= append(identifier, "||")
				break
			case TYPE_RES_TRUE:
				identifier= append(identifier, true)
				break
			case TYPE_RES_FALSE:
				identifier= append(identifier, false)
				break
			case TYPE_LP:
				if n.SideNode.SideNode.Type==TYPE_EQ||n.SideNode.SideNode.Type==TYPE_NEQ {
					identifier= append(identifier, "(")
					break
				}else {
					calculate(n.SideNode,data)
					n=n.SideNode.SideNode.SideNode
					break
				}

			case TYPE_NOT:
				identifier= append(identifier, "!")
				break
			case TYPE_RP:
				identifier= append(identifier, ")")
				break
			default:
				break
			}

			if n.SideNode!=nil {
				n=n.SideNode
			}

		} else {
			break
		}
	}



	return processBoolSlice(identifier)
}

func calculate(n *Node,data map[string]interface{}) {
	va := data[n.Value]
	_,Bool:=va.(int)
	var i int
	if Bool {
		i=va.(int)
	}
	//i:=va.(int)
	Rel := n.SideNode.Type

	switch Rel {
	case TYPE_PLUS:
		sideChar := n.SideNode.SideNode.Value
		selfNum := data[sideChar].(int)
		selfNum=selfNum+i
		n.SideNode.SideNode.SideNode.Value= strconv.Itoa(selfNum)
		n.SideNode.SideNode.SideNode.Type=TYPE_VAR
		n.SideNode.SideNode.SideNode.Bool=true
		break
	case TYPE_MUL:
		sideChar := n.SideNode.SideNode.Value
		selfNum := data[sideChar].(int)
		selfNum=selfNum*i
		n.SideNode.SideNode.SideNode.Value= strconv.Itoa(selfNum)
		n.SideNode.SideNode.SideNode.Type=TYPE_VAR
		n.SideNode.SideNode.SideNode.Bool=true
		break

	case TYPE_DIV:
		sideChar := n.SideNode.SideNode.Value
		selfNum := data[sideChar].(int)
		selfNum=selfNum/i
		n.SideNode.SideNode.SideNode.Value= strconv.Itoa(selfNum)
		n.SideNode.SideNode.SideNode.Type=TYPE_VAR
		n.SideNode.SideNode.SideNode.Bool=true
		break
	case TYPE_SUB:
		sideChar := n.SideNode.SideNode.Value
		selfNum := data[sideChar].(int)
		selfNum=i-selfNum
		n.SideNode.SideNode.SideNode.Value= strconv.Itoa(selfNum)
		n.SideNode.SideNode.SideNode.Type=TYPE_VAR
		n.SideNode.SideNode.SideNode.Bool=true
		break
	}



}


//对bool 切片处理
func processBoolSlice(slice []interface{}) bool {
	orNum:=0
	falseNum:=0
	for va:=range slice{
		flag:=true
		num:=0
		if slice[va]=="(" {
			if va>0&&slice[va-1]=="!" {
				flag=false
			}
			for true {
				if slice[va]==true&&slice[va+1]=="||" {
					num=va
					slice[va+1]=""
					slice[va+2]=""
				}
				if slice[va]==true&&slice[va+1]=="&&"&&slice[va+2]==false {
					num=va
					slice[va]=false
					slice[va+1]=""
					slice[va+2]=""
				}
				if slice[va]==true&&slice[va+1]=="&&"&&slice[va+2]==true {
					num=va
					slice[va]=true
					slice[va+1]=""
					slice[va+2]=""
				}
				if slice[va]==false&&slice[va+1]=="||"&&slice[va+2]==false {
					num=va
					slice[va]=false
					slice[va+1]=""
					slice[va+2]=""
				}
				if slice[va]==false&&slice[va+1]=="||"&&slice[va+2]==true {
					num=va
					slice[va]=true
					slice[va+1]=""
					slice[va+2]=""
				}
				if slice[va]==false&&slice[va+1]=="&&" {
					num=va
					slice[va]=false
					slice[va+1]=""
					slice[va+2]=""
				}

				va++
				if slice[va] == ")" {
					break
				}
			}
			if _,Bool:=slice[num].(bool);!flag&&Bool {
				slice[num]=!slice[num].(bool)
			}
		}
		if slice[va] == "||" {
			orNum++
		}


	}
	for va:=range slice {

		if _, flag := slice[va].(bool); flag {
			if slice[va].(bool)==false {
				falseNum++
				for i:=va;i< len(slice);i++ {
					if slice[i]=="&&" {
						return false
					}else if slice[i]=="||" {
						break
					}
				}
				for i:=va;i>=0;i-- {
					if slice[i]=="&&" {
						return false
					}else if slice[i]=="||" {
						break
					}
				}


			}
			if falseNum==orNum+1 {
				return false
			}
		}
	}
	return true
}


// 别的节点。。。。
