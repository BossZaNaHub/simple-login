package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kz-login/app/services"
	"github.com/kz-login/env"
)

type defaultHandler struct {
	cfg *env.Environment
	svc services.Service
}

type ClientHandler interface {
	MemberLogin(ctx *fiber.Ctx) error
	MemberLogout(ctx *fiber.Ctx) error
	MemberProfile(ctx *fiber.Ctx) error
	RefreshToken(ctx *fiber.Ctx) error
}

func NewHandler(cfg *env.Environment, svc services.Service) ClientHandler {
	return &defaultHandler{cfg: cfg, svc: svc}
}
