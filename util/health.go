package util

type Health int16

const (
	Weak        Health = 0
	Moderate    Health = 1
	Strong      Health = 2
	Interesting Health = 3
	Secure      Health = 4
)

func NewPasswordHealth(password string) Health {
	if len(password) < 8 {
		return Weak
	}

	hasUppercase := hasUppercase(password)
	hasLowercase := hasLowercase(password)
	hasDigit := hasDigit(password)
	hasSpecial := hasSpecial(password)

	switch {
	case len(password) < 10:
		if hasUppercase && hasLowercase && hasDigit {
			return Moderate
		}
	case len(password) < 12:
		if hasUppercase && hasLowercase && hasDigit && hasSpecial {
			return Strong
		}
	default:
		if hasUppercase && hasLowercase && hasDigit && hasSpecial && len(password) >= 16 {
			return Secure
		}
	}

	return Interesting
}

func hasUppercase(password string) bool {
	for _, char := range password {
		if 'A' <= char && char <= 'Z' {
			return true
		}
	}
	return false
}

func hasLowercase(password string) bool {
	for _, char := range password {
		if 'a' <= char && char <= 'z' {
			return true
		}
	}
	return false
}

func hasDigit(password string) bool {
	for _, char := range password {
		if '0' <= char && char <= '9' {
			return true
		}
	}
	return false
}

func hasSpecial(password string) bool {
	for _, char := range password {
		if isSpecial(char) {
			return true
		}
	}
	return false
}

func isSpecial(r rune) bool {
	return '!' <= r && r <= '/' || ':' <= r && r <= '@' || '[' <= r && r <= '`' || '{' <= r && r <= '~'
}
