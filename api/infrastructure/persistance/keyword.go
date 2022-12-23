package persistance

import (
	"fmt"
	"profileyou/api/domain/model/keyword"
	"profileyou/api/domain/model/user"
	"profileyou/api/domain/repository"
	"profileyou/api/infrastructure/dto"
	"time"

	"gorm.io/gorm"
)

const dbTimeout = time.Second * 3

type keywordPersistance struct {
	Conn *gorm.DB
}

func NewKeywordPersistance(conn *gorm.DB) repository.KeywordRepository {
	return &keywordPersistance{Conn: conn}
}

func (kp *keywordPersistance) GetKeyword(id string) (result *keyword.Keyword, err error) {

	var keyword dto.Keyword
	if result := kp.Conn.Where("keyword_id = ?", id).First(&keyword); result.Error != nil {
		err := result.Error
		return nil, err
	}

	// return &keyword, nil
	result_keyword, err := dto.AdaptKeyword(&keyword)
	if err != nil {
		return nil, err
	}

	return result_keyword, nil
}
func (kp *keywordPersistance) GetLastKeyword() (result *keyword.Keyword, err error) {

	var keyword dto.Keyword
	if result := kp.Conn.Last(&keyword); result.Error != nil {
		err := result.Error
		return nil, err
	}

	// return &keyword, nil
	result_keyword, err := dto.AdaptKeyword(&keyword)
	if err != nil {
		return nil, err
	}

	return result_keyword, nil
}

func (kp *keywordPersistance) GetKeywords() (result []*keyword.Keyword, err error) {

	var keywords []*dto.Keyword

	// 降順に変更
	if result := kp.Conn.Order("created_at desc").Find(&keywords); result.Error != nil {
		err := result.Error
		return nil, err
	}
	fmt.Println(result)

	result_keywords, err := dto.AdaptKeywords(keywords)
	if err != nil {
		return nil, err
	}
	fmt.Println(result_keywords)
	return result_keywords, nil
}

func (kp *keywordPersistance) Create(k *keyword.Keyword) error {

	converted_keyword := dto.ConvertKeyword(k)
	if result := kp.Conn.Create(converted_keyword); result.Error != nil {
		err := result.Error
		return err
	}
	return nil
}

func (kp *keywordPersistance) Update(k *keyword.Keyword) error {

	converted_keyword := dto.ConvertKeyword(k)
	if result := kp.Conn.Where("keyword_id = ?", converted_keyword.KeywordId).
		Updates(converted_keyword); result.Error != nil {
		err := result.Error
		return err
	}
	return nil
}

func (kp *keywordPersistance) Delete(k *keyword.Keyword) error {

	converted_keyword := dto.ConvertKeyword(k)
	if result := kp.Conn.Where("keyword_id = ?", converted_keyword.KeywordId).Delete(converted_keyword); result.Error != nil {
		err := result.Error
		return err
	}
	return nil
}

func (kp *keywordPersistance) GetUserByEmail(email string) (result *user.User, err error) {

	var user user.User
	if result := kp.Conn.Where("email = ?", email).First(&user); result.Error != nil {
		err := result.Error
		return nil, err
	}

	return &user, nil
}
