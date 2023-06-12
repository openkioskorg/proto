package proto

import "encoding/json"

// Result of a command or user interaction with the components
type Event struct {
	// Event name
	Event string `json:"event"`
	// Data related to the event whose type is component dependent
	Data json.RawMessage `json:"data,omitempty"`
}

// Instruction for components
type Cmd struct {
	// Command name
	Cmd string `json:"cmd"`
	// Data related to the command
	Data json.RawMessage `json:"data,omitempty"`
}
