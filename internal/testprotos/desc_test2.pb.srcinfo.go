// Code generated by protoc-gen-gosrcinfo. DO NOT EDIT.
// source: desc_test2.proto

package testprotos

import "github.com/jhump/protoreflect/desc/sourceinfo"
import "google.golang.org/protobuf/proto"
import "google.golang.org/protobuf/types/descriptorpb"

var srcInfo_desc_test2 = []byte{
	0x0a, 0x06, 0x12, 0x04, 0x00, 0x00, 0x2a, 0x01, 0x0a, 0x08, 0x0a, 0x01, 0x0c, 0x12, 0x03, 0x00,
	0x00, 0x12, 0x0a, 0x08, 0x0a, 0x01, 0x08, 0x12, 0x03, 0x02, 0x00, 0x48, 0x0a, 0x09, 0x0a, 0x02,
	0x08, 0x0b, 0x12, 0x03, 0x02, 0x00, 0x48, 0x0a, 0x08, 0x0a, 0x01, 0x02, 0x12, 0x03, 0x04, 0x00,
	0x13, 0x0a, 0x09, 0x0a, 0x02, 0x03, 0x00, 0x12, 0x03, 0x06, 0x00, 0x1a, 0x0a, 0x09, 0x0a, 0x02,
	0x03, 0x01, 0x12, 0x03, 0x07, 0x00, 0x21, 0x0a, 0x09, 0x0a, 0x02, 0x03, 0x02, 0x12, 0x03, 0x08,
	0x00, 0x25, 0x0a, 0x0a, 0x0a, 0x02, 0x04, 0x00, 0x12, 0x04, 0x0a, 0x00, 0x19, 0x01, 0x0a, 0x0a,
	0x0a, 0x03, 0x04, 0x00, 0x01, 0x12, 0x03, 0x0a, 0x08, 0x10, 0x0a, 0x0b, 0x0a, 0x04, 0x04, 0x00,
	0x02, 0x00, 0x12, 0x03, 0x0b, 0x08, 0x23, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00, 0x02, 0x00, 0x04,
	0x12, 0x03, 0x0b, 0x08, 0x10, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00, 0x02, 0x00, 0x06, 0x12, 0x03,
	0x0b, 0x11, 0x1c, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00, 0x02, 0x00, 0x01, 0x12, 0x03, 0x0b, 0x1d,
	0x1e, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00, 0x02, 0x00, 0x03, 0x12, 0x03, 0x0b, 0x21, 0x22, 0x0a,
	0x0b, 0x0a, 0x04, 0x04, 0x00, 0x02, 0x01, 0x12, 0x03, 0x0c, 0x08, 0x2a, 0x0a, 0x0c, 0x0a, 0x05,
	0x04, 0x00, 0x02, 0x01, 0x04, 0x12, 0x03, 0x0c, 0x08, 0x10, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00,
	0x02, 0x01, 0x06, 0x12, 0x03, 0x0c, 0x11, 0x23, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00, 0x02, 0x01,
	0x01, 0x12, 0x03, 0x0c, 0x24, 0x25, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00, 0x02, 0x01, 0x03, 0x12,
	0x03, 0x0c, 0x28, 0x29, 0x0a, 0x0c, 0x0a, 0x04, 0x04, 0x00, 0x08, 0x00, 0x12, 0x04, 0x0d, 0x08,
	0x10, 0x09, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00, 0x08, 0x00, 0x01, 0x12, 0x03, 0x0d, 0x0e, 0x11,
	0x0a, 0x0b, 0x0a, 0x04, 0x04, 0x00, 0x02, 0x02, 0x12, 0x03, 0x0e, 0x10, 0x31, 0x0a, 0x0c, 0x0a,
	0x05, 0x04, 0x00, 0x02, 0x02, 0x06, 0x12, 0x03, 0x0e, 0x10, 0x29, 0x0a, 0x0c, 0x0a, 0x05, 0x04,
	0x00, 0x02, 0x02, 0x01, 0x12, 0x03, 0x0e, 0x2a, 0x2c, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00, 0x02,
	0x02, 0x03, 0x12, 0x03, 0x0e, 0x2f, 0x30, 0x0a, 0x0b, 0x0a, 0x04, 0x04, 0x00, 0x02, 0x03, 0x12,
	0x03, 0x0f, 0x10, 0x2e, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00, 0x02, 0x03, 0x06, 0x12, 0x03, 0x0f,
	0x10, 0x26, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00, 0x02, 0x03, 0x01, 0x12, 0x03, 0x0f, 0x27, 0x29,
	0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00, 0x02, 0x03, 0x03, 0x12, 0x03, 0x0f, 0x2c, 0x2d, 0x0a, 0x0b,
	0x0a, 0x04, 0x04, 0x00, 0x02, 0x04, 0x12, 0x03, 0x11, 0x08, 0x31, 0x0a, 0x0c, 0x0a, 0x05, 0x04,
	0x00, 0x02, 0x04, 0x04, 0x12, 0x03, 0x11, 0x08, 0x10, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00, 0x02,
	0x04, 0x06, 0x12, 0x03, 0x11, 0x11, 0x2a, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00, 0x02, 0x04, 0x01,
	0x12, 0x03, 0x11, 0x2b, 0x2c, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00, 0x02, 0x04, 0x03, 0x12, 0x03,
	0x11, 0x2f, 0x30, 0x0a, 0x0b, 0x0a, 0x04, 0x04, 0x00, 0x02, 0x05, 0x12, 0x03, 0x12, 0x08, 0x41,
	0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00, 0x02, 0x05, 0x04, 0x12, 0x03, 0x12, 0x08, 0x10, 0x0a, 0x0c,
	0x0a, 0x05, 0x04, 0x00, 0x02, 0x05, 0x06, 0x12, 0x03, 0x12, 0x11, 0x27, 0x0a, 0x0c, 0x0a, 0x05,
	0x04, 0x00, 0x02, 0x05, 0x01, 0x12, 0x03, 0x12, 0x28, 0x29, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00,
	0x02, 0x05, 0x03, 0x12, 0x03, 0x12, 0x2c, 0x2d, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00, 0x02, 0x05,
	0x08, 0x12, 0x03, 0x12, 0x2e, 0x40, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00, 0x02, 0x05, 0x07, 0x12,
	0x03, 0x12, 0x39, 0x3f, 0x0a, 0x0b, 0x0a, 0x04, 0x04, 0x00, 0x02, 0x06, 0x12, 0x03, 0x13, 0x08,
	0x32, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00, 0x02, 0x06, 0x04, 0x12, 0x03, 0x13, 0x08, 0x10, 0x0a,
	0x0c, 0x0a, 0x05, 0x04, 0x00, 0x02, 0x06, 0x05, 0x12, 0x03, 0x13, 0x11, 0x17, 0x0a, 0x0c, 0x0a,
	0x05, 0x04, 0x00, 0x02, 0x06, 0x01, 0x12, 0x03, 0x13, 0x18, 0x19, 0x0a, 0x0c, 0x0a, 0x05, 0x04,
	0x00, 0x02, 0x06, 0x03, 0x12, 0x03, 0x13, 0x1c, 0x1d, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00, 0x02,
	0x06, 0x08, 0x12, 0x03, 0x13, 0x1e, 0x31, 0x0a, 0x0d, 0x0a, 0x06, 0x04, 0x00, 0x02, 0x06, 0x08,
	0x03, 0x12, 0x03, 0x13, 0x1f, 0x30, 0x0a, 0x0c, 0x0a, 0x04, 0x04, 0x00, 0x08, 0x01, 0x12, 0x04,
	0x14, 0x08, 0x18, 0x09, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00, 0x08, 0x01, 0x01, 0x12, 0x03, 0x14,
	0x0e, 0x11, 0x0a, 0x0b, 0x0a, 0x04, 0x04, 0x00, 0x02, 0x07, 0x12, 0x03, 0x15, 0x10, 0x1d, 0x0a,
	0x0c, 0x0a, 0x05, 0x04, 0x00, 0x02, 0x07, 0x05, 0x12, 0x03, 0x15, 0x10, 0x15, 0x0a, 0x0c, 0x0a,
	0x05, 0x04, 0x00, 0x02, 0x07, 0x01, 0x12, 0x03, 0x15, 0x16, 0x18, 0x0a, 0x0c, 0x0a, 0x05, 0x04,
	0x00, 0x02, 0x07, 0x03, 0x12, 0x03, 0x15, 0x1b, 0x1c, 0x0a, 0x0b, 0x0a, 0x04, 0x04, 0x00, 0x02,
	0x08, 0x12, 0x03, 0x16, 0x10, 0x1e, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00, 0x02, 0x08, 0x05, 0x12,
	0x03, 0x16, 0x10, 0x16, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00, 0x02, 0x08, 0x01, 0x12, 0x03, 0x16,
	0x17, 0x19, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00, 0x02, 0x08, 0x03, 0x12, 0x03, 0x16, 0x1c, 0x1d,
	0x0a, 0x0b, 0x0a, 0x04, 0x04, 0x00, 0x02, 0x09, 0x12, 0x03, 0x17, 0x10, 0x1f, 0x0a, 0x0c, 0x0a,
	0x05, 0x04, 0x00, 0x02, 0x09, 0x05, 0x12, 0x03, 0x17, 0x10, 0x16, 0x0a, 0x0c, 0x0a, 0x05, 0x04,
	0x00, 0x02, 0x09, 0x01, 0x12, 0x03, 0x17, 0x17, 0x19, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x00, 0x02,
	0x09, 0x03, 0x12, 0x03, 0x17, 0x1c, 0x1e, 0x0a, 0x0a, 0x0a, 0x02, 0x04, 0x01, 0x12, 0x04, 0x1b,
	0x00, 0x1d, 0x01, 0x0a, 0x0a, 0x0a, 0x03, 0x04, 0x01, 0x01, 0x12, 0x03, 0x1b, 0x08, 0x17, 0x0a,
	0x0b, 0x0a, 0x04, 0x04, 0x01, 0x02, 0x00, 0x12, 0x03, 0x1c, 0x08, 0x36, 0x0a, 0x0c, 0x0a, 0x05,
	0x04, 0x01, 0x02, 0x00, 0x04, 0x12, 0x03, 0x1c, 0x08, 0x10, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x01,
	0x02, 0x00, 0x06, 0x12, 0x03, 0x1c, 0x11, 0x2c, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x01, 0x02, 0x00,
	0x01, 0x12, 0x03, 0x1c, 0x2d, 0x31, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x01, 0x02, 0x00, 0x03, 0x12,
	0x03, 0x1c, 0x34, 0x35, 0x0a, 0x0a, 0x0a, 0x02, 0x04, 0x02, 0x12, 0x04, 0x1f, 0x00, 0x21, 0x01,
	0x0a, 0x0a, 0x0a, 0x03, 0x04, 0x02, 0x01, 0x12, 0x03, 0x1f, 0x08, 0x0f, 0x0a, 0x0b, 0x0a, 0x04,
	0x04, 0x02, 0x02, 0x00, 0x12, 0x03, 0x20, 0x08, 0x39, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x02, 0x02,
	0x00, 0x04, 0x12, 0x03, 0x20, 0x08, 0x10, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x02, 0x02, 0x00, 0x06,
	0x12, 0x03, 0x20, 0x11, 0x2c, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x02, 0x02, 0x00, 0x01, 0x12, 0x03,
	0x20, 0x2d, 0x34, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x02, 0x02, 0x00, 0x03, 0x12, 0x03, 0x20, 0x37,
	0x38, 0x0a, 0x09, 0x0a, 0x01, 0x07, 0x12, 0x04, 0x23, 0x00, 0x2a, 0x01, 0x0a, 0x09, 0x0a, 0x02,
	0x07, 0x00, 0x12, 0x03, 0x24, 0x08, 0x24, 0x0a, 0x0a, 0x0a, 0x03, 0x07, 0x00, 0x02, 0x12, 0x03,
	0x23, 0x07, 0x0f, 0x0a, 0x0a, 0x0a, 0x03, 0x07, 0x00, 0x04, 0x12, 0x03, 0x24, 0x08, 0x10, 0x0a,
	0x0a, 0x0a, 0x03, 0x07, 0x00, 0x06, 0x12, 0x03, 0x24, 0x11, 0x19, 0x0a, 0x0a, 0x0a, 0x03, 0x07,
	0x00, 0x01, 0x12, 0x03, 0x24, 0x1a, 0x1d, 0x0a, 0x0a, 0x0a, 0x03, 0x07, 0x00, 0x03, 0x12, 0x03,
	0x24, 0x20, 0x23, 0x0a, 0x0a, 0x0a, 0x02, 0x07, 0x01, 0x12, 0x04, 0x26, 0x08, 0x29, 0x09, 0x0a,
	0x0a, 0x0a, 0x03, 0x07, 0x01, 0x02, 0x12, 0x03, 0x23, 0x07, 0x0f, 0x0a, 0x0a, 0x0a, 0x03, 0x07,
	0x01, 0x04, 0x12, 0x03, 0x26, 0x08, 0x10, 0x0a, 0x0a, 0x0a, 0x03, 0x07, 0x01, 0x05, 0x12, 0x03,
	0x26, 0x11, 0x16, 0x0a, 0x0a, 0x0a, 0x03, 0x07, 0x01, 0x01, 0x12, 0x03, 0x26, 0x17, 0x1d, 0x0a,
	0x0a, 0x0a, 0x03, 0x07, 0x01, 0x03, 0x12, 0x03, 0x26, 0x20, 0x23, 0x0a, 0x0a, 0x0a, 0x02, 0x04,
	0x03, 0x12, 0x04, 0x26, 0x08, 0x29, 0x09, 0x0a, 0x0a, 0x0a, 0x03, 0x04, 0x03, 0x01, 0x12, 0x03,
	0x26, 0x17, 0x1d, 0x0a, 0x0a, 0x0a, 0x03, 0x07, 0x01, 0x06, 0x12, 0x03, 0x26, 0x17, 0x1d, 0x0a,
	0x0b, 0x0a, 0x04, 0x04, 0x03, 0x02, 0x00, 0x12, 0x03, 0x27, 0x10, 0x2e, 0x0a, 0x0c, 0x0a, 0x05,
	0x04, 0x03, 0x02, 0x00, 0x04, 0x12, 0x03, 0x27, 0x10, 0x18, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x03,
	0x02, 0x00, 0x05, 0x12, 0x03, 0x27, 0x19, 0x1e, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x03, 0x02, 0x00,
	0x01, 0x12, 0x03, 0x27, 0x1f, 0x26, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x03, 0x02, 0x00, 0x03, 0x12,
	0x03, 0x27, 0x29, 0x2d, 0x0a, 0x0b, 0x0a, 0x04, 0x04, 0x03, 0x02, 0x01, 0x12, 0x03, 0x28, 0x10,
	0x2f, 0x0a, 0x0c, 0x0a, 0x05, 0x04, 0x03, 0x02, 0x01, 0x04, 0x12, 0x03, 0x28, 0x10, 0x18, 0x0a,
	0x0c, 0x0a, 0x05, 0x04, 0x03, 0x02, 0x01, 0x05, 0x12, 0x03, 0x28, 0x19, 0x1f, 0x0a, 0x0c, 0x0a,
	0x05, 0x04, 0x03, 0x02, 0x01, 0x01, 0x12, 0x03, 0x28, 0x20, 0x27, 0x0a, 0x0c, 0x0a, 0x05, 0x04,
	0x03, 0x02, 0x01, 0x03, 0x12, 0x03, 0x28, 0x2a, 0x2e,
}

func init() {
	var si descriptorpb.SourceCodeInfo
	if err := proto.Unmarshal(srcInfo_desc_test2, &si); err != nil {
		panic(err)
	}
	sourceinfo.RegisterSourceInfo("desc_test2.proto", &si)
}