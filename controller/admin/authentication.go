package admin

import (
	"kang-sayur-backend/infrastructure/helper/validation"
	"kang-sayur-backend/model/domain/admin"

	request_body "kang-sayur-backend/model/web/request_body/admin"

	"github.com/gofiber/fiber/v2"
)

type authenticationController struct {
	adminService admin.AdminService
}

func (ac *adminController) AuthenticationController() authenticationController {
	return authenticationController{adminService: ac.AuthenticationController().adminService}
}

func (ac *authenticationController) Request(c *fiber.Ctx) error {
	body := request_body.LoginRequest{}
	if err := c.BodyParser(&body); err != nil {
		panic(err)
	}

	e, m := requestValidation(body.Email)
	if e {
		response := validation.Validation().Response.SendResponse(m)
		return c.Status(response.Status).JSON(response)
	}

	response := ac.adminService.RequestLogin(body)

	return c.Status(response.Status).JSON(response)
}

func (ac *authenticationController) Login(c *fiber.Ctx) error {
	body := &request_body.Login{}
	if err := c.BodyParser(body); err != nil {
		panic(err)
	}

	e, m := loginValidation(body.Email, body.Password)
	if e {
		response := validation.Validation().Response.SendResponse(m)
		return c.Status(response.Status).JSON(response)
	}

	response := ac.adminService.Login(*body)

	return c.Status(response.Status).JSON(response)
}

func requestValidation(email string) (bool, string) {
	v := validation.Validation()

	validity := v.Email.EmailValidity(email)
	required := v.Email.EmailRequired(email)

	if required {
		return required, "email tidak boleh kosong"
	}

	if validity {
		return validity, "email tidak valid"
	}

	return false, ""
}

func loginValidation(email string, password string) (bool, string) {
	v := validation.Validation()

	required := v.Email.EmailRequired(email)
	pwRequired := v.String.Required(password)

	if required {
		return required, "email tidak boleh kosong"
	}

	if pwRequired {
		return pwRequired, "password tidak boleh kosong"
	}

	return false, ""
}
