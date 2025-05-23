// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: service_analyzer.proto

package analyzer

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_analyzer_proto protoreflect.FileDescriptor

var file_service_analyzer_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a,
	0x65, 0x72, 0x1a, 0x0f, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x74, 0x6f, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0x9b, 0x01, 0x0a, 0x08, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72,
	0x12, 0x48, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x1e, 0x2e, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x7a, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x7a, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x04, 0x53, 0x74,
	0x6f, 0x70, 0x12, 0x1d, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x2e, 0x53, 0x74,
	0x6f, 0x70, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f,
	0x70, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6f, 0x72, 0x67, 0x77, 0x61, 0x74, 0x73, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var file_service_analyzer_proto_goTypes = []any{
	(*StartAnalyzerRequest)(nil),  // 0: analyzer.StartAnalyzerRequest
	(*StopAnalyzerRequest)(nil),   // 1: analyzer.StopAnalyzerRequest
	(*StartAnalyzerResponse)(nil), // 2: analyzer.StartAnalyzerResponse
	(*StopAnalyzerResponse)(nil),  // 3: analyzer.StopAnalyzerResponse
}
var file_service_analyzer_proto_depIdxs = []int32{
	0, // 0: analyzer.Analyzer.Start:input_type -> analyzer.StartAnalyzerRequest
	1, // 1: analyzer.Analyzer.Stop:input_type -> analyzer.StopAnalyzerRequest
	2, // 2: analyzer.Analyzer.Start:output_type -> analyzer.StartAnalyzerResponse
	3, // 3: analyzer.Analyzer.Stop:output_type -> analyzer.StopAnalyzerResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_analyzer_proto_init() }
func file_service_analyzer_proto_init() {
	if File_service_analyzer_proto != nil {
		return
	}
	file_rpc_start_proto_init()
	file_rpc_stop_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_service_analyzer_proto_rawDesc), len(file_service_analyzer_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_analyzer_proto_goTypes,
		DependencyIndexes: file_service_analyzer_proto_depIdxs,
	}.Build()
	File_service_analyzer_proto = out.File
	file_service_analyzer_proto_goTypes = nil
	file_service_analyzer_proto_depIdxs = nil
}
