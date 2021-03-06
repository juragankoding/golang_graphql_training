package domain

type Staffs struct {
	StaffID   int    `json:"staff_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Active    string `json:"active"`
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
