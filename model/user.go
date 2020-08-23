package models

import (
	"time"

	"github.com/apranta/event_models/basemodel"
)

type User struct {
	basemodel.BaseModel
	Avatar          string     `json:"avatar" gorm:"column:avatar"`
	Username        string     `json:"username" gorm:"column:username"`
	Password        string     `json:"-" gorm:"column:password"`
	ConfirmPassword string     `json:"-" gorm:"-"`
	Name            string     `json:"name" gorm:"column:name"`
	DateOfBirth     string     `json:"date_of_birth" gorm:"column:date_of_birth"`
	PhoneNumber     string     `json:"phone_number" gorm:"column:phone_number"`
	Email           string     `json:"email" gorm:"column:email"`
	Address         string     `json:"address" gorm:"column:address"`
	VerifyEmail     string     `json:"verify_email" gorm:"column:verify_email"`
	VerifyOtp       string     `json:"verify_otp" gorm:"column:verify_otp"`
	MaritalStatus   string     `json:"marital_status" gorm:"column:marital_status"`
	Status          string     `json:"status" gorm:"column:status"`
	Role            string     `json:"role" gorm:"column:role"`
	DeletedAt       *time.Time `json:"deleted_at" sql:"index"`
}

func (User) TableName() string {
	return "users"
}

func (p *User) Create() error {
	err := basemodel.Create(&p)
	return err
}

func (p *User) Save() error {
	err := basemodel.Save(&p)
	return err
}

func (p *User) Delete() error {
	err := basemodel.Delete(&p)
	return err
}

func (p *User) FindbyID(id int) error {
	err := basemodel.FindbyID(&p, id)
	return err
}

func (p *User) SingleFind(filter interface{}) error {
	err := basemodel.SingleFindFilter(&p, filter)
	return err
}

func (p *User) GetAll(filter interface{}) (result interface{}, err error) {
	dist := []User{}
	result, err = basemodel.GetAll(&dist, filter)

	return result, err
}

func (b *User) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result basemodel.PagedFindResult, err error) {
	Question := []User{}
	orders := []string{orderby}
	sorts := []string{sort}
	result, err = basemodel.PagedFindFilter(&Question, page, rows, orders, sorts, filter, []string{})

	return result, err
}

func (model *User) FindFilter(order string, sort string, limit int, offset int, filter interface{}) (result interface{}, err error) {
	cat := []User{}
	orders := []string{order}
	sorts := []string{sort}
	result, err = basemodel.FindFilter(&cat, orders, sorts, limit, offset, filter)
	return result, err
}
