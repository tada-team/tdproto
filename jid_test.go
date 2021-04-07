package tdproto

import (
	"encoding/json"
	"testing"
)

func TestNewJID(t *testing.T) {
	t.Run("marshall", func(t *testing.T) {
		res := new(struct {
			Val JID `json:"val,omitempty"`
		})

		b,  err := json.Marshal(res)
		if err != nil {
			t.Fatal(err)
		}

		if string(b) != "{}" {
			t.Error("must be empty, got:", string(b))
		}
	})

	t.Run("unmarshall", func(t *testing.T) {
		res := new(struct {
			Val JID `json:"val"`
		})

		data := `{"val": "t-63150419-c5c7-40fb-8aa2-ca6b613ca5a0"}`
		if err := json.Unmarshal([]byte(data), res); err != nil {
			t.Fatal(err)
		}

		if !res.Val.IsTask() {
			t.Error("must be task")
		}

		if res.Val.Uid() != "63150419-c5c7-40fb-8aa2-ca6b613ca5a0" {
			t.Error("invalid uid:", res.Val.Uid())
		}
	})
}
