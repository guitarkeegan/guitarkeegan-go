package forms

type ContactForm struct {
	Email   string `json:"email"`
	Message string `json:"message"`
}
