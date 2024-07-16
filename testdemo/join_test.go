package testdemo

import (
	"testing"
)

func TestTwoElements(t *testing.T) {
	phrases := []string{"a", "b"}
	want := "a and b"
	got := JoinWithCommas(phrases)
	if got != want {
		t.Errorf("JoinWithCommas(%#v) = \"%s\",want\"%s\"", phrases, got, want)
	}

}

func TestThreeElements(t *testing.T) {
	phrases := []string{"a", "b", "c"}
	if JoinWithCommas(phrases) != "a,b,and c" {
		t.Error("error")
	}
}
