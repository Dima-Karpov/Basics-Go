package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	const prec = 250 // Точность 250 бит в мантисее
	steps := int(math.Log2(prec))
	two := new(big.Float).SetPrec(prec).SetInt64(2) //инициализация
	half := new(big.Float).SetPrec(prec).SetFloat64(0.5)
	x := new(big.Float).SetPrec(prec).SetInt64(1)
	t := new(big.Float).SetPrec(prec)
	for i := 0; i < steps; i++ {
		t.Quo(two, x)  // t = 2.0/x[n]
		t.Add(x, t)    // t = x[n] + (2.0/x[n])
		x.Mul(half, t) // x[n+1] = 0,5 * t
	}

	fmt.Printf("sqrt(2) = %.85f\n", x)
	t.Mul(x, x)
	fmt.Printf("error = %+.85f\n", t.Sub(two, t)) // ошибка

}
