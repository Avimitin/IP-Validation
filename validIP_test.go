package validIP

import (
	"reflect"
	"testing"
)

func TestIsDigit(t *testing.T) {
	t.Run("number rune test", func(t *testing.T) {
		got := IsDigit('1')
		want := true
		if got != want {
			t.Errorf("Got %t want %t", got, want)
		}
	})

	t.Run("letter rune test", func(t *testing.T) {
		got := IsDigit('a')
		want := false
		if got != want {
			t.Errorf("Got %t want %t", got, want)
		}
	})
}

func TestIsValid(t *testing.T) {
	t.Run("192.168.1.1", func(t *testing.T) {
		got := isValidIP("192.168.1.1")
		want := true
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
}

func TestGetIPPart(t *testing.T) {
	t.Run("192.168.1.1", func(t *testing.T) {
		whole, err := getIPPart("192.168.1.1")
		if err != nil {
			t.Fatal(err)
		}
		wantedPart := [][]rune{
			{'1', '9', '2'},
			{'1', '6', '8'},
			{'1'},
			{'1'},
		}
		if !reflect.DeepEqual(whole, wantedPart) {
			t.Errorf("Got %v want %v", whole, wantedPart)
		}
	})
}
