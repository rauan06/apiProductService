package domain

type PaginationInfoSt struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"page_size"`
}

type ListParamsSt struct {
	Page     int64    `json:"page"`
	PageSize int64    `json:"page_size"`
	Sort     []string `json:"sort"`
}
