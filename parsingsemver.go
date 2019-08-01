package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`^v(\d+\.)?(\d+\.)?(\*|\d+)$`)

type semVer struct {
	major, minor, patch uint
}

func parseSemVer(s string) (out semVer, err error) {
	if !re.Match([]byte(s)) {
		return semVer{}, fmt.Errorf("invalid semantic version string: %v", s)
	}
	for i, num := range strings.Split(s[1:], ".") {
		n, err := strconv.Atoi(num)
		if err != nil {
			return semVer{}, fmt.Errorf("could not parse '%v' as semantic version: %v", s, err)
		}
		if n < 0 {
			return out, fmt.Errorf("invalid semantic version string: %v", s)
		}
		u := uint(n)
		switch i {
		case 0:
			out.major = u
		case 1:
			out.minor = u
		case 2:
			out.patch = u
		default:

			return out, fmt.Errorf("invalid semantic version string: %v", s)
		}
	}
	return out, nil
}

func main() {

	testCases := []struct {
		s       string
		matches bool
		semVer
	}{
		{
			"v1.0.2",
			true,
			semVer{major: 1, minor: 0, patch: 2},
		},

		{
			"v0",
			true,
			semVer{major: 0, minor: 0, patch: 0},
		},

		{
			"v2.2",
			true,
			semVer{major: 2, minor: 2, patch: 0},
		},

		{
			"v1.2.3.4",
			false,
			semVer{},
		},
	}

	for _, tc := range testCases {
		if re.Match([]byte(tc.s)) != tc.matches {
			log.Fatalf("failed on %+v\n", tc)
		}
		successfullyTestedForFailure := !tc.matches
		if successfullyTestedForFailure {
			continue
		}
		semVer, err := parseSemVer(tc.s)
		if err != nil {
			log.Fatalf("failed on '%+v': %v\n", tc, err)
		}
		if semVer != tc.semVer {
			log.Fatalf("expected '%+v' but received '%+v'\n", tc.semVer, semVer)
		}
	}

	fmt.Println("passed tests")
}
