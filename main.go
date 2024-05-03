package main

import (
	"basicsGo/electronic"
	"fmt"
)

func main() {
	apple := electronic.NewApplePhone("iPhone 12mini")
	samsung := electronic.NewAndroidPhone("Samsung", "Galaxy S21")
	radio := electronic.NewRadioPhone("Sony", "Some Model", 12)

	printCharacteristics(apple)
	printCharacteristics(samsung)
	printCharacteristics(radio)
}

func printCharacteristics(p electronic.Phone) {
	fmt.Printf("Brand: %s\n", p.Brand())
	fmt.Printf("Model: %s\n", p.Model())
	fmt.Printf("Type: %s\n", p.Type())

	if sp, ok := p.(electronic.Smartphone); ok {
		fmt.Printf("OS: %s\n", sp.OS())
	}

	if sp, ok := p.(electronic.StationPhone); ok {
		fmt.Printf("Buttons count: %d\n", sp.ButtonsCount())
	}
}
