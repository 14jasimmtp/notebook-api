package routes

import (
	api "github.com/14jasimmtp/notebook-api/pkg/apis"
	"github.com/gin-gonic/gin"
)

func Notebook(r *gin.Engine,svc *api.Service){
	r.GET("/notes/all",svc.ShowAllNotes)
	r.GET("/notes",svc.ShowOneNote)

	r.DELETE("/notes/delete",svc.DeleteNote)
	r.PUT("/notes/edit",svc.EditNote)
	r.POST("/notes/add",svc.AddNote)
}