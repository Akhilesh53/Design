package dtos

type GetUserRequestDto struct {
	UserId string `json:"user_id"`
}

func NewGetUserRequestDto(userId string) *GetUserRequestDto {
	return &GetUserRequestDto{
		UserId: userId,
	}
}

func NewDefaultGetUserRequestDto() *GetUserRequestDto {
	return &GetUserRequestDto{}
}

func (agr *GetUserRequestDto) GetUserId() string {
	return agr.UserId
}

func (agr *GetUserRequestDto) SetUserId(userId string) *GetUserRequestDto {
	agr.UserId = userId
	return agr
}
