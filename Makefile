

.PHNOY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/home.proto --service frontend -module github.com/Alanxtl/mymall_go/app/frontend -I ../../idl
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/auth_page.proto --service frontend -module github.com/Alanxtl/mymall_go/app/frontend -I ../../idl
