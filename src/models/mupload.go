package models

import "mime/multipart"

type Upload struct {
	Name   string                `form:"name" bson:"name" binding:"required"`
	Avatar *multipart.FileHeader `form:"avatar" bson:"avatar" binding:"required"`
}
