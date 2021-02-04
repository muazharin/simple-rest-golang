package controller

import (
	"fmt"
	"gin-rest-mongo/config"
	m_user "gin-rest-mongo/src/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const UserCollection = "user"

// type typetime time.Time

func MongoConfig() *mgo.Database {
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func GetAllUser(c *gin.Context) {
	db := *MongoConfig()
	fmt.Println("MONGO RUNNING: ", db)

	users := m_user.Users{}
	err := db.C(UserCollection).Find(bson.M{}).All(&users)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Error Get All User",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": &users,
	})
}
func GetUser(c *gin.Context) {
	db := *MongoConfig()
	fmt.Println("MONGO RUNNING: ", db)

	id := c.Param("id")
	idParse, errParse := strconv.Atoi(id)
	if errParse != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Parse Param",
		})
		return
	}
	user := m_user.User{}
	err := db.C(UserCollection).Find(bson.M{"id": &idParse}).One(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Get All User",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": &user,
	})
}
func CreateUser(c *gin.Context) {
	db := *MongoConfig()
	fmt.Println("MONGO RUNNING: ", db)

	user := m_user.User{}
	err := c.Bind(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Get Body",
		})
		return
	}

	user.CreateAt = time.Now().String()
	user.UpdateAt = time.Now().String()

	err = db.C(UserCollection).Insert(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Insert User",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Succes Insert User",
		"user":    &user,
	})
}

func UpdateUser(c *gin.Context) {
	db := *MongoConfig()
	fmt.Println("MONGO RUNNING: ", db)

	id := c.Param("id")
	idParse, errParse := strconv.Atoi(id)

	if errParse != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Parse Param",
		})
		return
	}

	user := m_user.User{}
	err := c.Bind(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Get Body",
		})
		return
	}

	user.ID = idParse
	user.UpdateAt = time.Now().String()

	err = db.C(UserCollection).Update(bson.M{"id": &idParse}, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Update User",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Succes Update User",
		"user":    &user,
	})
}
func DeleteUser(c *gin.Context) {
	db := *MongoConfig()
	fmt.Println("MONGO RUNNING: ", db)

	id := c.Param("id")
	idParse, errParse := strconv.Atoi(id)
	if errParse != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Parse Param",
		})
		return
	}

	err := db.C(UserCollection).Remove(bson.M{"id": &idParse})
	if err != nil {
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
