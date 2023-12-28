package db

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// Util test functions
func createUser(t *testing.T) CreateUserRow {
	arg := CreateUserParams{
		Username:  "foo",
		Email:     "foo@mail.com",
		Password:  "asdf",
		FirstName: "John",
		LastName:  "Smith",
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)

	_, err = uuid.Parse(user.ID.String())
	require.True(t, err == nil)

	return user
}

func getUserByID(t *testing.T, userID uuid.UUID) GetUserByIDRow {
	user, err := testQueries.GetUserByID(context.Background(), userID)
	require.NoError(t, err)
	return user
}

func doesUserExist(t *testing.T, userID uuid.UUID) bool {
	_, err := testQueries.GetUserByID(context.Background(), userID)
	return err != nil
}

func deleteUser(t *testing.T, userID uuid.UUID) {
	err := testQueries.DeleteUserByID(context.Background(), userID)
	require.NoError(t, err)
}

// Tests
func TestCreateUser(t *testing.T) {
	user := createUser(t)
	deleteUser(t, user.ID)
}

func TestDeleteUserByID(t *testing.T) {
	user := createUser(t)
	deleteUser(t, user.ID)
	require.True(t, doesUserExist(t, user.ID))
}

func TestGetUserByID(t *testing.T) {
	user1 := createUser(t)
	user2 := getUserByID(t, user1.ID)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.FirstName, user2.FirstName)
	require.Equal(t, user1.LastName, user2.LastName)

	deleteUser(t, user2.ID)
}

func TestGetUserByEmail(t *testing.T) {
	arg := CreateUserParams{
		Username:  "foo",
		Email:     "foo@mail.com",
		Password:  "asdf",
		FirstName: "John",
		LastName:  "Smith",
	}

	user1, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	user2, err := testQueries.GetUserByEmail(context.Background(), user1.Email)
	require.NoError(t, err)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, arg.Password, user2.Password)

	deleteUser(t, user2.ID)
}
