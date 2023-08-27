package main
import "fmt"

func main1() {


	fmt.Println("Please enter a tmeperature in Fahrenheit:")

	var userInput float64
	fmt.Scanf("%f", &userInput)

	fmt.Println("This temperature in Celsius is:")
	var temperatureInCelsius float64 = (userInput-32)*(5/9)
	fmt.Println(temperatureInCelsius)

}