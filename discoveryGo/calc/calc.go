package calc

import (
	"strconv"
	"strings"
)

// Eval returns the evaluation result of the given expr.
// The expression can have _, -, *, /, (, ) operators and
// decimal integers. Operators and operands shoudl be
// space delimited.
func Eval(expr string) int {
	var ops []string
	var nums []int

	pop := func() int {
		last := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		return last
	}

	reduce := func(higher string) {
		for len(ops) > 0 {
			op := ops[len(ops)-1]

			if strings.Index(higher, op) < 0 {
				// op is not a higher operand than current. return.
				return
			}

			ops = ops[:len(ops)-1]
			if op == "(" {
				// '(' is removed. return.
				return
			}

			b, a := pop(), pop()

			switch op {
			case "+":
				nums = append(nums, a+b)
			case "-":
				nums = append(nums, a-b)
			case "*":
				nums = append(nums, a*b)
			case "/":
				nums = append(nums, a/b)
			default:
				panic("error")
			}
		}
	}

	for _, token := range strings.Split(expr, " ") {
		switch token {
		case "(":
			ops = append(ops, token)
		case "+", "-":
			// apply the operands that have higher or equal order than +, -
			reduce("+-*/")
			ops = append(ops, token)
		case "*", "/":
			// apply the operands that have higher or equal order than *, /
			reduce("*/")
			ops = append(ops, token)
		case ")":
			// remove closing parenthesis after applying until opening parenthesis.
			reduce("+-*/(")
		default:
			num, _ := strconv.Atoi(token)
			nums = append(nums, num)
		}
	}

	reduce("+-*/")
	return nums[0]
}
