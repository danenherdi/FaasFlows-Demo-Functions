package function

// Point is a struct to represent a point in the map
type Point struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// UserInfo is a struct to represent the user information of a passenger
type UserInfo struct {
	ID                     uint     `json:"id"`
	FirstName              string   `json:"first_name"`
	LastName               string   `json:"last_name"`
	PhoneNumber            string   `json:"phone_number"`
	CurrentAddressLocation Point    `json:"current_address_location"`
	Addresses              []string `json:"addresses"`
}

// FriendsInfo is a struct to represent the number of friends of a user
type FriendsInfo struct {
	NumberOfFriends string `json:"number_of_friends"`
}

// Input is the argument of your flow function
type Input struct {
	UserID *uint `json:"user_id"`
}
