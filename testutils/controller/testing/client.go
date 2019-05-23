// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import grpc "google.golang.org/grpc"
import mock "github.com/stretchr/testify/mock"
import proto "github.com/opensds/opensds/pkg/model/proto"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Client) Close() {
	_m.Called()
}

// CollectMetrics provides a mock function with given fields: ctx, in, opts
func (_m *Client) CollectMetrics(ctx context.Context, in *proto.CollectMetricsOpts, opts ...grpc.CallOption) (*proto.GenericResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.GenericResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.CollectMetricsOpts, ...grpc.CallOption) *proto.GenericResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GenericResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.CollectMetricsOpts, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Connect provides a mock function with given fields: edp
func (_m *Client) Connect(edp string) error {
	ret := _m.Called(edp)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(edp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateFileShare provides a mock function with given fields: ctx, in, opts
func (_m *Client) CreateFileShare(ctx context.Context, in *proto.CreateFileShareOpts, opts ...grpc.CallOption) (*proto.GenericResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.GenericResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.CreateFileShareOpts, ...grpc.CallOption) *proto.GenericResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GenericResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.CreateFileShareOpts, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateFileShareAcl provides a mock function with given fields: ctx, in, opts
func (_m *Client) CreateFileShareAcl(ctx context.Context, in *proto.CreateFileShareAclOpts, opts ...grpc.CallOption) (*proto.GenericResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.GenericResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.CreateFileShareAclOpts, ...grpc.CallOption) *proto.GenericResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GenericResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.CreateFileShareAclOpts, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateReplication provides a mock function with given fields: ctx, in, opts
func (_m *Client) CreateReplication(ctx context.Context, in *proto.CreateReplicationOpts, opts ...grpc.CallOption) (*proto.GenericResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.GenericResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.CreateReplicationOpts, ...grpc.CallOption) *proto.GenericResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GenericResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.CreateReplicationOpts, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVolume provides a mock function with given fields: ctx, in, opts
func (_m *Client) CreateVolume(ctx context.Context, in *proto.CreateVolumeOpts, opts ...grpc.CallOption) (*proto.GenericResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.GenericResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.CreateVolumeOpts, ...grpc.CallOption) *proto.GenericResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GenericResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.CreateVolumeOpts, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVolumeAttachment provides a mock function with given fields: ctx, in, opts
func (_m *Client) CreateVolumeAttachment(ctx context.Context, in *proto.CreateVolumeAttachmentOpts, opts ...grpc.CallOption) (*proto.GenericResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.GenericResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.CreateVolumeAttachmentOpts, ...grpc.CallOption) *proto.GenericResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GenericResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.CreateVolumeAttachmentOpts, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVolumeGroup provides a mock function with given fields: ctx, in, opts
func (_m *Client) CreateVolumeGroup(ctx context.Context, in *proto.CreateVolumeGroupOpts, opts ...grpc.CallOption) (*proto.GenericResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.GenericResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.CreateVolumeGroupOpts, ...grpc.CallOption) *proto.GenericResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GenericResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.CreateVolumeGroupOpts, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVolumeSnapshot provides a mock function with given fields: ctx, in, opts
func (_m *Client) CreateVolumeSnapshot(ctx context.Context, in *proto.CreateVolumeSnapshotOpts, opts ...grpc.CallOption) (*proto.GenericResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.GenericResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.CreateVolumeSnapshotOpts, ...grpc.CallOption) *proto.GenericResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GenericResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.CreateVolumeSnapshotOpts, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateFileShareSnapshot provides a mock function with given fields: ctx, in, opts
func (_m *Client) CreateFileShareSnapshot(ctx context.Context, in *proto.CreateFileShareSnapshotOpts, opts ...grpc.CallOption) (*proto.GenericResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.GenericResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.CreateFileShareSnapshotOpts, ...grpc.CallOption) *proto.GenericResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GenericResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.CreateFileShareSnapshotOpts, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteFileShare provides a mock function with given fields: ctx, in, opts
func (_m *Client) DeleteFileShare(ctx context.Context, in *proto.DeleteFileShareOpts, opts ...grpc.CallOption) (*proto.GenericResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.GenericResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.DeleteFileShareOpts, ...grpc.CallOption) *proto.GenericResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GenericResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.DeleteFileShareOpts, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteReplication provides a mock function with given fields: ctx, in, opts
func (_m *Client) DeleteReplication(ctx context.Context, in *proto.DeleteReplicationOpts, opts ...grpc.CallOption) (*proto.GenericResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.GenericResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.DeleteReplicationOpts, ...grpc.CallOption) *proto.GenericResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GenericResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.DeleteReplicationOpts, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVolume provides a mock function with given fields: ctx, in, opts
func (_m *Client) DeleteVolume(ctx context.Context, in *proto.DeleteVolumeOpts, opts ...grpc.CallOption) (*proto.GenericResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.GenericResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.DeleteVolumeOpts, ...grpc.CallOption) *proto.GenericResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GenericResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.DeleteVolumeOpts, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVolumeAttachment provides a mock function with given fields: ctx, in, opts
func (_m *Client) DeleteVolumeAttachment(ctx context.Context, in *proto.DeleteVolumeAttachmentOpts, opts ...grpc.CallOption) (*proto.GenericResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.GenericResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.DeleteVolumeAttachmentOpts, ...grpc.CallOption) *proto.GenericResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GenericResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.DeleteVolumeAttachmentOpts, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVolumeGroup provides a mock function with given fields: ctx, in, opts
func (_m *Client) DeleteVolumeGroup(ctx context.Context, in *proto.DeleteVolumeGroupOpts, opts ...grpc.CallOption) (*proto.GenericResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.GenericResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.DeleteVolumeGroupOpts, ...grpc.CallOption) *proto.GenericResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GenericResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.DeleteVolumeGroupOpts, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteVolumeSnapshot provides a mock function with given fields: ctx, in, opts
func (_m *Client) DeleteVolumeSnapshot(ctx context.Context, in *proto.DeleteVolumeSnapshotOpts, opts ...grpc.CallOption) (*proto.GenericResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.GenericResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.DeleteVolumeSnapshotOpts, ...grpc.CallOption) *proto.GenericResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GenericResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.DeleteVolumeSnapshotOpts, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteFileShareSnapshot provides a mock function with given fields: ctx, in, opts
func (_m *Client) DeleteFileShareSnapshot(ctx context.Context, in *proto.DeleteFileShareSnapshotOpts, opts ...grpc.CallOption) (*proto.GenericResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.GenericResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.DeleteFileShareSnapshotOpts, ...grpc.CallOption) *proto.GenericResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GenericResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.DeleteFileShareSnapshotOpts, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisableReplication provides a mock function with given fields: ctx, in, opts
func (_m *Client) DisableReplication(ctx context.Context, in *proto.DisableReplicationOpts, opts ...grpc.CallOption) (*proto.GenericResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.GenericResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.DisableReplicationOpts, ...grpc.CallOption) *proto.GenericResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GenericResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.DisableReplicationOpts, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnableReplication provides a mock function with given fields: ctx, in, opts
func (_m *Client) EnableReplication(ctx context.Context, in *proto.EnableReplicationOpts, opts ...grpc.CallOption) (*proto.GenericResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.GenericResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.EnableReplicationOpts, ...grpc.CallOption) *proto.GenericResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GenericResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.EnableReplicationOpts, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExtendVolume provides a mock function with given fields: ctx, in, opts
func (_m *Client) ExtendVolume(ctx context.Context, in *proto.ExtendVolumeOpts, opts ...grpc.CallOption) (*proto.GenericResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.GenericResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.ExtendVolumeOpts, ...grpc.CallOption) *proto.GenericResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GenericResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.ExtendVolumeOpts, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FailoverReplication provides a mock function with given fields: ctx, in, opts
func (_m *Client) FailoverReplication(ctx context.Context, in *proto.FailoverReplicationOpts, opts ...grpc.CallOption) (*proto.GenericResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.GenericResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.FailoverReplicationOpts, ...grpc.CallOption) *proto.GenericResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GenericResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.FailoverReplicationOpts, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMetrics provides a mock function with given fields: ctx, in, opts
func (_m *Client) GetMetrics(ctx context.Context, in *proto.GetMetricsOpts, opts ...grpc.CallOption) (*proto.GenericResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.GenericResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.GetMetricsOpts, ...grpc.CallOption) *proto.GenericResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GenericResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.GetMetricsOpts, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUrls provides a mock function with given fields: ctx, in, opts
func (_m *Client) GetUrls(ctx context.Context, in *proto.NoParams, opts ...grpc.CallOption) (*proto.GenericResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.GenericResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.NoParams, ...grpc.CallOption) *proto.GenericResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GenericResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.NoParams, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateVolumeGroup provides a mock function with given fields: ctx, in, opts
func (_m *Client) UpdateVolumeGroup(ctx context.Context, in *proto.UpdateVolumeGroupOpts, opts ...grpc.CallOption) (*proto.GenericResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.GenericResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.UpdateVolumeGroupOpts, ...grpc.CallOption) *proto.GenericResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GenericResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.UpdateVolumeGroupOpts, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
