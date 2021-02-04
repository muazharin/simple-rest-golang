package models

type User struct {
	ID       int    `bson:"id"`
	Name     string `bson:"name"`
	Address  string `bson:"address"`
	Age      int    `bson:"age"`
	Image    string `bson:"image"`
	CreateAt string `bson:"create_at"`
	UpdateAt string `bson:"update_at"`
}

type Users []User
