package repository

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/dhimweray222/users/exception"
	"github.com/dhimweray222/users/model/domain"
	"github.com/jackc/pgx/v5"
)

func (repository *UserRepositoryImpl) CreateUser(ctx context.Context, db pgx.Tx, newUser domain.User) error {
	query := fmt.Sprintf(`
	INSERT INTO users (
		id,
		name,
		address,
		email,
		password,
		photos,
		creditcard_type,
		creditcard_number,
		creditcard_name,
		creditcard_expired,
		creditcard_cvv)
	VALUES ($1,$2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`)

	if _, err := db.Prepare(context.Background(), "create_user", query); err != nil {
		return exception.ErrorInternalServer("Something went wrong. Please try again later.")
	}

	if _, err := db.Exec(context.Background(), "create_user",
		newUser.ID,
		newUser.Name,
		newUser.Address,
		newUser.Email,
		newUser.Password,
		newUser.Photos,
		newUser.CreditType,
		newUser.CCNumber,
		newUser.CCName,
		newUser.CCExpired,
		newUser.CCV,
	); err != nil {
		// log.Println("query:", err)
		return exception.ErrorInternalServer("Something went wrong. Please try again later.")
	}

	return nil
}

func (repository *UserRepositoryImpl) FindUserById(ctx context.Context, db pgx.Tx, id string) (domain.User, error) {
	queryStr := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, "users")

	user, err := db.Query(context.Background(), queryStr, id)

	if err != nil {
		log.Println(err)
		return domain.User{}, err
	}

	defer user.Close()

	data, err := pgx.CollectOneRow(user, pgx.RowToStructByPos[domain.User])

	if err != nil {
		log.Println(err)
		return domain.User{}, err
	}

	return data, nil
}

func (repository *UserRepositoryImpl) FindAllUser(ctx context.Context, db pgx.Tx, q, page, limit, searchName, searchEmail string) ([]domain.User, error) {
	queryStr := fmt.Sprintf(`SELECT *  FROM "users"`)

	if q != "" {
		queryStr += fmt.Sprintf(" AND (LOWER(name) LIKE LOWER('%%%s%%') OR LOWER(email) LIKE LOWER('%%%s%%'))", q, q)
	}

	queryStr += ` ORDER BY name ASC`
	queryStr = QueryPage(page, limit, queryStr)
	user, err := db.Query(context.Background(), queryStr)
	if err != nil {
		log.Println(err)
		return []domain.User{}, err
	}

	defer user.Close()
	data, err := pgx.CollectRows(user, pgx.RowToStructByPos[domain.User])
	if err != nil {
		log.Println(data, err)
		return []domain.User{}, err
	}

	return data, nil
}

func (repository *UserRepositoryImpl) UpdateUser(ctx context.Context, db pgx.Tx, updatedUser domain.User) error {
	query := `
		UPDATE users
		SET
			name = $1,
			address = $2,
			email = $3,
			password = $4,
			photos = $5,
			creditcard_type = $6,
			creditcard_number = $7,
			creditcard_name = $8,
			creditcard_expired = $9,
			creditcard_cvv = $10
		WHERE id = $11`

	if _, err := db.Exec(ctx, query,
		updatedUser.Name,
		updatedUser.Address,
		updatedUser.Email,
		updatedUser.Password,
		updatedUser.Photos,
		updatedUser.CreditType,
		updatedUser.CCNumber,
		updatedUser.CCName,
		updatedUser.CCExpired,
		updatedUser.CCV,
		updatedUser.ID,
	); err != nil {
		// Handle the error, log, and return an appropriate response
		log.Println(err)
		return err
		// return exception.ErrorInternalServer("Something went wrong. Please try again later.")
	}

	return nil
}

// Query page limit
func QueryPage(page, limit, queryStr string) string {
	defaultLimit, _ := strconv.Atoi("10")
	if page != "" {
		pageInt, _ := strconv.Atoi(page)
		if limit != "" {
			limitInt, _ := strconv.Atoi(limit)
			queryStr += fmt.Sprintf(" OFFSET %d LIMIT %d", (pageInt-1)*limitInt, limitInt)
		} else {
			queryStr += fmt.Sprintf(" OFFSET  %d LIMIT %d", (pageInt-1)*defaultLimit, defaultLimit)
		}
	} else {
		if limit != "" {
			limitInt, _ := strconv.Atoi(limit)
			queryStr += fmt.Sprintf(" LIMIT %d", limitInt)
		}
	}

	return queryStr
}
