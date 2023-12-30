package exception

import "github.com/gofiber/fiber/v2"

func ErrorUnauthorize(message string) error {
	return fiber.NewError(fiber.StatusUnauthorized, message)
}

func ErrorNotFound(message string) error {
	return fiber.NewError(fiber.StatusNotFound, message)
}

func ErrorBadRequest(message string) error {
	return fiber.NewError(fiber.StatusBadRequest, message)
}

func ErrorInternalServer(message string) error {
	return fiber.NewError(fiber.StatusInternalServerError, message)
}

func ErrorKey(message string) error {
	return fiber.NewError(403, message)
}
