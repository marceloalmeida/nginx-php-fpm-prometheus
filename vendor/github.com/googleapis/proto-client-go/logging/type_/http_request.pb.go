// Code generated by protoc-gen-go.
// source: google/logging/type/http_request.proto
// DO NOT EDIT!

/*
Package type_ is a generated protocol buffer package.

It is generated from these files:
	google/logging/type/http_request.proto
	google/logging/type/log_severity.proto

It has these top-level messages:
	HttpRequest
*/
package type_

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/googleapis/proto-client-go/api"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// A common proto for logging HTTP requests.
//
type HttpRequest struct {
	// The request method. Examples: `"GET"`, `"HEAD"`, `"PUT"`, `"POST"`.
	RequestMethod string `protobuf:"bytes,1,opt,name=request_method,json=requestMethod" json:"request_method,omitempty"`
	// The scheme (http, https), the host name, the path and the query
	// portion of the URL that was requested.
	// Example: `"http://example.com/some/info?color=red"`.
	RequestUrl string `protobuf:"bytes,2,opt,name=request_url,json=requestUrl" json:"request_url,omitempty"`
	// The size of the HTTP request message in bytes, including the request
	// headers and the request body.
	RequestSize int64 `protobuf:"varint,3,opt,name=request_size,json=requestSize" json:"request_size,omitempty"`
	// The response code indicating the status of response.
	// Examples: 200, 404.
	Status int32 `protobuf:"varint,4,opt,name=status" json:"status,omitempty"`
	// The size of the HTTP response message sent back to the client, in bytes,
	// including the response headers and the response body.
	ResponseSize int64 `protobuf:"varint,5,opt,name=response_size,json=responseSize" json:"response_size,omitempty"`
	// The user agent sent by the client. Example:
	// `"Mozilla/4.0 (compatible; MSIE 6.0; Windows 98; Q312461; .NET CLR 1.0.3705)"`.
	UserAgent string `protobuf:"bytes,6,opt,name=user_agent,json=userAgent" json:"user_agent,omitempty"`
	// The IP address (IPv4 or IPv6) of the client that issued the HTTP
	// request. Examples: `"192.168.1.1"`, `"FE80::0202:B3FF:FE1E:8329"`.
	RemoteIp string `protobuf:"bytes,7,opt,name=remote_ip,json=remoteIp" json:"remote_ip,omitempty"`
	// The referer URL of the request, as defined in
	// [HTTP/1.1 Header Field Definitions](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html).
	Referer string `protobuf:"bytes,8,opt,name=referer" json:"referer,omitempty"`
	// Whether or not an entity was served from cache
	// (with or without validation).
	CacheHit bool `protobuf:"varint,9,opt,name=cache_hit,json=cacheHit" json:"cache_hit,omitempty"`
	// Whether or not the response was validated with the origin server before
	// being served from cache. This field is only meaningful if `cache_hit` is
	// True.
	ValidatedWithOriginServer bool `protobuf:"varint,10,opt,name=validated_with_origin_server,json=validatedWithOriginServer" json:"validated_with_origin_server,omitempty"`
}

func (m *HttpRequest) Reset()                    { *m = HttpRequest{} }
func (m *HttpRequest) String() string            { return proto.CompactTextString(m) }
func (*HttpRequest) ProtoMessage()               {}
func (*HttpRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*HttpRequest)(nil), "google.logging.type.HttpRequest")
}

