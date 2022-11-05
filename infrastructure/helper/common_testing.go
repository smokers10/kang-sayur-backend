package helper

import (
	"kang-sayur-backend/infrastructure/middleware"
	"kang-sayur-backend/model/web"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Expected struct {
	Message   string
	Status    int
	IsSuccess bool
	Data      interface{}
	Token     string
}

type ExpectedMiddlewareResponse struct {
	Message string `json:"message,omitempty"`
	Status  int    `json:"status,omitempty"`
	Is_pass bool   `json:"is_password,omitempty"`
	Reason  string `json:"reason,omitempty"`
	Claim   struct {
		Id    string `json:"id,omitempty"`
		Email string `json:"email,omitempty"`
	}
}

type Options struct {
	DataChecking  bool
	TokenChecking bool
}

var DefaultOption = &Options{
	DataChecking:  false,
	TokenChecking: false,
}

type unitTesting struct{}

func UnitTesting() *unitTesting {
	return &unitTesting{}
}

func (ut *unitTesting) CommonAssertion(t *testing.T, expected *Expected, actual *web.HTTPResponse, option *Options) {
	// response should not empty
	assert.NotEmpty(t, actual.Message)
	assert.NotEmpty(t, actual.Status)

	// response  expected detail must match w/ actual response
	assert.Equal(t, expected.Message, actual.Message)
	assert.Equal(t, expected.IsSuccess, actual.IsSuccess)

	// if token must be checked
	if option.TokenChecking {
		assert.NotEmpty(t, actual.Token)
	}

	// if data must checked
	if option.DataChecking {
		assert.NotEmpty(t, actual.Data)
	}
}

func (ut *unitTesting) MiddlewareAssertion(t *testing.T, expected *middleware.MiddlewareResponse, actual *middleware.MiddlewareResponse) {
	// response should not empty
	assert.NotEmpty(t, expected.Message)
	assert.NotEmpty(t, expected.Status)
	assert.NotEmpty(t, expected.Reason)

	// response  expected detail must match w/ actual response
	assert.Equal(t, expected.Message, actual.Message)
	assert.Equal(t, expected.Status, actual.Status)
	assert.Equal(t, expected.Is_pass, actual.Is_pass)
	assert.Equal(t, expected.Reason, actual.Reason)
	assert.Equal(t, expected.Claim, actual.Claim)

	// test claim
	if actual.Claim.Email != "" {
		assert.NotEmpty(t, expected.Claim.Email)
	}

	if actual.Claim.Id != "" {
		assert.NotEmpty(t, expected.Claim.Id)
	}
}
