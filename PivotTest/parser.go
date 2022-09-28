package mapq

import "errors"

// Parser 语法分析器
type Parser struct {
	lexer *Lexer
}

// 你的递归下降分析代码（如果你使用递归下降的话
 func (p *Parser) boolexp() (node Node, err error) {

	 defer func() {
		 if recover()!=nil {
			 err=errors.New("not implemented")
		 }
	 }()




 	panic("not implemented")
 }

  func (p *Parser) boolean() (node Node, err error) {
 	panic("not implemented")
 }
// 别的分析函数


// Parse 生成ast
func (p *Parser) Parse(str string) (n Node,err error) {
	p.lexer=new(Lexer)
	p.lexer.SetInput(str)
	scanCode:= p.lexer.Scan()

	return scanCode,nil

}
