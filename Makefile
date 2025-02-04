export ROOT_MOD=github.com/Alanxtl/mymall_go

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/home.proto --service frontend --module ${ROOT_MOD}/app/frontend -I ../../idl
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/auth_page.proto --service frontend -module ${ROOT_MOD}/app/frontend -I ../../idl
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/category_page.proto --service frontend -module ${ROOT_MOD}/app/frontend -I ../../idl
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/product_page.proto --service frontend -module ${ROOT_MOD}/app/frontend -I ../../idl
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/cart_page.proto --service frontend -module ${ROOT_MOD}/app/frontend -I ../../idl
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/checkout_page.proto --service frontend -module ${ROOT_MOD}/app/frontend -I ../../idl

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

.PHONY: gen-payment
gen-payment:
	@cd rpc_gen && cwgo client --type RPC --idl ../idl/payment.proto --service ${ROOT_MOD}/payment --module ${ROOT_MOD}/rpc_gen -I ../idl
	@cd app/payment && cwgo server --type RPC --idl ../../idl/payment.proto --service payment --module ${ROOT_MOD}/app/payment -I ../../idl --pass "-use github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen"

.PHONY: gen-checkout
gen-checkout:
	@cd rpc_gen && cwgo client --type RPC --idl ../idl/checkout.proto --service ${ROOT_MOD}/checkout --module ${ROOT_MOD}/rpc_gen -I ../idl
	@cd app/checkout && cwgo server --type RPC --idl ../../idl/checkout.proto --service checkout --module ${ROOT_MOD}/app/checkout -I ../../idl --pass "-use github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen"

.PHONY: gen-order
gen-order:
	@cd rpc_gen && cwgo client --type RPC --idl ../idl/order.proto --service ${ROOT_MOD}/order --module ${ROOT_MOD}/rpc_gen -I ../idl
	@cd app/order && cwgo server --type RPC --idl ../../idl/order.proto --service order --module ${ROOT_MOD}/app/order -I ../../idl --pass "-use github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen"
