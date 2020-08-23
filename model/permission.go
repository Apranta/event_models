package models

import basemodel "github.com/apranta/event_models"

type (
	// Permissions main type
	Permissions struct {
		basemodel.BaseModel
		Permission string `json:"permission" gorm:"column:permission"`
		RoleID     int    `json:"role_id" gorm:"column:role_id"`
	}
)

// Create new
func (model *Permissions) Create() error {
	err := basemodel.Create(&model)
	return err
}

// Save role
func (model *Permissions) Save() error {
	err := basemodel.Save(&model)
	return err
}

// Delete role
func (model *Permissions) Delete() error {
	err := basemodel.Delete(&model)
	return err
}

// FindbyID self explanatory
func (model *Permissions) FindbyID(id int) error {
	err := basemodel.FindbyID(&model, id)
	return err
}

// FilterSearchSingle use filter to find one role
func (model *Permissions) FilterSearchSingle(filter interface{}) error {
	err := basemodel.SingleFindFilter(&model, filter)
	return err
}

// PagedFindFilter use filter to find all matching role, return using paging format
func (model *Permissions) PagedFindFilter(page int, rows int, orderby []string, sort []string, filter interface{}) (basemodel.PagedFindResult, error) {
	role := []Permissions{}

	return basemodel.PagedFindFilter(&role, page, rows, orderby, sort, filter, []string{})
}

// FindFilter use filter to find all matching role
func (model *Permissions) FindFilter(limit int, offset int, orderby []string, sort []string, filter interface{}) (interface{}, error) {
	role := []Permissions{}

	return basemodel.FindFilter(&role, orderby, sort, limit, offset, filter)
}
