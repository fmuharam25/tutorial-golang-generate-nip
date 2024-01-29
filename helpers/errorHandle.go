package helpers

import (
	"regexp"
	"strconv"
	"strings"
)

// Check Gender
func checkGender(gender string) string {
	var alias string
	switch strings.ToLower(gender) {
	case "ikhwan":
		alias = "N"
	case "akhwat":
		alias = "T"
	default:
		alias = " "
	}

	return alias
}

// Check semester
func checkSemester(date string) string {
	var semester string
	pattern := "^" + `\d{4}-\d{2}-\d{2}` + "$"
	regex := regexp.MustCompile(pattern)

	split := strings.Split(date, "-")
	year := split[0][len(split[0])-2:]
	month, err := strconv.Atoi(split[1])

	if !regex.MatchString(date) {
		return " "
	}

	if month > 12 || month < 1 || err != nil {
		return " "
	}

	if month <= 5 {
		semester = "1"
	} else {
		semester = "2"
	}

	return year + semester
}

// Check NIp
func checkPrefix(nip string) string {
	split := strings.Split(nip, "-")
	prefix := split[0]

	if strings.ToLower(prefix[0:3]) != "arn" && strings.ToLower(prefix[0:3]) != "art" {
		return " "
	}

	return prefix
}

func checkId(nip string) int {
	split := strings.Split(nip, "-")
	id, err := strconv.Atoi(split[1])

	if err != nil {
		return -1
	}

	return id
}
