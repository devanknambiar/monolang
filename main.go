package main

func main() {
	input := `let alpha_1 = 54;
			return 100;
			`
	l := New(input)
	p := NewParser(l)
	program := p.ParseProgram()

	println("Total Number of Statements: ", len(program.Statements))

	for i, stmt := range program.Statements {
		println("Statement [",i,"] Token Literal: ",stmt.TokenLiteral())
	}
}