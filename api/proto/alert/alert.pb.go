// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: alert.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Alert struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AssetId       int32                  `protobuf:"varint,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	AlertType     string                 `protobuf:"bytes,2,opt,name=alert_type,json=alertType,proto3" json:"alert_type,omitempty"`
	Timestamp     int64                  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Severity      string                 `protobuf:"bytes,4,opt,name=severity,proto3" json:"severity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Alert) Reset() {
	*x = Alert{}
	mi := &file_alert_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Alert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alert) ProtoMessage() {}

func (x *Alert) ProtoReflect() protoreflect.Message {
	mi := &file_alert_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alert.ProtoReflect.Descriptor instead.
func (*Alert) Descriptor() ([]byte, []int) {
	return file_alert_proto_rawDescGZIP(), []int{0}
}

func (x *Alert) GetAssetId() int32 {
	if x != nil {
		return x.AssetId
	}
	return 0
}

func (x *Alert) GetAlertType() string {
	if x != nil {
		return x.AlertType
	}
	return ""
}

func (x *Alert) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Alert) GetSeverity() string {
	if x != nil {
		return x.Severity
	}
	return ""
}

type AlertRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AssetId       int32                  `protobuf:"varint,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	Type          string                 `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AlertRequest) Reset() {
	*x = AlertRequest{}
	mi := &file_alert_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertRequest) ProtoMessage() {}

func (x *AlertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alert_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertRequest.ProtoReflect.Descriptor instead.
func (*AlertRequest) Descriptor() ([]byte, []int) {
	return file_alert_proto_rawDescGZIP(), []int{1}
}

func (x *AlertRequest) GetAssetId() int32 {
	if x != nil {
		return x.AssetId
	}
	return 0
}

func (x *AlertRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type AlertResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Alerts        []*Alert               `protobuf:"bytes,1,rep,name=alerts,proto3" json:"alerts,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AlertResponse) Reset() {
	*x = AlertResponse{}
	mi := &file_alert_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AlertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertResponse) ProtoMessage() {}

func (x *AlertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_alert_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertResponse.ProtoReflect.Descriptor instead.
func (*AlertResponse) Descriptor() ([]byte, []int) {
	return file_alert_proto_rawDescGZIP(), []int{2}
}

func (x *AlertResponse) GetAlerts() []*Alert {
	if x != nil {
		return x.Alerts
	}
	return nil
}

var File_alert_proto protoreflect.FileDescriptor

const file_alert_proto_rawDesc = "" +
	"\n" +
	"\valert.proto\x12\x05alert\"{\n" +
	"\x05Alert\x12\x19\n" +
	"\basset_id\x18\x01 \x01(\x05R\aassetId\x12\x1d\n" +
	"\n" +
	"alert_type\x18\x02 \x01(\tR\talertType\x12\x1c\n" +
	"\ttimestamp\x18\x03 \x01(\x03R\ttimestamp\x12\x1a\n" +
	"\bseverity\x18\x04 \x01(\tR\bseverity\"=\n" +
	"\fAlertRequest\x12\x19\n" +
	"\basset_id\x18\x01 \x01(\x05R\aassetId\x12\x12\n" +
	"\x04type\x18\x02 \x01(\tR\x04type\"5\n" +
	"\rAlertResponse\x12$\n" +
	"\x06alerts\x18\x01 \x03(\v2\f.alert.AlertR\x06alerts2F\n" +
	"\fAlertService\x126\n" +
	"\tGetAlerts\x12\x13.alert.AlertRequest\x1a\x14.alert.AlertResponseB\x03Z\x01.b\x06proto3"

var (
	file_alert_proto_rawDescOnce sync.Once
	file_alert_proto_rawDescData []byte
)

func file_alert_proto_rawDescGZIP() []byte {
	file_alert_proto_rawDescOnce.Do(func() {
		file_alert_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_alert_proto_rawDesc), len(file_alert_proto_rawDesc)))
	})
	return file_alert_proto_rawDescData
}

var file_alert_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_alert_proto_goTypes = []any{
	(*Alert)(nil),         // 0: alert.Alert
	(*AlertRequest)(nil),  // 1: alert.AlertRequest
	(*AlertResponse)(nil), // 2: alert.AlertResponse
}
var file_alert_proto_depIdxs = []int32{
	0, // 0: alert.AlertResponse.alerts:type_name -> alert.Alert
	1, // 1: alert.AlertService.GetAlerts:input_type -> alert.AlertRequest
	2, // 2: alert.AlertService.GetAlerts:output_type -> alert.AlertResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_alert_proto_init() }
func file_alert_proto_init() {
	if File_alert_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_alert_proto_rawDesc), len(file_alert_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_alert_proto_goTypes,
		DependencyIndexes: file_alert_proto_depIdxs,
		MessageInfos:      file_alert_proto_msgTypes,
	}.Build()
	File_alert_proto = out.File
	file_alert_proto_goTypes = nil
	file_alert_proto_depIdxs = nil
}
