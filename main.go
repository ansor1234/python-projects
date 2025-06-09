package main
import ("fmt"
	"strings"
	"math"
)

func main() {
	

	symbol := "Maga"
	count_to := 15
	// Print the top half of the diamond
	for i := 1; i <= count_to; i++ {
		times := int(math.Ceil(float64(count_to) - float64(i)/2))
		// if i%2 != 0 {
			fmt.Printf(strings.Repeat(" ", times))
		
			for j := 1; j <= i; j++ {
				fmt.Printf(symbol)
			}
		fmt.Print("\n")
		// }
	}

	// Print the bottom half of the diamond
	for i:= count_to - 1; i >= 1; i-- {
		times := int(math.Ceil(float64(count_to) - float64(i)/2))
		// if i%2 != 0 {
			fmt.Printf(strings.Repeat(" ", times))
		
		for j := 1; j <= i; j++ {
			fmt.Printf(symbol)
		}
		fmt.Print("\n")
	// }
	}

}