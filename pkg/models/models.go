package models

import "gorm.io/gorm"

type Notebook struct {
	gorm.Model
	Title string `json:"title" gorm:"title"`
	Data  string `json:"Data" gorm:"data"`
}

type EditRequest struct{
	ID    int64  `json:"id" `
	Title string `json:"title"`
	Data  string `json:"Data"`
}

type EditResponse struct{
	Message string `json:"message"`
}

type AddRequest struct {
	Title string `json:"title"`
	Data  string `json:"Data"`
}

type AddResponse struct{
	Message string `json:"message"`
}

type DeleteRequest struct{
	Id uint `json:"id"`
}
type DeleteResponse struct{
	Message string `json:"message"`
}


