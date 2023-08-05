package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kz-login/app/models"
	"github.com/kz-login/pkg/errors"
	"strconv"
)

func (h *defaultHandler) MemberProfile(ctx *fiber.Ctx) error {
	uid := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["sub"].(string)
	userId, err := strconv.Atoi(uid)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors.ErrBadParameter.FiberMap())
	}

	user, vErr := h.svc.Profile(int64(userId))
	if vErr != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(vErr.FiberMap())
	}

	response := models.ClientMemberProfileResp{
		UserId:       user.Id,
		Name:         user.Name,
		Email:        user.Email,
		MobileNumber: user.MobileNumber,
		BirthOfDate:  user.BirthOfDate,
	}
	return ctx.JSON(response)
}
