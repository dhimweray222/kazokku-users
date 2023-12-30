package service

import (
	"io/ioutil"
	"log"
	"mime/multipart"

	"github.com/dhimweray222/users/exception"
	"github.com/dhimweray222/users/helper"
	"github.com/dhimweray222/users/model/domain"
	"github.com/dhimweray222/users/model/web"
	"github.com/dhimweray222/users/repository"
	"github.com/gofiber/fiber/v2"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

type UserService interface {
	CreateUser(ctx *fiber.Ctx, user web.UserRequest, form *multipart.Form) (web.UserRepsonse, error)
	FindUserById(ctx *fiber.Ctx, id string) (domain.User, error)
	FindAllUser(ctx *fiber.Ctx, search, page, limit, searchName, searchEmail string) ([]domain.User, error)
	UpdateUser(ctx *fiber.Ctx, user web.UserRequest, form *multipart.Form, id string) (domain.User, error)
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

func (service *UserServiceImpl) CreateUser(ctx *fiber.Ctx, user web.UserRequest, form *multipart.Form) (web.UserRepsonse, error) {

	if user.Name == "" {
		return web.UserRepsonse{}, exception.ErrorBadRequest("Please provide name fields")
	}
	if user.Address == "" {
		return web.UserRepsonse{}, exception.ErrorBadRequest("Please provide name address")
	}
	if user.Email == "" {
		return web.UserRepsonse{}, exception.ErrorBadRequest("Please provide email fields")
	}
	if user.Password == "" {
		return web.UserRepsonse{}, exception.ErrorBadRequest("Please provide Password fields")
	}
	if user.CreditType == "" {
		return web.UserRepsonse{}, exception.ErrorBadRequest("Please provide CreditCardType fields")
	}
	if user.CCNumber == "" {
		return web.UserRepsonse{}, exception.ErrorBadRequest("Please provide CreditCardNumber fields")
	}
	if user.CCName == "" {
		return web.UserRepsonse{}, exception.ErrorBadRequest("Please provide CreditCardName fields")
	}
	if user.CCExpired == "" {
		return web.UserRepsonse{}, exception.ErrorBadRequest("Please provide CreditCardExpired fields")
	}
	if user.CCV == "" {
		return web.UserRepsonse{}, exception.ErrorBadRequest("Please provide CreditCardCVV fields")
	}

	dataUser := domain.User{
		Name:       user.Name,
		Address:    user.Address,
		Email:      user.Email,
		Password:   user.Password,
		Photos:     user.Photos,
		CreditType: user.CreditType,
		CCNumber:   user.CCNumber,
		CCName:     user.CCName,
		CCExpired:  user.CCExpired,
		CCV:        user.CCV,
	}
	dataUser.GenerateIdKey()

	var images []string
	// Get the "photos" files from the form data
	photos := form.File["photos"]

	// Process each uploaded file
	for _, file := range photos {
		// Open the file
		src, err := file.Open()
		if err != nil {
			return web.UserRepsonse{}, err
		}
		defer src.Close()

		// Read the content of the file
		fileContent, err := ioutil.ReadAll(src)
		if err != nil {
			return web.UserRepsonse{}, err
		}

		imageURL, err := helper.UploadImage(fileContent, file.Filename, dataUser.ID.String())
		if err != nil {
			log.Println("upload", err)
			return web.UserRepsonse{}, exception.ErrorBadRequest("Could not upload image")
		}

		images = append(images, imageURL)

	}

	dataUser.Photos = images
	data, err := service.UserRepository.CreateUserImpl(ctx.Context(), dataUser)
	if err != nil {
		return web.UserRepsonse{}, err
	}

	return domain.ToUserResponse(data), nil
}

func (service *UserServiceImpl) FindUserById(ctx *fiber.Ctx, id string) (domain.User, error) {
	user, err := service.UserRepository.FindUserByIdTx(ctx.Context(), id)

	if err != nil || user.Email == "" {
		return domain.User{}, exception.ErrorNotFound("User not found.")
	}

	return user, nil
}

func (service *UserServiceImpl) FindAllUser(ctx *fiber.Ctx, search, page, limit, searchName, searchEmail string) ([]domain.User, error) {

	users, err := service.UserRepository.FindAllUserTx(ctx.Context(), search, page, limit, searchName, searchEmail)
	if err != nil {
		log.Println("service", err)
		return []domain.User{}, err
	}
	return users, nil
}

func (service *UserServiceImpl) UpdateUser(ctx *fiber.Ctx, user web.UserRequest, form *multipart.Form, id string) (domain.User, error) {

	if user.Name == "" {
		return domain.User{}, exception.ErrorBadRequest("Please provide name fields")
	}
	if user.Address == "" {
		return domain.User{}, exception.ErrorBadRequest("Please provide name address")
	}
	if user.Email == "" {
		return domain.User{}, exception.ErrorBadRequest("Please provide email fields")
	}
	if user.Password == "" {
		return domain.User{}, exception.ErrorBadRequest("Please provide Password fields")
	}
	if user.CreditType == "" {
		return domain.User{}, exception.ErrorBadRequest("Please provide CreditCardType fields")
	}
	if user.CCNumber == "" {
		return domain.User{}, exception.ErrorBadRequest("Please provide CreditCardNumber fields")
	}
	if user.CCName == "" {
		return domain.User{}, exception.ErrorBadRequest("Please provide CreditCardName fields")
	}
	if user.CCExpired == "" {
		return domain.User{}, exception.ErrorBadRequest("Please provide CreditCardExpired fields")
	}
	if user.CCV == "" {
		return domain.User{}, exception.ErrorBadRequest("Please provide CreditCardCVV fields")
	}
	data, err := service.UserRepository.FindUserByIdTx(ctx.Context(), id)
	if err != nil {
		return domain.User{}, err
	}
	dataUser := domain.User{
		ID:         data.ID,
		Name:       user.Name,
		Address:    user.Address,
		Email:      user.Email,
		Password:   user.Password,
		Photos:     user.Photos,
		CreditType: user.CreditType,
		CCNumber:   user.CCNumber,
		CCName:     user.CCName,
		CCExpired:  user.CCExpired,
		CCV:        user.CCV,
	}

	var images []string
	// Get the "photos" files from the form data
	photos := form.File["photos"]

	// Process each uploaded file
	for _, file := range photos {
		// Open the file
		src, err := file.Open()
		if err != nil {
			return domain.User{}, err
		}
		defer src.Close()

		// Read the content of the file
		fileContent, err := ioutil.ReadAll(src)
		if err != nil {
			return domain.User{}, err
		}

		imageURL, err := helper.UploadImage(fileContent, file.Filename, dataUser.ID.String())
		if err != nil {
			log.Println("upload", err)
			return domain.User{}, exception.ErrorBadRequest("Could not upload image")
		}

		images = append(images, imageURL)

	}

	dataUser.Photos = images
	err = service.UserRepository.UpdateUserImpl(ctx.Context(), dataUser)
	if err != nil {
		return domain.User{}, err
	}
	return dataUser, nil
}
