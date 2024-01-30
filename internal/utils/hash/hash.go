package hash

import "golang.org/x/crypto/bcrypt"

func GenerateHashPassword(pwd string) string {
	password, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return string(password)
	}
	return string(password)
}

func ComparePassword(hashPwd, pwd string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd))
	return err
}
