package router

import (
	"github.com/go-playground/validator/v10"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/kz-login/app/handlers"
	"github.com/kz-login/app/repositories"
	"github.com/kz-login/app/services"
	"github.com/kz-login/env"
	"github.com/kz-login/pkg/csredis"
	"github.com/kz-login/pkg/db"
	customjwt "github.com/kz-login/pkg/jwt"
	"log"
)

type validateErr struct {
	Field string      `json:"field"`
	Tag   string      `json:"tag"`
	Value interface{} `json:"value"`
}

type Options struct {
	Client db.Client
	CsJwt  customjwt.Client
	Rdc    csredis.Client
}

var validate = validator.New()

func Validate(data interface{}) []validateErr {
	var vErrs []validateErr

	errs := validate.Struct(data)
	if errs != nil {
		var vErr validateErr
		for _, err := range errs.(validator.ValidationErrors) {

			vErr.Field = err.Field()
			vErr.Tag = err.Tag()
			vErr.Value = err.Value()
			vErrs = append(vErrs, vErr)
		}
	}

	return vErrs
}

func NewRouter(cfg *env.Environment, opt *Options) *fiber.App {
	if cfg == nil {
		log.Fatal("no environment found")
	}

	app := fiber.New(fiber.Config{
		AppName: cfg.App.Name,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			log.Println("is error handler: ", ctx, err)
			return nil
			//return ctx.JSON(err)
			//vErr := errors.NewDefaultError(err)
			//return ctx.Status(fiber.StatusBadRequest).JSON(vErr)
		},
	})

	app.Get("/healthcheck", func(ctx *fiber.Ctx) error {
		return ctx.SendString("OK")
	})

	repo := repositories.NewRepository(opt.Client)
	service := services.NewService(repo, opt.CsJwt, opt.Rdc)
	handler := handlers.NewHandler(cfg, service)

	api := app.Group("/api")
	client := api.Group("/client")
	{
		client.Post("/login", handler.MemberLogin)
	}

	user := api.Group("/member")
	{
		user.Use(jwtware.New(jwtware.Config{
			SigningKey: jwtware.SigningKey{
				Key: []byte(cfg.JWT.JwtSecret),
			},
			ErrorHandler: func(ctx *fiber.Ctx, err error) error {
				//log.Println("error handler: ", ctx, err)
				return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"code": 401, "error": "Invalid or expired access token"})
			},
		}))
		user.Get("/profile", handler.MemberProfile)
		user.Post("/refresh_token", handler.RefreshToken)
		user.Get("/logout", handler.MemberLogout)
	}

	return app
}

//func protected(ctx *fiber.Ctx) error {
//	_ = ctx.Locals("user").(*jwt.Token)
//	claims := user.Claims.(jwt.MapClaims)
//	return nil
//}
