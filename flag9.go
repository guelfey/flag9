// Package flag9 provides Plan 9-like flag parsing.
//
// Though it is not as convenient to use as the flag package from the standard
// library, it provides more control over what portion of the arguments is
// actually parsed.
//
// See the example for a typical usage case.
package flag9

import (
	"os"
	"strings"
	"unicode/utf8"
)

// Args are a set of arguments being parsed.
type Args struct {
	s    []string
	c    rune
	cur  string
}

// NewArgs returns an Args structure that can be used to parse the given slice
// of strings.
func NewArgs(s []string) *Args {
	return &Args{s, utf8.RuneError, ""}
}

// Argc returns the current option character. If it is called before the first
// call to Next or after it returned false, the error rune is returned.
func (a *Args) Argc() rune {
	return a.c
}

// Argf tries to return the current option argument (the rest of the option
// string if it's not empty, or the next member in the slice). If none is
// present, the empty string and false are returned. Otherwise, the option
// argument and true are returned.
//
// It must not be called multiple times for the same argument.
func (a *Args) Argf() (string, bool) {
	cur := a.cur
	a.cur = ""
	if cur == "" {
		if len(a.s) > 0 {
			s := a.s[0]
			a.s = a.s[1:]
			return s, true
		}
		return "", false
	}
	return cur, true
}

// Argv returns the arguments that are not (yet) processed.
func (a *Args) Argv() []string {
	return a.s
}

// Next tries to read the next option character. If it is successfull, it
// returns true and Argc will return the option character. Otherwise, it returns
// false.
func (a *Args) Next() bool {
	if a.cur == "" {
		if len(a.s) == 0 {
			a.c = utf8.RuneError
			return false
		}
		switch {
		case a.s[0] == "--":
			a.s = a.s[1:]
			fallthrough
		case a.s[0] == "-" || !strings.HasPrefix(a.s[0], "-"):
			a.c = utf8.RuneError
			return false
		}
		a.cur = a.s[0][1:]
		a.s = a.s[1:]
	}
	c, size := utf8.DecodeRuneInString(a.cur)
	a.cur = a.cur[size:]
	a.c = c
	return true
}

var cmdline = &Args{os.Args[1:], utf8.RuneError, ""}

// Argc returns the current option character from the command-line arguments. If
// it is called before the first call to Next or after it returned false, the
// error rune is returned.
func Argc() rune {
	return cmdline.Argc()
}

// Argf tries to return the current option argument (the rest of the option
// string if it's not empty, or the next argument) from the command-line
// arguments. If none is present, the empty string and false are returned.
// Otherwise, the option argument and true are returned.
//
// It must not be called multiple times for the same argument.
func Argf() (string, bool) {
	return cmdline.Argf()
}

// Argv returns the command-line arguments that are not (yet) processed.
func Argv() []string {
	return cmdline.Argv()
}

// Next tries to read the next option character from the command-line arguments.
// If it is successfull, it returns true and Argc will return the option
// character. Otherwise, it returns false.
func Next() bool {
	return cmdline.Next()
}
