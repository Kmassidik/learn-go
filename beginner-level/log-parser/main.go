package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	file, err := os.Open("logfile.log")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	logPattern := `\[(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2})\] \[(INFO|ERROR|DEBUG|WARN)\] (.+)`
	re := regexp.MustCompile(logPattern)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Reading line:", line) // Debugging: print each line

		match := re.FindStringSubmatch(line)
		if match == nil {
			fmt.Println("No match for line:", line) // Debugging: log unmatched lines
			continue
		}
		fmt.Printf("Timestamp: %s | Level: %s | Message: %s\n", match[1], match[2], match[3])
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
