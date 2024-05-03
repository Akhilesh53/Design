package dtos

// CreateUserDto struct
type CreateUserDto struct {
	name  string
	email string
}

// NewCreateUserDto function
func NewCreateUserDto(name string, email string) *CreateUserDto {
	return &CreateUserDto{
		name:  name,
		email: email,
	}
}

// GetName function
func (cud *CreateUserDto) GetName() string {
	return cud.name
}

// GetEmail function
func (cud *CreateUserDto) GetEmail() string {
	return cud.email
}

// SetName function
func (cud *CreateUserDto) SetName(name string) *CreateUserDto {
	cud.name = name
	return cud
}

// SetEmail function
func (cud *CreateUserDto) SetEmail(email string) *CreateUserDto {
	cud.email = email
	return cud
}

// GetUserDto struct
type GetUserDto struct {
	userid uint16
}

// NewGetUserDto function
func NewGetUserDto(userid uint16) *GetUserDto {
	return &GetUserDto{
		userid: userid,
	}
}

// GetUserID function
func (gud *GetUserDto) GetUserID() uint16 {
	return gud.userid
}

// SetUserID function
func (gud *GetUserDto) SetUserID(userid uint16) *GetUserDto {
	gud.userid = userid
	return gud
}
