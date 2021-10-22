package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file: %s", *csvFileName))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("failed to parse the provided CSV file")
	}

	problems := parseLines(lines)
	fmt.Println(problems)
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
		var ans string
		fmt.Scanf("%s\n", &ans)
		if ans == p.answer {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))

}
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Printf(msg)
	os.Exit(1)
}
