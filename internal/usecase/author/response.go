package author

type AddResponse struct {
	id string
}

func (r *AddResponse) GetID() string {
	return r.id
}
