package quark

//Visitor which can visit AST nodes.
type Visitor interface {
	Visit(node AST) Visitor
}


//Calls visit on the visitor for the node and its children.
func Walk(v Visitor, node AST) {

}