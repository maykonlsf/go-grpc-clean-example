package author

type AddRequest struct {
	firstName string
	lastName  string
	pseudonym string
}

func NewAddRequest() *AddRequest {
	return &AddRequest{}
}

func (r *AddRequest) FirstName(s string) *AddRequest {
	r.firstName = s
	return r
}

func (r *AddRequest) LastName(s string) *AddRequest {
	r.lastName = s
	return r
}

func (r *AddRequest) Pseudonym(s string) *AddRequest {
	r.pseudonym = s
	return r
}
