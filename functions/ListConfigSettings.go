package functions

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func ListConfigSettings() {
	fmt.Print("\nCurrent configuration settings:\n\n")

	printNestedConfig(viper.AllSettings(), 0)
}

func printNestedConfig(settings map[string]interface{}, depth int) {
	for key, value := range settings {
		fmt.Printf("%s%s: ", strings.Repeat("    ", depth), key)
		switch val := value.(type) {
		case map[string]interface{}:
			fmt.Println()
			printNestedConfig(val, depth+1)
		default:
			fmt.Println(val)
			fmt.Println()
		}
	}
}
