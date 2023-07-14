package peerpb

import (
	"fmt"
	"strings"
)

//MessageSend

func (x *MessageSendReq) SetHeader(header *RequestHeader) {
	x.Header = header
}

func (x *MessageSendResp) SetHeader(header *ResponseHeader) {
	x.Header = header
}

//GetConvMessageSeq

func (x *GetConvMessageSeqReq) SetHeader(header *RequestHeader) {
	x.Header = header
}

func (x *GetConvMessageSeqResp) SetHeader(header *ResponseHeader) {
	x.Header = header
}

//SyncMessage

func (x *SyncMessageReq) SetHeader(header *RequestHeader) {
	x.Header = header
}

func (x *SyncMessageResp) SetHeader(header *ResponseHeader) {
	x.Header = header
}

const (
	SingleChatConversationSeparator = "&"
)

func GetSingleChatConversationId(id1 string, id2 string) string {
	if id1 < id2 {
		return fmt.Sprintf("%s%s%s", id1, SingleChatConversationSeparator, id2)
	} else {
		return fmt.Sprintf("%s%s%s", id2, SingleChatConversationSeparator, id1)
	}
}

func ParseSingleChatConversationId(convId string) (id1 string, id2 string) {
	split := strings.Split(convId, SingleChatConversationSeparator)
	if len(split) != 2 {
		return
	}
	id1 = split[0]
	id2 = split[1]
	return
}

func GetSingleChatOtherId(convId string, id1 string) (id2 string) {
	split := strings.Split(convId, SingleChatConversationSeparator)
	if len(split) != 2 {
		return
	}
	if split[0] == id1 {
		id2 = split[1]
	} else {
		id2 = split[0]
	}
	return
}
