package shopify

import (
	"context"
	"fmt"
	"net/http"
)

type OrderService service

type Order struct {
	ID             int64           `json:"id,omitempty"`
	Name           string          `json:"name,omitempty"`
	Email          string          `json:"email,omitempty"`
	ContactEmail   string          `json:"contact_email,omitempty"`
	Token          string          `json:"token,omitempty"`
	CheckoutToken  string          `json:"checkout_token,omitempty"`
	CustomerLocale string          `json:"customer_locale,omitempty"`
	OrderNumber    int             `json:"order_number,omitempty"`
	NoteAttributes []NoteAttribute `json:"note_attributes,omitempty"`
	LineItems      []LineItem      `json:"line_items,omitempty"`
	Customer       *struct {
		Name string `json:"name,omitempty"`
	} `json:"customer,omitempty"`
}

type NoteAttribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type OrderRequest struct {
	*Order `json:"order"`
}

func (s *OrderService) Get(ctx context.Context, orderID int64) (*Order, *http.Response, error) {
	req, err := s.client.NewRequest(
		"GET",
		fmt.Sprintf("/admin/orders/%d.json", orderID),
		nil,
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
