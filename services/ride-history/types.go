package function

type FlowOutput struct {
	Data     []byte `json:"data"`
	Function string `json:"function"`
}

type FlowInput struct {
	Args     Input                  `json:"args"`
	Children map[string]*FlowOutput `json:"children,omitempty"`
}
