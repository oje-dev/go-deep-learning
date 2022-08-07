package main

import (
	"reflect"
	"testing"
)

func TestTranspose(t *testing.T) {

}

func TestFlatten(t *testing.T) {
	m := NewMatrix(Rows(2), Columns(2))
	m.Elements[0].Elements[0] = 1.2
	m.Elements[0].Elements[1] = 1.4
	m.Elements[1].Elements[0] = 1.9
	m.Elements[1].Elements[1] = 1.6
	v := NewVector(Elements([]float64{1.2, 1.4, 1.9, 1.6}))

	if !reflect.DeepEqual(m.Flatten().Elements, v.Elements) {
		t.Errorf("Flatten operation incorrect, got: %s, want: %s", m.Flatten().String(), v.String())
	}

}

func TestEqual(t *testing.T) {
	m := NewMatrix(Rows(2), Columns(2))
	m.Elements[0].Elements[0] = 1.2
	m.Elements[0].Elements[1] = 1.4
	m.Elements[1].Elements[0] = 1.9
	m.Elements[1].Elements[1] = 1.6

	m2 := NewMatrix(Rows(2), Columns(2))
	m2.Elements[0].Elements[0] = 1.2
	m2.Elements[0].Elements[1] = 1.4
	m2.Elements[1].Elements[0] = 1.9
	m2.Elements[1].Elements[1] = 1.6

	if !m.Equal(m2) {
		t.Errorf("Matrices are not equal\nwant:%s\ngot:%s", m.String(), m2.String())
	}
}

func TestDotProduct(t *testing.T) {

}
func TestMultiply(t *testing.T) {

}

func TestAdd(t *testing.T) {

}

func TestSubtract(t *testing.T) {

}

func TestApplyFunction(t *testing.T) {

}

func TestScale(t *testing.T) {

}

func TestAddScalar(t *testing.T) {

}
