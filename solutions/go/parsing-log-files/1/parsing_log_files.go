package parsinglogfiles

import (
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<(\*|-|=|~)*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`".*(?i)password.*"`)
	count := 0
	for _, line := range lines {
		count += len(re.FindAllString(line, -1))
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line(\d)*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	// 匹配 "User" + 一个或多个空格 + 用户名
	pattern := regexp.MustCompile(`User\s+(\S+)`)

	result := make([]string, len(lines))

	for i, line := range lines {
		// 查找匹配并提取用户名
		matches := pattern.FindStringSubmatch(line)

		if matches != nil {
			// matches[0] 是完整匹配，matches[1] 是用户名
			username := matches[1]
			result[i] = "[USR] " + username + " " + line
		} else {
			// 没有匹配，保持原样
			result[i] = line
		}
	}

	return result
}
