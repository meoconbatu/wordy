package wordy

import "regexp"
import "strings"
import "strconv"

const testVersion = 1

func Answer(question string) (int, bool) {
	r, _ := regexp.Compile("What is (-?[0-9]+( (plus|minus|multiplied by|divided by){1}? -?[0-9]+){1,})+?")
	if !r.MatchString(question) {
		return 0, false
	}
	subMatches := r.FindStringSubmatch(question)
	elements := strings.Split(subMatches[1], " ")
	x, operation := 0, ""
	for _, e := range elements {
		if d, err := strconv.Atoi(e); err == nil {
			if operation == "" {
				x = d
			} else {
				switch {
				case operation == "plus":
					x = x + d
				case operation == "minus":
					x = x - d
				case operation == "multiplied":
					x = x * d
				case operation == "divided":
					x = x / d
				}
			}
		} else {
			if e != "by" {
				operation = e
			}
		}
	}
	return x, true
}
