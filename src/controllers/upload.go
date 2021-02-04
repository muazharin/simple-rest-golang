package controller

import (
	"fmt"
	"net/http"

	m_profile "gin-rest-mongo/src/models"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

const UploadCollection = "profile"

func CreateProfile(c *gin.Context) {
	db := *MongoConfig()
	form := m_profile.Upload{}
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Bad Request",
		})
		return
	}
	err := c.SaveUploadedFile(form.Avatar, form.Avatar.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unknown Error",
		})
		return
	}
	// form.Name = "Muaz"

	db.C(UploadCollection).Insert(&form)
	c.JSON(http.StatusOK, gin.H{
		"message": "Succes Update User",
		"user":    &form,
	})
}

func GetAllProfile(c *gin.Context) {
	db := *MongoConfig()

	profile := make([]m_profile.Upload, 0, 10)
	err := db.C(UploadCollection).Find(bson.M{}).All(&profile)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Get All Profile",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": len(profile),
		"data":  &profile,
	})
}
func DeleteProfile(c *gin.Context) {
	db := *MongoConfig()
	fmt.Println("MONGO RUNNING: ", db)

	name := c.Param("name")

	err := db.C(UploadCollection).Remove(bson.M{"name": &name})
	if err != nil {
		fmt.Println("error: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Delete User",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Succes Delete User",
	})
	return
}
