package entities

import "time"

type IGroup interface {
	GetId() int
	GetName() string
	GetParticipants() []IUser
	GetExpenses() []IExpense
	GetDescription() string
	GetCreatedBy() IUser
	GetCreatedAt() time.Time
	GetLastModified() time.Time
	SetId(int) IGroup
	SetName(string) IGroup
	SetParticipants([]IUser) IGroup
	AddParticipant(user IUser) IGroup
	SetExpenses([]IExpense) IGroup
	SetDescription(string) IGroup
	SetCreatedBy(IUser) IGroup
	SetCreatedAt(time.Time) IGroup
	SetLastModified(time.Time) IGroup
}

type group struct {
	id           int
	name         string
	description  string
	participants []IUser
	admins       []IUser
	expenses     []IExpense
	createdBy    IUser
	createdAt    time.Time
	lastModified time.Time
}

func NewDefaultGroup() IGroup {
	return &group{
		participants: make([]IUser, 0),
		admins:       make([]IUser, 0),
		expenses:     make([]IExpense, 0),
		createdAt:    time.Now(),
	}
}

func NewGroup(id int, name string, description string, createdBy IUser) IGroup {
	return &group{
		id:           id,
		name:         name,
		description:  description,
		participants: make([]IUser, 0),
		admins:       make([]IUser, 0),
		expenses:     make([]IExpense, 0),
		createdBy:    createdBy,
		createdAt:    time.Now(),
	}
}

func (g *group) GetId() int {
	return g.id
}

func (g *group) GetName() string {
	return g.name
}

func (g *group) GetParticipants() []IUser {
	return g.participants
}

func (g *group) GetExpenses() []IExpense {
	return g.expenses
}

func (g *group) GetDescription() string {
	return g.description
}

func (g *group) GetCreatedBy() IUser {
	return g.createdBy
}

func (g *group) GetCreatedAt() time.Time {
	return g.createdAt
}

func (g *group) SetId(id int) IGroup {
	g.id = id
	return g
}

func (g *group) SetName(name string) IGroup {
	g.name = name
	return g
}

func (g *group) SetParticipants(participants []IUser) IGroup {
	g.participants = participants
	return g
}

func (g *group) SetExpenses(expenses []IExpense) IGroup {
	g.expenses = expenses
	return g
}

func (g *group) SetDescription(description string) IGroup {
	g.description = description
	return g
}

func (g *group) SetCreatedBy(createdBy IUser) IGroup {
	g.createdBy = createdBy
	return g
}

func (g *group) SetCreatedAt(createdAt time.Time) IGroup {
	g.createdAt = createdAt
	return g
}

func (g *group) GetLastModified() time.Time {
	return g.lastModified
}

func (g *group) SetLastModified(lastModified time.Time) IGroup {
	g.lastModified = lastModified
	return g
}

func (g *group) AddParticipant(user IUser) IGroup {
	g.participants = append(g.participants, user)
	return g
}
