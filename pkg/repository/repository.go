package repository

import (
	"fmt"

	"github.com/14jasimmtp/notebook-api/pkg/models"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) NoteRepository {
	return &Repo{DB: db}
}

func (r *Repo) AddNote(note models.AddRequest) (string, error) {
	fmt.Println("repository")
	res := r.DB.Exec(`INSERT INTO notebooks(title,data)VALUES(?,?)`,note.Title,note.Data)
	if res.Error != nil {
		fmt.Println("error")
		return "", res.Error
	}
	fmt.Println("returned no issues")
	return "added", nil
}

func (r *Repo) EditNote(note models.EditRequest) (string, error) {
	if note.Title == "" {
		res := r.DB.Exec(`UPDATE notebooks SET data = ? WHERE id = ?`, note.Data, note.ID)
		if res.Error != nil {
			return "", res.Error
		}
		return "edited", nil
	}
	res := r.DB.Exec(`UPDATE notebooks SET title = ?,data = ? WHERE id = ?`, note.Title, note.Data, note.ID)
	if res.Error != nil {
		return "", res.Error
	}
	return "edited", nil

}

func (r *Repo) DeleteNote(note models.DeleteRequest) (string, error) {
	res := r.DB.Exec(`DELETE FROM notebooks WHERE id = ?`, note.Id)
	if res.Error != nil {
		return "", res.Error
	}
	return "deleted", res.Error
}

func (r *Repo) GetAllNotes() ([]models.Notebook, error) {
	var notes []models.Notebook

	res := r.DB.Raw(`SELECT * FROM notebooks`).Scan(&notes)
	if res.Error != nil {
		return nil, res.Error
	}
	return notes, nil
}

func (r *Repo) GetNoteWithId(Id uint) (*models.Notebook, error) {
	var note models.Notebook
	res := r.DB.Raw(`SELECT * FROM notebooks WHERE id = ?`, Id).Scan(&note)
	if res.Error != nil {
		return nil, res.Error
	}
	return &note, nil
}
