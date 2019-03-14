// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stellarstation/api/v1/transport.proto

package v1 // import "github.com/infostellarinc/go-stellarstation/api/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import monitoring "github.com/infostellarinc/go-stellarstation/api/v1/monitoring"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A type of framing of a binary payload used in satellite communicaation.
type Framing int32

const (
	// No framing done in the API. All payloads are assumed to be pre-framed and ready for
	// transmission to the satellite or API client with no additional processing.
	Framing_BITSTREAM Framing = 0
	// AX.25 (Amateur X.25) framing - https://en.wikipedia.org/wiki/AX.25
	Framing_AX25 Framing = 1
	// No framing or demodulation done in the API. Raw IQ data is sent to the API client with no
	// additional processing.
	Framing_IQ Framing = 2
	// A decoded PNG image frame.
	Framing_IMAGE_PNG Framing = 3
	// A decoded JPEG image.
	Framing_IMAGE_JPEG Framing = 4
	// Completely arbitrary, freeform text contained in a frame.
	Framing_FREE_TEXT_UTF8 Framing = 5
	// A waterfall diagram. This is actually for the whole plan and does not correspond to an
	// individual frame.
	Framing_WATERFALL Framing = 6
)

var Framing_name = map[int32]string{
	0: "BITSTREAM",
	1: "AX25",
	2: "IQ",
	3: "IMAGE_PNG",
	4: "IMAGE_JPEG",
	5: "FREE_TEXT_UTF8",
	6: "WATERFALL",
}
var Framing_value = map[string]int32{
	"BITSTREAM":      0,
	"AX25":           1,
	"IQ":             2,
	"IMAGE_PNG":      3,
	"IMAGE_JPEG":     4,
	"FREE_TEXT_UTF8": 5,
	"WATERFALL":      6,
}

func (x Framing) String() string {
	return proto.EnumName(Framing_name, int32(x))
}
func (Framing) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_transport_e71250fe687613f6, []int{0}
}

// A chunk or frame of telemetry data that has been received from a satellite.
type Telemetry struct {
	// The framing of this telemetry data. If `RAW`, this telemetry will be an arbitrarily sized
	// chunk of the bitstream.
	Framing Framing `protobuf:"varint,1,opt,name=framing,proto3,enum=stellarstation.api.v1.Framing" json:"framing,omitempty"`
	// The payload of this telemetry.
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// The actual downlink frequency, in Hz, used when receiving `data`, including adjustments for
	// doppler shift.
	DownlinkFrequencyHz uint64 `protobuf:"varint,3,opt,name=downlink_frequency_hz,json=downlinkFrequencyHz,proto3" json:"downlink_frequency_hz,omitempty"`
	// Timestamp when the first byte of `data` was received.
	TimeFirstByteReceived *timestamp.Timestamp `protobuf:"bytes,4,opt,name=time_first_byte_received,json=timeFirstByteReceived,proto3" json:"time_first_byte_received,omitempty"`
	// Timestamp when the last byte of `data` was received.
	TimeLastByteReceived *timestamp.Timestamp `protobuf:"bytes,5,opt,name=time_last_byte_received,json=timeLastByteReceived,proto3" json:"time_last_byte_received,omitempty"`
	// The binary header of the telemetry frame, if `framing` is not `RAW`.
	//
	// * AX25 - This is either Address + Control, or Address + Control + PID. The checksum is not
	//          returned.
	FrameHeader          []byte   `protobuf:"bytes,6,opt,name=frame_header,json=frameHeader,proto3" json:"frame_header,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Telemetry) Reset()         { *m = Telemetry{} }
func (m *Telemetry) String() string { return proto.CompactTextString(m) }
func (*Telemetry) ProtoMessage()    {}
func (*Telemetry) Descriptor() ([]byte, []int) {
	return fileDescriptor_transport_e71250fe687613f6, []int{0}
}
func (m *Telemetry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Telemetry.Unmarshal(m, b)
}
func (m *Telemetry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Telemetry.Marshal(b, m, deterministic)
}
func (dst *Telemetry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Telemetry.Merge(dst, src)
}
func (m *Telemetry) XXX_Size() int {
	return xxx_messageInfo_Telemetry.Size(m)
}
func (m *Telemetry) XXX_DiscardUnknown() {
	xxx_messageInfo_Telemetry.DiscardUnknown(m)
}

var xxx_messageInfo_Telemetry proto.InternalMessageInfo

func (m *Telemetry) GetFraming() Framing {
	if m != nil {
		return m.Framing
	}
	return Framing_BITSTREAM
}

func (m *Telemetry) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Telemetry) GetDownlinkFrequencyHz() uint64 {
	if m != nil {
		return m.DownlinkFrequencyHz
	}
	return 0
}

func (m *Telemetry) GetTimeFirstByteReceived() *timestamp.Timestamp {
	if m != nil {
		return m.TimeFirstByteReceived
	}
	return nil
}

func (m *Telemetry) GetTimeLastByteReceived() *timestamp.Timestamp {
	if m != nil {
		return m.TimeLastByteReceived
	}
	return nil
}

func (m *Telemetry) GetFrameHeader() []byte {
	if m != nil {
		return m.FrameHeader
	}
	return nil
}

// An event that occurred while processing the stream. A `StreamEvent` will have one of several
// types of event payloads corresponding to event types. Many of these payloads will be empty,
// depending on the context of an event, but are still returned as messages to allow future
// extension.
type StreamEvent struct {
	// The ID of the request this event corresponds to, copied from
	// `SatelliteStreamRequest.request_id` when present. If the event doesn't correspond to a request,
	// this is always empty.
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// A stream event payload.
	//
	// Types that are valid to be assigned to Event:
	//	*StreamEvent_CommandSent
	//	*StreamEvent_PlanMonitoringEvent
	Event                isStreamEvent_Event `protobuf_oneof:"Event"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *StreamEvent) Reset()         { *m = StreamEvent{} }
