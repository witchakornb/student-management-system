package user

type (
	// User struct
	User struct {
		UserID       string `json:"user_id" gorm:"primary_key"`
		Username     string `json:"username"`
		Password     string `json:"password"`
		Role         string `json:"role"`
		FirstNameTH  string `json:"first_name_TH"`
		LastNameTH   string `json:"last_name_TH"`
		FirstNameEN  string `json:"first_name_EN"`
		LastNameEN   string `json:"last_name_EN"`
		Email        string `json:"email"`
		Phone        string `json:"phone"`
		Department   string `json:"department"`
	}
)