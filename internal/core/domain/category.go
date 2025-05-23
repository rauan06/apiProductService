package domain

type CategoryMain struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Deleted   bool   `json:"deleted"`
	ID        string `json:"id"`
	Name      string `json:"name"`
}

type CategoryListRep struct {
	PaginationInfo PaginationInfoSt `json:"pagination_info"`
	Results        []CategoryMain   `json:"results"`
}

func (c *CategoryMain) IsValid() bool {
	if c.CreatedAt == "" || c.UpdatedAt == "" || c.Deleted || c.ID == "" || c.Name == "" {
		return false
	}

	return true
}
