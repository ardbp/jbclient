// Code generated by protoc-gen-go. DO NOT EDIT.
// source: job.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type JobStatus int32

const (
	JobStatus_UNDEFINED JobStatus = 0
	// Running job
	JobStatus_STARTED JobStatus = 1
	// Finished successfully
	JobStatus_FINISHED JobStatus = 2
	// Manually stop by stop api
	JobStatus_MANUAL_STOP JobStatus = 3
	// Error
	JobStatus_ERROR JobStatus = 4
)

var JobStatus_name = map[int32]string{
	0: "UNDEFINED",
	1: "STARTED",
	2: "FINISHED",
	3: "MANUAL_STOP",
	4: "ERROR",
}

var JobStatus_value = map[string]int32{
	"UNDEFINED":   0,
	"STARTED":     1,
	"FINISHED":    2,
	"MANUAL_STOP": 3,
	"ERROR":       4,
}

func (x JobStatus) String() string {
	return proto.EnumName(JobStatus_name, int32(x))
}

func (JobStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{0}
}

type StartJobRequest struct {
	// Command path to be run, required non empty string
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Command arguments
	Args []string `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	// Process environment
	Envs []string `protobuf:"bytes,3,rep,name=envs,proto3" json:"envs,omitempty"`
	// Directory to run the command
	Directory            string   `protobuf:"bytes,4,opt,name=directory,proto3" json:"directory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartJobRequest) Reset()         { *m = StartJobRequest{} }
func (m *StartJobRequest) String() string { return proto.CompactTextString(m) }
func (*StartJobRequest) ProtoMessage()    {}
func (*StartJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{0}
}

func (m *StartJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartJobRequest.Unmarshal(m, b)
}
func (m *StartJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartJobRequest.Marshal(b, m, deterministic)
}
func (m *StartJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartJobRequest.Merge(m, src)
}
func (m *StartJobRequest) XXX_Size() int {
	return xxx_messageInfo_StartJobRequest.Size(m)
}
func (m *StartJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartJobRequest proto.InternalMessageInfo

func (m *StartJobRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *StartJobRequest) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *StartJobRequest) GetEnvs() []string {
	if m != nil {
		return m.Envs
	}
	return nil
}

func (m *StartJobRequest) GetDirectory() string {
	if m != nil {
		return m.Directory
	}
	return ""
}

type StartJobResponse struct {
	// Job object
	Job                  *Job     `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartJobResponse) Reset()         { *m = StartJobResponse{} }
func (m *StartJobResponse) String() string { return proto.CompactTextString(m) }
func (*StartJobResponse) ProtoMessage()    {}
func (*StartJobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{1}
}

func (m *StartJobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartJobResponse.Unmarshal(m, b)
}
func (m *StartJobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartJobResponse.Marshal(b, m, deterministic)
}
func (m *StartJobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartJobResponse.Merge(m, src)
}
func (m *StartJobResponse) XXX_Size() int {
	return xxx_messageInfo_StartJobResponse.Size(m)
}
func (m *StartJobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartJobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartJobResponse proto.InternalMessageInfo

func (m *StartJobResponse) GetJob() *Job {
	if m != nil {
		return m.Job
	}
	return nil
}

type IDRequest struct {
	// Identifier of the job, UUID generated string
	JobId                string   `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IDRequest) Reset()         { *m = IDRequest{} }
func (m *IDRequest) String() string { return proto.CompactTextString(m) }
func (*IDRequest) ProtoMessage()    {}
func (*IDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{2}
}

func (m *IDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IDRequest.Unmarshal(m, b)
}
func (m *IDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IDRequest.Marshal(b, m, deterministic)
}
func (m *IDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDRequest.Merge(m, src)
}
func (m *IDRequest) XXX_Size() int {
	return xxx_messageInfo_IDRequest.Size(m)
}
func (m *IDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IDRequest proto.InternalMessageInfo

func (m *IDRequest) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

// StopResponse
type StopResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopResponse) Reset()         { *m = StopResponse{} }
func (m *StopResponse) String() string { return proto.CompactTextString(m) }
func (*StopResponse) ProtoMessage()    {}
func (*StopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{3}
}

func (m *StopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopResponse.Unmarshal(m, b)
}
func (m *StopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopResponse.Marshal(b, m, deterministic)
}
func (m *StopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopResponse.Merge(m, src)
}
func (m *StopResponse) XXX_Size() int {
	return xxx_messageInfo_StopResponse.Size(m)
}
func (m *StopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StopResponse proto.InternalMessageInfo

type Job struct {
	// Command path to be run, required non empty string
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Command arguments
	Args []string `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	// Process environment
	Envs []string `protobuf:"bytes,3,rep,name=envs,proto3" json:"envs,omitempty"`
	// Directory to run the command
	Directory string `protobuf:"bytes,4,opt,name=directory,proto3" json:"directory,omitempty"`
	// job status
	Status JobStatus `protobuf:"varint,5,opt,name=status,proto3,enum=api.JobStatus" json:"status,omitempty"`
	// id string identifier for job, UUID generated string
	JobId string `protobuf:"bytes,6,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	// exit code for job
	ExitCode             int32    `protobuf:"varint,7,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Job) Reset()         { *m = Job{} }
func (m *Job) String() string { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()    {}
func (*Job) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{4}
}

func (m *Job) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Job.Unmarshal(m, b)
}
func (m *Job) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Job.Marshal(b, m, deterministic)
}
func (m *Job) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Job.Merge(m, src)
}
func (m *Job) XXX_Size() int {
	return xxx_messageInfo_Job.Size(m)
}
func (m *Job) XXX_DiscardUnknown() {
	xxx_messageInfo_Job.DiscardUnknown(m)
}

var xxx_messageInfo_Job proto.InternalMessageInfo

func (m *Job) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Job) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Job) GetEnvs() []string {
	if m != nil {
		return m.Envs
	}
	return nil
}

func (m *Job) GetDirectory() string {
	if m != nil {
		return m.Directory
	}
	return ""
}

func (m *Job) GetStatus() JobStatus {
	if m != nil {
		return m.Status
	}
	return JobStatus_UNDEFINED
}

func (m *Job) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *Job) GetExitCode() int32 {
	if m != nil {
		return m.ExitCode
	}
	return 0
}

type JobOutput struct {
	// job output this will only provide stdout
	Output string `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	// job error if any this will only provide stderr
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobOutput) Reset()         { *m = JobOutput{} }
func (m *JobOutput) String() string { return proto.CompactTextString(m) }
func (*JobOutput) ProtoMessage()    {}
func (*JobOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{5}
}

func (m *JobOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobOutput.Unmarshal(m, b)
}
func (m *JobOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobOutput.Marshal(b, m, deterministic)
}
func (m *JobOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobOutput.Merge(m, src)
}
func (m *JobOutput) XXX_Size() int {
	return xxx_messageInfo_JobOutput.Size(m)
}
func (m *JobOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_JobOutput.DiscardUnknown(m)
}

var xxx_messageInfo_JobOutput proto.InternalMessageInfo

func (m *JobOutput) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *JobOutput) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterEnum("api.JobStatus", JobStatus_name, JobStatus_value)
	proto.RegisterType((*StartJobRequest)(nil), "api.StartJobRequest")
	proto.RegisterType((*StartJobResponse)(nil), "api.StartJobResponse")
	proto.RegisterType((*IDRequest)(nil), "api.IDRequest")
	proto.RegisterType((*StopResponse)(nil), "api.StopResponse")
	proto.RegisterType((*Job)(nil), "api.Job")
	proto.RegisterType((*JobOutput)(nil), "api.JobOutput")
}

