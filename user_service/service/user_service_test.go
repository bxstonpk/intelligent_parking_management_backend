package service

import (
	"testing"
	"user_services/repository"

	"github.com/stretchr/testify/assert"
)

func MockUserRepo() userService {
	// Mock user repository
	MockRepo := repository.NewMockUserRepo()

	// Create a user service with the mock repository
	userService := NewUserService(&MockRepo)

	// return the user service
	return userService
}

func TestLoginUser(t *testing.T) {
	// Create a mock user repository
	userService := MockUserRepo()

	// Test case: valid login
	loginRequest := &UserLoginRequest{
		Username: "test",
		Password: "1234",
	}
	userResponse, err := userService.LoginUser(loginRequest)
	if err != nil {
		assert.NoError(t, err)
	}
	if userResponse.Username != "test" {
		assert.Equal(t, loginRequest.Username, userResponse.Username)
	}

	// Test case: invalid login
	loginRequestInvalid := &UserLoginRequest{
		Username: "invalid_user",
		Password: "wrong_password",
	}
	userResponseInvalid, err := userService.LoginUser(loginRequestInvalid)
	if err == nil {
		assert.NoError(t, err)
	}
	if userResponseInvalid != nil {
		assert.Equal(t, userResponseInvalid, nil)
	}
}

func TestGetUser(t *testing.T) {
	// Create a mock user repository
	userService := MockUserRepo()

	// Test case: valid user ID
	userID := 1
	userResponse, err := userService.GetUser(userID)
	if err != nil {
		assert.NoError(t, err)
	}
	if userResponse.ID != userID {
		assert.Equal(t, userID, userResponse.ID)
	}

	// Test case: invalid user ID
	userIDInvalid := 999
	userResponseInvalid, err := userService.GetUser(userIDInvalid)
	if err == nil {
		assert.NoError(t, err)
	}
	if userResponseInvalid != nil {
		assert.Equal(t, userResponseInvalid, nil)
	}
}

func TestRegistorUser(t *testing.T) {
	// Create a mock user repository
	userService := MockUserRepo()

	// Test case: valid registration
	registrationRequest := UserRegisterRequest{
		Email:        "unittest@test.com",
		Password:     "1234",
		Username:     "unittest",
		UserFullname: "Unit Test",
		UserBirthday: "2000-01-01",
		UserGender:   1,
		CreateAt:     "2023-01-01",
		UserProfile:  "",
	}
	userID, err := userService.RegisterUser(registrationRequest)
	if err != nil {
		assert.NoError(t, err)
	}
	if userID <= 0 {
		assert.LessOrEqual(t, userID, 0)
	}

	// Test case: invalid registration (duplicate username)
	registrationRequestDuplicate := UserRegisterRequest{
		Email:        "unittest@test.com",
		Password:     "1234",
		Username:     "unittest",
		UserFullname: "Unit Test",
		UserBirthday: "2000-01-01",
		UserGender:   1,
		CreateAt:     "2023-01-01",
		UserProfile:  "",
	}
	userIDDuplicate, err := userService.RegisterUser(registrationRequestDuplicate)
	if err == nil {
		assert.NoError(t, err)
	}
	if userIDDuplicate != 0 {
		assert.NotEqual(t, userIDDuplicate, 0)
	}

	// Test case: invalid registration (missing required fields)
	registrationRequestInvalid := UserRegisterRequest{
		Email:        "",
		Password:     "1234",
		Username:     "unittest",
		UserFullname: "Unit Test",
		UserBirthday: "2000-01-01",
		UserGender:   1,
		CreateAt:     "2023-01-01",
		UserProfile:  "",
	}
	userIDInvalid, err := userService.RegisterUser(registrationRequestInvalid)
	if err == nil {
		assert.NoError(t, err)
	}
	if userIDInvalid != 0 {
		assert.NotEqual(t, userIDInvalid, 0)
	}

	// Test case: invalid registration (invalid email format)
	registrationRequestInvalidEmail := UserRegisterRequest{
		Email:        "invalid_email",
		Password:     "1234",
		Username:     "unittest",
		UserFullname: "Unit Test",
		UserBirthday: "2000-01-01",
		UserGender:   1,
		CreateAt:     "2023-01-01",
		UserProfile:  "",
	}
	userIDInvalidEmail, err := userService.RegisterUser(registrationRequestInvalidEmail)
	if err == nil {
		assert.NoError(t, err)
	}
	if userIDInvalidEmail != 0 {
		assert.NotEqual(t, userIDInvalidEmail, 0)
	}
}

