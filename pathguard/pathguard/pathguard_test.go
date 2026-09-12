package pathguard

import (
	"reflect"
	"testing"
)

func TestCanJump(t *testing.T) {
	cases := []struct { name string; in []uint; want bool }{
		{"reachable", []uint{2,3,1,1,4}, true},
		{"dead end", []uint{3,2,1,0,4}, false},
		{"single", []uint{0}, true},
		{"empty", []uint{}, false},
		{"exact", []uint{1,1,1}, true},
		{"out of bounds", []uint{4,1,1}, false},
		{"cycle", []uint{0,1}, false},
	}
	for _, tc := range cases { if got := CanJump(tc.in); got != tc.want { t.Errorf("%s: got %v, want %v", tc.name, got, tc.want) } }
}

func TestPath(t *testing.T) {
	got, ok := Path([]uint{2,3,1,1,4})
	want := []int{0,2,3,4}
	if !ok || !reflect.DeepEqual(got, want) { t.Errorf("got %v, %v; want %v, true", got, ok, want) }
}
