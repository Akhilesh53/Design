package dtos

type AddGroupRequestDto struct {
	UserId      string `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewAddGroupRequestDto(userId, name, description string) *AddGroupRequestDto {
	return &AddGroupRequestDto{
		UserId:      userId,
		Name:        name,
		Description: description,
	}
}

func NewDefaultGroupRequestDto() *AddGroupRequestDto {
	return &AddGroupRequestDto{}
}

func (agr *AddGroupRequestDto) GetUserId() string {
	return agr.UserId
}

func (agr *AddGroupRequestDto) GetName() string {
	return agr.Name
}

func (agr *AddGroupRequestDto) GetDescription() string {
	return agr.Description
}

func (agr *AddGroupRequestDto) SetUserId(userId string) *AddGroupRequestDto {
	agr.UserId = userId
	return agr
}

func (agr *AddGroupRequestDto) SetName(name string) *AddGroupRequestDto {
	agr.Name = name
	return agr
}

func (agr *AddGroupRequestDto) SetDescription(description string) *AddGroupRequestDto {
	agr.Description = description
	return agr
}
