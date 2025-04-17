package function

import (
	"encoding/json"
	"errors"
	"math/rand"
	"time"
)

// FlowInput is the input to the flow function handler function in the template project
func ExecFlow(request FlowInput) ([]byte, error) {
	// Check if the user_id is provided
	if request.Args.UserID == nil {
		return nil, errors.New("user_id is required")
	}

	// Add a sleep time for simulating database connection
	time.Sleep(time.Duration(rand.Intn(10)+5) * time.Millisecond)

	// Simulate a database query to get the user information
	userInfo := UserInfo{
		ID: *request.Args.UserID,
	}
	switch *request.Args.UserID {
	case 10:
		userInfo = UserInfo{
			FirstName:   "Danendra",
			LastName:    "Herdiansyah",
			PhoneNumber: "012345678901",
			CurrentAddressLocation: Point{
				Lat: 10.10,
				Lon: 40.40,
			},
			Addresses: []string{"CSL UI", "Fasilkom UI"},
		}
		break
	case 20:
		userInfo = UserInfo{
			FirstName:   "user",
			LastName:    "2",
			PhoneNumber: "012345678902",
			CurrentAddressLocation: Point{
				Lat: 22.11,
				Lon: 22.22,
			},
			Addresses: []string{"Depok"},
		}
		break
	case 30:
		userInfo = UserInfo{
			FirstName:   "user",
			LastName:    "3",
			PhoneNumber: "012345678903",
			CurrentAddressLocation: Point{
				Lat: 40.40,
				Lon: 10.10,
			},
			Addresses: []string{"Margonda"},
		}
		break
	case 40:
		userInfo = UserInfo{
			FirstName:   "user",
			LastName:    "4",
			PhoneNumber: "012345678904",
			CurrentAddressLocation: Point{
				Lat: 44.11,
				Lon: 44.22,
			},
			Addresses: []string{"Depok", "Pondok Cina"},
		}
		break
	default:
		return nil, errors.New("can not find the user")
	}

	// Return the user information as a JSON byte array
	return json.Marshal(userInfo)
}
