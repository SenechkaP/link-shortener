package di

import "advpractice/internal/user"

type IStatRepository interface {
	AddClick(uint)
}

type IUserRepository interface {
	Create(*user.User) (*user.User, error)
	GetByEmail(string) (*user.User, error)
}
