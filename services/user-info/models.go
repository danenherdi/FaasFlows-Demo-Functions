package function

// Point is a struct to represent a point in the map
type Point struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// UserInfo is a struct to represent information of a user in the system
type UserInfo struct {
	ID                     uint     `json:"id"`
	FirstName              string   `json:"first_name"`
	LastName               string   `json:"last_name"`
	PhoneNumber            string   `json:"phone_number"`
	CurrentAddressLocation Point    `json:"current_address_location"`
	Addresses              []string `json:"addresses"`
}

// Input is the argument of your flow function
type Input struct {
	UserID *uint `json:"user_id"`
}
