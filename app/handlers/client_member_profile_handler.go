package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kz-login/app/models"
)

func (h *defaultHandler) MemberProfile(ctx *fiber.Ctx) error {
	var req models.ClientMemberProfileReq
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	response := models.ClientMemberLoginResp{}
	return ctx.JSON(response)
}
