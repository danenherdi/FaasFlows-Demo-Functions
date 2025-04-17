package function

// Message is the output of your flow function
type Message struct {
	Text string `json:"text"`
}

// Input is the argument of your flow function
type Input struct {
	Name *string `json:"name"`
}
