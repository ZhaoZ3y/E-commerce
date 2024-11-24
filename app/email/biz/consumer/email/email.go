package email

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/nats-io/nats.go"
	thrifter "github.com/thrift-iterator/go"
	"gomall/app/email/infra/mq"
	"gomall/app/email/infra/notify"
	"gomall/rpc_gen/kitex_gen/email"
)

func ConsumerInit() {
	sub, err := mq.Nc.Subscribe("email", func(msg *nats.Msg) {
		var req email.EmailReq
		err := thrifter.Unmarshal(msg.Data, &req)
		if err != nil {
			klog.Error(err)
			return
		}

		noopEmail := notify.NewNoopEmail()
		_ = noopEmail.Send(&req)
	})
	if err != nil {
		panic(err)
	}

	server.RegisterShutdownHook(func() {
		sub.Unsubscribe()
		mq.Nc.Close()
	})
}
