package author

type Author struct {
	Id        uint32 `json:"id,omitempty"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
}
