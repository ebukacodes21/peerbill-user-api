// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: peerbill_user.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_peerbill_user_proto protoreflect.FileDescriptor

var file_peerbill_user_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x65, 0x65, 0x72, 0x62, 0x69, 0x6c, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x61, 0x69, 0x72, 0x74, 0x69, 0x6d, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x62, 0x75, 0x6e,
	0x64, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x72, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x5f, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x72, 0x70, 0x63,
	0x5f, 0x6d, 0x61, 0x64, 0x65, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76,
	0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xae, 0x07, 0x0a, 0x0c,
	0x50, 0x65, 0x65, 0x72, 0x62, 0x69, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x12, 0x89, 0x01, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x65, 0x73, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x52,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e,
	0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5a, 0x92, 0x41,
	0x3e, 0x12, 0x15, 0x47, 0x65, 0x74, 0x20, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x46, 0x69,
	0x61, 0x74, 0x20, 0x72, 0x61, 0x74, 0x65, 0x73, 0x1a, 0x25, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68,
	0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x67, 0x65, 0x74, 0x20, 0x43, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x2f, 0x46, 0x69, 0x61, 0x74, 0x20, 0x72, 0x61, 0x74, 0x65, 0x73, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x65, 0x74, 0x2d, 0x72, 0x61, 0x74, 0x65, 0x73, 0x12, 0xbb, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7c, 0x92, 0x41, 0x5c, 0x12, 0x24, 0x47,
	0x65, 0x74, 0x20, 0x44, 0x61, 0x74, 0x61, 0x20, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x20, 0x66, 0x6f,
	0x72, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x20, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x1a, 0x34, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50,
	0x49, 0x20, 0x74, 0x6f, 0x20, 0x67, 0x65, 0x74, 0x20, 0x44, 0x61, 0x74, 0x61, 0x20, 0x50, 0x6c,
	0x61, 0x6e, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x20, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a,
	0x01, 0x2a, 0x22, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x64, 0x61, 0x74,
	0x61, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x12, 0x89, 0x01, 0x0a, 0x0a, 0x42, 0x75, 0x79, 0x41, 0x69,
	0x72, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x69, 0x72, 0x74, 0x69,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x41,
	0x69, 0x72, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x52,
	0x92, 0x41, 0x34, 0x12, 0x10, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x20, 0x61, 0x69,
	0x72, 0x74, 0x69, 0x6d, 0x65, 0x1a, 0x20, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20,
	0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x20,
	0x61, 0x69, 0x72, 0x74, 0x69, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a,
	0x22, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x75, 0x79, 0x2d, 0x61, 0x69, 0x72, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x97, 0x01, 0x0a, 0x09, 0x47, 0x65, 0x6e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x6e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x6e, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5d, 0x92,
	0x41, 0x40, 0x12, 0x16, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x20, 0x65, 0x73, 0x63,
	0x72, 0x6f, 0x77, 0x20, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x1a, 0x26, 0x55, 0x73, 0x65, 0x20,
	0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x20, 0x65, 0x73, 0x63, 0x72, 0x6f, 0x77, 0x20, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2d, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x9d, 0x01, 0x0a,
	0x0b, 0x4d, 0x61, 0x64, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x70,
	0x62, 0x2e, 0x4d, 0x61, 0x64, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x61, 0x64, 0x65, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5d, 0x92,
	0x41, 0x3e, 0x12, 0x15, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x74, 0x72, 0x61,
	0x64, 0x65, 0x20, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x25, 0x55, 0x73, 0x65, 0x20, 0x74,
	0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x20, 0x74, 0x72, 0x61, 0x64, 0x65, 0x20, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6d, 0x61, 0x64, 0x65, 0x2d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x8d, 0x01, 0x0a,
	0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x70,
	0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4d, 0x92,
	0x41, 0x2e, 0x12, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x1a, 0x1d, 0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20,
	0x74, 0x6f, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x74, 0x92, 0x41,
	0x43, 0x12, 0x41, 0x0a, 0x0d, 0x50, 0x65, 0x65, 0x72, 0x62, 0x69, 0x6c, 0x6c, 0x20, 0x55, 0x73,
	0x65, 0x72, 0x22, 0x2b, 0x0a, 0x0d, 0x50, 0x65, 0x65, 0x72, 0x62, 0x69, 0x6c, 0x6c, 0x20, 0x75,
	0x73, 0x65, 0x72, 0x1a, 0x1a, 0x70, 0x65, 0x65, 0x72, 0x62, 0x69, 0x6c, 0x6c, 0x6f, 0x66, 0x66,
	0x69, 0x63, 0x69, 0x61, 0x6c, 0x40, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x32,
	0x03, 0x31, 0x2e, 0x30, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x62, 0x75, 0x6b, 0x61, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x32, 0x31, 0x2f, 0x70, 0x65,
	0x65, 0x72, 0x62, 0x69, 0x6c, 0x6c, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_peerbill_user_proto_goTypes = []any{
	(*RateRequest)(nil),         // 0: pb.RateRequest
	(*DataLookupRequest)(nil),   // 1: pb.DataLookupRequest
	(*AirtimeRequest)(nil),      // 2: pb.AirtimeRequest
	(*GenWalletRequest)(nil),    // 3: pb.GenWalletRequest
	(*MadePaymentRequest)(nil),  // 4: pb.MadePaymentRequest
	(*UpdateOrderRequest)(nil),  // 5: pb.UpdateOrderRequest
	(*RateResponse)(nil),        // 6: pb.RateResponse
	(*DataLookupResponse)(nil),  // 7: pb.DataLookupResponse
	(*AirtimeResponse)(nil),     // 8: pb.AirtimeResponse
	(*GenWalletResponse)(nil),   // 9: pb.GenWalletResponse
	(*MadePaymentResponse)(nil), // 10: pb.MadePaymentResponse
	(*UpdateOrderResponse)(nil), // 11: pb.UpdateOrderResponse
}
var file_peerbill_user_proto_depIdxs = []int32{
	0,  // 0: pb.PeerbillUser.GetRates:input_type -> pb.RateRequest
	1,  // 1: pb.PeerbillUser.GetDataPlans:input_type -> pb.DataLookupRequest
	2,  // 2: pb.PeerbillUser.BuyAirtime:input_type -> pb.AirtimeRequest
	3,  // 3: pb.PeerbillUser.GenWallet:input_type -> pb.GenWalletRequest
	4,  // 4: pb.PeerbillUser.MadePayment:input_type -> pb.MadePaymentRequest
	5,  // 5: pb.PeerbillUser.UpdateOrder:input_type -> pb.UpdateOrderRequest
	6,  // 6: pb.PeerbillUser.GetRates:output_type -> pb.RateResponse
	7,  // 7: pb.PeerbillUser.GetDataPlans:output_type -> pb.DataLookupResponse
	8,  // 8: pb.PeerbillUser.BuyAirtime:output_type -> pb.AirtimeResponse
	9,  // 9: pb.PeerbillUser.GenWallet:output_type -> pb.GenWalletResponse
	10, // 10: pb.PeerbillUser.MadePayment:output_type -> pb.MadePaymentResponse
	11, // 11: pb.PeerbillUser.UpdateOrder:output_type -> pb.UpdateOrderResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_peerbill_user_proto_init() }
func file_peerbill_user_proto_init() {
	if File_peerbill_user_proto != nil {
		return
	}
	file_airtime_proto_init()
	file_data_bundle_proto_init()
	file_rate_proto_init()
	file_rpc_gen_wallet_proto_init()
	file_rpc_made_payment_proto_init()
	file_rpc_update_order_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_peerbill_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_peerbill_user_proto_goTypes,
		DependencyIndexes: file_peerbill_user_proto_depIdxs,
	}.Build()
	File_peerbill_user_proto = out.File
	file_peerbill_user_proto_rawDesc = nil
	file_peerbill_user_proto_goTypes = nil
	file_peerbill_user_proto_depIdxs = nil
}
