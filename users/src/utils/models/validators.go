package models

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Custom validator function for struct-level validation
func PasswordFieldValidator(sl validator.StructLevel) {
	user := sl.Current().Interface().(User)

	// Custom password validation (at least one uppercase, one number, one special character)
	re := regexp.MustCompile(`^(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`)
	if !re.MatchString(user.Password) {
		sl.ReportError(user.Password, "Password", "password", "password_strength", "")
	}
}
