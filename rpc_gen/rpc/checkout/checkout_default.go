package checkout

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	check_out "gomall/rpc_gen/kitex_gen/check_out"
)

func CheccOut(ctx context.Context, req *check_out.CheckOutReq, callOptions ...callopt.Option) (resp *check_out.CheckOutResp, err error) {
	resp, err = defaultClient.CheccOut(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CheccOut call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
