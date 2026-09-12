package pathguard

import "fmt"

func CanJump(a []uint) bool {
	if len(a) == 0 { return false }
	seen := map[int]bool{}
	pos := 0
	for {
		if pos == len(a)-1 { return true }
		if seen[pos] { return false }
		seen[pos] = true
		next := pos + int(a[pos])
		if next < 0 || next >= len(a) { return false }
		pos = next
	}
}

func Path(a []uint) ([]int, bool) {
	if len(a) == 0 { return nil, false }
	var path []int
	seen := map[int]bool{}
	pos := 0
	for {
		if seen[pos] { return path, false }
		seen[pos] = true
		path = append(path, pos)
		if pos == len(a)-1 { return path, true }
		next := pos + int(a[pos])
		if next < 0 || next >= len(a) { return path, false }
		pos = next
	}
}

func Validate(a []uint) (string, bool) {
	p, ok := Path(a)
	if ok { return fmt.Sprintf("VALID: %v", p), true }
	return fmt.Sprintf("INVALID: %v", p), false
}
