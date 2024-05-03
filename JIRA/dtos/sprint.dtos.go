package dtos

import "pattern/JIRA/entities"

// create sprint dto
type CreateSprintDto struct {
	name      string
	startdate string
	enddate   string
	createdBy entities.User
}

// get sprint dto
type GetSprintDto struct {
	sprintId uint16
}

// new create sprint dto
func NewCreateSprintDto(name string, startdate string, enddate string, createdBy entities.User) *CreateSprintDto {
	return &CreateSprintDto{
		name:      name,
		startdate: startdate,
		enddate:   enddate,
		createdBy: createdBy,
	}
}

// new get sprint dto
func NewGetSprintDto(sprintId uint16) *GetSprintDto {
	return &GetSprintDto{
		sprintId: sprintId,
	}
}

// get name
func (csd *CreateSprintDto) GetName() string {
	return csd.name
}

// get start date
func (csd *CreateSprintDto) GetStartDate() string {
	return csd.startdate
}

// get end date
func (csd *CreateSprintDto) GetEndDate() string {
	return csd.enddate
}

// get created by
func (csd *CreateSprintDto) GetCreatedBy() entities.User {
	return csd.createdBy
}

// get sprint id
func (gsd *GetSprintDto) GetSprintID() uint16 {
	return gsd.sprintId
}

// set name
func (csd *CreateSprintDto) SetName(name string) *CreateSprintDto {
	csd.name = name
	return csd
}

// set start date
func (csd *CreateSprintDto) SetStartDate(startdate string) *CreateSprintDto {
	csd.startdate = startdate
	return csd
}

// set end date
func (csd *CreateSprintDto) SetEndDate(enddate string) *CreateSprintDto {
	csd.enddate = enddate
	return csd
}

// set created by
func (csd *CreateSprintDto) SetCreatedBy(createdBy entities.User) *CreateSprintDto {
	csd.createdBy = createdBy
	return csd
}

// set sprint id
func (gsd *GetSprintDto) SetSprintID(sprintId uint16) *GetSprintDto {
	gsd.sprintId = sprintId
	return gsd
}

// set sprint id
func (gsd *GetSprintDto) SetSprintId(sprintId uint16) *GetSprintDto {
	gsd.sprintId = sprintId
	return gsd
}
