package tdproto

import (
	"encoding/json"
	"log"
	"strings"
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
	JID() *JID
}

type JID struct {
	val string
}

func NewJID(val string) *JID { return &JID{val} }

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

func (jid JID) IsDirect() bool  { return strings.HasPrefix(jid.val, ContactPrefix) }
func (jid JID) IsGroup() bool   { return strings.HasPrefix(jid.val, GroupPrefix) }
func (jid JID) IsTask() bool    { return strings.HasPrefix(jid.val, TaskPrefix) }
func (jid JID) IsSection() bool { return strings.HasPrefix(jid.val, ContactSectionPrefix) }

func (jid JID) Empty() bool          { return jid.val == "" }
func (jid JID) Equal(other JID) bool { return jid.val == other.val }
func (jid JID) JID() *JID            { return &jid }
func (jid JID) String() string       { return jid.val }
func (jid JID) Valid() bool          { return jid.Uid() != "" }
func (jid JID) Value() string        { return jid.val }

func (jid JID) Uid() string {
	if jid.Empty() {
		return ""
	}
	bits := strings.SplitN(jid.val, "-", 2)
	if len(bits) != 2 || !ValidUid(bits[1]) {
		return ""
	}
	return bits[1]
}

func (jid JID) MarshalJSON() ([]byte, error) {
	return json.Marshal(jid.val)
}

func (jid *JID) UnmarshalJSON(data []byte) error {
	var val string
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	jid.val = val
	return nil
}
