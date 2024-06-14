package entity

type UserData struct {
	Id          *uint   `json:"id,omitempty"`
	FirstName   *string `json:"firstName,omitempty"`
	Surname     *string `json:"surname,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	Active      *bool   `json:"active,omitempty"`
	Email       *string `json:"email,omitempty"`
	IsVerified  *bool   `json:"isVerified,omitempty"`
}
