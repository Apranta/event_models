package models

import (
	"github.com/apranta/event_models/basemodel"
)

type (
	Ticket struct {
		basemodel.BaseModel
		EventID     int    `json:"event_id" gorm:"column:event_id"`
		Name        string `json:"name" gorm:"column:name"`
		Price       int    `json:"price" gorm:"column:price"`
		Description string `json:"description" gorm:"column:description"`
	}
)

func (Ticket) TableName() string {
	return "tickets"
}

// Create new
func (model *Ticket) Create() error {
	err := basemodel.Create(&model)
	return err
}

// Save ticket
func (model *Ticket) Save() error {
	err := basemodel.Save(&model)
	return err
}

// Delete ticket
func (model *Ticket) Delete() error {
	err := basemodel.Delete(&model)
	return err
}

// FindbyID self explanatory
func (model *Ticket) FindbyID(id int) error {
	err := basemodel.FindbyID(&model, id)
	return err
}

// FilterSearchSingle use filter to find one ticket
func (model *Ticket) FilterSearchSingle(filter interface{}) error {
	err := basemodel.SingleFindFilter(&model, filter)
	return err
}

// PagedFindFilter use filter to find all matching ticket, return using paging format
func (model *Ticket) PagedFindFilter(page int, rows int, orderby []string, sort []string, filter interface{}) (basemodel.PagedFindResult, error) {
	ticket := []Ticket{}

	return basemodel.PagedFindFilter(&ticket, page, rows, orderby, sort, filter, []string{})
}

// FindFilter use filter to find all matching ticket
func (model *Ticket) FindFilter(limit int, offset int, orderby []string, sort []string, filter interface{}) (interface{}, error) {
	ticket := []Ticket{}
	return basemodel.FindFilter(&ticket, orderby, sort, limit, offset, filter)
}

func (b *Ticket) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result basemodel.PagedFindResult, err error) {
	Question := []Ticket{}
	orders := []string{orderby}
	sorts := []string{sort}
	result, err = basemodel.PagedFindFilter(&Question, page, rows, orders, sorts, filter, []string{})

	return result, err
}
