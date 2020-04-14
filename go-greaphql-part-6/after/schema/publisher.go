package schema

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
	"github.com/graphql-go/graphql"
	"github.com/gufranmirza/go-graphql-swapi/go-greaphql-part-6/after/subscriptions"
)

func HumanPublisher() {
	subscriptions.Subscribers.Range(func(key, value interface{}) bool {
		subscriber, ok := value.(*subscriptions.Subscriber)
		if !ok {
			return true
		}
		payload := graphql.Do(graphql.Params{
			Schema:        Schema,
			RequestString: subscriber.RequestString,
		})
		message, err := json.Marshal(map[string]interface{}{
			"type":    "data",
			"id":      subscriber.OperationID,
			"payload": payload,
		})
		if err != nil {
			log.Printf("failed to marshal message: %v", err)
			return true
		}
		if err := subscriber.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
			if err == websocket.ErrCloseSent {
				subscriptions.Subscribers.Delete(key)
				return true
			}
			log.Printf("failed to write to ws connection: %v", err)
			return true
		}
		return true
	})
}
