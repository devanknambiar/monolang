package main

func main() {
	input := `let alpha_1 = 54;
			if (alpha_1 == 54) {
				let target = true;
			} else {
				let target = false;
			}
			alpha_1 != 100;
			`
	l := New(input)
	for {
		t := l.NextToken()
		println("Type: ", t.Type, "& Literal: ", t.Literal)
		if t.Type == EOF {
			break
		}
	}
}