package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kz-login/app/models"
	"github.com/kz-login/pkg/errors"
	"log"
)

func (h *defaultHandler) RefreshToken(ctx *fiber.Ctx) error {
	var req models.ClientMemberRefreshTokenReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors.ErrBadParameter.FiberMap())
	}

	t, err := h.svc.Refresh(req.RefreshToken)
	if err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.FiberMap())
	}

	response := models.ClientMemberRefreshTokenResp{
		AccessToken: t.AccessToken,
	}

	return ctx.JSON(response)
}
