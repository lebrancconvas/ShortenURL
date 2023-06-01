package forms

import "time"

type URL struct {
	ID						int 				`json:"id"`
	OriginalURL 	string 			`json:"original_url"`
	ShortURL 			string 			`json:"short_url"`
	CreatedAt 		time.Time 	`json:"created_at"`
	UpdatedAt 		time.Time 	`json:"updated_at"`
}
