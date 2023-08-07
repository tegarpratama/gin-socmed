package dto

type ResponseParams struct {
	StatusCode int
	Message    string
	Paginate   *Paginate
	Data       any
}

type FilterParams struct {
	Page   int
	Limit  int
	Offset int
	Search string
}
