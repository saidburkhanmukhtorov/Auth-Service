package postgres

import (
	"database/sql"
	"log"

	"auth/models"

	"github.com/google/uuid"
)

type UserRepo struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{Db: db}
}


func (us *UserRepo) CreateUser(req *models.UserReq) (*models.UserRes, error) {
	id := uuid.New().String()
	query := `
		insert into allusers(id,name,age,email,gender,password) values($1,$2,$3,$4,$5,$6) 
		returning id,name,age,email,gender,created_at,updated_at
	`
	resUser := models.UserRes{}

	err := us.Db.QueryRow(query, id, req.Name, req.Age, req.Email, req.Gender, req.Password).Scan(
			&resUser.ID, &resUser.Name, &resUser.Age, &resUser.Email, &resUser.Gender, &resUser.CreatedAt, &resUser.UpdatedAt,
		)
	if err != nil {
		log.Fatal("Error whith create newuser in database", err)
		return nil, err
	}
	return &resUser, nil
}

func (us *UserRepo) UpdateUserRepo(id string, req *models.UserReq) (*models.UserRes, error) {

	query := `
        update allusers set username=$1, age=$2, email=$3, gender=$4, password=$5, updated_at=now() where id=$6
        returning id, username, age, email, gender, password, created_at, updated_at
    `
	resUser := models.UserRes{}
	err := us.Db.QueryRow(query, req.Name, req.Age, req.Email, req.Gender, req.Password, id).Scan(
		&resUser.ID, &resUser.Name, &resUser.Age, &resUser.Email, &resUser.Gender, &resUser.CreatedAt, &resUser.UpdatedAt,
	)
	if err != nil {
		log.Println("Error updating user in database:", err)
		return nil, err
	}
	return &resUser, nil
}

func (us *UserRepo) GetUserRepo(id string) (*models.UserRes, error) {
	query := "select id, username, age, email, gender, created_at, updated_at from users where id=$1"
	resUser := models.UserRes{}
	err := us.Db.QueryRow(query, id).Scan(
		&resUser.ID, &resUser.Name, &resUser.Age, &resUser.Email, &resUser.Gender, &resUser.CreatedAt, &resUser.UpdatedAt,
	)
	if err != nil {
		log.Println("Error fetching user from database:", err)
		return nil, err
	}
	return &resUser, nil
}

func (us *UserRepo) GetAllUsersRepo() ([]models.UserRes, error) {
	query := "select id, username, age, email, gender, created_at, updated_at from users"
	rows, err := us.Db.Query(query)
	if err != nil {
		log.Println("Error fetching all users from database:", err)
		return nil, err
	}
	defer rows.Close()

	var users []models.UserRes
	for rows.Next() {
		var user models.UserRes
		err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Email, &user.Gender, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Println("Error scanning user row:", err)
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		log.Println("Error with user rows:", err)
		return nil, err
	}

	return users, nil
}


func (us *UserRepo) DeleteUserRepo(id string) error {
	query := "update users set deleted_at = extract(epoch from NOW()) where deleted_at = 0"
	_, err := us.Db.Exec(query, id)
	if err != nil {
		log.Println("Error deleting user from database:", err)
		return err
	}
	return nil
}