package Golang_Generics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestLength(t *testing.T) {
	var result string = Length[string]("Hanif")
	assert.Equal(t, result, "Hanif")

	var resultNumber int = Length[int](10)
	assert.Equal(t, resultNumber, 10)
}

func MultiParameter[T1 any, T2 any](param1 T1, param2 T2) {
	fmt.Println(param1)
	fmt.Println(param2)
}

func TestMultipleParameter(t *testing.T) {
	MultiParameter[string, int]("Hanif", 10)
	MultiParameter[int, string](10, "Hanif")
}

func IsSame[T comparable](value1, value2 any) bool {
	if value1 == value2 {
		return true
	} else {
		return false
	}
}

func TestIsSame(t *testing.T) {
	assert.True(t, IsSame[string]("Hanif", "Hanif"))
	assert.False(t, IsSame[string]("Hanif", "Hanif1"))
	assert.True(t, IsSame[int](10, 10))
	assert.False(t, IsSame[int](10, 11))
	assert.True(t, IsSame[bool](true, true))
	assert.False(t, IsSame[bool](true, false))
}

type Employee interface {
	GetName() string
}

func GetName[T Employee](parameter T) string {
	return parameter.GetName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

type VicePresident interface {
	GetName() string
	GetVicePresidentName() string
}

type MyVicePresident struct {
	Name string
}

func (m MyManager) GetName() string {
	return m.Name
}

func (m MyManager) GetManagerName() string {
	return m.Name
}

func (m MyVicePresident) GetName() string {
	return m.Name
}

func (m MyVicePresident) GetVicePresidentName() string {
	return m.Name
}

func TestGetName(t *testing.T) {
	assert.Equal(t, GetName[Manager](&MyManager{Name: "Hanif"}), "Hanif")
	assert.Equal(t, GetName[VicePresident](&MyVicePresident{Name: "Hanif"}), "Hanif")
}

type Age int

type Number interface {
	~int | int8 | int16 | int32 | int64 |
		float32 | float64
}

func Min[T Number](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestMin(t *testing.T) {
	assert.Equal(t, Min[int](10, 20), 10)
	assert.Equal(t, Min[int8](10, 20), int8(10))
	assert.Equal(t, Min[int16](10, 20), int16(10))
	assert.Equal(t, Min[int32](10, 20), int32(10))
	assert.Equal(t, Min[int64](10, 20), int64(10))
	assert.Equal(t, Min[float32](10.5, 20.5), float32(10.5))
	assert.Equal(t, Min[float64](10.5, 20.5), float64(10.5))
	assert.Equal(t, Min[Age](10, 20), Age(10))
}

func TestTypeInference(t *testing.T) {
	assert.Equal(t, Min(10, 20), 10)
	assert.Equal(t, Min(int64(10), int64(20)), int64(10))
	assert.Equal(t, Min(float32(10.5), float32(20.5)), float32(10.5))
	assert.Equal(t, Min(Age(10), Age(20)), Age(10))
}

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, value := range bag {
		fmt.Println(value)
	}
}

func TestBagString(t *testing.T) {
	names := Bag[string]{"Hanif", "Al", "Irsyad"}
	PrintBag(names)
}

func TestBagInt(t *testing.T) {
	number := Bag[int]{1, 2, 3, 4, 5}
	PrintBag(number)
}

type Data[T any] struct {
	First  T
	Second T
}

func (d Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d Data[T]) ChangeFirst(first T) T {
	d.First = first
	return d.First
}

func TestData(t *testing.T) {
	data := Data[string]{
		First:  "Hanif",
		Second: "Al",
	}

	fmt.Println(data)
}

func TestGenericMethod(t *testing.T) {
	data := Data[string]{
		First:  "Hanif",
		Second: "Al",
	}

	assert.Equal(t, "Budi", data.ChangeFirst("Budi"))
	assert.Equal(t, "Hello Hanif", data.SayHello("Hanif"))
}

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

func ChangeValue[T any](param GetterSetter[T], value T) T {
	param.SetValue(value)
	return param.GetValue()
}

type MyData[T any] struct {
	Value T
}

func (d *MyData[T]) GetValue() T {
	return d.Value
}

func (d *MyData[T]) SetValue(value T) {
	d.Value = value
}

func TestGenericInterface(t *testing.T) {
	myData := MyData[string]{}
	result := ChangeValue[string](&myData, "Hanif")

	assert.Equal(t, "Hanif", result)
	assert.Equal(t, "Hanif", myData.Value)
}

func FindMin[T interface{ int | int64 | float64 }](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestFindMin(t *testing.T) {
	assert.Equal(t, 10, FindMin(10, 20))
	assert.Equal(t, int64(10), FindMin(int64(10), int64(20)))
	assert.Equal(t, 10.5, FindMin(10.5, 20.5))
}

func GetFirst[T []E, E any](data T) E {
	first := data[0]
	return first
}

func TestGetFirst(t *testing.T) {
	names := []string{
		"Hanif",
		"Al",
		"Irsyad",
	}

	first := GetFirst[[]string, string](names)
	assert.Equal(t, "Hanif", first)
}
