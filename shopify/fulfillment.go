package shopify

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type FulfillmentService service

type Fulfillment struct {
	ID             int64      `json:"id"`
	LocationID     int64      `json:"location_id"`
	NotifyCustomer bool       `json:"notify_customer"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

type FulfillmentRequest struct {
	*Fulfillment `json:"fulfillment"`
}

func (s *FulfillmentService) Create(ctx context.Context, orderID int64, fulfillment *Fulfillment) (*Fulfillment, *http.Response, error) {
	req, err := s.client.NewRequest(
		"POST",
		fmt.Sprintf("/admin/orders/%d/fulfillments.json", orderID),
		&FulfillmentRequest{fulfillment},
	)
	if err != nil {
		return nil, nil, err
	}
	wrapper := &FulfillmentRequest{fulfillment}
	resp, err := s.client.Do(ctx, req, &wrapper)
	if err != nil {
		return nil, resp, err
	}

	return wrapper.Fulfillment, resp, nil
}