package function

import (
	"encoding/json"
	"errors"
	"fmt"
)

// FlowInput is the input to the flow function handler function in the template project
func ExecFlow(request FlowInput) ([]byte, error) {

	// Validate the input request
	if request.Args.UserID == nil {
		return nil, errors.New("user_id is required")
	}
	if request.Args.Origin == nil {
		return nil, errors.New("origin is required")
	}

	// Unmarshal the input request to the required struct to get last ride info
	var lastRide Ride
	lastRideResponse, ok := request.Children["last_ride_of_passenger"]
	if !ok || lastRideResponse == nil {
		return nil, errors.New("response of last_ride_of_passenger is required to process")
	}
	// Unmarshal lastRide from child function
	var lastRideWrapper map[string]json.RawMessage
	err := json.Unmarshal(lastRideResponse.Data, &lastRideWrapper)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal lastRide wrapper: %w", err)
	}
	err = json.Unmarshal(lastRideWrapper["response"], &lastRide)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal lastRide: %w", err)
	}
	fmt.Printf("%+v\n", lastRide)

	// Unmarshal the input request to the required struct to get user info
	var userInfo UserInfo
	userInfoResponse, ok := request.Children["user_info_of_passenger"]
	if !ok || userInfoResponse == nil {
		return nil, errors.New("response of user_info_of_passenger is required to process")
	}
	// Unmarshal userInfo from child function
	var userInfoWrapper map[string]json.RawMessage
	err = json.Unmarshal(userInfoResponse.Data, &userInfoWrapper)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal userInfo wrapper: %w", err)
	}
	err = json.Unmarshal(userInfoWrapper["response"], &userInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal userInfo: %w", err)
	}
	fmt.Printf("%+v\n", userInfo)

	var recommendation Recommendation
	// Repeat recommendation
	if request.Args.Origin.Lat == lastRide.Origin.Lat && request.Args.Origin.Lon == lastRide.Origin.Lon {
		recommendation = Recommendation{
			Type:           RecommendationRepeat,
			Recommendation: &lastRide.Destination,
			BannerText:     fmt.Sprintf("Dear %s, Here is your repeat recommendation.", userInfo.FirstName),
		}
		return json.Marshal(recommendation)
	}

	// Reverse recommendation
	if request.Args.Origin.Lat == lastRide.Destination.Lat && request.Args.Origin.Lon == lastRide.Destination.Lon {
		recommendation = Recommendation{
			Type:           RecommendationReverse,
			Recommendation: &lastRide.Origin,
			BannerText:     fmt.Sprintf("Dear %s, Here is your reverse recommendation.", userInfo.FirstName),
		}
		return json.Marshal(recommendation)
	}

	// No recommendation
	recommendation = Recommendation{
		Type:       RecommendationNothing,
		BannerText: fmt.Sprintf("Dear %s, There is no recommendation.", userInfo.FirstName),
	}

	return json.Marshal(recommendation)
}
