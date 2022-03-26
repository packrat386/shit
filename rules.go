package main

import (
	"os"
	"regexp"
)

var defaultRuleset = ruleset{
	rules: []rule{
		rule{
			match: executableMatcher,
			msg:   "executable files are inherently dangerous!!!",
			level: "criticalest",
		},
		rule{
			match: regexpMatcher(regexp.MustCompile(`\.py$`)),
			msg:   "Python is a programming language that can be used to execute malicious code!!!",
			level: "criticaler",
		},
		rule{
			match: regexpMatcher(regexp.MustCompile(`\.(js|jsx)$`)),
			msg:   "JavaScript is a programming language that can be used to execute malicious code!!!",
			level: "criticaler",
		},
		rule{
			match: regexpMatcher(regexp.MustCompile(`\.java$`)),
			msg:   "Java is a programming language that can be used to execute malicious code!!!",
			level: "criticaler",
		},
		rule{
			match: regexpMatcher(regexp.MustCompile(`\.(ts|tsx)$`)),
			msg:   "TypeScript is a programming language that can be used to execute malicious code!!!",
			level: "criticaler",
		},
		rule{
			match: regexpMatcher(regexp.MustCompile(`\.go$`)),
			msg:   "Go is a programming language that can be used to execute malicious code!!!",
			level: "criticaler",
		},
		rule{
			match: regexpMatcher(regexp.MustCompile(`\.(cc|cpp|hpp)$`)),
			msg:   "C++ is a programming language that can be used to execute malicious code!!!",
			level: "criticaler",
		},
		rule{
			match: regexpMatcher(regexp.MustCompile(`\.(rb|ru)$`)),
			msg:   "Ruby is a programming language that can be used to execute malicious code!!!",
			level: "criticaler",
		},
		rule{
			match: regexpMatcher(regexp.MustCompile(`\.(php|phtml)$`)),
			msg:   "PHP? Seriously?",
			level: "criticalest",
		},
		rule{
			match: regexpMatcher(regexp.MustCompile(`\.cs$`)),
			msg:   "C# is a programming language that can be used to execute malicious code!!!",
			level: "criticaler",
		},
		rule{
			match: regexpMatcher(regexp.MustCompile(`\.(c|h)$`)),
			msg:   "C is a programming language that can be used to execute malicious code!!!",
			level: "criticaler",
		},
		rule{
			match: nothingElseMatcher,
			msg:   "SHIT doesn't recognize this file, which is suspicious!!!",
			level: "critical",
		},
	},
}

type ruleset struct {
	rules []rule
}

func (rs ruleset) evaluate(filename string) []violation {
	v := make([]violation, 0)

	for _, r := range rs.rules {
		if r.match(filename, v) {
			v = append(v, violation{
				Filename: filename,
				Message:  r.msg,
				Level:    r.level,
			})
		}
	}

	return v
}

type rule struct {
	match fileMatcher
	msg   string
	level string
}

type violation struct {
	Filename string
	Message  string
	Level    string
}

type fileMatcher func(filename string, vv []violation) bool

func executableMatcher(filename string, vv []violation) bool {
	fi, err := os.Stat(filename)
	if err != nil {
		return false
	}

	return fi.Mode()&0111 != 0
}

func nothingElseMatcher(filename string, vv []violation) bool {
	return len(vv) == 0
}

func regexpMatcher(re *regexp.Regexp) fileMatcher {
	return func(filename string, vv []violation) bool {
		return re.MatchString(filename)
	}
}
