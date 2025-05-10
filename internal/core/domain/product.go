package domain

type ProductMain struct {
	CreatedAt   string       `json:"created_at"`
	UpdatedAt   string       `json:"updated_at"`
	Deleted     bool         `json:"deleted"`
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	CategoryID  string       `json:"category_id"`
	Category    CategoryMain `json:"category"`
}

type ProductListRep struct {
	PaginationInfo PaginationInfoSt `json:"pagination_info"`
	Results        []ProductMain    `json:"results"`
}
