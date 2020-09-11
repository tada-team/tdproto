package tdproto

import (
	"encoding/json"
	"testing"

	uuid "github.com/satori/go.uuid"
)

func BenchmarkSerialization(b *testing.B) {
	b.ReportAllocs()

	b.Run("json", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			event := NewClientMessageUpdated(ClientMessageUpdatedParams{
				Content:   MessageContent{Text: "test 123"},
				MessageId: uuid.NewV4().String(),
				Important: true,
			})

			data, err := json.Marshal(event)
			if err != nil {
				b.Fatal(err)
			}

			v := new(ClientMessageUpdated)
			if err := json.Unmarshal(data, v); err != nil {
				b.Fatal(err)
			}

			if v.Params.MessageId != event.Params.MessageId {
				b.Fatal("message id mismatchs")
			}
		}
	})
}
