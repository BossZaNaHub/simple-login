package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kz-login/app/models"
	"github.com/kz-login/pkg/errors"
)

func (h *defaultHandler) MemberLogin(ctx *fiber.Ctx) error {
	var req models.ClientMemberLoginReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors.ErrBadParameter.FiberMap())
	}

	user, token, vErr := h.svc.Login(&models.MemberLoginData{
		MobileNumber: req.MobileNumber,
		Password:     req.Password,
	})
	if vErr != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(vErr.FiberMap())
	}

	response := models.ClientMemberLoginResp{
		UserId:              user.Id,
		Name:                user.Name,
		Email:               user.Email,
		MobileNumber:        user.MobileNumber,
		AccessToken:         token.AccessToken,
		AccessTokenExpired:  token.AccessTokenExpired,
		RefreshToken:        token.RefreshToken,
		RefreshTokenExpired: token.RefreshTokenExpired,
	}

	return ctx.JSON(response)
}
