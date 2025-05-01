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

	// Simulate the database query to get the last ride of the user
	lastRide := Ride{
		PassengerID: *request.Args.UserID,
	}

	switch *request.Args.UserID {
	case 10:
		// In this case, the user with ID 10 has the last ride 10 minutes ago from the current time
		// with the origin at (10.10, 40.40) and the destination at (20.20, 30.30)
		lastRide = Ride{
			Time: time.Now().Add(-10 * time.Minute),
			Origin: Point{
				Lat: 10.10,
				Lon: 40.40,
			},
			Destination: Point{
				Lat: 20.20,
				Lon: 30.30,
			},
		}
		break
	case 20:
		// In this case, the user with ID 20 has the last ride 20 minutes ago from the current time
		// with the origin at (20.20, 30.30) and the destination at (30.30, 20.20)
		lastRide = Ride{
			Time: time.Now().Add(-20 * time.Minute),
			Origin: Point{
				Lat: 20.20,
				Lon: 30.30,
			},
			Destination: Point{
				Lat: 30.30,
				Lon: 20.20,
			},
		}
		break
	case 30:
		// In this case, the user with ID 30 has the last ride 30 minutes ago from the current time
		// with the origin at (30.30, 20.20) and the destination at (40.40, 10.10)
		lastRide = Ride{
			Time: time.Now().Add(-30 * time.Minute),
			Origin: Point{
				Lat: 30.30,
				Lon: 20.20,
			},
			Destination: Point{
				Lat: 40.40,
				Lon: 10.10,
			},
		}
		break
	case 40:
		// In this case, the user with ID 40 has the last ride 40 minutes ago from the current time
		// with the origin at (40.40, 10.10) and the destination at (10.10, 40.40)
		lastRide = Ride{
			Time: time.Now().Add(-40 * time.Minute),
			Origin: Point{
				Lat: 40.40,
				Lon: 10.10,
			},
			Destination: Point{
				Lat: 10.10,
				Lon: 40.40,
			},
		}
		break
	default:
		return nil, errors.New("can not find the user")
	}

	// Return the last ride of the user in JSON format
	return json.Marshal(map[string]interface{}{
		"response": lastRide,
	})
}
