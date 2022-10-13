package validation

import (
	response "kang-sayur-backend/model/web"
	"net"
	"net/http"
	"net/mail"
	"strings"
)

type validation struct {
	Email    EmailValidation
	String   StringValidation
	Array    ArrayValidation
	Response Response
}

type EmailValidation struct{}

type StringValidation struct{}

type ArrayValidation struct{}

type Response struct{}

func Validation() *validation { return &validation{} }

// email validation
func (ve *EmailValidation) EmailRequired(email string) bool { return email == "" }

func (ve *EmailValidation) EmailValidity(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return true
	}

	domainPart := strings.Split(email, "@")
	_, err = net.LookupMX(domainPart[1])
	return err != nil
}

// string validation
func (vs *StringValidation) Required(val string) bool { return val == "" }

func (vs *StringValidation) Length(val string, min int, max int) bool {
	if val == "" {
		return true
	}

	if len(val) < min {
		return true
	}

	if len(val) > max {
		return true
	}

	return false
}

func (vs *StringValidation) Match(val1 string, val2 string) bool { return val1 != val2 }

// array validation
func (va *ArrayValidation) NotEmpty(length int) bool { return length == 0 }

// response
func (r *Response) SendResponse(message string) response.HTTPResponse {
	return response.HTTPResponse{
		Message: message,
		Status:  http.StatusBadRequest,
	}
}
