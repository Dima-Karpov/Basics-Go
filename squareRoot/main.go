package squareRoot

import (
	"fmt"
	"math"
)

var (
	ErrZeroA       = fmt.Errorf("coefficient of zero")        // уравнение не является квадратным
	ErrNoRealRoots = fmt.Errorf("equation has no real roots") // у уравнения нет вещественных корней
)

// SolveQuadraticEquation finds real roots of equation defined with 3 real coefficients
// It returns 2 roots if no error encountered or default float64 values and error otherwise
func SolveQuadraticEquation(a, b, c float64) (x1, x2 float64, err error) {
	if a == 0 { // проверка на то, что у нас квадратное уравнение
		err = ErrZeroA // возвращаем ошибку
		// так как мы дали имена переменным в сигнатуре, то
		// мы вернём именно их; по умолчанию x1 = x2 = 0.0

		return
	}

	D := b*b - 4*a*c // вычисляем дискриминант
	if D < 0 {
		err = ErrNoRealRoots

		return
	}
	if D == 0 { // уравнение имеет два одинаковых корня
		x1 = -b / (2 * a)
		x2 = x1

		return // err == nil по умолчанию
	}
	dRoot := math.Sqrt(D)
	x1 = (-b + dRoot) / (2 * a)
	x2 = (-b - dRoot) / (2 * a)

	return
}
