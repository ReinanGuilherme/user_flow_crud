package security

import (
	"golang.org/x/crypto/bcrypt"
)

// ConverterSenhaParaHashBCrypt converte uma senha em texto puro para um hash bcrypt.
func ConverterSenhaParaHashBCrypt(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// VerificarSenhaComHashBCrypt compara uma senha em texto puro com um hash bcrypt.
func VerificarSenhaComHashBCrypt(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
