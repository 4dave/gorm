package controllers

import (
	"github.com/4dave/goapi/initializers"
	"github.com/4dave/goapi/models"
	"github.com/gin-gonic/gin"
)

func ContactCreate(c *gin.Context) {
	var body []models.ContactBody
	c.Bind(&body)
	var contacts []models.Contact
	for _, contact := range body {
		contacts = append(contacts, models.Contact{FirstName: contact.FirstName, LastName: contact.LastName, Address: contact.Address, Email: contact.Email, Phone: contact.Phone, City: contact.City, State: contact.State, Zip: contact.Zip})
	}
	result := initializers.DB.Create(&contacts)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"contacts": contacts,
	})
}

func ContactList(c *gin.Context) {
	var contacts []models.Contact
	googlemapsurl := "https://www.google.com/maps/search/?api=1&query="
	result := initializers.DB.Find(&contacts)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// combine address, city, state, zip and return google maps url
	for i, contact := range contacts {
		addressCombined := contact.Address + ", " + contact.City + ", " + contact.State + ", " + contact.Zip
		contacts[i].Map = googlemapsurl + addressCombined
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
