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

	// Check if the user_info_of_passenger is provided
	var userInfo UserInfo
	userInfoResponse, ok := request.Children["user_info_of_passenger"]
	if !ok || userInfoResponse == nil {
		return nil, errors.New("response of user_info_of_passenger is required to process")
	}

	// Unmarshal the response of user_info_of_passenger to userInfo
	err := json.Unmarshal(userInfoResponse.Data, &userInfo)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error in unmarshalling: %s", err))
	}
	fmt.Printf("%+v\n", userInfo)

	// Return the response of the function handler
	return json.Marshal(FriendsInfo{
		NumberOfFriends: userInfo.PhoneNumber,
	})
}
