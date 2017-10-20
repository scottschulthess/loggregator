// Code generated by protoc-gen-gogo.
// source: log_message.proto
// DO NOT EDIT!

package endtoend

import proto "github.com/gogo/protobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type LogMessage_MessageType int32

const (
	LogMessage_OUT LogMessage_MessageType = 1
	LogMessage_ERR LogMessage_MessageType = 2
)

var LogMessage_MessageType_name = map[int32]string{
	1: "OUT",
	2: "ERR",
}
var LogMessage_MessageType_value = map[string]int32{
	"OUT": 1,
	"ERR": 2,
}

func (x LogMessage_MessageType) Enum() *LogMessage_MessageType {
	p := new(LogMessage_MessageType)
	*p = x
	return p
}
func (x LogMessage_MessageType) String() string {
	return proto.EnumName(LogMessage_MessageType_name, int32(x))
}
func (x LogMessage_MessageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *LogMessage_MessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LogMessage_MessageType_value, data, "LogMessage_MessageType")
	if err != nil {
		return err
	}
	*x = LogMessage_MessageType(value)
	return nil
}

type LogMessage struct {
	Message          []byte                  `protobuf:"bytes,1,req,name=message" json:"message,omitempty"`
	MessageType      *LogMessage_MessageType `protobuf:"varint,2,req,name=message_type,enum=logmessage.LogMessage_MessageType" json:"message_type,omitempty"`
	Timestamp        *int64                  `protobuf:"zigzag64,3,req,name=timestamp" json:"timestamp,omitempty"`
	AppId            *string                 `protobuf:"bytes,4,req,name=app_id" json:"app_id,omitempty"`
	SourceId         *string                 `protobuf:"bytes,6,opt,name=source_id" json:"source_id,omitempty"`
	DrainUrls        []string                `protobuf:"bytes,7,rep,name=drain_urls" json:"drain_urls,omitempty"`
	SourceName       *string                 `protobuf:"bytes,8,opt,name=source_name" json:"source_name,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *LogMessage) Reset()         { *m = LogMessage{} }
func (m *LogMessage) String() string { return proto.CompactTextString(m) }
func (*LogMessage) ProtoMessage()    {}

func (m *LogMessage) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *LogMessage) GetMessageType() LogMessage_MessageType {
	if m != nil && m.MessageType != nil {
		return *m.MessageType
	}
	return 0
}

func (m *LogMessage) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *LogMessage) GetAppId() string {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return ""
}

func (m *LogMessage) GetSourceId() string {
	if m != nil && m.SourceId != nil {
		return *m.SourceId
	}
	return ""
}

func (m *LogMessage) GetDrainUrls() []string {
	if m != nil {
		return m.DrainUrls
	}
	return nil
}

func (m *LogMessage) GetSourceName() string {
	if m != nil && m.SourceName != nil {
		return *m.SourceName
	}
	return ""
}

type LogEnvelope struct {
	RoutingKey       *string     `protobuf:"bytes,1,req,name=routing_key" json:"routing_key,omitempty"`
	Signature        []byte      `protobuf:"bytes,2,req,name=signature" json:"signature,omitempty"`
	LogMessage       *LogMessage `protobuf:"bytes,3,req,name=log_message" json:"log_message,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *LogEnvelope) Reset()         { *m = LogEnvelope{} }
func (m *LogEnvelope) String() string { return proto.CompactTextString(m) }
func (*LogEnvelope) ProtoMessage()    {}

func (m *LogEnvelope) GetRoutingKey() string {
	if m != nil && m.RoutingKey != nil {
		return *m.RoutingKey
	}
	return ""
}

func (m *LogEnvelope) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *LogEnvelope) GetLogMessage() *LogMessage {
	if m != nil {
		return m.LogMessage
	}
	return nil
}

func init() {
	proto.RegisterEnum("logmessage.LogMessage_MessageType", LogMessage_MessageType_name, LogMessage_MessageType_value)
}
