package delivery

type IndexUser struct {
	Email string `json:"Email"`
}

type IndexUserAndPlan struct {
	Email string `json:"Email"`
}

type Save struct {
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required"`
	Phone      string `json:"phone" validate:"required"`
	Name       string `json:"name" validate:"required"`
	LastName   string `json:"lastName" validate:"required"`
	DataRegion string `json:"dataRegion" validate:"required"`
	PlanID     int    `json:"planId" validate:"required"`
}

type Change struct {
	Email      string `json:"email" validate:"required"`
	NewEmail   string `json:"newEmail"`
	Phone      string `json:"phone" validate:"required"`
	Name       string `json:"name" validate:"required"`
	LastName   string `json:"lastName" validate:"required"`
	DataRegion string `json:"dataRegion" validate:"required"`
}

type ChangePlan struct {
	Email      string `json:"email" validate:"required"`
	PlanID     int    `json:"planId" validate:"required"`
}

type ChangePass struct {
	Email      string `json:"email" validate:"required"`
	Password   string `json:"password" validate:"required"`
	NewPassword   string `json:"newPassword" validate:"required"`
}

type Destroy struct {
	Email      string `json:"email" validate:"required"`
	Password   string `json:"password" validate:"required"`
}