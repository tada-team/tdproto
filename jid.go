package tdproto

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/fxamacker/cbor/v2"
	"github.com/vmihailenco/msgpack/v5"
)

const (
	ContactPrefix        = "d-"
	GroupPrefix          = "g-"
	TaskPrefix           = "t-"
	ContactSectionPrefix = "sd-"
	GroupSectionPrefix   = "sg-"
	TaskSectionPrefix    = "st-"
)

type HasJid interface {
	JID() JID
}

type JID string

func (jid JID) ChatType() ChatType {
	switch {
	case jid.IsDirect():
		return DirectChatType
	case jid.IsGroup():
		return GroupChatType
	case jid.IsTask():
		return TaskChatType
	default:
		log.Fatalf("invalid chat type: %s", jid)
		return ""
	}
}

func (jid JID) IsDirect() bool  { return strings.HasPrefix(jid.String(), ContactPrefix) }
func (jid JID) IsGroup() bool   { return strings.HasPrefix(jid.String(), GroupPrefix) }
func (jid JID) IsTask() bool    { return strings.HasPrefix(jid.String(), TaskPrefix) }
func (jid JID) IsSection() bool { return strings.HasPrefix(jid.String(), ContactSectionPrefix) }

func (jid JID) Empty() bool          { return jid.String() == "" }
func (jid JID) JID() JID             { return jid }
func (jid JID) String() string       { return string(jid) }
func (jid JID) Valid() bool          { return jid.Uid() != "" }
func (jid JID) Value() string        { return jid.String() }

func (jid JID) Uid() string {
	if jid.Empty() {
		return ""
	}
	bits := strings.SplitN(jid.String(), "-", 2)
	if len(bits) != 2 || !ValidUid(bits[1]) {
		return ""
	}
	return bits[1]
}

func (jid JID) MarshalJSON() ([]byte, error) {
	return json.Marshal(jid.String())
}

func (jid *JID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*jid = JID(s)
	return nil
}

func (jid JID) MarshalMsgpack() ([]byte, error) {
	return msgpack.Marshal(jid.String())
}

func (jid *JID) UnmarshalMsgpack(data []byte) error {
	var s string
	if err := msgpack.Unmarshal(data, &s); err != nil {
		return err
	}
	* jid = JID(s)
	return nil
}

func (jid JID) MarshalCBOR() ([]byte, error) {
	return cbor.Marshal(jid.String())
}

func (jid *JID) UnmarshalCBOR(data []byte) error {
	var s string
	if err := cbor.Unmarshal(data, &s); err != nil {
		return err
	}
	* jid = JID(s)
	return nil
}

func (jid *JID) UnmarshalText(data []byte) error {
	*jid = JID(data)
	return nil
}
