package processor

func Chunk(data []int, size int) [][]int {
	if size <= 0 { return nil }
	out := make([][]int, 0, (len(data)+size-1)/size)
	for i := 0; i < len(data); i += size {
		end := i + size
		if end > len(data) { end = len(data) }
		out = append(out, data[i:end])
	}
	return out
}

func ConcatSlice(a, b []int) []int {
	out := make([]int, 0, len(a)+len(b))
	return append(out, append(a, b...)...)
}

func ConcatAlternate(a, b []int) []int {
	out := make([]int, 0, len(a)+len(b))
	first, second := a, b
	if len(b) > len(a) { first, second = b, a }
	for i := 0; i < len(first); i++ {
		out = append(out, first[i])
		if i < len(second) { out = append(out, second[i]) }
	}
	return out
}

func NormalizeLastLetter(s string) string {
	r := []rune(s)
	for i, c := range r {
		if c >= 'A' && c <= 'Z' { r[i] += 'a' - 'A' }
	}
	for i, c := range r {
		if c >= 'a' && c <= 'z' && (i == len(r)-1 || r[i+1] == ' ') { r[i] -= 'a' - 'A' }
	}
	return string(r)
}
