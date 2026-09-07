package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func PrimeFactors(n int) []int {
	if n < 2 {
		return nil
	}

	var factors []int
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}

func Hidden(s1, s2 string) bool {
	j := 0
	for i := range s2 {
		if j < len(s1) && s2[i] == s1[j] {
			j++
		}
	}
	return j == len(s1)
}

func Intersection(s1, s2 string) string {
	var result []byte
	seen := make(map[byte]bool)

	for i := 0; i < len(s1); i++ {
		c := s1[i]
		if seen[c] {
			continue
		}
		for j := 0; j < len(s2); j++ {
			if c == s2[j] {
				result = append(result, c)
				seen[c] = true
				break
			}
		}
	}
	return string(result)
}

func Union(s1, s2 string) string {
	var result []byte
	seen := make(map[byte]bool)

	for i := 0; i < len(s1); i++ {
		if !seen[s1[i]] {
			result = append(result, s1[i])
			seen[s1[i]] = true
		}
	}
	for i := 0; i < len(s2); i++ {
		if !seen[s2[i]] {
			result = append(result, s2[i])
			seen[s2[i]] = true
		}
	}
	return string(result)
}

func SaveAndMiss(s string, n int) string {
	if n <= 0 {
		return s
	}

	var result []byte
	save := true

	for i := 0; i < len(s); i += n {
		end := i + n
		if end > len(s) {
			end = len(s)
		}
		if save {
			result = append(result, s[i:end]...)
		}
		save = !save
	}
	return string(result)
}

func printUsage() {
	fmt.Println(`Security Toolkit

Usage:
  security-tool factor <positive-integer>
  security-tool hidden <pattern> <text>
  security-tool inter <string1> <string2>
  security-tool union <string1> <string2>
  security-tool mask <text> <chunk-size>

Commands:
  factor  Prime-factorize an integer.
  hidden  Check whether a pattern occurs as an ordered subsequence.
  inter   Find unique characters shared by two strings.
  union   Combine unique characters while preserving first appearance.
  mask    Save and miss alternating chunks of text.`)
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "factor":
		if len(os.Args) != 3 {
			printUsage()
			return
		}
		n, err := strconv.Atoi(os.Args[2])
		if err != nil || n < 2 {
			fmt.Println("error: expected a positive integer >= 2")
			return
		}
		factors := PrimeFactors(n)
		for i, f := range factors {
			if i > 0 {
				fmt.Print("*")
			}
			fmt.Print(f)
		}
		fmt.Println()

	case "hidden":
		if len(os.Args) != 4 {
			printUsage()
			return
		}
		if Hidden(os.Args[2], os.Args[3]) {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}

	case "inter":
		if len(os.Args) != 4 {
			printUsage()
			return
		}
		fmt.Println(Intersection(os.Args[2], os.Args[3]))

	case "union":
		if len(os.Args) != 4 {
			printUsage()
			return
		}
		fmt.Println(Union(os.Args[2], os.Args[3]))

	case "mask":
		if len(os.Args) != 4 {
			printUsage()
			return
		}
		n, err := strconv.Atoi(os.Args[3])
		if err != nil || n <= 0 {
			fmt.Println("error: chunk-size must be greater than 0")
			return
		}
		fmt.Println(SaveAndMiss(os.Args[2], n))

	default:
		printUsage()
	}
}