func (m *StreamEvent) String() string { return proto.CompactTextString(m) }
func (*StreamEvent) ProtoMessage()    {}
func (*StreamEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_transport_e71250fe687613f6, []int{1}
}
func (m *StreamEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamEvent.Unmarshal(m, b)
}
func (m *StreamEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamEvent.Marshal(b, m, deterministic)
}
func (dst *StreamEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamEvent.Merge(dst, src)
}
func (m *StreamEvent) XXX_Size() int {
	return xxx_messageInfo_StreamEvent.Size(m)
}
func (m *StreamEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamEvent.DiscardUnknown(m)
}

var xxx_messageInfo_StreamEvent proto.InternalMessageInfo

func (m *StreamEvent) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

type isStreamEvent_Event interface {
	isStreamEvent_Event()
}

type StreamEvent_CommandSent struct {
	CommandSent *StreamEvent_CommandSentFromGroundStation `protobuf:"bytes,2,opt,name=command_sent,json=commandSent,proto3,oneof"`
}

type StreamEvent_PlanMonitoringEvent struct {
	PlanMonitoringEvent *PlanMonitoringEvent `protobuf:"bytes,3,opt,name=plan_monitoring_event,json=planMonitoringEvent,proto3,oneof"`
}

func (*StreamEvent_CommandSent) isStreamEvent_Event() {}

func (*StreamEvent_PlanMonitoringEvent) isStreamEvent_Event() {}

func (m *StreamEvent) GetEvent() isStreamEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *StreamEvent) GetCommandSent() *StreamEvent_CommandSentFromGroundStation {
	if x, ok := m.GetEvent().(*StreamEvent_CommandSent); ok {
		return x.CommandSent
	}
	return nil
}

func (m *StreamEvent) GetPlanMonitoringEvent() *PlanMonitoringEvent {
	if x, ok := m.GetEvent().(*StreamEvent_PlanMonitoringEvent); ok {
		return x.PlanMonitoringEvent
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StreamEvent) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _StreamEvent_OneofMarshaler, _StreamEvent_OneofUnmarshaler, _StreamEvent_OneofSizer, []interface{}{
		(*StreamEvent_CommandSent)(nil),
		(*StreamEvent_PlanMonitoringEvent)(nil),
	}
}

func _StreamEvent_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*StreamEvent)
	// Event
	switch x := m.Event.(type) {
	case *StreamEvent_CommandSent:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CommandSent); err != nil {
			return err
		}
	case *StreamEvent_PlanMonitoringEvent:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PlanMonitoringEvent); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("StreamEvent.Event has unexpected type %T", x)
	}
	return nil
}

func _StreamEvent_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*StreamEvent)
	switch tag {
	case 2: // Event.command_sent
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamEvent_CommandSentFromGroundStation)
		err := b.DecodeMessage(msg)
		m.Event = &StreamEvent_CommandSent{msg}
		return true, err
	case 3: // Event.plan_monitoring_event
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PlanMonitoringEvent)
		err := b.DecodeMessage(msg)
		m.Event = &StreamEvent_PlanMonitoringEvent{msg}
		return true, err
	default:
		return false, nil
	}
}

