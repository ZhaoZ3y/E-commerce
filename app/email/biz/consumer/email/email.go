package email

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/nats-io/nats.go"
	thrifter "github.com/thrift-iterator/go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"gomall/app/email/infra/mq"
	"gomall/app/email/infra/notify"
	"gomall/rpc_gen/kitex_gen/email"
)

func ConsumerInit() {
	tracer := otel.Tracer("shop-nats-consumer")
	sub, err := mq.Nc.Subscribe("email", func(msg *nats.Msg) {
		var req email.EmailReq
		err := thrifter.Unmarshal(msg.Data, &req)
		if err != nil {
			klog.Error(err)
			return
		}
		ctx := context.Background()
		ctx = otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(msg.Header))
		_, span := tracer.Start(ctx, "shop-email-consumer")

		defer span.End()
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
