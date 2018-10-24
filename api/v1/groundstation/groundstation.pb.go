// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stellarstation/api/v1/groundstation/groundstation.proto

package groundstation

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	radio "github.com/infostellarinc/go-stellarstation/api/v1/radio"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request for the `ListPlans` method.
type ListPlansRequest struct {
	// The ID of the ground station to list plans for. The ID can be found on the StellarStation
	// Console page for the ground station.
	GroundStationId string `protobuf:"bytes,1,opt,name=ground_station_id,json=groundStationId,proto3" json:"ground_station_id,omitempty"`
	// The start time of the range of plans to list (inclusive). Only plans with an Acquisition of
	// Signal (AOS) at or after this time will be returned. It is an error for the duration between
	// `aos_after` and `aos_before` to be longer than 31 days.
	AosAfter *timestamp.Timestamp `protobuf:"bytes,2,opt,name=aos_after,json=aosAfter,proto3" json:"aos_after,omitempty"`
	// The end time of the range of plans to list (exclusive). Only plans with an Acquisition of
	// Signal (AOS) before this time will be returned. It is an error for the duration between
	// `aos_after` and `aos_before` to be longer than 31 days.
	AosBefore            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=aos_before,json=aosBefore,proto3" json:"aos_before,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListPlansRequest) Reset()         { *m = ListPlansRequest{} }
func (m *ListPlansRequest) String() string { return proto.CompactTextString(m) }
func (*ListPlansRequest) ProtoMessage()    {}
func (*ListPlansRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bce3778e6c67798, []int{0}
}

