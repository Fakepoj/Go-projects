package processor

import (
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	got := Chunk([]int{0,1,2,3,4,5,6,7}, 3)
	want := [][]int{{0,1,2},{3,4,5},{6,7}}
	if !reflect.DeepEqual(got, want) { t.Fatalf("got %v, want %v", got, want) }
}

func TestConcatSlice(t *testing.T) {
	got := ConcatSlice([]int{1,2,3}, []int{4,5,6})
	want := []int{1,2,3,4,5,6}
	if !reflect.DeepEqual(got, want) { t.Fatalf("got %v, want %v", got, want) }
}

func TestConcatAlternate(t *testing.T) {
	got := ConcatAlternate([]int{1,2,3}, []int{4,5,6,7})
	want := []int{4,1,5,2,6,3,7}
	if !reflect.DeepEqual(got, want) { t.Fatalf("got %v, want %v", got, want) }
}

func TestNormalizeLastLetter(t *testing.T) {
	if got := NormalizeLastLetter("First SMALL TesT"); got != "firsT smalL tesT" { t.Fatalf("got %q", got) }
}
