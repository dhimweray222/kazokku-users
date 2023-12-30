package controller

import (
	"log"
	"strconv"

	"github.com/dhimweray222/users/exception"
	"github.com/dhimweray222/users/model/web"
	"github.com/dhimweray222/users/service"
	"github.com/gofiber/fiber/v2"
)

type UserControllerImpl struct {
	UserService service.UserService
}

type UserController interface {
	NewUserRouter(app *fiber.App)
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) NewUserRouter(app *fiber.App) {
	user := app.Group("/users")
	user.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
			Code:    fiber.StatusOK,
			Status:  true,
			Message: "ok",
		})
	})

	user.Post("/", controller.CreateUser)
	user.Get("/:id", controller.GetUserById)
	user.Get("/", controller.GetAllUsers)
	user.Put("/:id", controller.UpdateUser)

}

func (controller *UserControllerImpl) CreateUser(ctx *fiber.Ctx) error {
	var user web.UserRequest
	key := ctx.Query("key")

	if key == "" {
		return exception.ErrorHandler(ctx, exception.ErrorKey("API key is missing"))
	}
	form, err := ctx.MultipartForm()
	if err != nil {
		return exception.ErrorBadRequest("Please provide photos fields")
	}
	if err := ctx.BodyParser(&user); err != nil {
		return exception.ErrorHandler(ctx, err)
	}
	log.Println(user)

	response, err := controller.UserService.CreateUser(ctx, user, form)
	if err != nil {
		return exception.ErrorHandler(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "success",
		Data:    response,
	})
}

func (controller *UserControllerImpl) GetUserById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	key := ctx.Query("key")

	if key == "" {
		return exception.ErrorHandler(ctx, exception.ErrorKey("API key is missing"))
	}
	response, err := controller.UserService.FindUserById(ctx, id)

	if err != nil {
		return exception.ErrorHandler(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "success",
		Data:    response,
	})
}

func (controller *UserControllerImpl) GetAllUsers(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	page := ctx.Query("page")
	limit := ctx.Query("limit")

	searchName := ctx.Query("search-name")
	searchEmail := ctx.Query("search-email")

	response, err := controller.UserService.FindAllUser(ctx, search, page, limit, searchName, searchEmail)
	if err != nil {
		return exception.ErrorHandler(ctx, err)
	}

	if page != "" {
		pageInt, _ := strconv.Atoi(page)
		return ctx.Status(fiber.StatusOK).JSON(web.WebResponseWithPage{
			Code:    fiber.StatusOK,
			Status:  true,
			Page:    pageInt,
			Count:   len(response),
			Message: "success",
			Data:    response,
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
			Code:    fiber.StatusOK,
			Status:  true,
			Message: "success",
			Data:    response,
		})
	}
}

func (controller *UserControllerImpl) UpdateUser(ctx *fiber.Ctx) error {
	var user web.UserRequest
	key := ctx.Query("key")
	id := ctx.Params("id")
	if key == "" {
		return exception.ErrorHandler(ctx, exception.ErrorKey("API key is missing"))
	}
	form, err := ctx.MultipartForm()
	if err != nil {
		return exception.ErrorBadRequest("Please provide photos fields")
	}
	if err := ctx.BodyParser(&user); err != nil {
		return exception.ErrorHandler(ctx, err)
	}
	log.Println("sini")
	response, err := controller.UserService.UpdateUser(ctx, user, form, id)
	if err != nil {
		return exception.ErrorHandler(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "success",
		Data:    response,
	})
}
