package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kz-login/app/models"
)

func (h *defaultHandler) RefreshToken(ctx *fiber.Ctx) error {
	var req models.ClientMemberRefreshTokenReq
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	response := models.ClientMemberRefreshTokenResp{}
	return ctx.JSON(response)
}
