package models

// Error model.
type Error struct {
	Field    string   `json:"field"`
	Messages []string `json:"messages"`
}
