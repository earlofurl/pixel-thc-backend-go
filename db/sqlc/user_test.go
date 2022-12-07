package db

import (
	"context"
	"github.com/gobuffalo/nulls"
	"pixel-thc-backend-go/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		HashedPassword: hashedPassword,
		Username:       util.RandomString(10),
		Email:          util.RandomValidEmail(),
		FirstName:      util.RandomString(10),
		LastName:       util.RandomString(10),
		Phone:          util.RandomString(10),
		Role:           "user",
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.Phone, user.Phone)
	require.Equal(t, arg.Role, user.Role)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	// TODO: don't return the password hash
	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.FirstName, user2.FirstName)
	require.Equal(t, user1.LastName, user2.LastName)
	require.Equal(t, user1.Phone, user2.Phone)
	require.Equal(t, user1.Role, user2.Role)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateUserOnlyFirstName(t *testing.T) {
	oldUser := createRandomUser(t)

	newFirstName := util.RandomOwner()
	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		Username:  oldUser.Username,
		FirstName: nulls.NewString(newFirstName),
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.FirstName, updatedUser.FirstName)
	require.Equal(t, newFirstName, updatedUser.FirstName)
	require.Equal(t, oldUser.LastName, updatedUser.LastName)
	require.Equal(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, oldUser.HashedPassword, updatedUser.HashedPassword)
	require.Equal(t, oldUser.Phone, updatedUser.Phone)
	require.Equal(t, oldUser.Role, updatedUser.Role)
}

func TestUpdateUserOnlyLastName(t *testing.T) {
	oldUser := createRandomUser(t)

	newLastName := util.RandomOwner()
	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		Username: oldUser.Username,
		LastName: nulls.NewString(newLastName),
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.LastName, updatedUser.LastName)
	require.Equal(t, newLastName, updatedUser.LastName)
	require.Equal(t, oldUser.FirstName, updatedUser.FirstName)
	require.Equal(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, oldUser.HashedPassword, updatedUser.HashedPassword)
	require.Equal(t, oldUser.Phone, updatedUser.Phone)
	require.Equal(t, oldUser.Role, updatedUser.Role)
}

func TestUpdateUserOnlyEmail(t *testing.T) {
	oldUser := createRandomUser(t)

	newEmail := util.RandomValidEmail()
	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		Username: oldUser.Username,
		Email:    nulls.NewString(newEmail),
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, newEmail, updatedUser.Email)
	require.Equal(t, oldUser.FirstName, updatedUser.FirstName)
	require.Equal(t, oldUser.LastName, updatedUser.LastName)
	require.Equal(t, oldUser.HashedPassword, updatedUser.HashedPassword)
	require.Equal(t, oldUser.Phone, updatedUser.Phone)
	require.Equal(t, oldUser.Role, updatedUser.Role)
}

func TestUpdateUserOnlyPassword(t *testing.T) {
	oldUser := createRandomUser(t)

	newPassword := util.RandomString(6)
	newHashedPassword, err := util.HashPassword(newPassword)
	require.NoError(t, err)

	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		Username:       oldUser.Username,
		HashedPassword: nulls.NewString(newHashedPassword),
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.HashedPassword, updatedUser.HashedPassword)
	require.Equal(t, newHashedPassword, updatedUser.HashedPassword)
	require.Equal(t, oldUser.FirstName, updatedUser.FirstName)
	require.Equal(t, oldUser.LastName, updatedUser.LastName)
	require.Equal(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, oldUser.Phone, updatedUser.Phone)
	require.Equal(t, oldUser.Role, updatedUser.Role)
}

func TestUpdateUserOnlyPhone(t *testing.T) {
	oldUser := createRandomUser(t)

	newPhone := util.RandomPhone()
	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		Username: oldUser.Username,
		Phone:    nulls.NewString(newPhone),
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.Phone, updatedUser.Phone)
	require.Equal(t, newPhone, updatedUser.Phone)
	require.Equal(t, oldUser.FirstName, updatedUser.FirstName)
	require.Equal(t, oldUser.LastName, updatedUser.LastName)
	require.Equal(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, oldUser.HashedPassword, updatedUser.HashedPassword)
	require.Equal(t, oldUser.Role, updatedUser.Role)
}

func TestUpdateUserOnlyRole(t *testing.T) {
	oldUser := createRandomUser(t)

	newRole := util.RandomRole()
	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		Username: oldUser.Username,
		Role:     nulls.NewString(newRole),
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.Role, updatedUser.Role)
	require.Equal(t, newRole, updatedUser.Role)
	require.Equal(t, oldUser.FirstName, updatedUser.FirstName)
	require.Equal(t, oldUser.LastName, updatedUser.LastName)
	require.Equal(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, oldUser.HashedPassword, updatedUser.HashedPassword)
	require.Equal(t, oldUser.Phone, updatedUser.Phone)
}

func TestUpdateUserAllFields(t *testing.T) {
	oldUser := createRandomUser(t)

	newFirstName := util.RandomOwner()
	newLastName := util.RandomOwner()
	newEmail := util.RandomValidEmail()
	newPassword := util.RandomString(6)
	newHashedPassword, err := util.HashPassword(newPassword)
	newPhone := util.RandomPhone()
	newRole := util.RandomRole()
	require.NoError(t, err)

	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		Username:       oldUser.Username,
		FirstName:      nulls.NewString(newFirstName),
		LastName:       nulls.NewString(newLastName),
		Email:          nulls.NewString(newEmail),
		HashedPassword: nulls.NewString(newHashedPassword),
		Phone:          nulls.NewString(newPhone),
		Role:           nulls.NewString(newRole),
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.HashedPassword, updatedUser.HashedPassword)
	require.Equal(t, newHashedPassword, updatedUser.HashedPassword)
	require.NotEqual(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, newEmail, updatedUser.Email)
	require.NotEqual(t, oldUser.FirstName, updatedUser.FirstName)
	require.Equal(t, newFirstName, updatedUser.FirstName)
	require.NotEqual(t, oldUser.LastName, updatedUser.LastName)
	require.Equal(t, newLastName, updatedUser.LastName)
	require.NotEqual(t, oldUser.Phone, updatedUser.Phone)
	require.Equal(t, newPhone, updatedUser.Phone)
	require.NotEqual(t, oldUser.Role, updatedUser.Role)
	require.Equal(t, newRole, updatedUser.Role)
}
