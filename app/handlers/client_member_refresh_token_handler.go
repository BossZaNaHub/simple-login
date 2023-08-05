package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kz-login/app/models"
	"github.com/kz-login/pkg/errors"
	"strconv"
)

func (h *defaultHandler) RefreshToken(ctx *fiber.Ctx) error {
	uid := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["sub"].(string)
	userId, err := strconv.Atoi(uid)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors.ErrBadParameter.FiberMap())
	}

	var req models.ClientMemberRefreshTokenReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors.ErrBadParameter.FiberMap())
	}

	t, vErr := h.svc.Refresh(int64(userId), req.RefreshToken)
	if vErr != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(vErr.FiberMap())
	}

	response := models.ClientMemberRefreshTokenResp{
		AccessToken: t.AccessToken,
	}

	return ctx.JSON(response)
}
