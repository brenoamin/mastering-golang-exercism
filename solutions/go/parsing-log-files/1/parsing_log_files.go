package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re, err := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
    if err != nil {
        panic(err)
    }

    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re, err := regexp.Compile(`<[~*=\-]*>`)
    if err != nil {
        panic(err)
    }
    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
    count := 0
    re, err := regexp.Compile(`(?i)"[^"]*password[^"]*"`)
    if err != nil {
        panic(err)
    }
    for _, line := range lines {
         if re.MatchString(line) {
             count ++
         }
    }
    return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
    clean := re.ReplaceAllString(text, "")

    return clean
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\S+)`)
	var result []string

    for _, line := range lines {
		match := re.FindStringSubmatch(line)        
        if match != nil {
            username := match[1]
            line = "[USR] " + username + " " + line
        }
        result = append(result, line)
    }
    return result
}
