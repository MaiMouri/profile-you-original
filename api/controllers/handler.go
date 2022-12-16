package contorllers

import (
	"fmt"
	"net/http"

	"profileyou/api/usecase"
	"profileyou/api/utils/errors"

	"github.com/gin-gonic/gin"
	// "gorm.io/driver/sqlite"
)

type keywordController struct {
	keywordUseCase usecase.KeywordUseCase
}

// likes to Usecase by "ku"
func NewKeywordController(ku usecase.KeywordUseCase) keywordController {
	return keywordController{
		keywordUseCase: ku,
	}

}

func (ku *keywordController) GetAllKeywordsGin(c *gin.Context) {
	keywords, err := ku.keywordUseCase.GetKeywords()
	fmt.Printf("keywords :%v\n", keywords)
	if err != nil {
		fmt.Println(err)
		apiErr := errors.NewBadRequestError("Get all Bad Request")
		c.IndentedJSON(apiErr.Status, apiErr)
		return
	}
	type ResultDataField struct {
		KeywordId   string
		Word        string
		Description string
		ImageUrl    string
	}
	var data []ResultDataField
	for _, keyword := range keywords {
		keywordId := string(keyword.GetKeywordId())
		word := string(keyword.GetWord())
		description := string(keyword.GetDescription())
		imageUrl := string(keyword.GetImageUrl())
		data = append(data, ResultDataField{KeywordId: keywordId, Word: word, Description: description, ImageUrl: imageUrl})
	}
	// c.HTML(200, "index.html", gin.H{"keywords": data})
	c.IndentedJSON(http.StatusOK, data)
}

func (ku *keywordController) GetKeyword(c *gin.Context) {
	id := c.Param("id")
	fmt.Printf("param id: %v\n", id)
	keyword, err := ku.keywordUseCase.GetKeyword(id)
	fmt.Printf("keyword id: %v\n", keyword)
	if err != nil {
		fmt.Println(err)
		apiErr := errors.NotFoundError("Tried to find the record but Not found")
		c.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	type ResultDataField struct {
		KeywordId   string
		Word        string
		Description string
		ImageUrl    string
	}

	data := ResultDataField{
		KeywordId:   string(keyword.GetKeywordId()),
		Word:        string(keyword.GetWord()),
		Description: string(keyword.GetDescription()),
		ImageUrl:    string(keyword.GetImageUrl()),
	}
	c.IndentedJSON(http.StatusOK, data)

}

func (ku *keywordController) Index(c *gin.Context) {
	keywords, err := ku.keywordUseCase.GetKeywords()
	fmt.Printf("keywords :%v\n", keywords)
	if err != nil {
		fmt.Println(err)
		apiErr := errors.NewBadRequestError("Index Bad Request")
		c.IndentedJSON(apiErr.Status, apiErr)
		return
	}
	// 20221213 - Deleted prior to implement ResultDataField struct
	// c.IndentedJSON(http.StatusOK, keywords)

	type ResultDataField struct {
		KeywordId   string
		Word        string
		Description string
		ImageUrl    string
	}
	var data []ResultDataField
	for _, keyword := range keywords {
		keywordId := string(keyword.GetKeywordId())
		word := string(keyword.GetWord())
		description := string(keyword.GetDescription())
		imageUrl := string(keyword.GetImageUrl())
		data = append(data, ResultDataField{KeywordId: keywordId, Word: word, Description: description, ImageUrl: imageUrl})
	}
	c.HTML(200, "index.html", gin.H{"keywords": data})
	// c.IndentedJSON(http.StatusOK, data)
}

func (ku *keywordController) DetailKeyword(c *gin.Context) {

	id := c.Param("id")
	keyword, err := ku.keywordUseCase.GetKeyword(id)
	if err != nil {
		fmt.Println(err)
		apiErr := errors.NotFoundError("Not found")
		c.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	type ResultDataField struct {
		KeywordId   string
		Word        string
		Description string
		ImageUrl    string
	}

	data := ResultDataField{KeywordId: string(keyword.GetKeywordId()), Word: string(keyword.GetWord()), Description: string(keyword.GetDescription()), ImageUrl: string(keyword.GetImageUrl())}
	c.IndentedJSON(http.StatusOK, data)
}

func (ku *keywordController) CreateKeyword(c *gin.Context) {
	type RequestDataField struct {
		Word        string `json:"Word" binding:"required"`
		Description string `json:"Description"`
		ImageUrl    string `json:"ImageUrl"`
		KeywordId   string `json:"KeywordId"`
	}

	var json RequestDataField

	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Printf("Error: %v\n", err)
		apiErr := errors.NewBadRequestError("Bad request on binding json")
		c.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	word := json.Word
	description := ""
	imageUrl := ""
	fmt.Printf("Receive a post: %s\n", word)

	err := ku.keywordUseCase.CreateKeyword(word, description, imageUrl)
	if err != nil {
		fmt.Println(err)
		apiErr := errors.InternalSeverError("Server Error when posting")
		c.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": json})
}

func (ku *keywordController) UpdateKeyword(c *gin.Context) {
	type RequestDataField struct {
		KeywordId   string `json:"KeywordId" binding:"required"`
		Word        string `json:"Word" binding:"required"`
		Description string `json:"Description"`
		ImageUrl    string `json:"ImageUrl"`
	}

	var json RequestDataField

	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println(err)
		apiErr := errors.NewBadRequestError("Bad request")
		c.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	id := json.KeywordId
	word := json.Word
	description := json.Description
	imageUrl := json.ImageUrl

	fmt.Printf("Updating a keyword id: %v", id)
	fmt.Printf("Updating a description: %v", description)
	keyword, err := ku.keywordUseCase.GetKeyword(id)
	if err != nil {
		fmt.Println(err)
		apiErr := errors.NotFoundError("Not found")
		c.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	err = ku.keywordUseCase.UpdateKeyword(id, word, description, imageUrl)
	if err != nil {
		fmt.Println(err)
		apiErr := errors.InternalSeverError("Server Error")
		c.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	c.IndentedJSON(http.StatusOK, keyword)
}

func (ku *keywordController) DeleteKeyword(c *gin.Context) {
	type RequestDataField struct {
		KeywordId string `json:"keywordId" binding:"required"`
	}
	var json RequestDataField

	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println(err)
		apiErr := errors.NewBadRequestError("Bad Request")
		c.IndentedJSON(apiErr.Status, apiErr)
		return
	}

	keyword_id := json.KeywordId

	fmt.Printf("Delete mode%v\n", keyword_id)
	err := ku.keywordUseCase.DeleteKeyword(keyword_id)
	if err != nil {
		fmt.Println(err)
		apiErr := errors.InternalSeverError("Server Error")
		c.IndentedJSON(apiErr.Status, apiErr)
		return
	}
	fmt.Printf("Delete json %v\n", json)
	c.IndentedJSON(http.StatusOK, gin.H{"data": keyword_id})

}
