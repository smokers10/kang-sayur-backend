package service

import "fmt"

func AdminPasswordEmail(password string) string {
	return fmt.Sprintf(`
		Password %s <br>
		<h1>
		Jangan sampai orang lain tau!!!
		</h1>
	`, password)
}
