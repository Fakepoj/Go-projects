package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"pathguard/pathguard"
)

func main() {
	if len(os.Args) < 2 { fmt.Println("Usage: pathguard <steps...>"); return }
	steps := make([]uint, len(os.Args)-1)
	for i, arg := range os.Args[1:] {
		n, err := strconv.ParseUint(arg, 10, 64)
		if err != nil { fmt.Printf("invalid input %q: expected a non-negative integer\n", arg); return }
		steps[i] = uint(n)
	}
	p, ok := pathguard.Path(steps)
	if ok { fmt.Printf("VALID: %s\n", format(p)) } else { fmt.Printf("INVALID: %s\n", format(p)) }
}

func format(p []int) string {
	s := make([]string, len(p))
	for i, n := range p { s[i] = strconv.Itoa(n) }
	return strings.Join(s, " -> ")
}
