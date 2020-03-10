package math

import "fmt"

// CalcOperation is a type that define allowed operations for calculator
type CalcOperation string;

// Calc is a function that return an operation function
func Calc(operation CalcOperation) func(a, b float64) float64 {
	// Place your solution here
	switch operation { 
	case "+":
		return func(a, b float64) float64{
			return (a+b);
		}
	case "-":
		return func(a, b float64) float64{
			return (a-b);
		}
	case "*":
		return func(a, b float64) float64{
			return (a*b);
		}
	case "/":
		return func(a, b float64) float64{
			return (a/b);
		}
	}
	return func(a, b float64) float64{
		fmt.Println("unsuported operation, 0.0 means nothing")
		return 0.0
	};
}