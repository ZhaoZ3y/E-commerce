// Code generated by thriftgo (0.3.18). DO NOT EDIT.

package home

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"gomall/app/frontend/hertz_gen/frontend/common"
)

type HomeService interface {
	Home(ctx context.Context, e *common.Empty) (r *common.Empty, err error)
}

type HomeServiceClient struct {
	c thrift.TClient
}

func NewHomeServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *HomeServiceClient {
	return &HomeServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewHomeServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *HomeServiceClient {
	return &HomeServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewHomeServiceClient(c thrift.TClient) *HomeServiceClient {
	return &HomeServiceClient{
		c: c,
	}
}

func (p *HomeServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *HomeServiceClient) Home(ctx context.Context, e *common.Empty) (r *common.Empty, err error) {
	var _args HomeServiceHomeArgs
	_args.E = e
	var _result HomeServiceHomeResult
	if err = p.Client_().Call(ctx, "Home", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type HomeServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      HomeService
}

func (p *HomeServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *HomeServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *HomeServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewHomeServiceProcessor(handler HomeService) *HomeServiceProcessor {
	self := &HomeServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("Home", &homeServiceProcessorHome{handler: handler})
	return self
}
func (p *HomeServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type homeServiceProcessorHome struct {
	handler HomeService
}

func (p *homeServiceProcessorHome) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := HomeServiceHomeArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("Home", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := HomeServiceHomeResult{}
	var retval *common.Empty
	if retval, err2 = p.handler.Home(ctx, args.E); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Home: "+err2.Error())
		oprot.WriteMessageBegin("Home", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("Home", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type HomeServiceHomeArgs struct {
	E *common.Empty `thrift:"e,1"`
}

func NewHomeServiceHomeArgs() *HomeServiceHomeArgs {
	return &HomeServiceHomeArgs{}
}

func (p *HomeServiceHomeArgs) InitDefault() {
}

var HomeServiceHomeArgs_E_DEFAULT *common.Empty

func (p *HomeServiceHomeArgs) GetE() (v *common.Empty) {
	if !p.IsSetE() {
		return HomeServiceHomeArgs_E_DEFAULT
	}
	return p.E
}

var fieldIDToName_HomeServiceHomeArgs = map[int16]string{
	1: "e",
}

func (p *HomeServiceHomeArgs) IsSetE() bool {
	return p.E != nil
}

func (p *HomeServiceHomeArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_HomeServiceHomeArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *HomeServiceHomeArgs) ReadField1(iprot thrift.TProtocol) error {
	_field := common.NewEmpty()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.E = _field
	return nil
}

func (p *HomeServiceHomeArgs) Write(oprot thrift.TProtocol) (err error) {

	var fieldId int16
	if err = oprot.WriteStructBegin("Home_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *HomeServiceHomeArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("e", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.E.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *HomeServiceHomeArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("HomeServiceHomeArgs(%+v)", *p)

}

type HomeServiceHomeResult struct {
	Success *common.Empty `thrift:"success,0,optional"`
}

func NewHomeServiceHomeResult() *HomeServiceHomeResult {
	return &HomeServiceHomeResult{}
}

func (p *HomeServiceHomeResult) InitDefault() {
}

var HomeServiceHomeResult_Success_DEFAULT *common.Empty

func (p *HomeServiceHomeResult) GetSuccess() (v *common.Empty) {
	if !p.IsSetSuccess() {
		return HomeServiceHomeResult_Success_DEFAULT
	}
	return p.Success
}

var fieldIDToName_HomeServiceHomeResult = map[int16]string{
	0: "success",
}

func (p *HomeServiceHomeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *HomeServiceHomeResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_HomeServiceHomeResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *HomeServiceHomeResult) ReadField0(iprot thrift.TProtocol) error {
	_field := common.NewEmpty()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Success = _field
	return nil
}

func (p *HomeServiceHomeResult) Write(oprot thrift.TProtocol) (err error) {

	var fieldId int16
	if err = oprot.WriteStructBegin("Home_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *HomeServiceHomeResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *HomeServiceHomeResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("HomeServiceHomeResult(%+v)", *p)

}
