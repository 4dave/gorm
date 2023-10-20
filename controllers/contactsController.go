package controllers

import (
	"github.com/4dave/goapi/initializers"
	"github.com/4dave/goapi/models"
	"github.com/gin-gonic/gin"
)

func ContactCreate(c *gin.Context) {
	var body models.ContactBody
	c.Bind(&body)
	contact := models.Contact{FirstName: body.FirstName, LastName: body.LastName, Address: body.Address, Email: body.Email, Phone: body.Phone}
	result := initializers.DB.Create(&contact)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"contact": contact,
	})
}

func ContactList(c *gin.Context) {
	var contacts []models.Contact
	result := initializers.DB.Find(&contacts)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"contacts": contacts,
	})
}

func ContactShow(c *gin.Context) {
	var contact models.Contact
	result := initializers.DB.First(&contact, c.Param("id"))

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"contact": contact,
	})
}

func ContactUpdate(c *gin.Context) {
	id := c.Param("id")
	var body models.ContactBody
	c.Bind(&body)
	var contact models.Contact
	result := initializers.DB.First(&contact, id)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Contact not found",
		})
		return
	}

	initializers.DB.Model(&contact).Updates(models.Contact{FirstName: body.FirstName, LastName: body.LastName, Address: body.Address, Email: body.Email, Phone: body.Phone})

	c.JSON(200, gin.H{
		"contact": contact,
	})
}

func ContactDelete(c *gin.Context) {
	var contact models.Contact
	result := initializers.DB.First(&contact, c.Param("id"))

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Contact not found",
		})
		return
	}

	initializers.DB.Delete(&contact)

	c.JSON(200, gin.H{
		"message": "Contact deleted",
	})
}
