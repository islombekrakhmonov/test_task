package postgres

import (
	"context"
	"task/models"

	"github.com/jackc/pgx/v4/pgxpool"
)


type phoneRepo struct {
	db *pgxpool.Pool
}

func NewPhoneRepo(db *pgxpool.Pool) *phoneRepo {
	return &phoneRepo{
		db: db,
	}
}


func (p *phoneRepo) Create(ctx context.Context, req *models.CreatePhoneRequest) (resp *models.CreatePhoneResponse, err error){
	resp = &models.CreatePhoneResponse{}

	query := `INSERT INTO phone (user_id, phone_number, description, is_fax) VALUES ($1, $2, $3, $4) RETURNING  phone_number, description, is_fax`

	err = p.db.QueryRow(ctx, query, req.UserId, req.PhoneNumber, req.Description, req.IsFax).Scan(&resp.PhoneNumber, &resp.Description, &resp.IsFax)
	if err != nil {
		return nil, err
	}

	return resp, nil
}


func (p *phoneRepo) GetByPhone(ctx context.Context,  req *models.GetByPhoneRequest) (resp []*models.GetByPhoneResponse,err error) {

	query := `SELECT user_id, phone_number, description, is_fax FROM phone WHERE phone_number LIKE '%' || $1 || '%'`

	rows, err := p.db.Query(ctx, query, req.PhoneNumber)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		r := &models.GetByPhoneResponse{}
		err := rows.Scan(&r.UserId, &r.PhoneNumber, &r.Description, &r.IsFax)
		if err != nil {
			return nil, err
		}
		resp = append(resp, r)
	}

	return resp, nil
}

func (p *phoneRepo) Update(ctx context.Context, req *models.UpdatePhoneRequest) (resp *models.UpdatePhoneResponse, err error) {
	resp = &models.UpdatePhoneResponse{}

	query := `UPDATE phone SET phone_number = $2, description = $3, is_fax = $4, updated_at = NOW()
	WHERE phone_id = $1 RETURNING  phone_number, description, is_fax`

	err = p.db.QueryRow(ctx, query, req.PhoneId, req.PhoneNumber, req.Description, req.IsFax).Scan(&resp.PhoneNumber, &resp.Description, &resp.IsFax)  // Возвращаем обновленные данные из БД, чтобы не делать дополнительный запрос. Из-за этого использовал QueryRow вместо Exec
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *phoneRepo) Delete(ctx context.Context, pKey *models.PhonePKey) (error) {
	query := `DELETE FROM "phone" WHERE phone_id = $1`
	
	_, err := p.db.Exec(ctx, query ,pKey.PhoneId)
	if err != nil {
		return err
	}

	return  nil
}

// Проверка на дубликат номера телефона, возвращает true если номер телефона уже существует
func (p *phoneRepo) CheckDuplicatePhoneNumber(ctx context.Context, phoneNumber string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM phone WHERE phone_number = $1)`

	var exists bool
	err := p.db.QueryRow(ctx, query, phoneNumber).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}