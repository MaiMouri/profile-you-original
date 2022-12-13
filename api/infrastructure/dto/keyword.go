package dto

import (
	"fmt"
	"profileyou/api/domain/model/keyword"
	"time"
)

type Keyword struct {
	ID          int
	KeywordId   string
	Word        string
	Description string
	ImageUrl    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func ConvertKeyword(k *keyword.Keyword) *Keyword {
	return &Keyword{
		KeywordId:   string(k.GetKeywordId()),
		Word:        string(k.GetWord()),
		Description: string(k.GetDescription()),
		ImageUrl:    string(k.GetImageUrl()),
	}
}

func AdaptKeyword(converted_keyword *Keyword) (*keyword.Keyword, error) {
	keyword, err := keyword.New(
		converted_keyword.KeywordId,
		converted_keyword.Word,
		converted_keyword.Description,
		converted_keyword.ImageUrl,
	)

	if err != nil {
		return nil, err
	}

	return keyword, nil
}

func AdaptKeywords(converted_keywords []*Keyword) ([]*keyword.Keyword, error) {
	var keywords []*keyword.Keyword

	for _, converted_keyword := range converted_keywords {
		keyword, err := keyword.New(
			converted_keyword.KeywordId,
			converted_keyword.Word,
			converted_keyword.Description,
			converted_keyword.ImageUrl,
		)
		fmt.Printf("DTO: CONVERTED KEYWORD 1 ERROR, %v\n", err)

		if err != nil {
			fmt.Printf("DTO: CONVERTED KEYWORDS ERROR, %v\n", err)
			return nil, err
		}
		keywords = append(keywords, keyword)
	}
	fmt.Printf("DTO: RETRIEVE KEYWORDS, %v\n", keywords)

	return keywords, nil
}
