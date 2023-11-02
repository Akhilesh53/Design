package dtos

type RegisterUserRequestDto struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func NewRegisterUserRequestDto(name, phone, password string) *RegisterUserRequestDto {
	return &RegisterUserRequestDto{
		Name:     name,
		Phone:    phone,
		Password: password,
	}
}

func (rurd *RegisterUserRequestDto) GetName() string {
	return rurd.Name
}

func (rurd *RegisterUserRequestDto) GetPhone() string {
	return rurd.Phone
}

func (rurd *RegisterUserRequestDto) GetPassword() string {
	return rurd.Password
}

func (rurd *RegisterUserRequestDto) SetName(name string) *RegisterUserRequestDto {
	rurd.Name = name
	return rurd
}

func (rurd *RegisterUserRequestDto) SetPhone(phone string) *RegisterUserRequestDto {
	rurd.Phone = phone
	return rurd
}

func (rurd *RegisterUserRequestDto) SetPassword(password string) *RegisterUserRequestDto {
	rurd.Password = password
	return rurd
}
