package main

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

type Car struct {
	BrandName      string
	ModelName      string
	DimensionsVal  Dimensions
	MaxSpeedVal    int
	EnginePowerVal int
}

func (c Car) Brand() string {
	return c.BrandName
}
func (c Car) Model() string {
	return c.ModelName
}
func (c Car) Dimensions() Dimensions {
	return c.DimensionsVal
}
func (c Car) MaxSpeed() int {
	return c.MaxSpeedVal
}
func (c Car) EnginePower() int {
	return c.EnginePowerVal
}

type carDimensions struct {
	LengthUnit Unit
	WidthUnit  Unit
	HeightUnit Unit
}

func (d carDimensions) Length() Unit {
	return d.LengthUnit
}
func (d carDimensions) Width() Unit {
	return d.WidthUnit
}
func (d carDimensions) Height() Unit {
	return d.HeightUnit
}

type BMW struct {
	Car
}

func NewBMW() *BMW {
	return &BMW{
		Car: Car{
			BrandName: "BMW",
			ModelName: "F10",
			DimensionsVal: carDimensions{
				LengthUnit: Unit{Value: 200, T: CM},
				WidthUnit:  Unit{Value: 65, T: CM},
				HeightUnit: Unit{Value: 68, T: CM},
			},
			MaxSpeedVal:    250,
			EnginePowerVal: 440,
		},
	}
}

type Mercedes struct {
	Car
}

func NewMercedes() *Mercedes {
	return &Mercedes{
		Car: Car{
			BrandName: "Mercedes",
			ModelName: "E-Class",
			DimensionsVal: carDimensions{
				LengthUnit: Unit{Value: 200, T: CM},
				WidthUnit:  Unit{Value: 65, T: CM},
				HeightUnit: Unit{Value: 68, T: CM},
			},
			MaxSpeedVal:    180,
			EnginePowerVal: 300,
		},
	}
}

type Dodge struct {
	Car
}

func NewDodge() *Dodge {
	return &Dodge{
		Car: Car{
			BrandName: "Dodge",
			ModelName: "Charger",
			DimensionsVal: carDimensions{
				LengthUnit: Unit{Value: 200, T: Inch},
				WidthUnit:  Unit{Value: 80, T: Inch},
				HeightUnit: Unit{Value: 60, T: Inch},
			},
			MaxSpeedVal:    200,
			EnginePowerVal: 370,
		},
	}
}

func main() {
	bmw := NewBMW()
	mercedes := NewMercedes()
	dodge := NewDodge()

	fmt.Println("BMW Dimensions (CM): ", bmw.Dimensions())
	fmt.Println("Mercedes Dimensions (CM): ", mercedes.Dimensions())
	fmt.Println("Dodge Dimensions (Inches): ", dodge.Dimensions())
}
