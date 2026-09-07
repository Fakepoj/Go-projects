package main

import (
	"reflect"
	"testing"
)

func TestPrimeFactors(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{"42", 42, []int{2, 3, 7}},
		{"225225", 225225, []int{3, 3, 5, 5, 7, 11, 13}},
		{"prime", 9539, []int{9539}},
		{"invalid", 1, nil},
	}

	for _, tt := range tests {
		if got := PrimeFactors(tt.n); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%s: got %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestHidden(t *testing.T) {
	tests := []struct {
		s1, s2 string
		want    bool
	}{
		{"123", "123", true},
		{"faya", "fgvvfdxcacpolhyghbreda", true},
		{"faya", "fgvvfdxcacpolhyghbred", false},
		{"DD", "DABC", false},
		{"", "anything", true},
	}

	for _, tt := range tests {
		if got := Hidden(tt.s1, tt.s2); got != tt.want {
			t.Errorf("Hidden(%q, %q) = %v, want %v", tt.s1, tt.s2, got, tt.want)
		}
	}
}

func TestIntersection(t *testing.T) {
	got := Intersection("padinton", "paqefwtdjetyiytjneytjoeyjnejeyj")
	if got != "padinto" {
		t.Errorf("got %q, want %q", got, "padinto")
	}
}

func TestUnion(t *testing.T) {
	got := Union("zpadinton", "paqefwtdjetyiytjneytjoeyjnejeyj")
	if got != "zpadintoqefwjy" {
		t.Errorf("got %q, want %q", got, "zpadintoqefwjy")
	}
}

func TestSaveAndMiss(t *testing.T) {
	if got := SaveAndMiss("123456789", 3); got != "123789" {
		t.Errorf("got %q, want %q", got, "123789")
	}
	if got := SaveAndMiss("hello", 0); got != "hello" {
		t.Errorf("got %q, want %q", got, "hello")
	}
}
