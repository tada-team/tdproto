// eXtra fast shortcuts
package tdevents

import "bytes"

func XClientPing(confirmId string) []byte {
	return xConcat( `{"event":"client.ping","confirm_id":"`, confirmId,`"}`)
}

func XClientConfirm(confirmId string) []byte {
	return xConcat(`{"event":"client.confirm","params":{"confirm_id":"`, confirmId, `"}}`)
}

func XServerConfirm(confirmId string) []byte {
	return xConcat(`{"event":"server.confirm","params":{"confirm_id":"`, confirmId, `"}}`)
}

func xConcat(begin, mid, end string) []byte {
	var b bytes.Buffer
	b.Grow(len(begin) + len(mid) + len(end))
	b.WriteString(begin)
	b.WriteString(mid)
	b.WriteString(end)
	return b.Bytes()
}
