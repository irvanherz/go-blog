package model

//Response struct
type Response struct {
	Status   int         `json:"status,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	PageInfo *PageInfo   `json:"pageInfo,omitempty"`
	Error    interface{} `json:"error,omitempty"`
}

//PageInfo struct
type PageInfo struct {
	TotalItems   int64 `json:"totalItems,omitempty"`
	TotalPages   int64 `json:"totalPages,omitempty"`
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`
	CurrentPage  int64 `json:"currentPage,omitempty"`
}
