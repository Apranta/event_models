package models

import (
	"time"

	"github.com/apranta/event_models/basemodel"
)

type (
	EventCategory struct {
		basemodel.BaseModel
		Name        string     `json:"name" gorm:"column:name"`
		Description string     `json:"description" gorm:"column:description"`
		Status      string     `json:"status" gorm:"column:status"`
		DeletedAt   *time.Time `json:"deleted_at" sql:"index"`
	}
)

func (EventCategory) TableName() string {
	return "event_category"
}

func (p *EventCategory) Create() error {
	err := basemodel.Create(&p)
	return err
}

func (p *EventCategory) Save() error {
	err := basemodel.Save(&p)
	return err
}

func (p *EventCategory) Delete() error {
	err := basemodel.Delete(&p)
	return err
}

func (p *EventCategory) FindbyID(id int) error {
	err := basemodel.FindbyID(&p, id)
	return err
}

func (p *EventCategory) SingleFind(filter interface{}) error {
	err := basemodel.SingleFindFilter(&p, filter)
	return err
}

func (p *EventCategory) GetAll(filter interface{}) (result interface{}, err error) {
	dist := []EventCategory{}
	result, err = basemodel.GetAll(&dist, filter)

	return result, err
}

func (b *EventCategory) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result basemodel.PagedFindResult, err error) {
	Question := []EventCategory{}
	orders := []string{orderby}
	sorts := []string{sort}
	result, err = basemodel.PagedFindFilter(&Question, page, rows, orders, sorts, filter, []string{})

	return result, err
}

func (model *EventCategory) FindFilter(order string, sort string, limit int, offset int, filter interface{}) (result interface{}, err error) {
	cat := []EventCategory{}
	orders := []string{order}
	sorts := []string{sort}
	result, err = basemodel.FindFilter(&cat, orders, sorts, limit, offset, filter)
	return result, err
}
