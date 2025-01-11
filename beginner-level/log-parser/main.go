package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type LogEntry struct {
	Timestamp string
	Level     string
	Message   string
}

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
	var entries []LogEntry

	for scanner.Scan() {
		line := scanner.Text()
		match := re.FindStringSubmatch(line)
		if match != nil {
			entry := LogEntry{
				Timestamp: match[1],
				Level:     match[2],
				Message:   match[3],
			}
			entries = append(entries, entry)
			fmt.Println(entry)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
