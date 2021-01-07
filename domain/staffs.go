package domain

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Staffs struct {
	StaffID   int    `json:"staff_id"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" 		validate:"required,email"`
	Phone     string `json:"phone" 		validate:"required"`
	Active    string `json:"active"		validate:"required"`
	StoreID   int    `json:"store_id"`
	ManagerID int    `json:"manager_id"`
	StoreData Stores `json:"store_data"`
	// ManagerData Staffs `json:"manager_data"`
}

type StaffsRepository interface {
	All() ([]*Staffs, error)
	Single(id int) (*Staffs, error)
	Insert(staffs Staffs) (int64, error)
	Update(staffs Staffs) (int64, error)
	Delete(id int) (int64, error)
}

type StaffsUseCase interface {
	All() ([]*Staffs, error)
	Single(id int) (*Staffs, error)
	Insert(staffs Staffs) (int64, error)
	Update(staffs Staffs) (int64, error)
	Delete(id int) (int64, error)
}

func (s *Staffs) Compare(staffs Staffs) bool {
	return s.StaffID == staffs.StaffID &&
		s.FirstName == staffs.FirstName &&
		s.LastName == staffs.LastName &&
		s.Email == staffs.Email &&
		s.Phone == staffs.Phone &&
		s.Active == staffs.Active &&
		s.StoreID == staffs.StoreID &&
		s.ManagerID == staffs.ManagerID
}

func (s *Staffs) Validate() error {
	var validate *validator.Validate

	validate = validator.New()

	err := validate.Struct(s)

	if err != nil {

		if _, ok := err.(*validator.InvalidValidationError); ok {
			return &validator.InvalidValidationError{}
		}

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}
	}

	return err
}