func _StreamEvent_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*StreamEvent)
	// Event
	switch x := m.Event.(type) {
	case *StreamEvent_CommandSent:
		s := proto.Size(x.CommandSent)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StreamEvent_PlanMonitoringEvent:
		s := proto.Size(x.PlanMonitoringEvent)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// An event indicating the commands in the request were sent by the ground station through its
// radio.
type StreamEvent_CommandSentFromGroundStation struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamEvent_CommandSentFromGroundStation) Reset() {
	*m = StreamEvent_CommandSentFromGroundStation{}
}
func (m *StreamEvent_CommandSentFromGroundStation) String() string { return proto.CompactTextString(m) }
func (*StreamEvent_CommandSentFromGroundStation) ProtoMessage()    {}
func (*StreamEvent_CommandSentFromGroundStation) Descriptor() ([]byte, []int) {
	return fileDescriptor_transport_e71250fe687613f6, []int{1, 0}
}
func (m *StreamEvent_CommandSentFromGroundStation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamEvent_CommandSentFromGroundStation.Unmarshal(m, b)
}
func (m *StreamEvent_CommandSentFromGroundStation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamEvent_CommandSentFromGroundStation.Marshal(b, m, deterministic)
}
func (dst *StreamEvent_CommandSentFromGroundStation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamEvent_CommandSentFromGroundStation.Merge(dst, src)
}
func (m *StreamEvent_CommandSentFromGroundStation) XXX_Size() int {
	return xxx_messageInfo_StreamEvent_CommandSentFromGroundStation.Size(m)
}
func (m *StreamEvent_CommandSentFromGroundStation) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamEvent_CommandSentFromGroundStation.DiscardUnknown(m)
}

var xxx_messageInfo_StreamEvent_CommandSentFromGroundStation proto.InternalMessageInfo

// A monitoring event that occurred during the execution of the plan. Information about the current
// configuration of the ground station and state of components is returned to provide information
// that can help to troubleshoot issues with the plan.
type PlanMonitoringEvent struct {
	// The ID of the plan being monitored.
	PlanId string `protobuf:"bytes,1,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	// Types that are valid to be assigned to Info:
	//	*PlanMonitoringEvent_GroundStationConfiguration
	//	*PlanMonitoringEvent_GroundStationState
	//	*PlanMonitoringEvent_GroundStationEvent
	Info                 isPlanMonitoringEvent_Info `protobuf_oneof:"Info"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *PlanMonitoringEvent) Reset()         { *m = PlanMonitoringEvent{} }
func (m *PlanMonitoringEvent) String() string { return proto.CompactTextString(m) }
func (*PlanMonitoringEvent) ProtoMessage()    {}
func (*PlanMonitoringEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_transport_e71250fe687613f6, []int{2}
}
func (m *PlanMonitoringEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlanMonitoringEvent.Unmarshal(m, b)
}
func (m *PlanMonitoringEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlanMonitoringEvent.Marshal(b, m, deterministic)
}
func (dst *PlanMonitoringEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlanMonitoringEvent.Merge(dst, src)
}
func (m *PlanMonitoringEvent) XXX_Size() int {
	return xxx_messageInfo_PlanMonitoringEvent.Size(m)
}
func (m *PlanMonitoringEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_PlanMonitoringEvent.DiscardUnknown(m)
}

var xxx_messageInfo_PlanMonitoringEvent proto.InternalMessageInfo

func (m *PlanMonitoringEvent) GetPlanId() string {
	if m != nil {
		return m.PlanId
	}
	return ""
}

type isPlanMonitoringEvent_Info interface {
	isPlanMonitoringEvent_Info()
}

type PlanMonitoringEvent_GroundStationConfiguration struct {
	GroundStationConfiguration *monitoring.GroundStationConfiguration `protobuf:"bytes,2,opt,name=ground_station_configuration,json=groundStationConfiguration,proto3,oneof"`
}

type PlanMonitoringEvent_GroundStationState struct {
	GroundStationState *monitoring.GroundStationState `protobuf:"bytes,3,opt,name=ground_station_state,json=groundStationState,proto3,oneof"`
}

type PlanMonitoringEvent_GroundStationEvent struct {
	GroundStationEvent *monitoring.GroundStationEvent `protobuf:"bytes,4,opt,name=ground_station_event,json=groundStationEvent,proto3,oneof"`
}

