package domains

type User struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	IDToken string `json:"id_token"`
}
