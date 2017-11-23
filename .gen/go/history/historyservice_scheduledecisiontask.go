// Code generated by thriftrw v1.6.0. DO NOT EDIT.
// @generated

package history

import (
	"errors"
	"fmt"
	"github.com/uber/cadence/.gen/go/shared"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type HistoryService_ScheduleDecisionTask_Args struct {
	ScheduleRequest *ScheduleDecisionTaskRequest `json:"scheduleRequest,omitempty"`
}

func (v *HistoryService_ScheduleDecisionTask_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.ScheduleRequest != nil {
		w, err = v.ScheduleRequest.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _ScheduleDecisionTaskRequest_Read(w wire.Value) (*ScheduleDecisionTaskRequest, error) {
	var v ScheduleDecisionTaskRequest
	err := v.FromWire(w)
	return &v, err
}

func (v *HistoryService_ScheduleDecisionTask_Args) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.ScheduleRequest, err = _ScheduleDecisionTaskRequest_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *HistoryService_ScheduleDecisionTask_Args) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	if v.ScheduleRequest != nil {
		fields[i] = fmt.Sprintf("ScheduleRequest: %v", v.ScheduleRequest)
		i++
	}
	return fmt.Sprintf("HistoryService_ScheduleDecisionTask_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *HistoryService_ScheduleDecisionTask_Args) Equals(rhs *HistoryService_ScheduleDecisionTask_Args) bool {
	if !((v.ScheduleRequest == nil && rhs.ScheduleRequest == nil) || (v.ScheduleRequest != nil && rhs.ScheduleRequest != nil && v.ScheduleRequest.Equals(rhs.ScheduleRequest))) {
		return false
	}
	return true
}

func (v *HistoryService_ScheduleDecisionTask_Args) MethodName() string {
	return "ScheduleDecisionTask"
}

func (v *HistoryService_ScheduleDecisionTask_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var HistoryService_ScheduleDecisionTask_Helper = struct {
	Args           func(scheduleRequest *ScheduleDecisionTaskRequest) *HistoryService_ScheduleDecisionTask_Args
	IsException    func(error) bool
	WrapResponse   func(error) (*HistoryService_ScheduleDecisionTask_Result, error)
	UnwrapResponse func(*HistoryService_ScheduleDecisionTask_Result) error
}{}

func init() {
	HistoryService_ScheduleDecisionTask_Helper.Args = func(scheduleRequest *ScheduleDecisionTaskRequest) *HistoryService_ScheduleDecisionTask_Args {
		return &HistoryService_ScheduleDecisionTask_Args{ScheduleRequest: scheduleRequest}
	}
	HistoryService_ScheduleDecisionTask_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *shared.BadRequestError:
			return true
		case *shared.InternalServiceError:
			return true
		case *shared.EntityNotExistsError:
			return true
		case *ShardOwnershipLostError:
			return true
		default:
			return false
		}
	}
	HistoryService_ScheduleDecisionTask_Helper.WrapResponse = func(err error) (*HistoryService_ScheduleDecisionTask_Result, error) {
		if err == nil {
			return &HistoryService_ScheduleDecisionTask_Result{}, nil
		}
		switch e := err.(type) {
		case *shared.BadRequestError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for HistoryService_ScheduleDecisionTask_Result.BadRequestError")
			}
			return &HistoryService_ScheduleDecisionTask_Result{BadRequestError: e}, nil
		case *shared.InternalServiceError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for HistoryService_ScheduleDecisionTask_Result.InternalServiceError")
			}
			return &HistoryService_ScheduleDecisionTask_Result{InternalServiceError: e}, nil
		case *shared.EntityNotExistsError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for HistoryService_ScheduleDecisionTask_Result.EntityNotExistError")
			}
			return &HistoryService_ScheduleDecisionTask_Result{EntityNotExistError: e}, nil
		case *ShardOwnershipLostError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for HistoryService_ScheduleDecisionTask_Result.ShardOwnershipLostError")
			}
			return &HistoryService_ScheduleDecisionTask_Result{ShardOwnershipLostError: e}, nil
		}
		return nil, err
	}
	HistoryService_ScheduleDecisionTask_Helper.UnwrapResponse = func(result *HistoryService_ScheduleDecisionTask_Result) (err error) {
		if result.BadRequestError != nil {
			err = result.BadRequestError
			return
		}
		if result.InternalServiceError != nil {
			err = result.InternalServiceError
			return
		}
		if result.EntityNotExistError != nil {
			err = result.EntityNotExistError
			return
		}
		if result.ShardOwnershipLostError != nil {
			err = result.ShardOwnershipLostError
			return
		}
		return
	}
}

