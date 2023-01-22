package publisher

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"time"
)

var NC, _ = nats.Connect(nats.DefaultURL)

func Publish() {

	for i := 0; ; i++ {
		new_order := generateOneOrder()
		json_order, err := json.Marshal(new_order)
		if err != nil {
			fmt.Printf("ERROR: can't marshal order to json: %v\n", err)
			return
		}

		err = NC.Publish("order", json_order)
		if err != nil {
			fmt.Printf("ERROR: can't publish message: %v\n", err)
			return
		}

		time.Sleep(10 * time.Second)
		if i%10 == 0 && i != 0 {
			time.Sleep(10 * time.Minute)
		}
	}
}
