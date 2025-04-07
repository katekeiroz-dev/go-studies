package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64 = 1000
	years := 10.0
	expectReturnRate := 5.5 //shortcut way to declare and assign a variable

	/* another way to write the code would be :
	   investment , expectReturnRate , year := 1000, 5.5 , 10
	   in this case the value type are inferred by go
	*/

	fmt.Print("How much would like to invest ?")
	fmt.Scan(&investmentAmount)

	futureValue := investmentAmount * math.Pow(1+expectReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

}
