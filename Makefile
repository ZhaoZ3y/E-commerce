.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/auth_page.thrift --service frontend --module gomall/app/frontend -I ../../idl/