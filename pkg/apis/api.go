package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/14jasimmtp/notebook-api/pkg/models"
	"github.com/14jasimmtp/notebook-api/pkg/repository"
	"github.com/gin-gonic/gin"
)

type Service struct {
	r repository.NoteRepository
}

func NewApi(R repository.NoteRepository) *Service{
	return &Service{r: R}
}

func (us *Service) AddNote(c *gin.Context) {
	var note models.AddRequest

	if c.ShouldBindJSON(&note) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error ocurred"})
	}
	fmt.Println("hi")
	res, err := us.r.AddNote(note)
	if err != nil {
		fmt.Println("error")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println("after addnote")
	c.JSON(http.StatusOK, res)
}

func (us *Service) DeleteNote(c *gin.Context) {
	var Id models.DeleteRequest

	if err := c.BindJSON(&Id); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	res, err := us.r.DeleteNote(Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, res)
}

func (us *Service) EditNote(c *gin.Context) {
	var note models.EditRequest

	if err := c.BindJSON(&note); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	res, err := us.r.EditNote(note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, res)
}

func (us *Service) ShowAllNotes(c *gin.Context) {
	res, err := us.r.GetAllNotes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, res)
}

func (us *Service) ShowOneNote(c *gin.Context) {
	id := c.Query("id")
	idUint, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	res, err := us.r.GetNoteWithId(uint(idUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, res)
}
