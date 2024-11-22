// Code generated by thriftgo (0.3.18). DO NOT EDIT.

package order_page

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"gomall/app/frontend/hertz_gen/frontend/common"
)

type OrderPageService interface {
	OrderList(ctx context.Context, req *common.Empty) (r *common.Empty, err error)
}

type OrderPageServiceClient struct {
	c thrift.TClient
}

func NewOrderPageServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *OrderPageServiceClient {
	return &OrderPageServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewOrderPageServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *OrderPageServiceClient {
	return &OrderPageServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewOrderPageServiceClient(c thrift.TClient) *OrderPageServiceClient {
	return &OrderPageServiceClient{
		c: c,
	}
}

func (p *OrderPageServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *OrderPageServiceClient) OrderList(ctx context.Context, req *common.Empty) (r *common.Empty, err error) {
	var _args OrderPageServiceOrderListArgs
	_args.Req = req
	var _result OrderPageServiceOrderListResult
	if err = p.Client_().Call(ctx, "OrderList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type OrderPageServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      OrderPageService
}

func (p *OrderPageServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *OrderPageServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *OrderPageServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewOrderPageServiceProcessor(handler OrderPageService) *OrderPageServiceProcessor {
	self := &OrderPageServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("OrderList", &orderPageServiceProcessorOrderList{handler: handler})
	return self
}
func (p *OrderPageServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
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

type orderPageServiceProcessorOrderList struct {
	handler OrderPageService
}

func (p *orderPageServiceProcessorOrderList) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := OrderPageServiceOrderListArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("OrderList", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := OrderPageServiceOrderListResult{}
	var retval *common.Empty
	if retval, err2 = p.handler.OrderList(ctx, args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing OrderList: "+err2.Error())
		oprot.WriteMessageBegin("OrderList", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("OrderList", thrift.REPLY, seqId); err2 != nil {
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

type OrderPageServiceOrderListArgs struct {
	Req *common.Empty `thrift:"req,1"`
}

func NewOrderPageServiceOrderListArgs() *OrderPageServiceOrderListArgs {
	return &OrderPageServiceOrderListArgs{}
}

func (p *OrderPageServiceOrderListArgs) InitDefault() {
}

var OrderPageServiceOrderListArgs_Req_DEFAULT *common.Empty

func (p *OrderPageServiceOrderListArgs) GetReq() (v *common.Empty) {
	if !p.IsSetReq() {
		return OrderPageServiceOrderListArgs_Req_DEFAULT
	}
	return p.Req
}

var fieldIDToName_OrderPageServiceOrderListArgs = map[int16]string{
	1: "req",
}

func (p *OrderPageServiceOrderListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *OrderPageServiceOrderListArgs) Read(iprot thrift.TProtocol) (err error) {

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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_OrderPageServiceOrderListArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *OrderPageServiceOrderListArgs) ReadField1(iprot thrift.TProtocol) error {
	_field := common.NewEmpty()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Req = _field
	return nil
}

func (p *OrderPageServiceOrderListArgs) Write(oprot thrift.TProtocol) (err error) {

	var fieldId int16
	if err = oprot.WriteStructBegin("OrderList_args"); err != nil {
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

func (p *OrderPageServiceOrderListArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Req.Write(oprot); err != nil {
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

func (p *OrderPageServiceOrderListArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("OrderPageServiceOrderListArgs(%+v)", *p)

}

type OrderPageServiceOrderListResult struct {
	Success *common.Empty `thrift:"success,0,optional"`
}

func NewOrderPageServiceOrderListResult() *OrderPageServiceOrderListResult {
	return &OrderPageServiceOrderListResult{}
}

func (p *OrderPageServiceOrderListResult) InitDefault() {
}

var OrderPageServiceOrderListResult_Success_DEFAULT *common.Empty

func (p *OrderPageServiceOrderListResult) GetSuccess() (v *common.Empty) {
	if !p.IsSetSuccess() {
		return OrderPageServiceOrderListResult_Success_DEFAULT
	}
	return p.Success
}

var fieldIDToName_OrderPageServiceOrderListResult = map[int16]string{
	0: "success",
}

func (p *OrderPageServiceOrderListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *OrderPageServiceOrderListResult) Read(iprot thrift.TProtocol) (err error) {

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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_OrderPageServiceOrderListResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *OrderPageServiceOrderListResult) ReadField0(iprot thrift.TProtocol) error {
	_field := common.NewEmpty()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Success = _field
	return nil
}

func (p *OrderPageServiceOrderListResult) Write(oprot thrift.TProtocol) (err error) {

	var fieldId int16
	if err = oprot.WriteStructBegin("OrderList_result"); err != nil {
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

func (p *OrderPageServiceOrderListResult) writeField0(oprot thrift.TProtocol) (err error) {
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

func (p *OrderPageServiceOrderListResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("OrderPageServiceOrderListResult(%+v)", *p)

}