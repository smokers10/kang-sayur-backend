package customer

import "fmt"

func verificationEmail(otp string, name string) string {
	return fmt.Sprintf(`
		<h1> Hallo, %s </h2> <br>
		Kode OTP Anda <b> %s </b>
	`, name, otp)
}

func forgotPasswordEmail(code string, name string) string {
	return fmt.Sprintf(`
	<h1> Hallo, %s </h2> <br>
	Kode reset password anda Anda <b> %s </b>
	`, name, code)
}
