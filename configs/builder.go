package configs

import (
	"fmt"
	"github.com/galagoshin-com/GoLogger/logger"
	"strconv"
	"strings"
)

func buildStr(data map[string]any) string {
	res := ""
	for key, value := range data {
		res += fmt.Sprintf("%v: %v\n", key, value)
	}
	return res
}

func buildMap(str string) (uint, map[string]any) {
	res := make(map[string]any)
	lines := strings.Split(str, "\n")
	version := ""
	for _, line := range lines {
		if line != "" {
			if string([]rune(line)[0]) != "#" {
				kv := strings.Split(line, ": ")
				if num, err := strconv.Atoi(kv[1]); err == nil {
					res[kv[0]] = num
				} else if double, err2 := strconv.ParseFloat(kv[1], 10); err2 == nil {
					res[kv[0]] = double
				} else {
					res[kv[0]] = kv[1]
				}
			} else {
				s := strings.Split(str, "v=")
				if len(s) > 1 {
					version = strings.Split(s[1], "\n")[0]
				}
			}
		}
	}
	rv, err := strconv.Atoi(version)
	if err != nil {
		logger.Error(err)
		return 0, make(map[string]any)
	}
	return uint(rv), res
}