type HistoryService_ScheduleDecisionTask_Result struct {
	BadRequestError         *shared.BadRequestError      `json:"badRequestError,omitempty"`
	InternalServiceError    *shared.InternalServiceError `json:"internalServiceError,omitempty"`
	EntityNotExistError     *shared.EntityNotExistsError `json:"entityNotExistError,omitempty"`
	ShardOwnershipLostError *ShardOwnershipLostError     `json:"shardOwnershipLostError,omitempty"`
}

func (v *HistoryService_ScheduleDecisionTask_Result) ToWire() (wire.Value, error) {
	var (
		fields [4]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.BadRequestError != nil {
		w, err = v.BadRequestError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.InternalServiceError != nil {
		w, err = v.InternalServiceError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	if v.EntityNotExistError != nil {
		w, err = v.EntityNotExistError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 3, Value: w}
		i++
	}
	if v.ShardOwnershipLostError != nil {
		w, err = v.ShardOwnershipLostError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 4, Value: w}
		i++
	}
	if i > 1 {
		return wire.Value{}, fmt.Errorf("HistoryService_ScheduleDecisionTask_Result should have at most one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *HistoryService_ScheduleDecisionTask_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.BadRequestError, err = _BadRequestError_Read(field.Value)
				if err != nil {
					return err
				}
			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.InternalServiceError, err = _InternalServiceError_Read(field.Value)
				if err != nil {
					return err
				}
			}
		case 3:
			if field.Value.Type() == wire.TStruct {
				v.EntityNotExistError, err = _EntityNotExistsError_Read(field.Value)
				if err != nil {
					return err
				}
			}
		case 4:
			if field.Value.Type() == wire.TStruct {
				v.ShardOwnershipLostError, err = _ShardOwnershipLostError_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.BadRequestError != nil {
		count++
	}
	if v.InternalServiceError != nil {
		count++
	}
	if v.EntityNotExistError != nil {
		count++
	}
	if v.ShardOwnershipLostError != nil {
		count++
	}
	if count > 1 {
		return fmt.Errorf("HistoryService_ScheduleDecisionTask_Result should have at most one field: got %v fields", count)
	}
	return nil
}

func (v *HistoryService_ScheduleDecisionTask_Result) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [4]string
	i := 0
	if v.BadRequestError != nil {
		fields[i] = fmt.Sprintf("BadRequestError: %v", v.BadRequestError)
		i++
	}
	if v.InternalServiceError != nil {
		fields[i] = fmt.Sprintf("InternalServiceError: %v", v.InternalServiceError)
		i++
	}
	if v.EntityNotExistError != nil {
		fields[i] = fmt.Sprintf("EntityNotExistError: %v", v.EntityNotExistError)
		i++
	}
	if v.ShardOwnershipLostError != nil {
		fields[i] = fmt.Sprintf("ShardOwnershipLostError: %v", v.ShardOwnershipLostError)
		i++
	}
	return fmt.Sprintf("HistoryService_ScheduleDecisionTask_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *HistoryService_ScheduleDecisionTask_Result) Equals(rhs *HistoryService_ScheduleDecisionTask_Result) bool {
	if !((v.BadRequestError == nil && rhs.BadRequestError == nil) || (v.BadRequestError != nil && rhs.BadRequestError != nil && v.BadRequestError.Equals(rhs.BadRequestError))) {
		return false
	}
	if !((v.InternalServiceError == nil && rhs.InternalServiceError == nil) || (v.InternalServiceError != nil && rhs.InternalServiceError != nil && v.InternalServiceError.Equals(rhs.InternalServiceError))) {
		return false
	}
	if !((v.EntityNotExistError == nil && rhs.EntityNotExistError == nil) || (v.EntityNotExistError != nil && rhs.EntityNotExistError != nil && v.EntityNotExistError.Equals(rhs.EntityNotExistError))) {
		return false
	}
	if !((v.ShardOwnershipLostError == nil && rhs.ShardOwnershipLostError == nil) || (v.ShardOwnershipLostError != nil && rhs.ShardOwnershipLostError != nil && v.ShardOwnershipLostError.Equals(rhs.ShardOwnershipLostError))) {
		return false
	}
	return true
}

func (v *HistoryService_ScheduleDecisionTask_Result) MethodName() string {
	return "ScheduleDecisionTask"
}

func (v *HistoryService_ScheduleDecisionTask_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