var fileDescriptor0 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x92, 0xcd, 0x4e, 0xe3, 0x30,
	0x14, 0x85, 0xd5, 0x76, 0xda, 0x26, 0x6e, 0x67, 0x34, 0xf2, 0x48, 0x33, 0x1e, 0x28, 0xe2, 0x4f,
	0x20, 0x36, 0x4d, 0x16, 0x7d, 0x00, 0x44, 0x57, 0x65, 0x81, 0xa8, 0x52, 0x21, 0x24, 0x36, 0x56,
	0x9a, 0x5e, 0x1c, 0x4b, 0x69, 0x1c, 0x6c, 0xa7, 0x08, 0xde, 0x98, 0xb7, 0xe0, 0xc6, 0x49, 0x10,
	0x48, 0x2c, 0x7d, 0xee, 0x77, 0xce, 0x75, 0x72, 0x4c, 0xce, 0x85, 0x52, 0x22, 0x83, 0x30, 0x53,
	0x42, 0xc8, 0x5c, 0x84, 0xf6, 0xa5, 0x80, 0x30, 0xb5, 0xb6, 0xe0, 0x1a, 0x9e, 0x4a, 0x30, 0x36,
	0x28, 0xb4, 0xb2, 0x8a, 0xfe, 0xa9, 0xb9, 0xa0, 0xe1, 0x82, 0x8a, 0xdb, 0x9b, 0x34, 0xe6, 0xb8,
	0x90, 0x61, 0x9c, 0xe7, 0xca, 0xc6, 0x56, 0xaa, 0xdc, 0xd4, 0x96, 0x93, 0xb7, 0x2e, 0x19, 0x2d,
	0x30, 0x29, 0xaa, 0x83, 0xe8, 0x19, 0xf9, 0xd5, 0x64, 0xf2, 0x2d, 0xd8, 0x54, 0x6d, 0x58, 0xe7,
	0xa8, 0x73, 0xe1, 0x47, 0x3f, 0x1b, 0xf5, 0xc6, 0x89, 0xf4, 0x90, 0x8c, 0x5a, 0xac, 0xd4, 0x19,
	0xeb, 0x3a, 0x86, 0x34, 0xd2, 0x9d, 0xce, 0xe8, 0x31, 0x19, 0xb7, 0x80, 0x91, 0xaf, 0xc0, 0x7a,
	0x48, 0xf4, 0xa2, 0xd6, 0xb4, 0x42, 0x89, 0xfe, 0x25, 0x03, 0x83, 0x97, 0x29, 0x0d, 0xfb, 0x81,
	0xc3, 0x7e, 0xd4, 0x9c, 0xe8, 0x29, 0xc1, 0x65, 0xa6, 0xc0, 0x3b, 0x42, 0xed, 0xed, 0x3b, 0xef,
	0xb8, 0x15, 0x9d, 0xf9, 0x80, 0x90, 0xd2, 0x80, 0xe6, 0xb1, 0x80, 0xdc, 0xb2, 0x81, 0xdb, 0xef,
	0x57, 0xca, 0x55, 0x25, 0xd0, 0x7d, 0xe2, 0x6b, 0xd8, 0x2a, 0x0b, 0x5c, 0x16, 0x6c, 0xe8, 0xa6,
	0x5e, 0x2d, 0x5c, 0x17, 0x94, 0x91, 0xa1, 0x86, 0x47, 0xd0, 0xa0, 0x99, 0xe7, 0x46, 0xed, 0xb1,
	0xb2, 0x25, 0x71, 0x92, 0x02, 0x4f, 0xa5, 0x65, 0x3e, 0xce, 0xbc, 0xc8, 0x73, 0xc2, 0x42, 0x5a,
	0x7a, 0x49, 0x26, 0xbb, 0x38, 0x93, 0x9b, 0xd8, 0xc2, 0x86, 0x3f, 0x4b, 0x9b, 0x72, 0xa5, 0x25,
	0xfe, 0x67, 0x8e, 0x5b, 0x77, 0x98, 0x45, 0x1c, 0xff, 0xff, 0x83, 0xb9, 0x47, 0xe4, 0xd6, 0x11,
	0x2b, 0x07, 0xcc, 0xd7, 0xe4, 0x5f, 0xa2, 0xb6, 0xc1, 0x37, 0x25, 0xcd, 0x7f, 0x7f, 0xea, 0x60,
	0x59, 0x15, 0xb3, 0xec, 0x3c, 0xcc, 0x04, 0x06, 0x94, 0xeb, 0x00, 0x3d, 0x61, 0xed, 0xc1, 0x0a,
	0x4d, 0xe8, 0x7a, 0x9b, 0x26, 0x99, 0xc4, 0xef, 0x9c, 0x0a, 0xf5, 0xe5, 0x51, 0xf0, 0xf5, 0xc0,
	0x8d, 0x67, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb9, 0xe5, 0x79, 0xd4, 0x33, 0x02, 0x00, 0x00,
}
