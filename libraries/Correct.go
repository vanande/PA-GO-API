package libraries

import (
	"net/mail"
	"strconv"
	"strings"
)

func isValidName(name string) bool {
	return len(name) > 0 && len(name) <= 45
}

func isValidAge(age string) bool {
	_, err := strconv.Atoi(age)
	return err == nil
}

func isValidSexe(sexe string) bool {
	return len(sexe) == 1 && (sexe == "M" || sexe == "F")
}

func isValidPhone(phone string) bool {
	if len(phone) != 10 {
		return false
	}
	for _, char := range phone {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func isValidId(id string) bool {
	_, err := strconv.Atoi(id)
	return err == nil
}

// isValidField returns true if the field is valid, false otherwise.
// Example: isValidField("John", "nom") returns true and isValidField("John", "age") returns false.
func Correct(fieldValue string, fieldType string) bool {
	fieldValue = strings.TrimSpace(fieldValue)
	switch fieldType {
	case "nom", "prenom":
		return isValidName(fieldValue)
	case "age":
		return isValidAge(fieldValue)
	case "sexe":
		return isValidSexe(fieldValue)
	case "telephone":
		return isValidPhone(fieldValue)
	case "idCLIENT", "idEQUIPE":
		return isValidId(fieldValue)
	case "email":
		return isValidEmail(fieldValue)
	default:
		return false
	}
}
