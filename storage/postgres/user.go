package postgres

import (
	"database/sql"
	"errors"
	"log"

	t "github.com/Project_Restaurant/Auth-Service/api/token"
	"github.com/Project_Restaurant/Auth-Service/models"
	"github.com/google/uuid"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) Register(user models.UserRegister) (*models.LoginRes, error) {
	var us models.LoginRes
	id := uuid.NewString()
	hashPassword, err := t.HashPassword(user.Password)
	if err != nil {
		panic(err)
	}

	err = u.db.QueryRow(`
		INSERT INTO
			users(
				id, 
				username, 
				password, 
				email
		)VALUES (
			 	$1, 
				$2, 
				$3, 
				$4
			) 
		RETURNING 
			id,
			username`, id, user.Name, hashPassword, user.Email).
		Scan(&us.ID, &us.Name)
	if err != nil {
		return nil, err
	}
	return &us, nil
}

func (u *UserRepo) Login(user models.UserLogin) (*models.LoginRes, error) {
	res := models.LoginRes{}

	var hashedPassword string
	log.Println("Hello")
	err := u.db.QueryRow(`
		SELECT 
			id, 
			username, 
			password
		FROM
			users 
		WHERE 
			username = $1 
		AND
			deleted_at = 0`, user.Name).Scan(&res.ID, &res.Name, &hashedPassword)
	if err != nil {
		return nil, err
	}

	if !t.CheckPasswordHash(user.Password, hashedPassword) {
		return nil, errors.New("wrong password")
	}
	return &res, nil
}

func (u *UserRepo) GetByUsername(username string) (*models.User, error) {
	user := models.User{}
	err := u.db.QueryRow(`
		SELECT 
			id, 
			username, 
			email, 
			created_at, 
			updated_at, 
			deleted_at 
		FROM
			users 
		WHERE 
			username = $1
		AND 
			deleted_at = 0`, username).
		Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
