package shopify

import (
	"context"
	"fmt"
	"net/http"
)

type OrderService service

type Order struct {
	ID             int64           `json:"id,omitempty"`
	NoteAttributes []NoteAttribute `json:"note_attributes,omitempty"`
}

type NoteAttribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type OrderRequest struct {
	*Order `json:"order"`
}

func (s *OrderService) AddNoteAttributes(ctx context.Context, orderID int64, attributes map[string]string) (*Order, *http.Response, error) {
	var attrs []NoteAttribute
	for k, v := range attributes {
		attrs = append(attrs, NoteAttribute{
			Name:  k,
			Value: v,
		})
	}
	req, err := s.client.NewRequest(
		"PUT",
		fmt.Sprintf("/admin/orders/%d.json", orderID),
		&OrderRequest{&Order{
			ID:             orderID,
			NoteAttributes: attrs,
		}},
	)
	if err != nil {
		return nil, nil, err
	}
	wrapper := &OrderRequest{&Order{}}
	resp, err := s.client.Do(ctx, req, &wrapper)
	if err != nil {
		return nil, resp, err
	}

	return wrapper.Order, resp, nil
}
