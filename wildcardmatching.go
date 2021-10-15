package wildcardmatching

import (
	"strings"
)

func areSame(s, p string) bool {
	for i := 0; i < len(p); i++ {
		if p[i] == '?' {
			continue
		}
		if p[i] != s[i] {
			return false
		}
	}
	return true
}

func backtrackIt(s string, p []string, sIndex, pIndex int, seen map[string]bool, firstHadStart, lastHadStar bool) bool {

	if pIndex == len(p) {
		return sIndex >= len(s) || lastHadStar
	}

	if sIndex >= len(s) {
		return false
	}

	if _, saw := seen[s[sIndex:]]; saw {
		return false
	}

	seen[s[sIndex:]] = true

	pWord := p[pIndex]

	if sIndex+len(pWord) <= len(s) && pIndex == 0 && !firstHadStart {
		return areSame(s[sIndex:sIndex+len(pWord)], pWord) && backtrackIt(s, p, sIndex+len(pWord), pIndex+1, seen, firstHadStart, lastHadStar)
	}

	for sIndex+len(pWord) <= len(s) {
		if areSame(s[sIndex:sIndex+len(pWord)], pWord) && backtrackIt(s, p, sIndex+len(pWord), pIndex+1, seen, firstHadStart, lastHadStar) {
			return true
		}
		sIndex++
	}

	return false
}

func isMatch(s string, p string) bool {

	if !strings.Contains(p, "*") {
		return len(s) == len(p) && areSame(s, p)
	}

	lastHadStar := p[len(p)-1] == '*'
	firstHadStart := p[0] == '*'
	p = strings.Trim(p, "*")
	if len(p) == 0 {
		return true
	}

	patternArr := strings.Split(p, "*")
	if len(patternArr) == 0 {
		return true
	}

	arr := make([]string, 0, len(patternArr))

	for _, str := range patternArr {
		if len(str) > 0 {
			arr = append(arr, str)
		}
	}
	return backtrackIt(s, arr, 0, 0, make(map[string]bool), firstHadStart, lastHadStar)
}
