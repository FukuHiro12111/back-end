package keyword

import (
	"fmt"

	"github.com/HackU-2020-vol4/back-end/db"
	"github.com/HackU-2020-vol4/back-end/entity"

	"github.com/gin-gonic/gin"
)

type Service struct{}

type Keyword entity.Keyword

//index
func (s Service) GetByPublisher(publisherID string) ([]Keyword, error) {
	db := db.GetDB()
	var k []Keyword
	if err := db.Where("publisher_id = ?", publisherID).Find(&k).Error; err != nil {
		return k, err
	}
	defer db.Close()
	return k, nil
}

// create
func (s Service) CreateModel(c *gin.Context) (Keyword, error) {
	db := db.GetDB()
	var k Keyword
	k.PublisherID = c.Param("publisherID")
	fmt.Println(k.PublisherID)
	if err := c.BindJSON(&k); err != nil {
		return k, err
	}
	if err := db.Create(&k).Error; err != nil {
		return k, err
	}
	defer db.Close()
	return k, nil
}

//delete
func (s Service) DeleteByID(id string) error {
	db := db.GetDB()
	var k Keyword
	if err := db.Where("id = ?", id).Delete(&k).Error; err != nil {
		return err
	}
	defer db.Close()
	return nil
}

//update
func (s Service) UpdateByID(id string, c *gin.Context) (Keyword, error) {
	db := db.GetDB()
	var k Keyword

	if err := db.Where("id = ?", id).First(&k).Error; err != nil {
		return k, err
	}

	if err := c.BindJSON(&k); err != nil {
		return k, err
	}

	db.Save(&k)
	defer db.Close()
	return k, nil
}
