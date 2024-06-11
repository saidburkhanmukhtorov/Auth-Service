package postgres

import (
	"database/sql"

	"github.com/Project_Restaurant/Auth-Service/models"
	t "github.com/Project_Restaurant/Auth-Service/token"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

// Create User return id, username
// login username, password  return id username
// get by username return User struct

func (u *UserRepo) Register(user models.UserRegister) (models.LoginRes, error) {
	var us models.LoginRes
	hashPassword, err := t.HashPassword(user.Password)
	if err != nil {
		panic(err)
	}
	err = u.db.QueryRow("insert into users(username, password, email) values ($1, $2, $3) returning id, username", user.Name, hashPassword, user.Email).
		Scan(&us.ID, &us.Name)
	if err != nil {
		return us, err
	}
	return us, nil
}

func (u *UserRepo) Login(user models.UserLogin) (models.LoginRes, error) {
	res := models.LoginRes{}

	hashPassword, err := t.HashPassword(user.Password)
	if err != nil {
		panic(err)
	}

	err = u.db.QueryRow("select id, username from users where username = $1 and password = $2", user.Name, hashPassword).Scan(&res.ID, &res.Name)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (u *UserRepo) GetByUsername(username string) (models.User, error) {
	user := models.User{}
	err := u.db.QueryRow("select id, username, email, created_at, updated_at, deleted_at from users where username = $1", username).
	Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}
