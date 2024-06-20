package postgres

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Project_Restaurant/Auth-Service/models"
	"github.com/Project_Restaurant/Auth-Service/storage/postgres"
)

func TestUserRepo_Register(t *testing.T) {
	db, err := postgres.ConnectDb()
	require.NoError(t, err)
	defer db.Close()

	userRepo := postgres.NewUserRepo(db)

	user := models.UserRegister{
		Name:     uuid.NewString(),
		Email:    uuid.NewString(),
		Password: uuid.NewString(),
	}

	loginRes, err := userRepo.Register(user)
	require.NoError(t, err)
	assert.NotEmpty(t, loginRes.ID)
	assert.Equal(t, loginRes.Name, user.Name)

	// Check if user is registered in the database
	var registeredUser models.User
	err = db.QueryRow(`SELECT id, username, email FROM users WHERE username = $1`, user.Name).
		Scan(&registeredUser.ID, &registeredUser.Name, &registeredUser.Email)
	require.NoError(t, err)
	assert.Equal(t, registeredUser.ID, loginRes.ID)
	assert.Equal(t, registeredUser.Name, user.Name)
	assert.Equal(t, registeredUser.Email, user.Email)
}

func TestUserRepo_Login(t *testing.T) {
	db, err := postgres.ConnectDb()
	require.NoError(t, err)
	defer db.Close()

	userRepo := postgres.NewUserRepo(db)

	// Register a user for testing
	user := models.UserRegister{
		Name:     uuid.NewString(),
		Email:    uuid.NewString(),
		Password: uuid.NewString(),
	}
	_, err = userRepo.Register(user)
	require.NoError(t, err)

	// Login with correct credentials
	loginRes, err := userRepo.Login(models.UserLogin{
		Name:     user.Name,
		Password: user.Password,
	})
	require.NoError(t, err)
	assert.NotEmpty(t, loginRes.ID)
	assert.Equal(t, loginRes.Name, user.Name)

	// Login with incorrect password
	_, err = userRepo.Login(models.UserLogin{
		Name:     user.Name,
		Password: uuid.NewString(),
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "wrong password")
}

func TestUserRepo_GetByUsername(t *testing.T) {
	db, err := postgres.ConnectDb()
	require.NoError(t, err)
	defer db.Close()

	userRepo := postgres.NewUserRepo(db)

	// Register a user for testing
	user := models.UserRegister{
		Name:     uuid.NewString(),
		Email:    uuid.NewString(),
		Password: uuid.NewString(),
	}
	_, err = userRepo.Register(user)
	require.NoError(t, err)

	// Get user by username
	fetchedUser, err := userRepo.GetByUsername(user.Name)
	require.NoError(t, err)
	assert.Equal(t, fetchedUser.Name, user.Name)
	assert.Equal(t, fetchedUser.Email, user.Email)
}
