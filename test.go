package main
import (
	"fmt"
	"os"
	"strconv"
)
func main() {
	fmt.Println(sumOfSquares([]string{"3", "-1", "1", "14"}))
}
func sumOfSquares(numArr []string) int {
	i, err := strconv.Atoi(numArr[0])

	rest := numArr[1:]

	//Error checking
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return 0
	}
