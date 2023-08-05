package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kz-login/pkg/errors"
	"strconv"
)

func (h *defaultHandler) MemberLogout(ctx *fiber.Ctx) error {
	uid := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["sub"].(string)
	userId, err := strconv.Atoi(uid)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors.ErrBadParameter.FiberMap())
	}

	h.svc.Logout(int64(userId))

	return ctx.JSON(fiber.Map{})
}