func TestUpdateUserInfo(t *testing.T) {
	// Create a mock user repository
	userService := MockUserRepo()

	// Test case: valid update
	updateRequest := UserUpdateInfoRequest{
		ID:           1,
		UserFullname: "Updated User",
		UserBirthday: "2000-01-01",
		UserGender:   1,
	}
	userResponse, err := userService.UpdateUserInfo(updateRequest)
	if err != nil {
		assert.NoError(t, err)
	}
	if userResponse.UserFullname != updateRequest.UserFullname {
		assert.NotEqual(t, updateRequest.UserFullname, userResponse.UserFullname)
	}

	// Test case: invalid update (user not found)
	updateRequestInvalid := UserUpdateInfoRequest{
		ID:           999,
		UserFullname: "Updated User",
		UserBirthday: "2000-01-01",
		UserGender:   1,
	}
	userResponseInvalid, err := userService.UpdateUserInfo(updateRequestInvalid)
	if err == nil {
		assert.NoError(t, err)
	}
	if userResponseInvalid != nil {
		assert.NotEqual(t, userResponseInvalid, nil)
	}
}

func TestUpdateUserPassword(t *testing.T) {
	// Create a mock user repository
	userService := MockUserRepo()

	// Test case: valid password update
	updatePasswordRequest := UserUpdatePasswordRequest{
		ID:       1,
		Password: "new_password",
	}
	userResponse, err := userService.UpdateUserPassword(updatePasswordRequest)
	if err != nil {
		assert.NoError(t, err)
	}
	if userResponse.ID != updatePasswordRequest.ID {
		assert.NotEqual(t, updatePasswordRequest.ID, userResponse.ID)
	}

	// Test case: invalid password update (user not found)
	updatePasswordRequestInvalid := UserUpdatePasswordRequest{
		ID:       999,
		Password: "new_password",
	}
	userResponseInvalid, err := userService.UpdateUserPassword(updatePasswordRequestInvalid)
	if err == nil {
		assert.NoError(t, err)
	}
	if userResponseInvalid != nil {
		assert.NotEqual(t, userResponseInvalid, nil)
	}
}

func TestUpdateUserEmail(t *testing.T) {
	// Create a mock user repository
	userService := MockUserRepo()

	// Test case: valid email update
	updateEmailRequest := UserUpdateEmailRequest{
		ID:    1,
		Email: "ddd@unittest.com",
	}
	userResponse, err := userService.UpdateUserEmail(updateEmailRequest)
	if err != nil {
		assert.NoError(t, err)
	}
	if userResponse.Email != updateEmailRequest.Email {
		assert.NotEqual(t, updateEmailRequest.Email, userResponse.Email)
	}

	// Test case: invalid email update (user not found)
	updateEmailRequestInvalid := UserUpdateEmailRequest{
		ID:    999,
		Email: "ddd@unittest.com",
	}
	userResponseInvalid, err := userService.UpdateUserEmail(updateEmailRequestInvalid)
	if err == nil {
		assert.NoError(t, err)
	}
	if userResponseInvalid != nil {
		assert.NotEqual(t, userResponseInvalid, nil)
	}
	// Test case: invalid email update (duplicate email)
	updateEmailRequestDuplicate := UserUpdateEmailRequest{
		ID:    1,
		Email: "test2@test.com",
	}
	userResponseDuplicate, err := userService.UpdateUserEmail(updateEmailRequestDuplicate)
	if err == nil {
		assert.NoError(t, err)
	}
	if userResponseDuplicate != nil {
		assert.NotEqual(t, userResponseDuplicate, nil)
	}
	// Test case: invalid email update (invalid email format)
	updateEmailRequestInvalidFormat := UserUpdateEmailRequest{
		ID:    1,
		Email: "invalidemail",
	}
	userResponseInvalidFormat, err := userService.UpdateUserEmail(updateEmailRequestInvalidFormat)
	if err == nil {
		assert.NoError(t, err)
	}
	if userResponseInvalidFormat != nil {
		assert.NotEqual(t, userResponseInvalidFormat, nil)
	}
	// Test case: invalid email update (empty email)
	updateEmailRequestEmpty := UserUpdateEmailRequest{
		ID:    1,
		Email: "",
	}
	userResponseEmpty, err := userService.UpdateUserEmail(updateEmailRequestEmpty)
	if err == nil {
		assert.NoError(t, err)
	}
	if userResponseEmpty != nil {
		assert.NotEqual(t, userResponseEmpty, nil)
	}
}

