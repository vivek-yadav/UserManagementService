package models

// This function is used for login api calls.
// this(User) is set with email and password before the call.
// It returns error with cause.
func (this User) FindLoginUserWithEmail() (User, error) {
	return User{}, nil
}

// This function is used for login api calls.
// this(User) is set with username and password before the call.
// It returns error with cause.
func (this User) FindLoginUserWithUsername() (User, error) {
	return User{}, nil
}

// This function is used to check for authorization of the user to access specific url
// if allowed it returns success else returns the reason of error
func (this User) IsAuth(AppToken string, accessLevel int8, url string) (User, error) {
	return User{}, nil
}

// This function is used to fetch detils of a user by Id
func (this User) GetById() (User, error) {
	this.Name = "Vivek"
	return this, nil
}

// This function is used to fetch all the users
func (this Users) GetList() (Users, error) {
	a1 := User{Name: "Vivek"}
	a2 := User{Name: "Anurag"}
	a3 := User{Name: "Lax"}
	this = Users([]User{a1, a2, a3})
	return this, nil
}