func (m *ListPlansRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPlansRequest.Unmarshal(m, b)
}
func (m *ListPlansRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPlansRequest.Marshal(b, m, deterministic)
}
func (m *ListPlansRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPlansRequest.Merge(m, src)
}
func (m *ListPlansRequest) XXX_Size() int {
	return xxx_messageInfo_ListPlansRequest.Size(m)
}
func (m *ListPlansRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPlansRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPlansRequest proto.InternalMessageInfo

func (m *ListPlansRequest) GetGroundStationId() string {
	if m != nil {
		return m.GroundStationId
	}
	return ""
}

func (m *ListPlansRequest) GetAosAfter() *timestamp.Timestamp {
	if m != nil {
		return m.AosAfter
	}
	return nil
}

func (m *ListPlansRequest) GetAosBefore() *timestamp.Timestamp {
	if m != nil {
		return m.AosBefore
	}
	return nil
}

// A response from the `ListPlans` method.
type ListPlansResponse struct {
	// The requested list of plans for the ground station.
	Plan                 []*Plan  `protobuf:"bytes,1,rep,name=plan,proto3" json:"plan,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPlansResponse) Reset()         { *m = ListPlansResponse{} }
func (m *ListPlansResponse) String() string { return proto.CompactTextString(m) }
func (*ListPlansResponse) ProtoMessage()    {}
func (*ListPlansResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bce3778e6c67798, []int{1}
}

func (m *ListPlansResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPlansResponse.Unmarshal(m, b)
}
func (m *ListPlansResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPlansResponse.Marshal(b, m, deterministic)
}
func (m *ListPlansResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPlansResponse.Merge(m, src)
}
func (m *ListPlansResponse) XXX_Size() int {
	return xxx_messageInfo_ListPlansResponse.Size(m)
}
func (m *ListPlansResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPlansResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPlansResponse proto.InternalMessageInfo

func (m *ListPlansResponse) GetPlan() []*Plan {
	if m != nil {
		return m.Plan
	}
	return nil
}

// A scheduled pass. The plan will be executed on its ground station to communicate with its satellite
// during a time range between AOS and LOS, unless explicitly cancelled.
type Plan struct {
	// The ID of this plan.
	PlanId string `protobuf:"bytes,1,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	// The TLE for the satellite in this plan.
	Tle *Tle `protobuf:"bytes,2,opt,name=tle,proto3" json:"tle,omitempty"`
	// The time of AOS between the ground station and satellite in this plan.
	AosTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=aos_time,json=aosTime,proto3" json:"aos_time,omitempty"`
	// The time of LOS between the ground station and satellite in this plan.
	LosTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=los_time,json=losTime,proto3" json:"los_time,omitempty"`
	// Configuration of the radio device used for downlinking from the satellite. Ground stations will
	// need to configure reception during the plan to match this device.
	DownlinkRadioDevice *radio.RadioDeviceConfiguration `protobuf:"bytes,5,opt,name=downlink_radio_device,json=downlinkRadioDevice,proto3" json:"downlink_radio_device,omitempty"`
	// Configuration of the radio device used for uplinking to the satellite. Ground stations will
	// need to configure transmission during the plan to match this device.
	UplinkRadioDevice    *radio.RadioDeviceConfiguration `protobuf:"bytes,6,opt,name=uplink_radio_device,json=uplinkRadioDevice,proto3" json:"uplink_radio_device,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *Plan) Reset()         { *m = Plan{} }
func (m *Plan) String() string { return proto.CompactTextString(m) }
func (*Plan) ProtoMessage()    {}
func (*Plan) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bce3778e6c67798, []int{2}
}

func (m *Plan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Plan.Unmarshal(m, b)
}
func (m *Plan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Plan.Marshal(b, m, deterministic)
}
func (m *Plan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Plan.Merge(m, src)
}
func (m *Plan) XXX_Size() int {
	return xxx_messageInfo_Plan.Size(m)
}
func (m *Plan) XXX_DiscardUnknown() {
	xxx_messageInfo_Plan.DiscardUnknown(m)
}

var xxx_messageInfo_Plan proto.InternalMessageInfo

func (m *Plan) GetPlanId() string {
	if m != nil {
		return m.PlanId
	}
	return ""
}

func (m *Plan) GetTle() *Tle {
	if m != nil {
		return m.Tle
	}
	return nil
}

func (m *Plan) GetAosTime() *timestamp.Timestamp {
	if m != nil {
		return m.AosTime
	}
	return nil
}

func (m *Plan) GetLosTime() *timestamp.Timestamp {
	if m != nil {
		return m.LosTime
	}
	return nil
}

func (m *Plan) GetDownlinkRadioDevice() *radio.RadioDeviceConfiguration {
	if m != nil {
		return m.DownlinkRadioDevice
	}
	return nil
}

func (m *Plan) GetUplinkRadioDevice() *radio.RadioDeviceConfiguration {
	if m != nil {
		return m.UplinkRadioDevice
	}
	return nil
}

// Unparsed TLE data for a satellite - https://en.wikipedia.org/wiki/Two-line_element_set
type Tle struct {
	// The first line of the TLE. Not a title line.
	Line_1 string `protobuf:"bytes,1,opt,name=line_1,json=line1,proto3" json:"line_1,omitempty"`
	// The second line of the TLE.
	Line_2               string   `protobuf:"bytes,2,opt,name=line_2,json=line2,proto3" json:"line_2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tle) Reset()         { *m = Tle{} }
func (m *Tle) String() string { return proto.CompactTextString(m) }
func (*Tle) ProtoMessage()    {}
func (*Tle) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bce3778e6c67798, []int{3}
}

func (m *Tle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tle.Unmarshal(m, b)
}
func (m *Tle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tle.Marshal(b, m, deterministic)
}
func (m *Tle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tle.Merge(m, src)
}
func (m *Tle) XXX_Size() int {
	return xxx_messageInfo_Tle.Size(m)
}
func (m *Tle) XXX_DiscardUnknown() {
	xxx_messageInfo_Tle.DiscardUnknown(m)
}

var xxx_messageInfo_Tle proto.InternalMessageInfo

func (m *Tle) GetLine_1() string {
	if m != nil {
		return m.Line_1
	}
	return ""
}

func (m *Tle) GetLine_2() string {
	if m != nil {
		return m.Line_2
	}
	return ""
}

// A time window during which a ground station is unavailable e.g. for local maintenance.
type UnavailabilityWindow struct {
	// The ID of the unavailability window.
	WindowId string `protobuf:"bytes,1,opt,name=window_id,json=windowId,proto3" json:"window_id,omitempty"`
	// Start time of the unavailabilty window.
	StartTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// End time of the unavailability window.
	EndTime              *timestamp.Timestamp `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UnavailabilityWindow) Reset()         { *m = UnavailabilityWindow{} }
func (m *UnavailabilityWindow) String() string { return proto.CompactTextString(m) }
func (*UnavailabilityWindow) ProtoMessage()    {}
func (*UnavailabilityWindow) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bce3778e6c67798, []int{4}
}

func (m *UnavailabilityWindow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnavailabilityWindow.Unmarshal(m, b)
}
func (m *UnavailabilityWindow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnavailabilityWindow.Marshal(b, m, deterministic)
}
func (m *UnavailabilityWindow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnavailabilityWindow.Merge(m, src)
}
func (m *UnavailabilityWindow) XXX_Size() int {
	return xxx_messageInfo_UnavailabilityWindow.Size(m)
}
func (m *UnavailabilityWindow) XXX_DiscardUnknown() {
	xxx_messageInfo_UnavailabilityWindow.DiscardUnknown(m)
}

var xxx_messageInfo_UnavailabilityWindow proto.InternalMessageInfo

func (m *UnavailabilityWindow) GetWindowId() string {
	if m != nil {
		return m.WindowId
	}
	return ""
}

func (m *UnavailabilityWindow) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *UnavailabilityWindow) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

// A request for a list of unavailability windows for the specified ground station that
// fall within the given time range.
type ListUnavailabilityWindowsRequest struct {
	// ID of the ground station for which to retrieve unavailability windows.
	GroundStationId string `protobuf:"bytes,1,opt,name=ground_station_id,json=groundStationId,proto3" json:"ground_station_id,omitempty"`
	// Start time.
	StartTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// End time.
	EndTime              *timestamp.Timestamp `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListUnavailabilityWindowsRequest) Reset()         { *m = ListUnavailabilityWindowsRequest{} }
func (m *ListUnavailabilityWindowsRequest) String() string { return proto.CompactTextString(m) }
func (*ListUnavailabilityWindowsRequest) ProtoMessage()    {}
func (*ListUnavailabilityWindowsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bce3778e6c67798, []int{5}
}

func (m *ListUnavailabilityWindowsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUnavailabilityWindowsRequest.Unmarshal(m, b)
}
func (m *ListUnavailabilityWindowsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUnavailabilityWindowsRequest.Marshal(b, m, deterministic)
}
func (m *ListUnavailabilityWindowsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUnavailabilityWindowsRequest.Merge(m, src)
}
func (m *ListUnavailabilityWindowsRequest) XXX_Size() int {
	return xxx_messageInfo_ListUnavailabilityWindowsRequest.Size(m)
}
func (m *ListUnavailabilityWindowsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUnavailabilityWindowsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUnavailabilityWindowsRequest proto.InternalMessageInfo

func (m *ListUnavailabilityWindowsRequest) GetGroundStationId() string {
	if m != nil {
		return m.GroundStationId
	}
	return ""
}

func (m *ListUnavailabilityWindowsRequest) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *ListUnavailabilityWindowsRequest) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

// A response containing unavailability windows for the requested ground station.
type ListUnavailabilityWindowsResponse struct {
	// A list of unavailability windows, sorted in ascending order of the start time.
	Window               []*UnavailabilityWindow `protobuf:"bytes,1,rep,name=window,proto3" json:"window,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ListUnavailabilityWindowsResponse) Reset()         { *m = ListUnavailabilityWindowsResponse{} }
func (m *ListUnavailabilityWindowsResponse) String() string { return proto.CompactTextString(m) }
func (*ListUnavailabilityWindowsResponse) ProtoMessage()    {}
func (*ListUnavailabilityWindowsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bce3778e6c67798, []int{6}
}

func (m *ListUnavailabilityWindowsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUnavailabilityWindowsResponse.Unmarshal(m, b)
}
func (m *ListUnavailabilityWindowsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUnavailabilityWindowsResponse.Marshal(b, m, deterministic)
}
func (m *ListUnavailabilityWindowsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUnavailabilityWindowsResponse.Merge(m, src)
}
func (m *ListUnavailabilityWindowsResponse) XXX_Size() int {
	return xxx_messageInfo_ListUnavailabilityWindowsResponse.Size(m)
}
func (m *ListUnavailabilityWindowsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUnavailabilityWindowsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUnavailabilityWindowsResponse proto.InternalMessageInfo

func (m *ListUnavailabilityWindowsResponse) GetWindow() []*UnavailabilityWindow {
	if m != nil {
		return m.Window
	}
	return nil
}

// A request for adding a new unavailability window for the specified ground station.
type AddUnavailabilityWindowRequest struct {
	// ID of the ground station to add a new unavailability window.
	GroundStationId string `protobuf:"bytes,1,opt,name=ground_station_id,json=groundStationId,proto3" json:"ground_station_id,omitempty"`
	// Start time of the unavailabilty window.
	StartTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// End time of the unavailability window.
	EndTime              *timestamp.Timestamp `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AddUnavailabilityWindowRequest) Reset()         { *m = AddUnavailabilityWindowRequest{} }
func (m *AddUnavailabilityWindowRequest) String() string { return proto.CompactTextString(m) }
func (*AddUnavailabilityWindowRequest) ProtoMessage()    {}
func (*AddUnavailabilityWindowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bce3778e6c67798, []int{7}
}

func (m *AddUnavailabilityWindowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddUnavailabilityWindowRequest.Unmarshal(m, b)
}
func (m *AddUnavailabilityWindowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddUnavailabilityWindowRequest.Marshal(b, m, deterministic)
}
func (m *AddUnavailabilityWindowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddUnavailabilityWindowRequest.Merge(m, src)
}
func (m *AddUnavailabilityWindowRequest) XXX_Size() int {
	return xxx_messageInfo_AddUnavailabilityWindowRequest.Size(m)
}
func (m *AddUnavailabilityWindowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddUnavailabilityWindowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddUnavailabilityWindowRequest proto.InternalMessageInfo

func (m *AddUnavailabilityWindowRequest) GetGroundStationId() string {
	if m != nil {
		return m.GroundStationId
	}
	return ""
}

func (m *AddUnavailabilityWindowRequest) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *AddUnavailabilityWindowRequest) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

// A response from the 'AddUnavailabilityWindow' method.
type AddUnavailabilityWindowResponse struct {
	// ID of the new window.
	WindowId             string   `protobuf:"bytes,1,opt,name=window_id,json=windowId,proto3" json:"window_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddUnavailabilityWindowResponse) Reset()         { *m = AddUnavailabilityWindowResponse{} }
func (m *AddUnavailabilityWindowResponse) String() string { return proto.CompactTextString(m) }
func (*AddUnavailabilityWindowResponse) ProtoMessage()    {}
func (*AddUnavailabilityWindowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bce3778e6c67798, []int{8}
}

func (m *AddUnavailabilityWindowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddUnavailabilityWindowResponse.Unmarshal(m, b)
}
func (m *AddUnavailabilityWindowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddUnavailabilityWindowResponse.Marshal(b, m, deterministic)
}
func (m *AddUnavailabilityWindowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddUnavailabilityWindowResponse.Merge(m, src)
}
func (m *AddUnavailabilityWindowResponse) XXX_Size() int {
	return xxx_messageInfo_AddUnavailabilityWindowResponse.Size(m)
}
func (m *AddUnavailabilityWindowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddUnavailabilityWindowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddUnavailabilityWindowResponse proto.InternalMessageInfo

func (m *AddUnavailabilityWindowResponse) GetWindowId() string {
	if m != nil {
		return m.WindowId
	}
	return ""
}

// A request for deleting an existing unavailability window for the specified ground station.
type DeleteUnavailabilityWindowRequest struct {
	// ID of the unavailability window to delete.
	WindowId             string   `protobuf:"bytes,1,opt,name=window_id,json=windowId,proto3" json:"window_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUnavailabilityWindowRequest) Reset()         { *m = DeleteUnavailabilityWindowRequest{} }
func (m *DeleteUnavailabilityWindowRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteUnavailabilityWindowRequest) ProtoMessage()    {}
func (*DeleteUnavailabilityWindowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bce3778e6c67798, []int{9}
}

func (m *DeleteUnavailabilityWindowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUnavailabilityWindowRequest.Unmarshal(m, b)
}
func (m *DeleteUnavailabilityWindowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUnavailabilityWindowRequest.Marshal(b, m, deterministic)
}
func (m *DeleteUnavailabilityWindowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUnavailabilityWindowRequest.Merge(m, src)
}
func (m *DeleteUnavailabilityWindowRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteUnavailabilityWindowRequest.Size(m)
}
func (m *DeleteUnavailabilityWindowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUnavailabilityWindowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUnavailabilityWindowRequest proto.InternalMessageInfo

func (m *DeleteUnavailabilityWindowRequest) GetWindowId() string {
	if m != nil {
		return m.WindowId
	}
	return ""
}

// A response to the request for deleting an existing unavailability window.
type DeleteUnavailabilityWindowResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUnavailabilityWindowResponse) Reset()         { *m = DeleteUnavailabilityWindowResponse{} }
func (m *DeleteUnavailabilityWindowResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteUnavailabilityWindowResponse) ProtoMessage()    {}
func (*DeleteUnavailabilityWindowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bce3778e6c67798, []int{10}
}

func (m *DeleteUnavailabilityWindowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUnavailabilityWindowResponse.Unmarshal(m, b)
}
func (m *DeleteUnavailabilityWindowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUnavailabilityWindowResponse.Marshal(b, m, deterministic)
}
func (m *DeleteUnavailabilityWindowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUnavailabilityWindowResponse.Merge(m, src)
}
func (m *DeleteUnavailabilityWindowResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteUnavailabilityWindowResponse.Size(m)
}
func (m *DeleteUnavailabilityWindowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUnavailabilityWindowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUnavailabilityWindowResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ListPlansRequest)(nil), "stellarstation.api.v1.groundstation.ListPlansRequest")
	proto.RegisterType((*ListPlansResponse)(nil), "stellarstation.api.v1.groundstation.ListPlansResponse")
	proto.RegisterType((*Plan)(nil), "stellarstation.api.v1.groundstation.Plan")
	proto.RegisterType((*Tle)(nil), "stellarstation.api.v1.groundstation.Tle")
	proto.RegisterType((*UnavailabilityWindow)(nil), "stellarstation.api.v1.groundstation.UnavailabilityWindow")
	proto.RegisterType((*ListUnavailabilityWindowsRequest)(nil), "stellarstation.api.v1.groundstation.ListUnavailabilityWindowsRequest")
	proto.RegisterType((*ListUnavailabilityWindowsResponse)(nil), "stellarstation.api.v1.groundstation.ListUnavailabilityWindowsResponse")
	proto.RegisterType((*AddUnavailabilityWindowRequest)(nil), "stellarstation.api.v1.groundstation.AddUnavailabilityWindowRequest")
	proto.RegisterType((*AddUnavailabilityWindowResponse)(nil), "stellarstation.api.v1.groundstation.AddUnavailabilityWindowResponse")
	proto.RegisterType((*DeleteUnavailabilityWindowRequest)(nil), "stellarstation.api.v1.groundstation.DeleteUnavailabilityWindowRequest")
	proto.RegisterType((*DeleteUnavailabilityWindowResponse)(nil), "stellarstation.api.v1.groundstation.DeleteUnavailabilityWindowResponse")
}

