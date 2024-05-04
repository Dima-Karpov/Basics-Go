package auto

import "fmt"

type UnitType string

const (
	Inch UnitType = "inch"
	CM   UnitType = "cm"
)

type Unit struct {
	Value float64
	T     UnitType
}

func (u Unit) Get(t UnitType) float64 {
	value := u.Value

	if t != u.T {
		if t == Inch && u.T == CM {
			value /= 2.54 // cm in inch
		} else if t == CM && u.T == Inch {
			value *= 2.54 // inch in cm
		} else {
			fmt.Println("Unsupported conversion")
		}
	}

	return value
}

type Dimensions interface {
	Length() Unit
	Width() Unit
	Height() Unit
}

type Auto interface {
	Brand() string
	Model() string
	Dimensions() Dimensions
	MaxSpeed() int
	EnginePower() int
}

type InchesDimensions struct {
	LengthUnit Unit
	WidthUnit  Unit
	HeightUnit Unit
}

func (d InchesDimensions) Length() Unit {
	return d.LengthUnit
}
func (d InchesDimensions) Width() Unit {
	return d.WidthUnit
}
func (d InchesDimensions) Height() Unit {
	return d.HeightUnit
}

type CMDDimensions struct {
	LengthUnit Unit
	WidthUnit  Unit
	HeightUnit Unit
}

func (d CMDDimensions) Length() Unit {
	return d.LengthUnit
}
func (d CMDDimensions) Width() Unit {
	return d.WidthUnit
}
func (d CMDDimensions) Height() Unit {
	return d.HeightUnit
}

type BMW struct{}

func (b BMW) Brand() string {
	return "BMW"
}
func (b BMW) Model() string {
	return "F10"
}
func (b BMW) Dimensions() Dimensions {
	return CMDDimensions{
		LengthUnit: Unit{Value: 200, T: CM},
		WidthUnit:  Unit{Value: 65, T: CM},
		HeightUnit: Unit{Value: 68, T: CM},
	}
}
func (b BMW) MaxSpeed() int {
	return 250
}
func (b BMW) EnginePower() int {
	return 440
}

type Mercedes struct{}

func (m Mercedes) Brand() string {
	return "Mercedes"
}
func (m Mercedes) Model() string {
	return "E-Class"
}
func (m Mercedes) Dimensions() Dimensions {
	return CMDDimensions{
		LengthUnit: Unit{Value: 200, T: CM},
		WidthUnit:  Unit{Value: 65, T: CM},
		HeightUnit: Unit{Value: 68, T: CM},
	}
}
func (m Mercedes) MaxSpeed() int {
	return 180
}
func (m Mercedes) EnginePower() int {
	return 300
}

type Dodge struct{}

func (d Dodge) Brand() string {
	return "Dodge"
}
func (d Dodge) Model() string {
	return "Charger"
}
func (d Dodge) Dimensions() Dimensions {
	return InchesDimensions{
		LengthUnit: Unit{Value: 200, T: Inch},
		WidthUnit:  Unit{Value: 80, T: Inch},
		HeightUnit: Unit{Value: 60, T: Inch},
	}
}
func (d Dodge) MaxSpeed() int {
	return 200
}
func (d Dodge) EnginePower() int {
	return 370
}

func main() {
	bmw := BMW{}
	mercedes := Mercedes{}
	dodge := Dodge{}

	fmt.Println("BMW Dimensions (CM): ", bmw.Dimensions())
	fmt.Println("Mercedes Dimensions (CM): ", mercedes.Dimensions())
	fmt.Println("Dodge Dimensions (Inches): ", dodge.Dimensions())
}
