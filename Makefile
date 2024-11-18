export MOD = gomall

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/category_page.thrift --service frontend --module ${MOD}/app/frontend -I ../../idl/
.PHONY: gen-user
gen-user:
	@cd rpc_gen && cwgo client --type RPC --service user --module ${MOD}/rpc_gen --I ../idl --idl ../idl/user.thrift
	@cd app/user && cwgo server --type RPC --service user --module ${MOD}/app/user -I ../../idl --idl ../../idl/user.thrift && rmdir /s /q kitex_gen
.PHONY: gen-product
gen-product:
	@cd rpc_gen && cwgo client --type RPC --service product --module ${MOD}/rpc_gen --I ../idl --idl ../idl/product.thrift
	@cd app/product && cwgo server --type RPC --service product --module ${MOD}/app/product -I ../../idl --idl ../../idl/product.thrift && rmdir /s /q kitex_gen
