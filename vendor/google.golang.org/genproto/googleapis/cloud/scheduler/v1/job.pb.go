// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/scheduler/v1/job.proto

package scheduler // import "google.golang.org/genproto/googleapis/cloud/scheduler/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import status "google.golang.org/genproto/googleapis/rpc/status"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// State of the job.
type Job_State int32

const (
	// Unspecified state.
	Job_STATE_UNSPECIFIED Job_State = 0
	// The job is executing normally.
	Job_ENABLED Job_State = 1
	// The job is paused by the user. It will not execute. A user can
	// intentionally pause the job using
	// [PauseJobRequest][google.cloud.scheduler.v1.PauseJobRequest].
	Job_PAUSED Job_State = 2
	// The job is disabled by the system due to error. The user
	// cannot directly set a job to be disabled.
	Job_DISABLED Job_State = 3
	// The job state resulting from a failed [CloudScheduler.UpdateJob][google.cloud.scheduler.v1.CloudScheduler.UpdateJob]
	// operation. To recover a job from this state, retry
	// [CloudScheduler.UpdateJob][google.cloud.scheduler.v1.CloudScheduler.UpdateJob] until a successful response is received.
	Job_UPDATE_FAILED Job_State = 4
)

var Job_State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "ENABLED",
	2: "PAUSED",
	3: "DISABLED",
	4: "UPDATE_FAILED",
}
var Job_State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"ENABLED":           1,
	"PAUSED":            2,
	"DISABLED":          3,
	"UPDATE_FAILED":     4,
}

func (x Job_State) String() string {
	return proto.EnumName(Job_State_name, int32(x))
}
func (Job_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_job_2db48d5da5a77145, []int{0, 0}
}