func init() { proto.RegisterFile("job.proto", fileDescriptor_f32c477d91a04ead) }

var fileDescriptor_f32c477d91a04ead = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0x93, 0xa6, 0x49, 0x9b, 0xe9, 0xd2, 0x0d, 0xa3, 0x5d, 0x14, 0x15, 0x0e, 0x95, 0x85,
	0x50, 0xc5, 0x21, 0x42, 0x45, 0x1c, 0xf6, 0x58, 0x91, 0x2c, 0xa4, 0x82, 0x16, 0x92, 0xee, 0xb9,
	0x4a, 0x1a, 0x0b, 0x52, 0xc4, 0x4e, 0x70, 0xdc, 0x15, 0x3c, 0x1b, 0xaf, 0xc1, 0x03, 0x21, 0x3b,
	0xe9, 0x1f, 0x55, 0x5c, 0xb9, 0x8d, 0x7f, 0xe3, 0x19, 0x7f, 0xe3, 0x6f, 0xc0, 0xdd, 0x52, 0x1e,
	0x54, 0x82, 0x24, 0xa1, 0x95, 0x55, 0x25, 0xfb, 0x06, 0x97, 0xa9, 0xcc, 0x84, 0x9c, 0x53, 0x9e,
	0xf0, 0x1f, 0x3b, 0x5e, 0x4b, 0x44, 0xe8, 0x56, 0x99, 0xfc, 0xea, 0x9b, 0x63, 0x73, 0xe2, 0x26,
	0x3a, 0x56, 0x2c, 0x13, 0x5f, 0x6a, 0xbf, 0x33, 0xb6, 0x14, 0x53, 0xb1, 0x62, 0xfc, 0xfe, 0xa1,
	0xf6, 0xad, 0x86, 0xa9, 0x18, 0x9f, 0x81, 0x5b, 0x94, 0x82, 0x6f, 0x24, 0x89, 0x5f, 0x7e, 0x57,
	0x37, 0x38, 0x02, 0x16, 0x80, 0x77, 0x7c, 0xac, 0xae, 0xe8, 0xbe, 0xe6, 0x38, 0x02, 0x6b, 0x4b,
	0xb9, 0x7e, 0x6c, 0x30, 0xed, 0x07, 0x59, 0x55, 0x06, 0x2a, 0xad, 0x20, 0x63, 0xe0, 0xc6, 0xe1,
	0x5e, 0xd6, 0x35, 0x38, 0x5b, 0xca, 0xd7, 0x65, 0xd1, 0x0a, 0xb3, 0xb7, 0x94, 0xc7, 0x05, 0x1b,
	0xc2, 0x45, 0x2a, 0xa9, 0xda, 0xf7, 0x63, 0xbf, 0x4d, 0xb0, 0xe6, 0x94, 0xff, 0xbf, 0x29, 0xf0,
	0x05, 0x38, 0xb5, 0xcc, 0xe4, 0xae, 0xf6, 0xed, 0xb1, 0x39, 0x19, 0x4e, 0x87, 0x7b, 0xd1, 0xa9,
	0xa6, 0x49, 0x9b, 0x3d, 0x11, 0xec, 0x9c, 0x08, 0xc6, 0xa7, 0xe0, 0xf2, 0x9f, 0xa5, 0x5c, 0x6f,
	0xa8, 0xe0, 0x7e, 0x6f, 0x6c, 0x4e, 0xec, 0xa4, 0xaf, 0xc0, 0x5b, 0x2a, 0x38, 0xbb, 0x01, 0x77,
	0x4e, 0xf9, 0x72, 0x27, 0xab, 0x9d, 0xc4, 0x27, 0xe0, 0x90, 0x8e, 0xda, 0x21, 0xda, 0x13, 0x5e,
	0x81, 0xcd, 0x85, 0x20, 0xe1, 0x77, 0x9a, 0xbe, 0xfa, 0xf0, 0xf2, 0xb3, 0x2e, 0x6d, 0x34, 0xe0,
	0x23, 0x70, 0xef, 0x16, 0x61, 0x74, 0x1b, 0x2f, 0xa2, 0xd0, 0x33, 0x70, 0x00, 0xbd, 0x74, 0x35,
	0x4b, 0x56, 0x51, 0xe8, 0x99, 0x78, 0x01, 0xfd, 0xdb, 0x78, 0x11, 0xa7, 0xef, 0xa3, 0xd0, 0xeb,
	0xe0, 0x25, 0x0c, 0x3e, 0xce, 0x16, 0x77, 0xb3, 0x0f, 0xeb, 0x74, 0xb5, 0xfc, 0xe4, 0x59, 0xe8,
	0x82, 0x1d, 0x25, 0xc9, 0x32, 0xf1, 0xba, 0xd3, 0x3f, 0x26, 0x80, 0xea, 0xc9, 0xc5, 0x43, 0xb9,
	0xe1, 0x78, 0x03, 0xfd, 0xbd, 0x7d, 0x78, 0xa5, 0x87, 0x3e, 0x5b, 0x9d, 0xd1, 0xf5, 0x19, 0x6d,
	0x3d, 0x31, 0x30, 0x80, 0x9e, 0x72, 0x49, 0x55, 0x36, 0xdf, 0x75, 0xf0, 0x75, 0xf4, 0xb8, 0xad,
	0x39, 0xf1, 0xd0, 0xc0, 0xe7, 0xe0, 0xbc, 0xe3, 0xf2, 0x5f, 0xd7, 0x0f, 0x2b, 0xc2, 0x0c, 0x7c,
	0xa3, 0x96, 0x57, 0xf0, 0xec, 0xfb, 0xf1, 0xcf, 0xce, 0xaf, 0x1f, 0xcc, 0x69, 0xf2, 0xcc, 0x78,
	0x65, 0xe6, 0x8e, 0xde, 0xff, 0xd7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x35, 0x28, 0x32,
	0x0c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// JobServiceClient is the client API for JobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JobServiceClient interface {
	StartJob(ctx context.Context, in *StartJobRequest, opts ...grpc.CallOption) (*StartJobResponse, error)
	StopJob(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*StopResponse, error)
	GetJob(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*Job, error)
	StreamJobOutput(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (JobService_StreamJobOutputClient, error)
}

