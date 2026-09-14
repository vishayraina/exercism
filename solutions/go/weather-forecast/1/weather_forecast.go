// Package weather: dumbest package ever.
package weather

var (
    // CurrentCondition: describes current weather conditions.
	CurrentCondition string 
    // CurrentLocation: stores current location city.
	CurrentLocation  string
)

// Forecast: gives exact forecast of modiji's nosehairs.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
