package models

import (
	"time"

	"github.com/apranta/event_models/basemodel"
)

type (
	TicketOrder struct {
		basemodel.BaseModel
		TicketID    int       `json:"ticket_id" gorm:"column:ticket_id"`
		UserID      int       `json:"user_id" gorm:"column:user_id"`
		Amount      int       `json:"amount" gorm:"column:amount"`
		OrderStatus int       `json:"order_status" gorm:"column:order_status"`
		PaidCreated time.Time `json:"paid_created" gorm:"column:paid_created"`
	}
)

func (TicketOrder) TableName() string {
	return "ticket_orders"
}

// Create new
func (model *TicketOrder) Create() error {
	err := basemodel.Create(&model)
	return err
}

// Save ticket order
func (model *TicketOrder) Save() error {
	err := basemodel.Save(&model)
	return err
}

// Delete ticket order
func (model *TicketOrder) Delete() error {
	err := basemodel.Delete(&model)
	return err
}

// FindbyID self explanatory
func (model *TicketOrder) FindbyID(id int) error {
	err := basemodel.FindbyID(&model, id)
	return err
}

// FilterSearchSingle use filter to find one ticket order
func (model *TicketOrder) FilterSearchSingle(filter interface{}) error {
	err := basemodel.SingleFindFilter(&model, filter)
	return err
}

// PagedFindFilter use filter to find all matching ticket order, return using paging format
func (model *TicketOrder) PagedFindFilter(page int, rows int, orderby []string, sort []string, filter interface{}) (basemodel.PagedFindResult, error) {
	ticketOrder := []TicketOrder{}

	return basemodel.PagedFindFilter(&ticketOrder, page, rows, orderby, sort, filter, []string{})
}

// FindFilter use filter to find all matching ticket order
func (model *TicketOrder) FindFilter(limit int, offset int, orderby []string, sort []string, filter interface{}) (interface{}, error) {
	ticketOrder := []TicketOrder{}
	return basemodel.FindFilter(&ticketOrder, orderby, sort, limit, offset, filter)
}

func (b *TicketOrder) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result basemodel.PagedFindResult, err error) {
	Question := []TicketOrder{}
	orders := []string{orderby}
	sorts := []string{sort}
	result, err = basemodel.PagedFindFilter(&Question, page, rows, orders, sorts, filter, []string{})

	return result, err
}
