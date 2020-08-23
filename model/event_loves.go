package models

import basemodel "github.com/apranta/event_models"

type (
	EventLoves struct {
		basemodel.BaseModel
		UserID  int `json:"user_id" gorm:"column:user_id"`
		EventID int `json:"event_id" gorm:"column:event_id"`
	}
)

func (EventLoves) TableName() string {
	return "ticket_orders"
}

// Create new
func (model *EventLoves) Create() error {
	err := basemodel.Create(&model)
	return err
}

// Save ticket order
func (model *EventLoves) Save() error {
	err := basemodel.Save(&model)
	return err
}

// Delete ticket order
func (model *EventLoves) Delete() error {
	err := basemodel.Delete(&model)
	return err
}

// FindbyID self explanatory
func (model *EventLoves) FindbyID(id int) error {
	err := basemodel.FindbyID(&model, id)
	return err
}

// FilterSearchSingle use filter to find one ticket order
func (model *EventLoves) FilterSearchSingle(filter interface{}) error {
	err := basemodel.SingleFindFilter(&model, filter)
	return err
}

// PagedFindFilter use filter to find all matching ticket order, return using paging format
func (model *EventLoves) PagedFindFilter(page int, rows int, orderby []string, sort []string, filter interface{}) (basemodel.PagedFindResult, error) {
	ticketOrder := []EventLoves{}

	return basemodel.PagedFindFilter(&ticketOrder, page, rows, orderby, sort, filter, []string{})
}

// FindFilter use filter to find all matching ticket order
func (model *EventLoves) FindFilter(limit int, offset int, orderby []string, sort []string, filter interface{}) (interface{}, error) {
	ticketOrder := []EventLoves{}
	return basemodel.FindFilter(&ticketOrder, orderby, sort, limit, offset, filter)
}

func (b *EventLoves) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result basemodel.PagedFindResult, err error) {
	Question := []EventLoves{}
	orders := []string{orderby}
	sorts := []string{sort}
	result, err = basemodel.PagedFindFilter(&Question, page, rows, orders, sorts, filter, []string{})

	return result, err
}
