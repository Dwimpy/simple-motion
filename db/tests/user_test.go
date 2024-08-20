package tests

import (
	"context"
	db "github.com/Dwimpy/simple-motion/db/sqlc"
	"github.com/Dwimpy/simple-motion/utils"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func createRandomUser(t *testing.T) db.User {

	hashedPassword, err := utils.GenerateHashPassword(utils.RandomString(8))

	if err != nil {
		log.Fatal("Unable to generate hashed password")
	}

	arg := db.CreateUserParams{
		Email:          utils.RandomEmail(),
		HashedPassword: hashedPassword,
	}

	user, err := testQuery.CreateUser(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, user)

	assert.Equal(t, user.Email, arg.Email)
	assert.Equal(t, user.HashedPassword, arg.HashedPassword)
	assert.WithinDuration(t, user.CreatedAt, time.Now(), 2*time.Second)
	assert.WithinDuration(t, user.UpdatedAt, time.Now(), 2*time.Second)

	return user
}

func TestCreateUser(t *testing.T) {

	user1 := createRandomUser(t)
	user2, err := testQuery.CreateUser(context.Background(), db.CreateUserParams{
		Email:          user1.Email,
		HashedPassword: user1.HashedPassword,
	})
	assert.Error(t, err)
	assert.Empty(t, user2)

}
