package notify

import (
	"github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/email"
	"github.com/kr/pretty"
)

type NoopEmail struct{}

func (e *NoopEmail) Send(req *email.EmailReq) error {
	_, err := pretty.Printf("NoopEmail.Send: %+v\n", req)
	if err != nil {
		return err
	}
	return nil
}

func NewNoopEmail() NoopEmail {
	return NoopEmail{}
}
