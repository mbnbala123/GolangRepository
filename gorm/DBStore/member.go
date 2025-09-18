package DBStore

type Date struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

type Member struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name" gorm:"size:100;not null"`
	LastName  string `json:"last_name" gorm:"size:100;not null"`
	Email     string `json:"email" gorm:"size:100;unique;not null"`
	Phone     string `json:"phone" gorm:"size:15;not null"`
	Address   string `json:"address" gorm:"size:255;not null"`
	StartDate string `json:"start_date" gorm:"type:date;not null"`
	EndDate   string `json:"end_date" gorm:"type:date;not null"`
}

func (m *Member) SaveMember() (bool, error) {
	db := CreateConnection()
	//defer db.Close()

	if err := db.Create(m).Error; err != nil {
		return false, err
	}
	return true, nil
}
