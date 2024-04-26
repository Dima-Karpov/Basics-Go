package circle

import "fmt"

type Circle struct {
	r      float64
	x0, y0 float64
}

//var c1 Circle
//c2 := Circle{}
//c3 := Circle{
//	r: 1,
//	x0: 2,
//	y0: 3
//}
//c4 := Circle{1,2,3}
//c5 := new(Circle)
//c6 := &Circle{r: 1, x0: 2, y0: 3}

func changeCircleParams(circle *Circle) {
	circle.r = 5
	circle.x0 = 10
	circle.y0 = 15
}

func main() {
	c := Circle{r: 3, x0: 1, y0: 2}

	fmt.Printf("%+v\n", c)

	changeCircleParams(&c)

	fmt.Printf("%+v\n", c)

}
