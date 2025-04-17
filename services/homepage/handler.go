package function

import (
	"encoding/json"
	"errors"
	"fmt"
)

// FlowInput is the input to the flow function handler function in the template project
func ExecFlow(request FlowInput) ([]byte, error) {

	// Validate the input request object of user_id and origin
	if request.Args.UserID == nil {
		return nil, errors.New("user_id is required")
	}
	if request.Args.Origin == nil {
		return nil, errors.New("origin is required")
	}

	// Validate the input request object of children response of ride_recommendation
	var recommendation Recommendation
	recommendationResponse, ok := request.Children["ride_recommendation"]
	if !ok || recommendationResponse == nil {
		return nil, errors.New("response of ride_recommendation is required to process")
	}

	// Unmarshal the response of ride_recommendation to recommendation object
	err := json.Unmarshal(recommendationResponse.Data, &recommendation)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error in unmarshalling: %s", err))
	}
	fmt.Printf("%+v\n", recommendation)

	// Validate the input request object of children response of user_info_of_passenger
	var userInfo UserInfo
	userInfoResponse, ok := request.Children["user_info_of_passenger"]
	if !ok || userInfoResponse == nil {
		return nil, errors.New("response of user_info_of_passenger is required to process")
	}

	// Unmarshal the response of user_info_of_passenger to userInfo object
	err = json.Unmarshal(userInfoResponse.Data, &userInfo)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error in unmarshalling: %s", err))
	}
	fmt.Printf("%+v\n", userInfo)

	// Prepare the response object to be returned by the flow function
	var homepageDetails HomepageDetails
	if recommendation.Type != RecommendationNothing {
		homepageDetails = HomepageDetails{
			IsAnythingRecommended:    true,
			RecommendationBannerText: &recommendation.BannerText,
			RecommendationType:       &recommendation.Type,
		}
	}

	// Fill the user details in the response object to be returned by the flow function
	homepageDetails.UserCurrentLocation = userInfo.CurrentAddressLocation
	homepageDetails.UserAddresses = userInfo.Addresses
	fmt.Printf("%+v\n", homepageDetails)

	return json.Marshal(homepageDetails)
}
