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

func TestIsAllDigit(t *testing.T) {
	t.Run("test 1.1.1.1", func(t *testing.T) {
		if !IsAllDigit([3]rune{'1'}) {
			t.Errorf("fuck")
		}
	})
}

func TestGetIPPart(t *testing.T) {
	test := func(ip string, wantedPart [4][3]rune) {
		t.Run(ip, func(t *testing.T) {
			whole, ok := getIPPart(ip)
			if ok == 0 {
				t.Fatal("unwanted part")
			}
			if !reflect.DeepEqual(whole, wantedPart) {
				t.Errorf("Got %v want %v", whole, wantedPart)
			}
		})
	}

	test("192.168.1.1", [4][3]rune{
		{'1', '9', '2'},
		{'1', '6', '8'},
		{'1'},
		{'1'},
	})

	test("1.1.1.1", [4][3]rune{
		{'1'},
		{'1'},
		{'1'},
		{'1'},
	})

	test("1.2.3.4", [4][3]rune{
		{'1'},
		{'2'},
		{'3'},
		{'4'},
	})

}

func TestIsValid(t *testing.T) {
	testIP := func(t *testing.T, ip string, want bool) {
		t.Run(ip, func(t *testing.T) {
			got := IPIsValid(ip)
			if got != want {
				t.Errorf("%s's validation should be %t, but got %t", ip, want, got)
			}
		})
	}

	testIP(t, "255.168.1.1", true)
	testIP(t, "1.2.3.4", true)
	testIP(t, "114.514.19.810", false)
	testIP(t, "1.2.3", false)
	testIP(t, "1.2.3.4.5", false)
	testIP(t, "01.02.03.045", false)
	testIP(t, "abc.gh.def.ijk", false)
	testIP(t, "12.255.56.1", true)
	testIP(t, "12-255-56-1", false)
}
