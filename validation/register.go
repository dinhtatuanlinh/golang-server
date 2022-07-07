package validation

import "regexp"

type validate interface {
	ValidatePassword(password, repassword string) bool
	IsEmailValid(email string) bool
}

type Validation struct {
}

func (v Validation) ValidatePassword(password, repassword string) bool {
	var emailRegex = regexp.MustCompile("^(?=.*[A-Z].*[A-Z])(?=.*[!@#$&*])(?=.*[0-9].*[0-9])(?=.*[a-z].*[a-z].*[a-z]).{8}$")

	if password != repassword {
		return false
	}

	if (len(password) < 6 && len(password) > 20){
		return false
	}

	return emailRegex.MatchString(password)
}

func (v Validation) IsEmailValid(e string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if len(e) < 10 && len(e) > 100 {
		return false
	}
	return emailRegex.MatchString(e)
}

func (v Validation) IsUsernameValid(e string) bool{
	if len(e) < 6 && len(e) > 100 {
		return false
	}

	return true
}