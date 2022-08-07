package main

import "fmt"

type Matrix struct {
	Rows     int
	Columns  int
	Elements []Vector
	Shape    func() string
	String   func() string

	Transpose     func() *Matrix
	Flatten       func() *Vector
	Equal         func(*Matrix) bool
	DotProduct    func(*Matrix) *Matrix
	Multiply      func(*Matrix) *Matrix
	Add           func(*Matrix) *Matrix
	Subtract      func(*Matrix) *Matrix
	ApplyFunction func(func(float64) float64) *Matrix
	Scale         func(float64) *Matrix
	AddScalar     func(*Vector) *Matrix
}

type MatrixOption func(*Matrix)

func (m *Matrix) Init() {
	m.Shape = func() string {
		return fmt.Sprintf("(%d,%d)", m.Rows, m.Columns)
	}

	m.String = func() string {
		outputString := "\n"
		for i, element := range m.Elements {
			for _, element := range element.Elements {
				outputString += fmt.Sprintf("%.1f ", element)
			}
			if i != len(m.Elements)-1 {
				outputString += "\n"
			}
		}
		return outputString
	}

	m.Transpose = func() *Matrix {
		matrix := NewMatrix(Rows(m.Columns), Columns(m.Rows))
		for i := 0; i < m.Rows; i++ {
			for j := 0; j < m.Columns; j++ {
				matrix.Elements[j].Elements[i] = m.Elements[i].Elements[j]
			}
		}
		return matrix
	}

	m.Flatten = func() *Vector {
		var flattened []float64
		for _, element := range m.Elements {
			for _, element := range element.Elements {
				flattened = append(flattened, element)
			}
		}
		return NewVector(Elements(flattened))
	}

	m.Equal = func(m2 *Matrix) bool {
		if m.Rows != m2.Rows || m.Columns != m2.Columns {
			return false
		}
		equal := true
		for i := 0; i < m.Rows; i++ {
			for j := 0; j < m.Columns; j++ {
				if m.Elements[i].Elements[j] != m2.Elements[i].Elements[j] {
					equal = false
				}
			}
		}
		return equal
	}

	m.DotProduct = func(m2 *Matrix) *Matrix {

		return m
	}

	m.Multiply = func(m2 *Matrix) *Matrix {
		return m
	}

	m.Add = func(m2 *Matrix) *Matrix {
		return m
	}

	m.Subtract = func(m2 *Matrix) *Matrix {
		return m
	}

	m.ApplyFunction = func(func(n float64) float64) *Matrix {
		return m
	}

	m.Scale = func(n float64) *Matrix {
		return m
	}

	m.AddScalar = func(v *Vector) *Matrix {
		return m
	}

}

func Rows(rows int) MatrixOption {
	return func(m *Matrix) {
		m.Rows = rows
	}
}

func Columns(columns int) MatrixOption {
	return func(m *Matrix) {
		m.Columns = columns
	}
}

func NewMatrix(options ...MatrixOption) *Matrix {
	m := &Matrix{}
	m.Init()
	for _, f := range options {
		f(m)
	}
	for i := 0; i < m.Rows; i++ {
		column := make([]float64, m.Columns)
		m.Elements = append(m.Elements, *NewVector(Elements(column)))
	}
	return m
}
