package publisher

import (
	"time"
	"wb_l0/domain"

	"github.com/google/uuid"
)

func generateOneOrder() *domain.Order {
	id := uuid.New()

	order := domain.Order{
		OrderUID:    id.String(),
		TrackNumber: "111-111",
		Entry:       "aaa",
		Delivery: domain.Delivery{
			Name:    "bbb",
			Phone:   "+0101",
			Zip:     "dd",
			City:    "dd",
			Address: "t",
			Region:  "ttt",
			Email:   "hhhh",
		},
		Payment: domain.Payment{
			Transaction:  "yyy",
			RequestID:    "yyy",
			Currency:     "yyy",
			Provider:     "yyy",
			Amount:       1,
			PaymentDt:    1,
			Bank:         "yyy",
			DeliveryCost: 1,
			GoodsTotal:   1,
			CustomFee:    1,
		},
		Items: []domain.Item{
			{
				ChrtID:      1,
				TrackNumber: "aaa",
				Price:       1,
				Rid:         "bbb",
				Name:        "ccc",
				Sale:        1,
				Size:        "ddd",
				TotalPrice:  1,
				NmID:        21454,
				Brand:       "aaa",
				Status:      1234,
			},
		},
		Locale:            "ddd",
		InternalSignature: "ccc",
		CustomerID:        "aaa",
		DeliveryService:   "ddd",
		Shardkey:          "ccc",
		SmID:              123,
		DateCreated:       time.Now(),
		OofShard:          "ddd",
	}

	return &order
}
