package tdproto

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"

	uuid "github.com/satori/go.uuid"
)

// BenchmarkSerialization/json-12         	  223381	      5204 ns/op	    1834 B/op	      21 allocs/op
// BenchmarkSerialization/jsoniter-12     	  394774	      3073 ns/op	    1546 B/op	      18 allocs/op
func BenchmarkSerialization(b *testing.B) {
	b.Run("json", func(b *testing.B) {
		b.ReportAllocs()
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

	b.Run("jsoniter", func(b *testing.B) {
		b.ReportAllocs()
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
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
