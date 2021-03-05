package keywordAssociation

import (
	"github.com/HackU-2020-vol4/back-end/db"
	"github.com/HackU-2020-vol4/back-end/entity"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Service struct{}

type keywordAssociation entity.KeywordAssociation

//index
func (s Service) GetbyKeyword(keywordID string) ([]keywordAssociation, error) {
	db := db.GetDB()
	var ka []keywordAssociation
	if err := db.Where("keyword_id = ?", keywordID).Find(&ka).Error; err != nil {
		return ka, err
	}
	return ka, nil
}

func (s Service) CreateModel(c *gin.Context) (keywordAssociation, error) {
	db := db.GetDB()
	var ka keywordAssociation
	ka.PublisherID = c.Param("publisherID")
	keyword_id := c.Param("keywordID")
	//型の変換
	convertedStrInt64, _ := strconv.ParseInt(keyword_id, 10, 64)
	ka.KeywordID = uint(convertedStrInt64)
	if err := c.BindJSON(&ka); err != nil {

	}
	if err := db.Create(&ka).Error; err != nil {
		return ka, err
	}
	return ka, nil
}

func (s Service) DeleteByID(id string) error {
	db := db.GetDB()
	var ka keywordAssociation
	if err := db.Where("id = ?", id).Delete(&ka).Error; err != nil {
		return err
	}
	return nil
}

func (s Service) UpdateByID(id string, c *gin.Context) (keywordAssociation, error) {
	db := db.GetDB()
	var ka keywordAssociation
	if err := db.Where("id = ?", id).First(&ka).Error; err != nil {
		return ka, err
	}
	if err := c.BindJSON(&ka); err != nil {
		return ka, err
	}
	db.Save(&ka)
	return ka, nil
}