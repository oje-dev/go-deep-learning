package main

import "fmt"

func main() {
	matrix := NewMatrix(Rows(2), Columns(3))
	matrix.Elements[0].Elements[0] = 1.2
	matrix.Elements[1].Elements[0] = 2.3
	matrix.Elements[0].Elements[1] = 4.3
	matrix.Elements[1].Elements[1] = 1.2
	matrix.Elements[0].Elements[2] = 3.5
	matrix.Elements[1].Elements[2] = 9.2

	fmt.Println(matrix.String())
	fmt.Println(matrix.Shape())
	matrix.Transpose()
	fmt.Println(matrix.Shape())
	vector := matrix.Flatten()
	fmt.Println(vector.String())
}
