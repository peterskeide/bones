package entities

type User struct {
	Id       int
	Email    string
	Password string
}

func UserEquals(a, b []User) bool {
	if len(a) != len(b) {
		return false
	}

	for i, user := range a {
		if b[i] != user {
			return false
		}
	}

	return true
}
