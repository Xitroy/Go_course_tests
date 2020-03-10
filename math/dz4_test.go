package math

import (    
	"testing" 
)

// Predefined operations
const (
	Add      = "+"
	Subtract = "-"
	Multiply = "*"
	Divide   = "/"
)

type CalcTest struct{
	a float64;
	b float64;
	op CalcOperation;
	exp float64;	
}


// returns 4 test cases, [0] = "+", [1] = "-", [2] = "*", [3] = "/"
func CalcTestsFabric(a,b float64) [4]CalcTest {
	return [4]CalcTest {
		{a: a, b: b, op: Add, exp: a+b},
		{a: a, b: b, op: Subtract, exp: a-b},
		{a: a, b: b, op: Multiply, exp: a*b},
		{a: a, b: b, op: Divide, exp: a/b},
	}
}

var tests = [][4]CalcTest {
	CalcTestsFabric(1,2),
	CalcTestsFabric(1,-2),
	CalcTestsFabric(-1,-2),
	CalcTestsFabric(-1,2),

	CalcTestsFabric(100,200),
	CalcTestsFabric(100,-200),
	CalcTestsFabric(-100,-200),
	CalcTestsFabric(-100,200),

	CalcTestsFabric(0,5),
	CalcTestsFabric(5,0),
	CalcTestsFabric(-5,0),
	CalcTestsFabric(5,-0),

	//вот тут тесты поломаются. Работа с нулями подвела :)
	CalcTestsFabric(0,0),
	CalcTestsFabric(-0,-0),
}

func TestCalcAdd(t *testing.T) {
    for _, curTests := range tests {
    	for _, curTest := range curTests{
    		t.Log(curTest)
    		if Calc(curTest.op)(curTest.a, curTest.b) != curTest.exp {
    		t.Log("breaks on", curTest)
    		t.Errorf("%f%s%f should be %f", curTest.a, curTest.op, curTest.b, curTest.exp)
    		}
    	}
    	
    }
}