type jobServiceClient struct {
	cc *grpc.ClientConn
}

func NewJobServiceClient(cc *grpc.ClientConn) JobServiceClient {
	return &jobServiceClient{cc}
}

func (c *jobServiceClient) StartJob(ctx context.Context, in *StartJobRequest, opts ...grpc.CallOption) (*StartJobResponse, error) {
	out := new(StartJobResponse)
	err := c.cc.Invoke(ctx, "/api.JobService/StartJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) StopJob(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*StopResponse, error) {
	out := new(StopResponse)
	err := c.cc.Invoke(ctx, "/api.JobService/StopJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) GetJob(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := c.cc.Invoke(ctx, "/api.JobService/GetJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) StreamJobOutput(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (JobService_StreamJobOutputClient, error) {
	stream, err := c.cc.NewStream(ctx, &_JobService_serviceDesc.Streams[0], "/api.JobService/StreamJobOutput", opts...)
	if err != nil {
		return nil, err
	}
	x := &jobServiceStreamJobOutputClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type JobService_StreamJobOutputClient interface {
	Recv() (*JobOutput, error)
	grpc.ClientStream
}

type jobServiceStreamJobOutputClient struct {
	grpc.ClientStream
}

func (x *jobServiceStreamJobOutputClient) Recv() (*JobOutput, error) {
	m := new(JobOutput)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// JobServiceServer is the server API for JobService service.
type JobServiceServer interface {
	StartJob(context.Context, *StartJobRequest) (*StartJobResponse, error)
	StopJob(context.Context, *IDRequest) (*StopResponse, error)
	GetJob(context.Context, *IDRequest) (*Job, error)
	StreamJobOutput(*IDRequest, JobService_StreamJobOutputServer) error
}

// UnimplementedJobServiceServer can be embedded to have forward compatible implementations.
type UnimplementedJobServiceServer struct {
}

func (*UnimplementedJobServiceServer) StartJob(ctx context.Context, req *StartJobRequest) (*StartJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartJob not implemented")
}
func (*UnimplementedJobServiceServer) StopJob(ctx context.Context, req *IDRequest) (*StopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopJob not implemented")
}
func (*UnimplementedJobServiceServer) GetJob(ctx context.Context, req *IDRequest) (*Job, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJob not implemented")
}
func (*UnimplementedJobServiceServer) StreamJobOutput(req *IDRequest, srv JobService_StreamJobOutputServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamJobOutput not implemented")
}

func RegisterJobServiceServer(s *grpc.Server, srv JobServiceServer) {
	s.RegisterService(&_JobService_serviceDesc, srv)
}

func _JobService_StartJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).StartJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.JobService/StartJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).StartJob(ctx, req.(*StartJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_StopJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).StopJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.JobService/StopJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).StopJob(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_GetJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).GetJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.JobService/GetJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).GetJob(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_StreamJobOutput_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(IDRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(JobServiceServer).StreamJobOutput(m, &jobServiceStreamJobOutputServer{stream})
}

type JobService_StreamJobOutputServer interface {
	Send(*JobOutput) error
	grpc.ServerStream
}

type jobServiceStreamJobOutputServer struct {
	grpc.ServerStream
}

func (x *jobServiceStreamJobOutputServer) Send(m *JobOutput) error {
	return x.ServerStream.SendMsg(m)
}

var _JobService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.JobService",
	HandlerType: (*JobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartJob",
			Handler:    _JobService_StartJob_Handler,
		},
		{
			MethodName: "StopJob",
			Handler:    _JobService_StopJob_Handler,
		},
		{
			MethodName: "GetJob",
			Handler:    _JobService_GetJob_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamJobOutput",
			Handler:       _JobService_StreamJobOutput_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "job.proto",
}
