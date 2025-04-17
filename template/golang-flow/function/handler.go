package function

import (
	"encoding/json"
	"errors"
	"fmt"
)

// FlowInput is the input to the flow function handler function in the template project
func ExecFlow(request FlowInput) ([]byte, error) {
	// Validate input
	if request.Args.Name == nil {
		return nil, errors.New("name is required")
	}

	// Create message and return as JSON string
	message := Message{Text: fmt.Sprintf("Hello %s!", *request.Args.Name)}
	return json.Marshal(message)
}
