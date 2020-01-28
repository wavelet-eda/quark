package quark

//Visitor which can visit AST nodes.
type Visitor interface {
	Visit(node AST) Visitor
}


func visitParamList(params []*ParameterDef, v Visitor) {
	for _, p := range params {
		p.Accept(v)
	}
}