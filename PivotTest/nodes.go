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
func (n *Node) Eval(data map[string]interface{}) interface{} {

	var identifier []interface{}

	for true {
		if n.SideNode != nil {

			Type:=n.Type
			switch Type {
			case TYPE_VAR:
				va := data[n.Value]

				_, b :=va.(int)
				if !b {
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
				}else {
					i:=va.(int)
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
						//sideChar := n.SideNode.SideNode.Value
						//		selfNum := data[sideChar].(int)

						break
					case TYPE_MUL:

						break
					case TYPE_DIV:
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
				identifier= append(identifier, "(")
				break
			case TYPE_NOT:
				identifier= append(identifier, "!")
				break
			case TYPE_RP:
				identifier= append(identifier, ")")
				break
			default:
				break
			}


			n=n.SideNode
		} else {
			break
		}
	}



	return processBoolSlice(identifier)
}

//对bool 切片处理
func processBoolSlice(slice []interface{}) bool {

	for va:=range slice{
		if slice[va]==false{

			if va==0 {
				return false
			}
			if va>0&&va+1< len(slice)&&(slice[va-1]=="&&"||slice[va+1]=="&&") {
				return false
			}

			if va+1< len(slice)&&(slice[va+1]=="&&") {
				return false
			}
			if va>0&&(slice[va-1]=="&&") {
				return false
			}


		}
		if slice[va]=="||"{
			if slice[va-1]==false&&slice[va+1]==false {
				return false
			}
		}




	}

	return true
}


// 别的节点。。。。
