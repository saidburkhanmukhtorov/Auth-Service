package postgres

import (
	"database/sql"

	"github.com/Project_Restaurant/Auth-Service/models"
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
	err := u.db.QueryRow("insert into users (username, password, email) values ($1, $2, &3)", user.Name, user.Password, user.Email).Scan(&us.ID, &us.Name)
	if err != nil {
		return us, err
	}
	return us, nil
}

func (u *UserRepo) Login(user models.UserLogin) (models.LoginRes, error) {
	res := models.LoginRes{}
	err := u.db.QueryRow("select id, username from users where username = $1 and password = $2", user.Name, user.Password).Scan(&res.ID, &res.Name)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (u *UserRepo) GetByUsername(username string) (models.User, error) {
	user := models.User{}
	err := u.db.QueryRow("select * from users where username = $1", username).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}
