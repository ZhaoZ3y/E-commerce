.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demp_proto && cwgo server -I ../../idl --type RPC --module demo_proto --service demo_proto --idl ../../idl/echo.proto

.PHONY: gen-demo-thrift
gen-demo-proto:
	@cd demo/demp_thrift && cwgo server -I ../../idl --type RPC --module demo_thrift --service demo_thrift --idl ../../idl/echo.thrift

.PHONY: gen-frontend
gen-frontend:
	cwgo server --type HTTP --idl ../../idl/frontend/home.thrift --service frontend --module gomall/app/frontend -I ../../idl