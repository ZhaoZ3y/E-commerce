export MOD = gomall

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/order_page.thrift --service frontend --module ${MOD}/app/frontend -I ../../idl/
.PHONY: gen-user
gen-user:
	@cd rpc_gen && cwgo client --type RPC --service user --module ${MOD}/rpc_gen --I ../idl --idl ../idl/user.thrift
	@cd app/user && cwgo server --type RPC --service user --module ${MOD}/app/user --pass "-use gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/user.thrift && rmdir /s /q kitex_gen
.PHONY: gen-product
gen-product:
	@cd rpc_gen && cwgo client --type RPC --service product --module ${MOD}/rpc_gen --I ../idl --idl ../idl/product.thrift
	@cd app/product && cwgo server --type RPC --service product --module ${MOD}/app/product --pass "-use gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/product.thrift && rmdir /s /q kitex_gen
.PHONY: gen-cart
gen-cart:
	@cd rpc_gen && cwgo client --type RPC --service cart --module ${MOD}/rpc_gen --I ../idl --idl ../idl/cart.thrift
	@cd app/cart && cwgo server --type RPC --service cart --module ${MOD}/app/cart --pass "-use gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/cart.thrift && rmdir /s /q kitex_gen
.PHONY: gen-payment
gen-payment:
	@cd rpc_gen && cwgo client --type RPC --service payment --module ${MOD}/rpc_gen --I ../idl --idl ../idl/payment.thrift
	@cd app/payment && cwgo server --type RPC	 --service payment --module ${MOD}/app/payment --pass "-use gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/payment.thrift && rmdir /s /q kitex_gen
.PHONY: gen-checkout
gen-checkout:
	@cd rpc_gen && cwgo client --type RPC --service checkout --module ${MOD}/rpc_gen --I ../idl --idl ../idl/checkout.thrift
	@cd app/checkout && cwgo server --type RPC --service checkout --module ${MOD}/app/checkout --pass "-use gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/checkout.thrift && rmdir /s /q kitex_gen
.PHONY: gen-order
gen-order:
	@cd rpc_gen && cwgo client --type RPC --service order --module ${MOD}/rpc_gen --I ../idl --idl ../idl/order.thrift
	@cd app/order && cwgo server --type RPC --service order --module ${MOD}/app/order --pass "-use gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/order.thrift && rmdir /s /q kitex_gen
.PHONY: gen-email
gen-email:
	@cd rpc_gen && cwgo client --type RPC --service email --module ${MOD}/rpc_gen --I ../idl --idl ../idl/email.thrift
	@cd app/email && cwgo server --type RPC --service email --module ${MOD}/app/email --pass "-use gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/email.thrift && rmdir /s /q kitex_gen
.PHONY: tidy-all
tidy-all:
	@cd app/frontend && go mod tidy
	@cd app/user && go mod tidy
	@cd app/product && go mod tidy
	@cd app/cart && go mod tidy
	@cd app/payment && go mod tidy
	@cd app/checkout && go mod tidy
	@cd app/order && go mod tidy
	@cd app/email && go mod tidy
.PHONY: build-frontend
build-frontend:
	docker build -f ./deploy/Dockerfile.frontend -t frontend:${v} .
.PHONY: build-svc
build-svc:
	docker build -f ./deploy/Dockerfile.svc -t ${svc}:${v} --build-arg SVC=${svc} .
