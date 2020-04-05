package model

import (
	"github.com/asaskevich/govalidator"
	"regexp"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)

	const Username = "^[a-zA-Z0-9]+$"

	var rxUsername = regexp.MustCompile(Username)

	govalidator.TagMap["username"] = govalidator.Validator(func(str string) bool {
		return rxUsername.MatchString(str)
	})
}
