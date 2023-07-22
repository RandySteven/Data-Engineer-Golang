package utils

import (
	"Week2ofGolangDataEngineeringSyllabus/models"
	"encoding/base64"
)

type SecurityUtil struct{}

var securityInstance SecurityUtil

func (su *SecurityUtil) GetInstance() *SecurityUtil {
	return &securityInstance
}

func (su *SecurityUtil) UserRequestValidation(user *models.User) map[string]bool {
	requestValidation := make(map[string]bool)

	return requestValidation
}

func (su *SecurityUtil) PasswordEncryption(userPassword string) string {
	encoding := base64.StdEncoding.EncodeToString([]byte(userPassword))
	return encoding
}
