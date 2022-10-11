package service

import "fmt"

func AdminPasswordEmail(password string) string {
	return fmt.Sprintf(`
		OTP %s \n
		Jangan sampai orang lain tau!!!
	`, password)
}