func (*PlanMonitoringEvent_GroundStationConfiguration) isPlanMonitoringEvent_Info() {}

func (*PlanMonitoringEvent_GroundStationState) isPlanMonitoringEvent_Info() {}

func (*PlanMonitoringEvent_GroundStationEvent) isPlanMonitoringEvent_Info() {}

func (m *PlanMonitoringEvent) GetInfo() isPlanMonitoringEvent_Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *PlanMonitoringEvent) GetGroundStationConfiguration() *monitoring.GroundStationConfiguration {
	if x, ok := m.GetInfo().(*PlanMonitoringEvent_GroundStationConfiguration); ok {
		return x.GroundStationConfiguration
	}
	return nil
}

func (m *PlanMonitoringEvent) GetGroundStationState() *monitoring.GroundStationState {
	if x, ok := m.GetInfo().(*PlanMonitoringEvent_GroundStationState); ok {
		return x.GroundStationState
	}
	return nil
}

func (m *PlanMonitoringEvent) GetGroundStationEvent() *monitoring.GroundStationEvent {
	if x, ok := m.GetInfo().(*PlanMonitoringEvent_GroundStationEvent); ok {
		return x.GroundStationEvent
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PlanMonitoringEvent) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PlanMonitoringEvent_OneofMarshaler, _PlanMonitoringEvent_OneofUnmarshaler, _PlanMonitoringEvent_OneofSizer, []interface{}{
		(*PlanMonitoringEvent_GroundStationConfiguration)(nil),
		(*PlanMonitoringEvent_GroundStationState)(nil),
		(*PlanMonitoringEvent_GroundStationEvent)(nil),
	}
}

func _PlanMonitoringEvent_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PlanMonitoringEvent)
	// Info
	switch x := m.Info.(type) {
	case *PlanMonitoringEvent_GroundStationConfiguration:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GroundStationConfiguration); err != nil {
			return err
		}
	case *PlanMonitoringEvent_GroundStationState:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GroundStationState); err != nil {
			return err
		}
	case *PlanMonitoringEvent_GroundStationEvent:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GroundStationEvent); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("PlanMonitoringEvent.Info has unexpected type %T", x)
	}
	return nil
}

func _PlanMonitoringEvent_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PlanMonitoringEvent)
	switch tag {
	case 2: // Info.ground_station_configuration
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(monitoring.GroundStationConfiguration)
		err := b.DecodeMessage(msg)
		m.Info = &PlanMonitoringEvent_GroundStationConfiguration{msg}
		return true, err
	case 3: // Info.ground_station_state
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(monitoring.GroundStationState)
		err := b.DecodeMessage(msg)
		m.Info = &PlanMonitoringEvent_GroundStationState{msg}
		return true, err
	case 4: // Info.ground_station_event
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(monitoring.GroundStationEvent)
		err := b.DecodeMessage(msg)
		m.Info = &PlanMonitoringEvent_GroundStationEvent{msg}
		return true, err
	default:
		return false, nil
	}
}

func _PlanMonitoringEvent_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PlanMonitoringEvent)
	// Info
	switch x := m.Info.(type) {
	case *PlanMonitoringEvent_GroundStationConfiguration:
		s := proto.Size(x.GroundStationConfiguration)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PlanMonitoringEvent_GroundStationState:
		s := proto.Size(x.GroundStationState)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PlanMonitoringEvent_GroundStationEvent:
		s := proto.Size(x.GroundStationEvent)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Telemetry)(nil), "stellarstation.api.v1.Telemetry")
	proto.RegisterType((*StreamEvent)(nil), "stellarstation.api.v1.StreamEvent")
	proto.RegisterType((*StreamEvent_CommandSentFromGroundStation)(nil), "stellarstation.api.v1.StreamEvent.CommandSentFromGroundStation")
	proto.RegisterType((*PlanMonitoringEvent)(nil), "stellarstation.api.v1.PlanMonitoringEvent")
	proto.RegisterEnum("stellarstation.api.v1.Framing", Framing_name, Framing_value)
}

func init() {
	proto.RegisterFile("stellarstation/api/v1/transport.proto", fileDescriptor_transport_e71250fe687613f6)
}