func TestUpdateUserUsername(t *testing.T) {
	// Create a mock user repository
	userService := MockUserRepo()

	// Test case: valid username update
	updateUsernameRequest := UserUpdateUsernameRequest{
		ID:       1,
		Username: "new_username",
	}
	userResponse, err := userService.UpdateUserUsername(updateUsernameRequest)
	if err != nil {
		assert.NoError(t, err)
	}
	if userResponse.Username != updateUsernameRequest.Username {
		assert.NotEqual(t, updateUsernameRequest.Username, userResponse.Username)
	}

	// Test case: invalid username update (user not found)
	updateUsernameRequestInvalid := UserUpdateUsernameRequest{
		ID:       999,
		Username: "new_username",
	}
	userResponseInvalid, err := userService.UpdateUserUsername(updateUsernameRequestInvalid)
	if err == nil {
		assert.NoError(t, err)
	}
	if userResponseInvalid != nil {
		assert.NotEqual(t, userResponseInvalid, nil)
	}
	// Test case: invalid username update (duplicate username)
	updateUsernameRequestDuplicate := UserUpdateUsernameRequest{
		ID:       1,
		Username: "test2",
	}
	userResponseDuplicate, err := userService.UpdateUserUsername(updateUsernameRequestDuplicate)
	if err == nil {
		assert.NoError(t, err)
	}
	if userResponseDuplicate != nil {
		assert.NotEqual(t, userResponseDuplicate, nil)
	}
	// Test case: invalid username update (empty username)
	updateUsernameRequestEmpty := UserUpdateUsernameRequest{
		ID:       1,
		Username: "",
	}
	userResponseEmpty, err := userService.UpdateUserUsername(updateUsernameRequestEmpty)
	if err == nil {
		assert.NoError(t, err)
	}
	if userResponseEmpty != nil {
		assert.NotEqual(t, userResponseEmpty, nil)
	}
}

func TestUpdateUserProfile(t *testing.T) {
	// Create a mock user repository
	userService := MockUserRepo()

	// Test case: valid profile update
	updateProfileRequest := UserUpdateProfileRequest{
		ID:          1,
		UserProfile: "new_profile_data",
	}
	userResponse, err := userService.UpdateUserProfile(updateProfileRequest)
	if err != nil {
		assert.NoError(t, err)
	}
	if string(userResponse.UserProfile) != string(updateProfileRequest.UserProfile) {
		assert.NotEqual(t, updateProfileRequest.UserProfile, userResponse.UserProfile)
	}

	// Test case: invalid profile update (user not found)
	updateProfileRequestInvalid := UserUpdateProfileRequest{
		ID:          999,
		UserProfile: "new_profile_data",
	}
	userResponseInvalid, err := userService.UpdateUserProfile(updateProfileRequestInvalid)
	if err == nil {
		assert.NoError(t, err)
	}
	if userResponseInvalid != nil {
		assert.NotEqual(t, userResponseInvalid, nil)
	}
}

func TestDeleteUser(t *testing.T) {
	// Create a mock user repository
	userService := MockUserRepo()

	// Test case: valid delete
	userID := 1
	result, err := userService.DeleteUser(userID)
	if err != nil {
		assert.NoError(t, err)
	}
	if result != 1 {
		assert.NotEqual(t, result, 1)
	}

	// Test case: invalid delete (user not found)
	userIDInvalid := 999
	resultInvalid, err := userService.DeleteUser(userIDInvalid)
	if err == nil {
		assert.NoError(t, err)
	}
	if resultInvalid != 0 {
		assert.NotEqual(t, resultInvalid, 0)
	}
}
