package models

import (
	"time"

	"github.com/apranta/event_models/basemodel"
)

//Event Modeling Table Event
type (
	Event struct {
		basemodel.BaseModel
		DeletedAt    *time.Time `json:"deleted_at" sql:"index"`
		Title        string     `json:"title" gorm:"column:title"`
		Cover        string     `json:"cover" gorm:"column:cover"`
		Camera       string     `json:"camera" gorm:"column:camera"`
		CategoryID   int        `json:"category_id" gorm:"column:category_id"`
		CategoryName string     `json:"category_name" gorm:"column:category_name"`
		Description  string     `json:"description" gorm:"column:description"`
		Tags         string     `json:"tags" gorm:"column:tags"`
		Price        int        `json:"price" gorm:"column:price"`
		CreatedByID  int        `json:"created_by_id" gorm:"column:created_by"`
		CreatedBy    string     `json:"created_by" gorm:"column:created_by_name`
		StartedAt    time.Time  `json:"started_at" gorm:"column:started_at"`
		EndedAt      time.Time  `json:"ended_at" gorm:"column:ended_at"`
	}
)

func (Event) TableName() string {
	return "events"
}
func (p *Event) AfterFind() error {
	user := User{}
	_ = user.FindbyID(p.CreatedByID)
	p.CreatedBy = user.Name
	return nil
}
func (p *Event) Create() error {
	err := basemodel.Create(&p)
	return err
}

func (p *Event) Save() error {
	err := basemodel.Save(&p)
	return err
}

func (p *Event) Delete() error {
	err := basemodel.Delete(&p)
	return err
}

func (p *Event) FindbyID(id int) error {
	err := basemodel.FindbyID(&p, id)
	return err
}

func (p *Event) SingleFind(filter interface{}) error {
	err := basemodel.SingleFindFilter(&p, filter)
	return err
}

func (p *Event) GetAll(filter interface{}) (result interface{}, err error) {
	dist := []Event{}
	result, err = basemodel.GetAll(&dist, filter)

	return result, err
}

func (b *Event) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result basemodel.PagedFindResult, err error) {
	Question := []Event{}
	orders := []string{orderby}
	sorts := []string{sort}
	result, err = basemodel.PagedFindFilter(&Question, page, rows, orders, sorts, filter, []string{})

	return result, err
}

func (model *Event) FindFilter(order string, sort string, limit int, offset int, filter interface{}) (result interface{}, err error) {
	cat := []Event{}
	orders := []string{order}
	sorts := []string{sort}
	result, err = basemodel.FindFilter(&cat, orders, sorts, limit, offset, filter)
	return result, err
}
