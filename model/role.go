package models

import (
	"github.com/apranta/event_models/basemodel"
)

type (
	// Role main type
	Role struct {
		basemodel.BaseModel
		Name        string `json:"name" gorm:"column:name"`
		Description string `json:"description" gorm:"column:description"`
		Status      string `json:"status" gorm:"column:status" sql:"DEFAULT:active"`
	}
)

func (Role) TableName() string {
	return "roles"
}

// Create new
func (model *Role) Create() error {
	err := basemodel.Create(&model)
	return err
}

// Save role
func (model *Role) Save() error {
	err := basemodel.Save(&model)
	return err
}

// Delete role
func (model *Role) Delete() error {
	err := basemodel.Delete(&model)
	return err
}

// FindbyID self explanatory
func (model *Role) FindbyID(id int) error {
	err := basemodel.FindbyID(&model, id)
	return err
}

// FilterSearchSingle use filter to find one role
func (model *Role) FilterSearchSingle(filter interface{}) error {
	err := basemodel.SingleFindFilter(&model, filter)
	return err
}

// PagedFindFilter use filter to find all matching role, return using paging format
func (model *Role) PagedFindFilter(page int, rows int, orderby []string, sort []string, filter interface{}) (basemodel.PagedFindResult, error) {
	role := []Role{}

	return basemodel.PagedFindFilter(&role, page, rows, orderby, sort, filter, []string{})
}

// FindFilter use filter to find all matching role
func (model *Role) FindFilter(limit int, offset int, orderby []string, sort []string, filter interface{}) (interface{}, error) {
	role := []Role{}

	return basemodel.FindFilter(&role, orderby, sort, limit, offset, filter)
}

func (b *Role) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result basemodel.PagedFindResult, err error) {
	Question := []Role{}
	orders := []string{orderby}
	sorts := []string{sort}
	result, err = basemodel.PagedFindFilter(&Question, page, rows, orders, sorts, filter, []string{})

	return result, err
}
