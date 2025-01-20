export ROOT_MOD=github.com/Alanxtl/mymall_go

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/home.proto --service frontend --module ${ROOT_MOD}/app/frontend -I ../../idl
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/auth_page.proto --service frontend -module ${ROOT_MOD}/app/frontend -I ../../idl
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/category_page.proto --service frontend -module ${ROOT_MOD}/app/frontend -I ../../idl
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/product_page.proto --service frontend -module ${ROOT_MOD}/app/frontend -I ../../idl

.PHONY: gen-user
gen-user:
	@cd rpc_gen && cwgo client --type RPC --idl ../idl/user.proto --service ${ROOT_MOD}/user --module ${ROOT_MOD}/rpc_gen -I ../idl
	@cd app/user && cwgo server --type RPC --idl ../../idl/user.proto --service user --module ${ROOT_MOD}/app/user -I ../../idl --pass "-use github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen"

.PHONY: gen-product
gen-product:
	@cd rpc_gen && cwgo client --type RPC --idl ../idl/product.proto --service ${ROOT_MOD}/product --module ${ROOT_MOD}/rpc_gen -I ../idl
	@cd app/product && cwgo server --type RPC --idl ../../idl/product.proto --service product --module ${ROOT_MOD}/app/product -I ../../idl --pass "-use github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen"

.PHONY: gen-cart
gen-cart:
	@cd rpc_gen && cwgo client --type RPC --idl ../idl/cart.proto --service ${ROOT_MOD}/cart --module ${ROOT_MOD}/rpc_gen -I ../idl
	@cd app/cart && cwgo server --type RPC --idl ../../idl/cart.proto --service cart --module ${ROOT_MOD}/app/cart -I ../../idl --pass "-use github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen"
