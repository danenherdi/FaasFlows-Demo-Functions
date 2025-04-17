package function

import "time"

// Recommendation is a type to represent the recommendation of the system
const (
	RecommendationRepeat  = "repeat_ride"
	RecommendationReverse = "reverse_ride"
	RecommendationNothing = "no_ride"
)

// Point is a struct to represent a point in the map
type Point struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// Ride is a struct to represent a ride in the system
type Ride struct {
	PassengerID uint      `json:"passengerID"`
	Time        time.Time `json:"time"`
	Origin      Point     `json:"origin"`
	Destination Point     `json:"destination"`
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

// Recommendation is a struct to represent the recommendation of the system
type Recommendation struct {
	Type           string `json:"type"`
	Recommendation *Point `json:"recommendation"`
	BannerText     string `json:"banner_text"`
}

// Input is the argument of your flow function
type Input struct {
	UserID *uint  `json:"user_id"`
	Origin *Point `json:"origin"`
}
