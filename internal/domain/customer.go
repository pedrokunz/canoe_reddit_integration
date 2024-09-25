package domain

type Customer struct {
	ID           string        `json:"id"`
	Name		 string        `json:"name"`
	Integrations []Integration `json:"integrations"`
}
