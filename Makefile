export MOD = gomall

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/cart_page.thrift --service frontend --module ${MOD}/app/frontend -I ../../idl/
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
