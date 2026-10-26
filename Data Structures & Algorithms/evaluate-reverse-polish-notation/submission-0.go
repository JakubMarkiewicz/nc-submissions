func evalRPN(tokens []string) int {
	stack := []int{}

	for _, v := range tokens {
		str := string(v)	
		vInt, err := strconv.Atoi(str)

		if err == nil {
			stack = append(stack, vInt)
			continue
		}

		l,r := stack[len(stack)-2], stack[len(stack)-1]	
		res := 0
		switch str {
			case "+": 
				res = l +r
			case "-": 
				res = l - r
			case "*": 
				res = l * r
			case "/": 
				res = l / r
			default: 
				panic(v)
		}
		stack = append(stack[0:len(stack)-2], res)
	}

	return stack[0]
}

