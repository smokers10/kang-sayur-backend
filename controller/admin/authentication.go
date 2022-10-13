package admin

import (
	validation "kang-sayur-backend/infrastructure/helper/validation"
	"kang-sayur-backend/model/domain/admin"
	request_body "kang-sayur-backend/model/web/request_body/admin"
	service "kang-sayur-backend/service/admin"

	"github.com/gofiber/fiber/v2"
)

type adminController struct {
	adminService admin.AdminService
}

func (mc *mainController) Authentication() adminController {
	service := service.AdminService(&mc.injector.Repository.Admin, &mc.injector.JWT, &mc.injector.Encryption, &mc.injector.SMTP, &mc.injector.Identifier)
	return adminController{adminService: service}
}

func (ac *adminController) Request(c *fiber.Ctx) error {
	body := request_body.LoginRequest{}
	if err := c.BodyParser(&body); err != nil {
		panic(err)
	}

	e, m := RequestValidation(body.Email)
	if e {
		response := validation.Validation().Response.SendResponse(m)
		return c.Status(response.Status).JSON(response)
	}

	response := ac.adminService.RequestLogin(body)

	return c.Status(response.Status).JSON(response)
}

func (ac *adminController) Login(c *fiber.Ctx) error {
	body := &request_body.Login{}
	if err := c.BodyParser(body); err != nil {
		panic(err)
	}

	e, m := LoginValidation(body.Email, body.Password)
	if e {
		response := validation.Validation().Response.SendResponse(m)
		return c.Status(response.Status).JSON(response)
	}

	response := ac.adminService.Login(*body)

	return c.Status(response.Status).JSON(response)
}

// ==== VALIDATION ====

func RequestValidation(email string) (bool, string) {
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

func LoginValidation(email string, password string) (bool, string) {
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
