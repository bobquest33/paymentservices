package ogone_alias_direct_link

type Endpoint struct {
}

func (ep *Endpoint) Response(r *Request) (*Response, error) {
	if err := r.isValid(); err != nil {
		return nil, err
	}

	// FIXME XXX: implementation
	return &Response{"OK", ""}, nil
}