func init() {
	proto.RegisterFile("stellarstation/api/v1/groundstation/groundstation.proto", fileDescriptor_9bce3778e6c67798)
}

var fileDescriptor_9bce3778e6c67798 = []byte{
	// 694 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0x5d, 0x6b, 0x13, 0x4d,
	0x14, 0x66, 0xdf, 0xb4, 0x69, 0x73, 0x7a, 0xf1, 0xda, 0x69, 0x4b, 0xe3, 0x0a, 0x36, 0x5d, 0x85,
	0x46, 0xc1, 0x5d, 0x92, 0x52, 0x4b, 0x05, 0xa5, 0x5f, 0xb6, 0x14, 0xbc, 0xa8, 0xdb, 0x8a, 0xe0,
	0xcd, 0x32, 0xc9, 0x4e, 0xe2, 0xe0, 0x64, 0x66, 0xdd, 0x9d, 0xa4, 0xe8, 0x9d, 0x7f, 0xc2, 0x1b,
	0x6f, 0xf5, 0x46, 0xf0, 0xc6, 0x1f, 0x20, 0xfe, 0x34, 0x99, 0x9d, 0x49, 0x9b, 0x68, 0xb2, 0x59,
	0x0d, 0x08, 0xde, 0x84, 0x9e, 0xb3, 0xe7, 0x79, 0xce, 0x73, 0xbe, 0x76, 0x0b, 0xdb, 0x89, 0x24,
	0x8c, 0xe1, 0x38, 0x91, 0x58, 0x52, 0xc1, 0x3d, 0x1c, 0x51, 0xaf, 0x57, 0xf3, 0xda, 0xb1, 0xe8,
	0xf2, 0xb0, 0xef, 0x1c, 0xb2, 0xdc, 0x28, 0x16, 0x52, 0xa0, 0x5b, 0xc3, 0x40, 0x17, 0x47, 0xd4,
	0xed, 0xd5, 0xdc, 0xa1, 0x50, 0x7b, 0xad, 0x2d, 0x44, 0x9b, 0x11, 0x2f, 0x85, 0x34, 0xba, 0x2d,
	0x4f, 0xd2, 0x0e, 0x49, 0x24, 0xee, 0x44, 0x9a, 0xc5, 0xde, 0x18, 0x9d, 0x3e, 0xc6, 0x21, 0x15,
	0xfa, 0x57, 0x07, 0x3a, 0x5f, 0x2d, 0xb8, 0xf6, 0x84, 0x26, 0xf2, 0x94, 0x61, 0x9e, 0xf8, 0xe4,
	0x75, 0x97, 0x24, 0x12, 0xdd, 0x85, 0x45, 0x9d, 0x2f, 0x30, 0xf8, 0x80, 0x86, 0x65, 0xab, 0x62,
	0x55, 0x4b, 0xfe, 0xff, 0xfa, 0xc1, 0x99, 0xf6, 0x9f, 0x84, 0x68, 0x1b, 0x4a, 0x58, 0x24, 0x01,
	0x6e, 0x49, 0x12, 0x97, 0xff, 0xab, 0x58, 0xd5, 0x85, 0xba, 0xed, 0x6a, 0x79, 0x6e, 0x5f, 0x9e,
	0x7b, 0xde, 0x97, 0xe7, 0xcf, 0x63, 0x91, 0xec, 0xa9, 0x58, 0xb4, 0x03, 0xa0, 0x80, 0x0d, 0xd2,
	0x12, 0x31, 0x29, 0x17, 0x26, 0x22, 0x55, 0x9a, 0xfd, 0x34, 0xd8, 0xf1, 0x61, 0x71, 0x40, 0x73,
	0x12, 0x09, 0x9e, 0x10, 0xf4, 0x10, 0x66, 0x22, 0x86, 0x79, 0xd9, 0xaa, 0x14, 0xaa, 0x0b, 0xf5,
	0x3b, 0x6e, 0x8e, 0x3e, 0xba, 0x8a, 0xc1, 0x4f, 0x61, 0xce, 0x87, 0x02, 0xcc, 0x28, 0x13, 0xad,
	0xc2, 0x9c, 0x72, 0x5c, 0x95, 0x5c, 0x54, 0xe6, 0x49, 0x88, 0x1e, 0x40, 0x41, 0x32, 0x62, 0x6a,
	0xac, 0xe6, 0xe2, 0x3f, 0x67, 0xc4, 0x57, 0x20, 0xb4, 0x05, 0xaa, 0xf0, 0x40, 0x8d, 0x29, 0x47,
	0xa9, 0x73, 0x58, 0x24, 0xca, 0x52, 0x30, 0xd6, 0x87, 0xcd, 0x4c, 0x86, 0x31, 0x03, 0xa3, 0xb0,
	0x12, 0x8a, 0x0b, 0xce, 0x28, 0x7f, 0x15, 0xa4, 0xc3, 0x0e, 0x42, 0xd2, 0xa3, 0x4d, 0x52, 0x9e,
	0x4d, 0x39, 0xb6, 0xc6, 0x68, 0xd7, 0x7b, 0xe1, 0xab, 0xdf, 0xc3, 0x34, 0xfe, 0x40, 0xf0, 0x16,
	0x6d, 0x77, 0xe3, 0x34, 0xca, 0x5f, 0xea, 0x73, 0x0e, 0x44, 0x20, 0x02, 0x4b, 0xdd, 0xe8, 0xd7,
	0x44, 0xc5, 0x69, 0x12, 0x2d, 0x6a, 0xc6, 0x81, 0xe7, 0xce, 0x26, 0x14, 0xce, 0x19, 0x41, 0x2b,
	0x50, 0x64, 0x94, 0x93, 0xa0, 0x66, 0x46, 0x33, 0xab, 0xac, 0xda, 0xa5, 0xbb, 0x9e, 0x0e, 0xc7,
	0xb8, 0xeb, 0xce, 0x27, 0x0b, 0x96, 0x9f, 0x71, 0xdc, 0xc3, 0x94, 0xe1, 0x06, 0x65, 0x54, 0xbe,
	0x79, 0x4e, 0x79, 0x28, 0x2e, 0xd0, 0x0d, 0x28, 0x5d, 0xa4, 0x7f, 0x5d, 0x0d, 0x79, 0x5e, 0x3b,
	0x4e, 0x42, 0xb5, 0x97, 0x89, 0xc4, 0xb1, 0xd4, 0x5d, 0x9f, 0xbc, 0xd1, 0xa5, 0x34, 0xba, 0x3f,
	0x2e, 0xc2, 0xc3, 0xdc, 0x53, 0x26, 0x3c, 0x54, 0x96, 0xf3, 0xdd, 0x82, 0x8a, 0xda, 0xe7, 0x51,
	0x5a, 0xff, 0xe8, 0x26, 0xff, 0x7e, 0x09, 0x3d, 0x58, 0xcf, 0xa8, 0xc0, 0x5c, 0xe8, 0x53, 0x28,
	0xea, 0x2e, 0x9b, 0x1b, 0xdd, 0xc9, 0x75, 0x43, 0xa3, 0x38, 0x7d, 0x43, 0xe4, 0x7c, 0xb3, 0xe0,
	0xe6, 0x5e, 0x18, 0x8e, 0x8c, 0xf9, 0x27, 0x1a, 0xf7, 0x08, 0xd6, 0xc6, 0xea, 0x37, 0x6d, 0xcb,
	0xda, 0x56, 0x67, 0x17, 0xd6, 0x0f, 0x09, 0x23, 0x92, 0x64, 0xb5, 0x20, 0x93, 0xe1, 0x36, 0x38,
	0x59, 0x0c, 0x5a, 0x44, 0xfd, 0xfd, 0x2c, 0x2c, 0x1f, 0x0f, 0x76, 0xeb, 0x8c, 0xc4, 0xe9, 0x0b,
	0xe0, 0xa3, 0x05, 0xab, 0x63, 0x2a, 0x40, 0x07, 0xb9, 0x06, 0x9c, 0x3d, 0x3f, 0xfb, 0x70, 0x3a,
	0x12, 0xd3, 0xc4, 0x2f, 0x16, 0xd8, 0xe3, 0xcb, 0x44, 0x47, 0xb9, 0x92, 0x4c, 0xec, 0xb4, 0x7d,
	0x3c, 0x35, 0x8f, 0xd1, 0xfb, 0x16, 0x4a, 0x97, 0x9f, 0x38, 0xb4, 0x95, 0x8b, 0xf5, 0xe7, 0xcf,
	0xb8, 0x7d, 0xff, 0x77, 0x61, 0x26, 0xf7, 0x67, 0x0b, 0xae, 0x8f, 0xbd, 0x66, 0xf4, 0x38, 0x37,
	0x6b, 0xd6, 0xfb, 0xcc, 0x3e, 0x9a, 0x96, 0x46, 0x8b, 0xdd, 0x7f, 0x67, 0xc1, 0x46, 0x53, 0x74,
	0xf2, 0xb0, 0xed, 0xa3, 0xa1, 0x0d, 0x3e, 0x55, 0x67, 0x79, 0x6a, 0xbd, 0xd8, 0x6d, 0x53, 0xf9,
	0xb2, 0xdb, 0x70, 0x9b, 0xa2, 0xe3, 0x51, 0xde, 0x12, 0x86, 0x89, 0xf2, 0xa6, 0xd7, 0x16, 0xf7,
	0x72, 0xfc, 0x1f, 0xd7, 0x28, 0xa6, 0x17, 0xbe, 0xf9, 0x23, 0x00, 0x00, 0xff, 0xff, 0xdb, 0xf2,
	0x86, 0x94, 0xf5, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GroundStationServiceClient is the client API for GroundStationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GroundStationServiceClient interface {
	// Adds a new unavailability window to the requested ground station.
	//
	// Existing plans that overlap the unavailability window will be canceled. However, there
	// are cases when the plan cannot be canceled. When this happens, the request will be closed
	// with a 'FAILED_PRECONDITION' status.
	//
	// The request will be closed with an `INVALID_ARGUMENT` status if `ground_station_id`,
	// `start_time`, or `end_time` are missing, or 'end_time' is not after 'start_time'.
	AddUnavailabilityWindow(ctx context.Context, in *AddUnavailabilityWindowRequest, opts ...grpc.CallOption) (*AddUnavailabilityWindowResponse, error)
	// Deletes an existing unavailability window of the requested ground station.
	//
	// The request will be closed with an `INVALID_ARGUMENT` status if `window_id` is missing
	// or invalid.
	DeleteUnavailabilityWindow(ctx context.Context, in *DeleteUnavailabilityWindowRequest, opts ...grpc.CallOption) (*DeleteUnavailabilityWindowResponse, error)
	// Lists the plans for a particular ground station.
	//
	// The request will be closed with an `INVALID_ARGUMENT` status if `ground_station_id`,
	// `aos_after`, or `aos_before` are missing, or the duration between the two times is longer than
	// 31 days.
	ListPlans(ctx context.Context, in *ListPlansRequest, opts ...grpc.CallOption) (*ListPlansResponse, error)
	// Returns a list of unavailability windows for the requested ground station.
	//
	// The request will be closed with an `INVALID_ARGUMENT` status if `ground_station_id`,
	// `start_time`, or `end_time` are missing, or 'end_time' is not after 'start_time'.
	ListUnavailabilityWindows(ctx context.Context, in *ListUnavailabilityWindowsRequest, opts ...grpc.CallOption) (*ListUnavailabilityWindowsResponse, error)
}

type groundStationServiceClient struct {
	cc *grpc.ClientConn
}

func NewGroundStationServiceClient(cc *grpc.ClientConn) GroundStationServiceClient {
	return &groundStationServiceClient{cc}
}

func (c *groundStationServiceClient) AddUnavailabilityWindow(ctx context.Context, in *AddUnavailabilityWindowRequest, opts ...grpc.CallOption) (*AddUnavailabilityWindowResponse, error) {
	out := new(AddUnavailabilityWindowResponse)
	err := c.cc.Invoke(ctx, "/stellarstation.api.v1.groundstation.GroundStationService/AddUnavailabilityWindow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groundStationServiceClient) DeleteUnavailabilityWindow(ctx context.Context, in *DeleteUnavailabilityWindowRequest, opts ...grpc.CallOption) (*DeleteUnavailabilityWindowResponse, error) {
	out := new(DeleteUnavailabilityWindowResponse)
	err := c.cc.Invoke(ctx, "/stellarstation.api.v1.groundstation.GroundStationService/DeleteUnavailabilityWindow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groundStationServiceClient) ListPlans(ctx context.Context, in *ListPlansRequest, opts ...grpc.CallOption) (*ListPlansResponse, error) {
	out := new(ListPlansResponse)
	err := c.cc.Invoke(ctx, "/stellarstation.api.v1.groundstation.GroundStationService/ListPlans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groundStationServiceClient) ListUnavailabilityWindows(ctx context.Context, in *ListUnavailabilityWindowsRequest, opts ...grpc.CallOption) (*ListUnavailabilityWindowsResponse, error) {
	out := new(ListUnavailabilityWindowsResponse)
	err := c.cc.Invoke(ctx, "/stellarstation.api.v1.groundstation.GroundStationService/ListUnavailabilityWindows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroundStationServiceServer is the server API for GroundStationService service.
type GroundStationServiceServer interface {
	// Adds a new unavailability window to the requested ground station.
	//
	// Existing plans that overlap the unavailability window will be canceled. However, there
	// are cases when the plan cannot be canceled. When this happens, the request will be closed
	// with a 'FAILED_PRECONDITION' status.
	//
	// The request will be closed with an `INVALID_ARGUMENT` status if `ground_station_id`,
	// `start_time`, or `end_time` are missing, or 'end_time' is not after 'start_time'.
	AddUnavailabilityWindow(context.Context, *AddUnavailabilityWindowRequest) (*AddUnavailabilityWindowResponse, error)
	// Deletes an existing unavailability window of the requested ground station.
	//
	// The request will be closed with an `INVALID_ARGUMENT` status if `window_id` is missing
	// or invalid.
	DeleteUnavailabilityWindow(context.Context, *DeleteUnavailabilityWindowRequest) (*DeleteUnavailabilityWindowResponse, error)
	// Lists the plans for a particular ground station.
	//
	// The request will be closed with an `INVALID_ARGUMENT` status if `ground_station_id`,
	// `aos_after`, or `aos_before` are missing, or the duration between the two times is longer than
	// 31 days.
	ListPlans(context.Context, *ListPlansRequest) (*ListPlansResponse, error)
	// Returns a list of unavailability windows for the requested ground station.
	//
	// The request will be closed with an `INVALID_ARGUMENT` status if `ground_station_id`,
	// `start_time`, or `end_time` are missing, or 'end_time' is not after 'start_time'.
	ListUnavailabilityWindows(context.Context, *ListUnavailabilityWindowsRequest) (*ListUnavailabilityWindowsResponse, error)
}

func RegisterGroundStationServiceServer(s *grpc.Server, srv GroundStationServiceServer) {
	s.RegisterService(&_GroundStationService_serviceDesc, srv)
}

func _GroundStationService_AddUnavailabilityWindow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUnavailabilityWindowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroundStationServiceServer).AddUnavailabilityWindow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellarstation.api.v1.groundstation.GroundStationService/AddUnavailabilityWindow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroundStationServiceServer).AddUnavailabilityWindow(ctx, req.(*AddUnavailabilityWindowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroundStationService_DeleteUnavailabilityWindow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUnavailabilityWindowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroundStationServiceServer).DeleteUnavailabilityWindow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellarstation.api.v1.groundstation.GroundStationService/DeleteUnavailabilityWindow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroundStationServiceServer).DeleteUnavailabilityWindow(ctx, req.(*DeleteUnavailabilityWindowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroundStationService_ListPlans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPlansRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroundStationServiceServer).ListPlans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellarstation.api.v1.groundstation.GroundStationService/ListPlans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroundStationServiceServer).ListPlans(ctx, req.(*ListPlansRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroundStationService_ListUnavailabilityWindows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUnavailabilityWindowsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroundStationServiceServer).ListUnavailabilityWindows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellarstation.api.v1.groundstation.GroundStationService/ListUnavailabilityWindows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroundStationServiceServer).ListUnavailabilityWindows(ctx, req.(*ListUnavailabilityWindowsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GroundStationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stellarstation.api.v1.groundstation.GroundStationService",
	HandlerType: (*GroundStationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUnavailabilityWindow",
			Handler:    _GroundStationService_AddUnavailabilityWindow_Handler,
		},
		{
			MethodName: "DeleteUnavailabilityWindow",
			Handler:    _GroundStationService_DeleteUnavailabilityWindow_Handler,
		},
		{
			MethodName: "ListPlans",
			Handler:    _GroundStationService_ListPlans_Handler,
		},
		{
			MethodName: "ListUnavailabilityWindows",
			Handler:    _GroundStationService_ListUnavailabilityWindows_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stellarstation/api/v1/groundstation/groundstation.proto",
}
