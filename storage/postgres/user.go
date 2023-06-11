package postgres

import (
	"context"
	"database/sql"
	"task/models"
	"task/security"

	"github.com/jackc/pgx/v4/pgxpool"
)


type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}


func (u *userRepo) Create(ctx context.Context, req *models.CreateUserRequest) (resp *models.CreateUserResponse, err error) {
	resp = &models.CreateUserResponse{}

	// Хешируем пароль
	hashedPassword, err := security.HashPassword(&req.Password)
	if err != nil {
		return nil, err
	}

	query := `
		INSERT INTO "user" (login, password, name, age) VALUES ($1, $2, $3, $4)
		RETURNING user_id, name, age`
// Тут я использовал QeuryRow, вместо Exec, чтобы получить данные из БД
	err = u.db.QueryRow(
		ctx, query, req.Login, hashedPassword, req.Name,req.Age,
		).Scan(
		&resp.UserId, &resp.Name, &resp.Age,
	)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (u *userRepo) GetByName(ctx context.Context, req *models.GetUserByNameRequest) (resp *models.GetUserByNameResponse, err error) {

	resp = &models.GetUserByNameResponse{}

	query := `
	SELECT user_id, name, age FROM "user" WHERE name = $1`

	row := u.db.QueryRow(ctx, query, req.Name)
	err = row.Scan(&resp.UserId, &resp.Name, &resp.Age)
	if err != nil {
		return nil, err
	}

	return resp, nil
}


func (u *userRepo) Login(ctx context.Context, login *models.AuthUserRequest) (resp *models.AuthUserResponse, err error) {
	resp = &models.AuthUserResponse{}


	query := `SELECT user_id, login, password FROM "user" WHERE login = $1`

	row := u.db.QueryRow(ctx, query, login.Login)
	err = row.Scan(&resp.UserId, &resp.Login, &resp.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return resp, nil
}


