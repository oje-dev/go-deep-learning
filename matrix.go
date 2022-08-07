package main

import "fmt"

type Matrix struct {
	Rows     int
	Columns  int
	Elements []Vector
	Shape    func() string
	String   func() string

	Transpose     func()
	Flatten       func() *Vector
	DotProduct    func(*Matrix) *Matrix
	Multiply      func(*Matrix) *Matrix
	Add           func(*Matrix) *Matrix
	Subtract      func(*Matrix) *Matrix
	ApplyFunction func(func(float64) float64) *Matrix
	Scale         func(float64) *Matrix
	AddScalar     func(*Vector)
}

type MatrixOption func(*Matrix)

func (m *Matrix) Init() {
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

	m.Transpose = func() {
		var temp int
		temp = m.Rows
		m.Rows = m.Columns
		m.Columns = temp

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

	m.Shape = func() string {
		return fmt.Sprintf("(%d,%d)", m.Rows, m.Columns)
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
