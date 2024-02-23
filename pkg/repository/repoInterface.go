package repository

import "github.com/14jasimmtp/notebook-api/pkg/models"

type NoteRepository interface {
	AddNote(note models.AddRequest) (string, error)
	EditNote(note models.EditRequest) (string, error)
	DeleteNote(note models.DeleteRequest) (string, error)
	GetAllNotes() ([]models.Notebook, error)
	GetNoteWithId(id uint) (*models.Notebook, error)
}