// Configuration for a job.
// The maximum allowed size for a job is 100KB.
type Job struct {
	// Optionally caller-specified in [CreateJob][google.cloud.scheduler.v1.CloudScheduler.CreateJob], after
	// which it becomes output only.
	//
	// The job name. For example:
	// `projects/PROJECT_ID/locations/LOCATION_ID/jobs/JOB_ID`.
	//
	// * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]),
	//    hyphens (-), colons (:), or periods (.).
	//    For more information, see
	//    [Identifying
	//    projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects)
	// * `LOCATION_ID` is the canonical ID for the job's location.
	//    The list of available locations can be obtained by calling
	//    [ListLocations][google.cloud.location.Locations.ListLocations].
	//    For more information, see https://cloud.google.com/about/locations/.
	// * `JOB_ID` can contain only letters ([A-Za-z]), numbers ([0-9]),
	//    hyphens (-), or underscores (_). The maximum length is 500 characters.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optionally caller-specified in [CreateJob][google.cloud.scheduler.v1.CloudScheduler.CreateJob] or
	// [UpdateJob][google.cloud.scheduler.v1.CloudScheduler.UpdateJob].
	//
	// A human-readable description for the job. This string must not contain
	// more than 500 characters.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Required.
	//
	// Delivery settings containing destination and parameters.
	//
	// Types that are valid to be assigned to Target:
	//	*Job_PubsubTarget
	//	*Job_AppEngineHttpTarget
	//	*Job_HttpTarget
	Target isJob_Target `protobuf_oneof:"target"`
	// Required, except when used with [UpdateJob][google.cloud.scheduler.v1.CloudScheduler.UpdateJob].
	//
	// Describes the schedule on which the job will be executed.
	//
	// The schedule can be either of the following types:
	//
	// * [Crontab](http://en.wikipedia.org/wiki/Cron#Overview)
	// * English-like
	// [schedule](https://cloud.google.com/scheduler/docs/configuring/cron-job-schedules)
	//
	// As a general rule, execution `n + 1` of a job will not begin
	// until execution `n` has finished. Cloud Scheduler will never
	// allow two simultaneously outstanding executions. For example,
	// this implies that if the `n+1`th execution is scheduled to run at
	// 16:00 but the `n`th execution takes until 16:15, the `n+1`th
	// execution will not start until `16:15`.
	// A scheduled start time will be delayed if the previous
	// execution has not ended when its scheduled time occurs.
	//
	// If [retry_count][google.cloud.scheduler.v1.RetryConfig.retry_count] > 0 and a job attempt fails,
	// the job will be tried a total of [retry_count][google.cloud.scheduler.v1.RetryConfig.retry_count]
	// times, with exponential backoff, until the next scheduled start
	// time.
	Schedule string `protobuf:"bytes,20,opt,name=schedule,proto3" json:"schedule,omitempty"`
	// Specifies the time zone to be used in interpreting
	// [schedule][google.cloud.scheduler.v1.Job.schedule]. The value of this field must be a time
	// zone name from the [tz database](http://en.wikipedia.org/wiki/Tz_database).
	//
	// Note that some time zones include a provision for
	// daylight savings time. The rules for daylight saving time are
	// determined by the chosen tz. For UTC use the string "utc". If a
	// time zone is not specified, the default will be in UTC (also known
	// as GMT).
	TimeZone string `protobuf:"bytes,21,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty"`
	// Output only. The creation time of the job.
	UserUpdateTime *timestamp.Timestamp `protobuf:"bytes,9,opt,name=user_update_time,json=userUpdateTime,proto3" json:"user_update_time,omitempty"`
	// Output only. State of the job.
	State Job_State `protobuf:"varint,10,opt,name=state,proto3,enum=google.cloud.scheduler.v1.Job_State" json:"state,omitempty"`
	// Output only. The response from the target for the last attempted execution.
	Status *status.Status `protobuf:"bytes,11,opt,name=status,proto3" json:"status,omitempty"`
	// Output only. The next time the job is scheduled. Note that this may be a
	// retry of a previously failed attempt or the next execution time
	// according to the schedule.
	ScheduleTime *timestamp.Timestamp `protobuf:"bytes,17,opt,name=schedule_time,json=scheduleTime,proto3" json:"schedule_time,omitempty"`
	// Output only. The time the last job attempt started.
	LastAttemptTime *timestamp.Timestamp `protobuf:"bytes,18,opt,name=last_attempt_time,json=lastAttemptTime,proto3" json:"last_attempt_time,omitempty"`
	// Settings that determine the retry behavior.
	RetryConfig          *RetryConfig `protobuf:"bytes,19,opt,name=retry_config,json=retryConfig,proto3" json:"retry_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Job) Reset()         { *m = Job{} }
func (m *Job) String() string { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()    {}
func (*Job) Descriptor() ([]byte, []int) {
	return fileDescriptor_job_2db48d5da5a77145, []int{0}
}
func (m *Job) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Job.Unmarshal(m, b)
}
func (m *Job) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Job.Marshal(b, m, deterministic)
}
func (dst *Job) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Job.Merge(dst, src)
}
func (m *Job) XXX_Size() int {
	return xxx_messageInfo_Job.Size(m)
}
func (m *Job) XXX_DiscardUnknown() {
	xxx_messageInfo_Job.DiscardUnknown(m)
}

var xxx_messageInfo_Job proto.InternalMessageInfo

func (m *Job) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Job) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type isJob_Target interface {
	isJob_Target()
}

type Job_PubsubTarget struct {
	PubsubTarget *PubsubTarget `protobuf:"bytes,4,opt,name=pubsub_target,json=pubsubTarget,proto3,oneof"`
}

type Job_AppEngineHttpTarget struct {
	AppEngineHttpTarget *AppEngineHttpTarget `protobuf:"bytes,5,opt,name=app_engine_http_target,json=appEngineHttpTarget,proto3,oneof"`
}

type Job_HttpTarget struct {
	HttpTarget *HttpTarget `protobuf:"bytes,6,opt,name=http_target,json=httpTarget,proto3,oneof"`
}

func (*Job_PubsubTarget) isJob_Target() {}

func (*Job_AppEngineHttpTarget) isJob_Target() {}

func (*Job_HttpTarget) isJob_Target() {}

func (m *Job) GetTarget() isJob_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *Job) GetPubsubTarget() *PubsubTarget {
	if x, ok := m.GetTarget().(*Job_PubsubTarget); ok {
		return x.PubsubTarget
	}
	return nil
}

func (m *Job) GetAppEngineHttpTarget() *AppEngineHttpTarget {
	if x, ok := m.GetTarget().(*Job_AppEngineHttpTarget); ok {
		return x.AppEngineHttpTarget
	}
	return nil
}

func (m *Job) GetHttpTarget() *HttpTarget {
	if x, ok := m.GetTarget().(*Job_HttpTarget); ok {
		return x.HttpTarget
	}
	return nil
}

func (m *Job) GetSchedule() string {
	if m != nil {
		return m.Schedule
	}
	return ""
}

func (m *Job) GetTimeZone() string {
	if m != nil {
		return m.TimeZone
	}
	return ""
}

func (m *Job) GetUserUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UserUpdateTime
	}
	return nil
}

func (m *Job) GetState() Job_State {
	if m != nil {
		return m.State
	}
	return Job_STATE_UNSPECIFIED
}

func (m *Job) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Job) GetScheduleTime() *timestamp.Timestamp {
	if m != nil {
		return m.ScheduleTime
	}
	return nil
}

func (m *Job) GetLastAttemptTime() *timestamp.Timestamp {
	if m != nil {
		return m.LastAttemptTime
	}
	return nil
}

func (m *Job) GetRetryConfig() *RetryConfig {
	if m != nil {
		return m.RetryConfig
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Job) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Job_OneofMarshaler, _Job_OneofUnmarshaler, _Job_OneofSizer, []interface{}{
		(*Job_PubsubTarget)(nil),
		(*Job_AppEngineHttpTarget)(nil),
		(*Job_HttpTarget)(nil),
	}
}

func _Job_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Job)
	// target
	switch x := m.Target.(type) {
	case *Job_PubsubTarget:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PubsubTarget); err != nil {
			return err
		}
	case *Job_AppEngineHttpTarget:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AppEngineHttpTarget); err != nil {
			return err
		}
	case *Job_HttpTarget:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HttpTarget); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Job.Target has unexpected type %T", x)
	}
	return nil
}

func _Job_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Job)
	switch tag {
	case 4: // target.pubsub_target
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PubsubTarget)
		err := b.DecodeMessage(msg)
		m.Target = &Job_PubsubTarget{msg}
		return true, err
	case 5: // target.app_engine_http_target
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AppEngineHttpTarget)
		err := b.DecodeMessage(msg)
		m.Target = &Job_AppEngineHttpTarget{msg}
		return true, err
	case 6: // target.http_target
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HttpTarget)
		err := b.DecodeMessage(msg)
		m.Target = &Job_HttpTarget{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Job_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Job)
	// target
	switch x := m.Target.(type) {
	case *Job_PubsubTarget:
		s := proto.Size(x.PubsubTarget)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Job_AppEngineHttpTarget:
		s := proto.Size(x.AppEngineHttpTarget)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Job_HttpTarget:
		s := proto.Size(x.HttpTarget)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Settings that determine the retry behavior.
//
// By default, if a job does not complete successfully (meaning that
// an acknowledgement is not received from the handler, then it will be retried
// with exponential backoff according to the settings in [RetryConfig][google.cloud.scheduler.v1.RetryConfig].
type RetryConfig struct {
	// The number of attempts that the system will make to run a job using the
	// exponential backoff procedure described by
	// [max_doublings][google.cloud.scheduler.v1.RetryConfig.max_doublings].
	//
	// The default value of retry_count is zero.
	//
	// If retry_count is zero, a job attempt will *not* be retried if
	// it fails. Instead the Cloud Scheduler system will wait for the
	// next scheduled execution time.
	//
	// If retry_count is set to a non-zero number then Cloud Scheduler
	// will retry failed attempts, using exponential backoff,
	// retry_count times, or until the next scheduled execution time,
	// whichever comes first.
	//
	// Values greater than 5 and negative values are not allowed.
	RetryCount int32 `protobuf:"varint,1,opt,name=retry_count,json=retryCount,proto3" json:"retry_count,omitempty"`
	// The time limit for retrying a failed job, measured from time when an
	// execution was first attempted. If specified with
	// [retry_count][google.cloud.scheduler.v1.RetryConfig.retry_count], the job will be retried until both
	// limits are reached.
	//
	// The default value for max_retry_duration is zero, which means retry
	// duration is unlimited.
	MaxRetryDuration *duration.Duration `protobuf:"bytes,2,opt,name=max_retry_duration,json=maxRetryDuration,proto3" json:"max_retry_duration,omitempty"`
	// The minimum amount of time to wait before retrying a job after
	// it fails.
	//
	// The default value of this field is 5 seconds.
	MinBackoffDuration *duration.Duration `protobuf:"bytes,3,opt,name=min_backoff_duration,json=minBackoffDuration,proto3" json:"min_backoff_duration,omitempty"`
	// The maximum amount of time to wait before retrying a job after
	// it fails.
	//
	// The default value of this field is 1 hour.
	MaxBackoffDuration *duration.Duration `protobuf:"bytes,4,opt,name=max_backoff_duration,json=maxBackoffDuration,proto3" json:"max_backoff_duration,omitempty"`
	// The time between retries will double `max_doublings` times.
	//
	// A job's retry interval starts at
	// [min_backoff_duration][google.cloud.scheduler.v1.RetryConfig.min_backoff_duration], then doubles
	// `max_doublings` times, then increases linearly, and finally
	// retries retries at intervals of
	// [max_backoff_duration][google.cloud.scheduler.v1.RetryConfig.max_backoff_duration] up to
	// [retry_count][google.cloud.scheduler.v1.RetryConfig.retry_count] times.
	//
	// For example, if [min_backoff_duration][google.cloud.scheduler.v1.RetryConfig.min_backoff_duration] is
	// 10s, [max_backoff_duration][google.cloud.scheduler.v1.RetryConfig.max_backoff_duration] is 300s, and
	// `max_doublings` is 3, then the a job will first be retried in 10s. The
	// retry interval will double three times, and then increase linearly by
	// 2^3 * 10s.  Finally, the job will retry at intervals of
	// [max_backoff_duration][google.cloud.scheduler.v1.RetryConfig.max_backoff_duration] until the job has
	// been attempted [retry_count][google.cloud.scheduler.v1.RetryConfig.retry_count] times. Thus, the
	// requests will retry at 10s, 20s, 40s, 80s, 160s, 240s, 300s, 300s, ....
	//
	// The default value of this field is 5.
	MaxDoublings         int32    `protobuf:"varint,5,opt,name=max_doublings,json=maxDoublings,proto3" json:"max_doublings,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetryConfig) Reset()         { *m = RetryConfig{} }
func (m *RetryConfig) String() string { return proto.CompactTextString(m) }
func (*RetryConfig) ProtoMessage()    {}
func (*RetryConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_job_2db48d5da5a77145, []int{1}
}
func (m *RetryConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetryConfig.Unmarshal(m, b)
}
func (m *RetryConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetryConfig.Marshal(b, m, deterministic)
}
func (dst *RetryConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetryConfig.Merge(dst, src)
}
func (m *RetryConfig) XXX_Size() int {
	return xxx_messageInfo_RetryConfig.Size(m)
}
func (m *RetryConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RetryConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RetryConfig proto.InternalMessageInfo

func (m *RetryConfig) GetRetryCount() int32 {
	if m != nil {
		return m.RetryCount
	}
	return 0
}

func (m *RetryConfig) GetMaxRetryDuration() *duration.Duration {
	if m != nil {
		return m.MaxRetryDuration
	}
	return nil
}

func (m *RetryConfig) GetMinBackoffDuration() *duration.Duration {
	if m != nil {
		return m.MinBackoffDuration
	}
	return nil
}

func (m *RetryConfig) GetMaxBackoffDuration() *duration.Duration {
	if m != nil {
		return m.MaxBackoffDuration
	}
	return nil
}

func (m *RetryConfig) GetMaxDoublings() int32 {
	if m != nil {
		return m.MaxDoublings
	}
	return 0
}

func init() {
	proto.RegisterType((*Job)(nil), "google.cloud.scheduler.v1.Job")
	proto.RegisterType((*RetryConfig)(nil), "google.cloud.scheduler.v1.RetryConfig")
	proto.RegisterEnum("google.cloud.scheduler.v1.Job_State", Job_State_name, Job_State_value)
}

func init() {
	proto.RegisterFile("google/cloud/scheduler/v1/job.proto", fileDescriptor_job_2db48d5da5a77145)
}

var fileDescriptor_job_2db48d5da5a77145 = []byte{
	// 693 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0x5f, 0x53, 0xda, 0x4a,
	0x18, 0xc6, 0x45, 0x81, 0x83, 0x6f, 0xc0, 0x03, 0xab, 0x9e, 0x13, 0xe9, 0x1f, 0x19, 0x6d, 0xad,
	0xd3, 0x8b, 0x64, 0xb4, 0x77, 0xed, 0x45, 0x07, 0x0c, 0x56, 0x6c, 0xc7, 0x61, 0x02, 0xcc, 0x74,
	0xbc, 0xc9, 0x6c, 0xc2, 0x12, 0xd3, 0x92, 0xdd, 0x9d, 0x64, 0xe3, 0xd0, 0x7e, 0xb0, 0x7e, 0x99,
	0x7e, 0x99, 0x4e, 0x76, 0x13, 0x44, 0xad, 0x70, 0xc7, 0xbe, 0xef, 0xf3, 0xfc, 0xf6, 0xc9, 0xbe,
	0xbb, 0xc0, 0xa1, 0xcf, 0x98, 0x3f, 0x25, 0xa6, 0x37, 0x65, 0xc9, 0xd8, 0x8c, 0xbd, 0x1b, 0x32,
	0x4e, 0xa6, 0x24, 0x32, 0x6f, 0x4f, 0xcc, 0x6f, 0xcc, 0x35, 0x78, 0xc4, 0x04, 0x43, 0x7b, 0x4a,
	0x64, 0x48, 0x91, 0x31, 0x17, 0x19, 0xb7, 0x27, 0xcd, 0xe7, 0x99, 0x1f, 0xf3, 0xc0, 0xc4, 0x94,
	0x32, 0x81, 0x45, 0xc0, 0x68, 0xac, 0x8c, 0xcd, 0xa3, 0xa7, 0xe9, 0x02, 0x47, 0x3e, 0x11, 0x99,
	0xee, 0x65, 0xa6, 0x93, 0x2b, 0x37, 0x99, 0x98, 0xe3, 0x24, 0x92, 0xa0, 0xac, 0xbf, 0xff, 0xb0,
	0x2f, 0x82, 0x90, 0xc4, 0x02, 0x87, 0x3c, 0x13, 0xfc, 0x9f, 0x09, 0x22, 0xee, 0x99, 0xb1, 0xc0,
	0x22, 0xc9, 0x12, 0x1c, 0xfc, 0x2e, 0xc3, 0xc6, 0x25, 0x73, 0x11, 0x82, 0x22, 0xc5, 0x21, 0xd1,
	0x0b, 0xad, 0xc2, 0xf1, 0xa6, 0x2d, 0x7f, 0xa3, 0x16, 0x68, 0x63, 0x12, 0x7b, 0x51, 0xc0, 0xd3,
	0xad, 0xf4, 0x75, 0xd9, 0x5a, 0x2c, 0xa1, 0x2b, 0xa8, 0xf1, 0xc4, 0x8d, 0x13, 0xd7, 0x51, 0x71,
	0xf5, 0x62, 0xab, 0x70, 0xac, 0x9d, 0xbe, 0x31, 0x9e, 0x3c, 0x10, 0xa3, 0x2f, 0xf5, 0x43, 0x29,
	0xbf, 0x58, 0xb3, 0xab, 0x7c, 0x61, 0x8d, 0x08, 0xfc, 0x87, 0x39, 0x77, 0x08, 0xf5, 0x03, 0x4a,
	0x9c, 0x1b, 0x21, 0x78, 0x0e, 0x2e, 0x49, 0xb0, 0xb1, 0x04, 0xdc, 0xe6, 0xbc, 0x2b, 0x7d, 0x17,
	0x42, 0xf0, 0x39, 0x7f, 0x1b, 0x3f, 0x2e, 0xa3, 0x0b, 0xd0, 0x16, 0xd9, 0x65, 0xc9, 0x7e, 0xbd,
	0x84, 0x7d, 0x0f, 0x09, 0x37, 0x77, 0xa4, 0x26, 0x54, 0x72, 0xa1, 0xbe, 0x23, 0xcf, 0x67, 0xbe,
	0x46, 0xcf, 0x60, 0x33, 0x1d, 0x83, 0xf3, 0x93, 0x51, 0xa2, 0xef, 0xaa, 0x66, 0x5a, 0xb8, 0x66,
	0x94, 0x20, 0x0b, 0xea, 0x49, 0x4c, 0x22, 0x27, 0xe1, 0x63, 0x2c, 0x88, 0x93, 0xd6, 0xf5, 0x4d,
	0x99, 0xa3, 0x99, 0xe7, 0xc8, 0x87, 0x69, 0x0c, 0xf3, 0x61, 0xda, 0x5b, 0xa9, 0x67, 0x24, 0x2d,
	0x69, 0x11, 0xbd, 0x87, 0x52, 0x3a, 0x4d, 0xa2, 0x43, 0xab, 0x70, 0xbc, 0x75, 0xfa, 0x6a, 0xc9,
	0x27, 0x5c, 0x32, 0xd7, 0x18, 0xa4, 0x5a, 0x5b, 0x59, 0xd0, 0x5b, 0x28, 0xab, 0x9b, 0xa0, 0x6b,
	0x72, 0x5f, 0x94, 0x9b, 0x23, 0xee, 0x49, 0x65, 0x12, 0xdb, 0x99, 0x02, 0x7d, 0x84, 0x5a, 0x0e,
	0x53, 0x51, 0x1b, 0x2b, 0xa3, 0x56, 0x73, 0x83, 0x0c, 0x7a, 0x0e, 0x8d, 0x29, 0x8e, 0x85, 0x83,
	0x85, 0x20, 0x21, 0x17, 0x0a, 0x82, 0x56, 0x42, 0xfe, 0x4d, 0x4d, 0x6d, 0xe5, 0x91, 0x9c, 0x1e,
	0x54, 0x23, 0x22, 0xa2, 0x1f, 0x8e, 0xc7, 0xe8, 0x24, 0xf0, 0xf5, 0x6d, 0x89, 0x38, 0x5a, 0xf2,
	0xdd, 0x76, 0x2a, 0x3f, 0x93, 0x6a, 0x5b, 0x8b, 0xee, 0x16, 0x07, 0x5f, 0xa1, 0x24, 0xcf, 0x03,
	0xed, 0x42, 0x63, 0x30, 0x6c, 0x0f, 0xbb, 0xce, 0xe8, 0x6a, 0xd0, 0xef, 0x9e, 0xf5, 0xce, 0x7b,
	0x5d, 0xab, 0xbe, 0x86, 0x34, 0xf8, 0xa7, 0x7b, 0xd5, 0xee, 0x7c, 0xe9, 0x5a, 0xf5, 0x02, 0x02,
	0x28, 0xf7, 0xdb, 0xa3, 0x41, 0xd7, 0xaa, 0xaf, 0xa3, 0x2a, 0x54, 0xac, 0xde, 0x40, 0x75, 0x36,
	0x50, 0x03, 0x6a, 0xa3, 0xbe, 0x95, 0xda, 0xcf, 0xdb, 0xbd, 0xb4, 0x54, 0xec, 0x54, 0xa0, 0xac,
	0x6e, 0xd6, 0xc1, 0xaf, 0x75, 0xd0, 0x16, 0x02, 0xa0, 0x7d, 0xd0, 0xf2, 0xf8, 0x09, 0x15, 0xf2,
	0xb1, 0x95, 0x6c, 0xc8, 0x52, 0x25, 0x54, 0xa0, 0x4f, 0x80, 0x42, 0x3c, 0x73, 0x94, 0x28, 0x7f,
	0xe4, 0xf2, 0xe5, 0x69, 0xa7, 0x7b, 0x8f, 0x0e, 0xca, 0xca, 0x04, 0x76, 0x3d, 0xc4, 0x33, 0xb9,
	0x4f, 0x5e, 0x41, 0x9f, 0x61, 0x27, 0x0c, 0xa8, 0xe3, 0x62, 0xef, 0x3b, 0x9b, 0x4c, 0xee, 0x50,
	0x1b, 0xab, 0x50, 0x28, 0x0c, 0x68, 0x47, 0xb9, 0xee, 0xc1, 0xf0, 0xec, 0x31, 0xac, 0xb8, 0x1a,
	0x86, 0x67, 0x0f, 0x61, 0x87, 0x50, 0x4b, 0x61, 0x63, 0x96, 0xb8, 0xd3, 0x80, 0xfa, 0xb1, 0x7c,
	0xda, 0x25, 0xbb, 0x1a, 0xe2, 0x99, 0x95, 0xd7, 0x3a, 0x0c, 0x5e, 0x78, 0x2c, 0x7c, 0x7a, 0xac,
	0x9d, 0xca, 0x25, 0x73, 0xfb, 0xe9, 0x7e, 0xfd, 0xc2, 0x75, 0x27, 0x93, 0xf9, 0x6c, 0x8a, 0xa9,
	0x6f, 0xb0, 0xc8, 0x37, 0x7d, 0x42, 0x65, 0x1a, 0x53, 0xb5, 0x30, 0x0f, 0xe2, 0xbf, 0xfc, 0xc9,
	0x7e, 0x98, 0x2f, 0xdc, 0xb2, 0x94, 0xbf, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x56, 0x25,
	0x9f, 0xf0, 0x05, 0x00, 0x00,
}
