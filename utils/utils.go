package utils

import "fmt"

// GetPercentage returns a string in the "dd.dd%" format calculated from a and b
func GetPercentage(a, b int) string {
	return fmt.Sprintf("%.2f%%", float64(a)*100.0/float64(b))
}
