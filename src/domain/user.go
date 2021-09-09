package domain

type UserData struct {
	ID       int64
	Name     string
	LastName string
	Region   string
}

type UserAccount struct {
	ID       int64
	Email    string
	Phone    string
	Password string
	Data     UserData
	Plan     int
}