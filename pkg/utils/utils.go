package utils

import (
	"fmt"
	"regexp"
	"strings"
)

var cleanRegex = regexp.MustCompile(`[^a-zA-Z0-9., ]+`)

func CleanText(s string) string {
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)
	return cleanRegex.ReplaceAllString(s, "")
}

func PrintProgressBar(i, total int) {
	percent := 100 * float64(i) / float64(total)
	filledLength := int(50 * i / total)
	end := ">"
	if i == total {
		end = "="
	}
	bar := strings.Repeat("=", filledLength) + end + strings.Repeat("-", (50-filledLength))
	fmt.Printf("\r%s [%s] %f%% %s", "Progress", bar, percent, "Complete")
	if i == total {
		fmt.Println()
	}
}

func KeysString(m map[string]string) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return strings.Join(keys, ", ")
}
