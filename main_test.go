package main

import "testing"

type SumTable struct {
	x 		int
	y 		int
	result 	int
}

func TestSum(t *testing.T){

	data := []SumTable{
		{1, 2, 3},
		{5, 5, 10},
		{5, 5, 10},
	}

	for _, value := range data {
		total := Sum(value.x, value.y)

		if total != value.result {
			t.Error("Resultado invalido")
		}
	}

}