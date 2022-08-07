package main

import (
	"fmt"
)

type Vector struct {
	Elements []float64

	Size   func() int
	Shape  func() string
	String func() string
}

type VectorOption func(*Vector)

func (v *Vector) Init() {
	v.Size = func() int {
		return len(v.Elements)
	}

	v.Shape = func() string {
		return fmt.Sprintf("(%d,)", v.Size())
	}

	v.String = func() string {
		out := "\n"
		for _, element := range v.Elements {
			out += fmt.Sprintf("%.2f ", element)
		}
		return out
	}
}
func Elements(elements []float64) VectorOption {
	return func(v *Vector) {
		v.Elements = elements
	}
}

func NewVector(options ...VectorOption) *Vector {
	v := &Vector{}
	v.Init()
	for _, f := range options {
		f(v)
	}
	return v
}

// func MultiplyVectors(vA *Vector, vB *Vector) float64 {
// 	if vA.Size() != vB.Size() {
// 		log.Fatalf("vectors must be in the same vector space, v1: %s v2: %s", vA.Shape(), vB.Shape())
// 	}
// 	var result float64
// 	for i, element := range vA.Elements {
// 		result += element * vB.Elements[i]
// 	}
// 	return result
// }
