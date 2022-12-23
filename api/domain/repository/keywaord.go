package repository

import (
	"profileyou/api/domain/model/keyword"
	"profileyou/api/domain/model/user"
)

type KeywordRepository interface {
	GetKeyword(id string) (result *keyword.Keyword, err error)
	GetLastKeyword() (result *keyword.Keyword, err error)
	GetKeywords() (result []*keyword.Keyword, err error)
	Create(k *keyword.Keyword) error
	Update(k *keyword.Keyword) error
	Delete(k *keyword.Keyword) error
	GetUserByEmail(email string) (result *user.User, err error)
}