var fileDescriptor_transport_e71250fe687613f6 = []byte{
	// 671 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x6e, 0x1a, 0x3d,
	0x10, 0x65, 0x09, 0x81, 0x8f, 0x21, 0x1f, 0x42, 0x4e, 0x50, 0x28, 0x4a, 0x53, 0x1a, 0xa9, 0x12,
	0x8a, 0xd4, 0x45, 0xa1, 0xad, 0x94, 0x8b, 0x4a, 0x15, 0x44, 0xcb, 0x4f, 0x95, 0x54, 0x64, 0xd9,
	0xaa, 0x51, 0x6f, 0x5c, 0xb3, 0x6b, 0x16, 0xab, 0xbb, 0x36, 0x35, 0x86, 0x8a, 0x5c, 0xf5, 0x05,
	0xfa, 0x08, 0x7d, 0xaf, 0x3e, 0x4e, 0xb5, 0x5e, 0x48, 0x48, 0x02, 0xad, 0xda, 0x2b, 0xdb, 0x33,
	0x73, 0xce, 0xcc, 0x9c, 0xb1, 0x0d, 0xcf, 0x26, 0x8a, 0x06, 0x01, 0x91, 0x13, 0x45, 0x14, 0x13,
	0xbc, 0x46, 0xc6, 0xac, 0x36, 0x3b, 0xa9, 0x29, 0x49, 0xf8, 0x64, 0x2c, 0xa4, 0x32, 0xc7, 0x52,
	0x28, 0x81, 0x8a, 0x77, 0xc3, 0x4c, 0x32, 0x66, 0xe6, 0xec, 0xa4, 0xfc, 0xc4, 0x17, 0xc2, 0x0f,
	0x68, 0x4d, 0x07, 0x0d, 0xa6, 0xc3, 0x9a, 0x62, 0x21, 0x9d, 0x28, 0x12, 0x8e, 0x63, 0x5c, 0xf9,
	0x64, 0x3d, 0x7d, 0x28, 0x38, 0x53, 0x42, 0x32, 0xee, 0xaf, 0x6c, 0x63, 0xc8, 0xd1, 0xcf, 0x24,
	0x64, 0x1d, 0x1a, 0xd0, 0x90, 0x2a, 0x39, 0x47, 0xa7, 0x90, 0x19, 0x4a, 0x12, 0x32, 0xee, 0x97,
	0x8c, 0x8a, 0x51, 0xcd, 0xd7, 0x0f, 0xcd, 0xb5, 0xa5, 0x98, 0xad, 0x38, 0xca, 0x5e, 0x86, 0x23,
	0x04, 0x29, 0x8f, 0x28, 0x52, 0x4a, 0x56, 0x8c, 0xea, 0x8e, 0xad, 0xf7, 0xa8, 0x0e, 0x45, 0x4f,
	0x7c, 0xe5, 0x01, 0xe3, 0x9f, 0xf1, 0x50, 0xd2, 0x2f, 0x53, 0xca, 0xdd, 0x39, 0x1e, 0x5d, 0x97,
	0xb6, 0x2a, 0x46, 0x35, 0x65, 0xef, 0x2e, 0x9d, 0xad, 0xa5, 0xaf, 0x73, 0x8d, 0xfa, 0x50, 0x8a,
	0xba, 0xc2, 0x43, 0x26, 0x27, 0x0a, 0x0f, 0xe6, 0x8a, 0x62, 0x49, 0x5d, 0xca, 0x66, 0xd4, 0x2b,
	0xa5, 0x2a, 0x46, 0x35, 0x57, 0x2f, 0x9b, 0xb1, 0x0c, 0xe6, 0x52, 0x06, 0xd3, 0x59, 0xca, 0x60,
	0x17, 0x23, 0x6c, 0x2b, 0x82, 0x36, 0xe7, 0x8a, 0xda, 0x0b, 0x20, 0xba, 0x84, 0x7d, 0x4d, 0x1a,
	0x90, 0x07, 0x9c, 0xdb, 0x7f, 0xe4, 0xdc, 0x8b, 0xa0, 0xe7, 0xe4, 0x1e, 0xe5, 0x53, 0xd8, 0x89,
	0x5a, 0xa7, 0x78, 0x44, 0x89, 0x47, 0x65, 0x29, 0xad, 0xfb, 0xce, 0x69, 0x5b, 0x47, 0x9b, 0x8e,
	0x7e, 0x24, 0x21, 0xd7, 0x57, 0x92, 0x92, 0xd0, 0x9a, 0x51, 0xae, 0xd0, 0x63, 0x00, 0xdd, 0xe8,
	0x44, 0x61, 0xe6, 0x69, 0x7d, 0xb3, 0x76, 0x76, 0x61, 0xe9, 0x7a, 0xc8, 0x83, 0x1d, 0x57, 0x84,
	0x21, 0xe1, 0x1e, 0x9e, 0x50, 0xae, 0xb4, 0x92, 0xb9, 0xfa, 0x9b, 0x0d, 0x03, 0x58, 0x21, 0x36,
	0xcf, 0x62, 0x58, 0x9f, 0x72, 0xd5, 0x92, 0x22, 0x6c, 0x4b, 0x31, 0xe5, 0x5e, 0x3f, 0x8e, 0xef,
	0x24, 0xec, 0x9c, 0x7b, 0xeb, 0x47, 0x9f, 0xa0, 0x38, 0x0e, 0x08, 0xc7, 0xb7, 0x17, 0x01, 0xd3,
	0x88, 0x44, 0xcf, 0x24, 0x57, 0x3f, 0xde, 0x90, 0xae, 0x17, 0x10, 0x7e, 0x71, 0x03, 0xd1, 0x69,
	0x3b, 0x09, 0x7b, 0x77, 0xfc, 0xd0, 0x5c, 0x3e, 0x84, 0x83, 0xdf, 0x15, 0xd4, 0xcc, 0xc0, 0xb6,
	0x0e, 0x3c, 0xfa, 0xbe, 0x05, 0xbb, 0x6b, 0x78, 0xd1, 0x3e, 0x64, 0x74, 0x89, 0x37, 0x22, 0xa5,
	0xa3, 0x63, 0xd7, 0x43, 0xdf, 0x0c, 0x38, 0xf0, 0x35, 0x17, 0x5e, 0x94, 0x87, 0x5d, 0xc1, 0x87,
	0xcc, 0x9f, 0x4a, 0x7d, 0x5a, 0x48, 0xf6, 0x7a, 0x43, 0x0f, 0x2b, 0x77, 0xff, 0x4e, 0x45, 0x67,
	0xab, 0x1c, 0x9d, 0x84, 0x5d, 0xf6, 0x37, 0x7a, 0xd1, 0x08, 0xf6, 0xee, 0x55, 0x10, 0xad, 0x74,
	0xa1, 0xde, 0xcb, 0xbf, 0xcc, 0x1c, 0x2d, 0xb4, 0x93, 0xb0, 0x91, 0xff, 0xc0, 0xba, 0x26, 0x53,
	0x3c, 0xa7, 0xd4, 0x3f, 0x65, 0x5a, 0x4e, 0xec, 0x6e, 0x26, 0x6d, 0x6d, 0xa6, 0x21, 0xd5, 0xe5,
	0x43, 0x71, 0x1c, 0x40, 0x66, 0xf1, 0xac, 0xd1, 0xff, 0x90, 0x6d, 0x76, 0x9d, 0xbe, 0x63, 0x5b,
	0x8d, 0x8b, 0x42, 0x02, 0xfd, 0x07, 0xa9, 0xc6, 0x55, 0xfd, 0x55, 0xc1, 0x40, 0x69, 0x48, 0x76,
	0x2f, 0x0b, 0xc9, 0x28, 0xa0, 0x7b, 0xd1, 0x68, 0x5b, 0xb8, 0xf7, 0xae, 0x5d, 0xd8, 0x42, 0x79,
	0x80, 0xf8, 0xf8, 0xb6, 0x67, 0xb5, 0x0b, 0x29, 0x84, 0x20, 0xdf, 0xb2, 0x2d, 0x0b, 0x3b, 0xd6,
	0x95, 0x83, 0xdf, 0x3b, 0xad, 0xd3, 0xc2, 0x76, 0x04, 0xf9, 0xd0, 0x70, 0x2c, 0xbb, 0xd5, 0x38,
	0x3f, 0x2f, 0xa4, 0x9b, 0x04, 0x1e, 0xb9, 0x22, 0x5c, 0xdf, 0x46, 0x33, 0xef, 0x2c, 0x7f, 0xc4,
	0x5e, 0xf4, 0x20, 0x7b, 0xc6, 0xc7, 0xba, 0xcf, 0xd4, 0x68, 0x3a, 0x30, 0x5d, 0x11, 0xd6, 0x18,
	0x1f, 0x8a, 0x05, 0x8e, 0x71, 0xb7, 0xe6, 0x8b, 0xe7, 0x6b, 0xff, 0xbd, 0x41, 0x5a, 0xbf, 0xe6,
	0x17, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x32, 0xf6, 0x8c, 0x76, 0x05, 0x00, 0x00,
}
