package validIP

var (
	validDigit = []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
)

type ipPart struct {
	part [3]rune
	bit  int
}

func IsDigit(c rune) bool {
	for _, r := range validDigit {
		if c == r {
			return true
		}
	}
	return false
}

func IsAllDigit(part [3]rune) bool {
	for i := 0; i < len(part); i++ {
		if part[i] == 0 {
			break
		}
		if !IsDigit(part[i]) {
			return false
		}
	}
	return true
}

func getIPPart(ip string) ([4]*ipPart, int32) {
	var whole [4]*ipPart
	var partPointer int
	var wholePointer int
	for i := 0; i < 4; i++ {
		whole[i] = &ipPart{
			part: [3]rune{},
			bit:  0,
		}
	}

	for _, r := range ip {
		if !IsDigit(r) && r != '.' {
			return whole, 0
		}
		if r == '.' {
			partPointer = 0
			wholePointer++
			continue
		}
		if wholePointer == 4 {
			return whole, 0
		}
		whole[wholePointer].part[partPointer] = r
		partPointer++
		whole[wholePointer].bit++
	}
	return whole, 1
}

func IsZero(r rune) bool {
	return r == '0'
}

func IPIsValid(ip string) bool {
	if part, ok := getIPPart(ip); ok == 1 {
		return ipIsValid(part)
	}
	return false
}

func ipIsValid(ipParts [4]*ipPart) bool {
	if len(ipParts) < 4 {
		return false
	}
	for _, part := range ipParts {
		switch part.bit {
		case 0:
			return false
		case 1:
			if !IsAllDigit(part.part) {
				return false
			}
		case 2:
			if !IsAllDigit(part.part) || IsZero(part.part[0]) {
				return false
			}
		case 3:
			if !IsAllDigit(part.part) || IsZero(part.part[0]) {
				return false
			}
			if part.part[0] != '1' && part.part[0] != '2' {
				return false
			}
			if part.part[0] == '2' {
				if !lessThanFive(part.part[1]) {
					return false
				}
				if part.part[1] == '5' {
					if !lessThanFive(part.part[2]) {
						return false
					}
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
