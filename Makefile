export MOD = gomall

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/checkout_page.thrift --service frontend --module ${MOD}/app/frontend -I ../../idl/
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
	@cd app/payment && cwgo server --type RPC --service payment --module ${MOD}/app/payment --pass "-use gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/payment.thrift && rmdir /s /q kitex_gen
.PHONY: gen-checkout
gen-checkout:
	@cd rpc_gen && cwgo client --type RPC --service checkout --module ${MOD}/rpc_gen --I ../idl --idl ../idl/checkout.thrift
	@cd app/checkout && cwgo server --type RPC --service checkout --module ${MOD}/app/checkout --pass "-use gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/checkout.thrift && rmdir /s /q kitex_gen
