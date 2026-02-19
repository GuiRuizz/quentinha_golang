package users_domain

import (
	"golang.org/x/crypto/bcrypt"
)

func (ud *userDomain) EncryptPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(ud.password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	ud.password = string(hashedPassword)
	return nil
}
