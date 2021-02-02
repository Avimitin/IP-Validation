package validIP

import "errors"

var validDigit = []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func IsDigit(c rune) bool {
	for _, r := range validDigit {
		if c == r {
			return true
		}
	}
	return false
}

func getIPPart(ip string) ([4][3]rune, error) {
	var part [3]rune
	var whole [4][3]rune
	var partPointer int
	var wholePointer int
	for _, r := range ip {
		if r == '.' {
			partPointer = 0
			if wholePointer == 4 {
				return whole, errors.New("too more dot")
			}
			whole[wholePointer] = part
			wholePointer++
			continue
		}
		if partPointer == 3 {
			return whole, errors.New("too large part")
		}
		part[partPointer] = r
		partPointer++
	}
	return whole, nil
}

func isValidIP(ip string) bool {
	whole, err := getIPPart(ip)
	if err != nil {
		return false
	}
	for i := 0; i < 4; i++ {
		switch len(whole[i]) {
		case 0:
			return false
		case 1:
			continue
		case 2:
			if whole[i][0] == '0' {
				return false
			}
		case 3:
			if whole[i][0] == '0' {
				return false
			}
			if !lessThanFive(whole[i][1]) {
				return false
			}
			if whole[i][1] == '5' {
				if !lessThanFive(whole[i][2]) {
					return false
				}
			}
		}
	}
	return true
}

func lessThanFive(r rune) bool {
	for _, i := range []rune{'1', '2', '3', '4', '5', '0'} {
		if r == i {
			return true
		}
	}
	return false
}
