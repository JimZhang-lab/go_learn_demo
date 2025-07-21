package dto

type CommonIDDTO struct {
	Id int64 `json:"id" form:"id" uri:"id"`
}

// 分页对应 DTO

type PaginateDTO struct {
	Page  int `json:"page,omitempty" form:"page"`
	Limit int `json:"limit,omitempty" form:"limit"`
}

func (m *PaginateDTO) GetPage() int {
	if m.Page <= 0 {
		return 1
	}
	return m.Page
}
func (m *PaginateDTO) GetLimit() int {
	if m.Limit <= 0 {
		return 10
	}
	return m.Limit
}
