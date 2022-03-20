/*Package weather is a package for 
forecasting the weather information of a city.*/
package weather

//CurrentCondition : gets the current condition of a city.
var CurrentCondition string

//CurrentLocation : gets location data (City).
var CurrentLocation string

//Forecast : returns the weather information about the specified city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
