package controllers

import (
	"net/http"
	"strconv"

	"github.com/aabdullahgungor/mybookcase/entities"
	"github.com/aabdullahgungor/mybookcase/models"
	"github.com/gin-gonic/gin"
)

type PublisherController struct{
}

func (p PublisherController) GetAll(c *gin.Context)  {
	var publisherModel models.PublisherModel
	publishers, _ := publisherModel.GetAll()
	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, publishers)
}

func (p PublisherController) GetById(c *gin.Context)  {
	str_id := c.Param("id")
	int_id, errId := strconv.Atoi(str_id)
	if errId != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}
	var publisherModel models.PublisherModel
	publisher, err := publisherModel.GetById(int_id)
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": err.Error(),
		})
	}

	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, publisher)
}

func (p PublisherController) Create(c *gin.Context)  {
	
	var publisher entities.Publisher
	err := c.ShouldBindJSON(&publisher)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	var publisherModel models.PublisherModel
	err = publisherModel.Create(&publisher)
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": "cannot edit publisher: " + err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message":"Publisher has been created","publisher_id": publisher.ID})
}

func (p PublisherController) Edit(c *gin.Context)  {
	
	var publisher entities.Publisher
	err := c.ShouldBindJSON(&publisher)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	var publisherModel models.PublisherModel
	err = publisherModel.Edit(&publisher)
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": "cannot edit publisher: " + err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"message":"Publisher has been edited","publisher_id": publisher.ID})
}

func (p PublisherController) Delete(c *gin.Context)  {

	str_id := c.Param("id")
	int_id, errId := strconv.Atoi(str_id)

	if errId  != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	var publisherModel models.PublisherModel

	err := publisherModel.Delete(int_id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete publisher: " + err.Error(),
		})
		return
	}


	c.IndentedJSON(http.StatusAccepted, gin.H{"message":"Publisher has been deleted","publisher_id": int_id})
	
}