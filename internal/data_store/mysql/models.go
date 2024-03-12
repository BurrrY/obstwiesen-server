package mysql

type GetMeadow struct {
	ID   string  `json:"id"`
	Name string  `json:"name"`
	Area *string `json:"area"`
}
