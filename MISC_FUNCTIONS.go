package main

import (
	"strconv"
	"strings"
)

func F64_TO_INT(NUM float64) int {
	STRING_FLOAT := strconv.FormatFloat(NUM, 'f', -1, 64)
	CLEAN_STRING := strings.ReplaceAll(STRING_FLOAT, ".", "")
	RESULT, ERR := strconv.Atoi(CLEAN_STRING)
	ERRHANDLE(ERR)
	return RESULT
}

// ^ explanation: json.unmarshal fucking sucks and makes ints become a float64
