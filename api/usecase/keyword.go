package usecase

import (
	"fmt"
	"profileyou/api/domain/model/keyword"
	"profileyou/api/domain/model/user"
	"profileyou/api/domain/repository"
)

type KeywordUseCase interface {
	GetKeyword(id string) (result *keyword.Keyword, err error)
	GetLastKeyword() (result *keyword.Keyword, err error)
	GetKeywords() (result []*keyword.Keyword, err error)
	CreateKeyword(word string, description string, imgaeUrl string) error
	UpdateKeyword(id string, word string, description string, imageUrl string) error
	DeleteKeyword(id string) error
	GetUserForAuth(email string) (result *user.User, err error)
}

type keywordUseCase struct {
	keywordRepository repository.KeywordRepository
}

func NewKeywordUseCase(kr repository.KeywordRepository) KeywordUseCase {
	return &keywordUseCase{
		keywordRepository: kr,
	}
}

func (ku *keywordUseCase) GetKeyword(id string) (result *keyword.Keyword, err error) {
	keyword, err := ku.keywordRepository.GetKeyword(id)
	if err != nil {
		return nil, err
	}

	return keyword, nil
}

func (ku *keywordUseCase) GetLastKeyword() (result *keyword.Keyword, err error) {
	keyword, err := ku.keywordRepository.GetLastKeyword()
	if err != nil {
		return nil, err
	}

	return keyword, nil
}

func (ku *keywordUseCase) GetKeywords() (result []*keyword.Keyword, err error) {
	keywords, err := ku.keywordRepository.GetKeywords()
	if err != nil {
		return nil, err
	}

	return keywords, nil
}

func (ku *keywordUseCase) CreateKeyword(word string, description string, imageUrl string) error {
	keyword, err := keyword.Create(word, description, imageUrl)
	if err != nil {
		return err
	}

	err = ku.keywordRepository.Create(keyword)
	if err != nil {
		return err
	}

	return nil
}

func (ku *keywordUseCase) UpdateKeyword(id string, word string, description string, imageUrl string) error {
	current_keyword, err := ku.keywordRepository.GetKeyword(id)
	if err != nil {
		return err
	}

	keywordId := current_keyword.GetKeywordId()
	update_keyword, err := keyword.New(keywordId, word, description, imageUrl)
	if err != nil {
		return err
	}
	err = ku.keywordRepository.Update(update_keyword)
	if err != nil {
		return err
	}

	return nil
}

func (ku *keywordUseCase) DeleteKeyword(id string) error {
	keyword, err := ku.keywordRepository.GetKeyword(id)
	if err != nil {
		return err
	}

	err = ku.keywordRepository.Delete(keyword)
	if err != nil {
		return err
	}

	return nil
}

func (ku *keywordUseCase) GetUserForAuth(email string) (result *user.User, err error) {
	fmt.Printf("User email %v has requested to log in!", email)
	current_user, err := ku.keywordRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return current_user, nil
}
