package repository

import "errors"

type mockUserRepo struct {
	users []User
}

func NewMockUserRepo() mockUserRepo {
	users := []User{
		{
			ID:           1,
			Email:        "test@test.com",
			Password:     "1234",
			Username:     "test",
			UserFullname: "Test User",
			UserBirthday: "2000-01-01",
			UserGender:   1,
			CreateAt:     "2023-01-01",
			UpdateAt:     "",
			DeleteAt:     "",
			IsDelete:     "0",
			UserProfile:  nil,
		},
		{
			ID:           2,
			Email:        "test2@test.com",
			Password:     "1234",
			Username:     "test2",
			UserFullname: "Test User 2",
			UserBirthday: "2000-01-01",
			UserGender:   1,
			CreateAt:     "2023-01-01",
			UpdateAt:     "2023-01-01",
			DeleteAt:     "2023-02-01",
			IsDelete:     "1",
			UserProfile:  nil,
		},
	}

	return mockUserRepo{users: users}
}

func (r mockUserRepo) LoginUser(username string, password string) (*User, error) {
	for _, user := range r.users {
		if (user.Username == username || user.Email == username) && user.Password == password {
			return &user, nil
		}
	}
	return nil, errors.New("User not found")
}

func (r mockUserRepo) GetUser(id int) (*User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("User not found")
}

func (r *mockUserRepo) RegisterUser(user *User) (int, error) {
	for _, u := range r.users {
		if u.Username == user.Username || u.Email == user.Email {
			return 0, errors.New("User already exists")
		}
	}

	user.ID = len(r.users) + 1
	r.users = append(r.users, *user)
	return user.ID, nil
}

func (r mockUserRepo) UpdateUserInfo(user *User) (*User, error) {
	for i, u := range r.users {
		if u.ID == user.ID {
			r.users[i] = *user
			return user, nil
		}
	}
	return nil, errors.New("User not found")
}

func (r mockUserRepo) UpdateUserPassword(user *User) (*User, error) {
	for i, u := range r.users {
		if u.ID == user.ID {
			r.users[i].Password = user.Password
			return &r.users[i], nil
		}
	}
	return nil, errors.New("User not found")
}

func (r mockUserRepo) UpdateUserEmail(user User) (*User, error) {
	// Check if email already exists
	for _, u := range r.users {
		if u.Email == user.Email && u.ID != user.ID {
			return nil, errors.New("email already exists")
		}
	}
	// Check email format xxx@xxx.com
	if user.Email == "" {
		return nil, errors.New("email cannot be empty")
	}
	// Update email
	for i, u := range r.users {
		if u.ID == user.ID {
			r.users[i].Email = user.Email
			return &r.users[i], nil
		}
	}
	return nil, errors.New("User not found")
}

func (r mockUserRepo) UpdateUserUsername(user User) (*User, error) {
	// Check if username already exists
	for _, u := range r.users {
		if u.Username == user.Username && u.ID != user.ID {
			return nil, errors.New("username already exists")
		}
	}
	// Check if username is empty
	if user.Username == "" {
		return nil, errors.New("username cannot be empty")
	}
	// Update username
	for i, u := range r.users {
		if u.ID == user.ID {
			r.users[i].Username = user.Username
			return &r.users[i], nil
		}
	}
	return nil, errors.New("User not found")
}

func (r mockUserRepo) UpdateUserProfile(user User) (*User, error) {
	for i, u := range r.users {
		if u.ID == user.ID {
			r.users[i].UserProfile = user.UserProfile
			return &r.users[i], nil
		}
	}
	return nil, errors.New("User not found")
}

func (r mockUserRepo) DeleteUser(id int) (int, error) {
	for i, u := range r.users {
		if u.ID == id {
			r.users[i].IsDelete = "1"
			return 1, nil
		}
	}
	return 0, errors.New("User not found")
}
