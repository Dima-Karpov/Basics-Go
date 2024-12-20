package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <input_file> <output_file>")
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	lines, err := readFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}
	results := processLines(lines)
	err = writeFile(outputFile, results)
	if err != nil {
		fmt.Printf("Error writing to output file: %v\n", err)
		return
	}
}

// Read input file
func readFile(fileName string) ([]string, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return []string{}, err
	}
	lines := strings.Split(string(content), "\n")
	return lines, nil
}

// Write file
func writeFile(fileName string, lines []string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func processLines(lines []string) []string {
	regex := regexp.MustCompile(`^(\d+)\s*([+\-*/])\s*(\d+)\s*=\??$`)
	results := make([]string, len(lines)) // Array to preserve order
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i, line := range lines {
		wg.Add(1)
		go func(i int, line string) {
			defer wg.Done()
			line = strings.TrimSpace(line)
			if line == "" {
				results[i] = ""
				return
			}

			if matches := regex.FindStringSubmatch(line); matches != nil {
				operand1, _ := strconv.Atoi(matches[1])
				operand2, _ := strconv.Atoi(matches[3])
				operator := matches[2]

				var result int
				var outputLine string
				switch operator {
				case "+":
					result = operand1 + operand2
					outputLine = fmt.Sprintf("%s%d", strings.TrimSuffix(line, "?"), result)
				case "-":
					result = operand1 - operand2
					outputLine = fmt.Sprintf("%s%d", strings.TrimSuffix(line, "?"), result)
				case "*":
					result = operand1 * operand2
					outputLine = fmt.Sprintf("%s%d", strings.TrimSuffix(line, "?"), result)
				case "/":
					if operand2 != 0 {
						result = operand1 / operand2
						outputLine = fmt.Sprintf("%s%d", strings.TrimSuffix(line, "?"), result)
					} else {
						outputLine = fmt.Sprintf("%sundefined (division by zero)", strings.TrimSuffix(line, "?"))
					}
				default:
					outputLine = line
				}

				mu.Lock()
				results[i] = outputLine
				mu.Unlock()
			} else {
				results[i] = ""
			}
		}(i, line)
	}

	wg.Wait()
	// Removing empty lines from the result
	var finalResult []string
	for _, result := range results {
		if result != "" {
			finalResult = append(finalResult, result)
		}
	}

	return finalResult
}
