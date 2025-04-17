package function

import (
	"encoding/json"
	"errors"
	"fmt"
)

// FlowInput is the input to the flow function handler function in the template project
func ExecFlow(request FlowInput) ([]byte, error) {

	// Check if the user_id is provided
	if request.Args.UserID == nil {
		return nil, errors.New("user_id is required")
	}

	// check if last_ride_of_passenger is provided in the response of the previous function in the flow
	var lastRide Ride
	lastRideResponse, ok := request.Children["last_ride_of_passenger"]
	if !ok || lastRideResponse == nil {
		return nil, errors.New("response of last_ride_of_passenger is required to process")
	}

	// Unmarshal the response of the last_ride_of_passenger function to get the last ride of the passenger
	err := json.Unmarshal(lastRideResponse.Data, &lastRide)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error in unmarshalling: %s", err))
	}
	fmt.Printf("%+v\n", lastRide)

	// Return the ride history of the passenger with the last ride repeated 3 times
	return json.Marshal(RideHistory{Rides: []RideSummary{
		{
			Time:        lastRide.Time,
			Destination: lastRide.Destination,
		},
		{
			Time:        lastRide.Time,
			Destination: lastRide.Origin,
		},
		{
			Time:        lastRide.Time,
			Destination: lastRide.Destination,
		},
	}})
}
