package main

import "fmt"

func main() {
	fmt.Println("Choose temperature converter:")
	fmt.Println("1. Celcius to Fahrenheit")
	fmt.Println("2. Celcius to Kelvin")
	fmt.Println("3. Fahrenheit to Celcius")
	fmt.Println("4. Fahrenheit to Kelvin")
	fmt.Println("5. Kelvin to Celcius")
	fmt.Println("6. Kelvin to Fahrenheit")
	fmt.Print("Your choice: ")

	var choice int8
	for {
		fmt.Print("Your choice: ")
		_, err := fmt.Scanln(&choice)
		if err != nil || choice < 1 || choice > 6 {
			fmt.Println("Option doesn't exists.")
			continue
		}

		break
	}
}

func CelciusToFahrenheit(celcius float64) float64 {
	fahrenheit := (9.0 / 5.0 * celcius) + 32
	return fahrenheit
}

func CelciusToKelvin(celcius float64) float64 {
	kelvin := celcius + 273.15
	return kelvin
}

func FahrenheitToCelcius(fahrenheit float64) float64 {
	celcius := (fahrenheit - 32) * 5.0 / 9.0
	return celcius
}

func FahrenheitToKelvin(fahrenheit float64) float64 {
	kelvin := (fahrenheit + 459.67) * 5.0 / 9.0
	return kelvin
}

func KelvinToCelcius(kelvin float64) float64 {
	celcius := kelvin - 273.15
	return celcius
}

func KelvinToFahrenheit(kelvin float64) float64 {
	fahrenheit := (kelvin * (9.0 / 5.0)) - 459.67
	return fahrenheit
}
