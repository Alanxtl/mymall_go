package service

import (
	"context"

	common "github.com/Alanxtl/mymall_go/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()

	var resp = make(map[string]any)

	resp["items"] = []map[string]any{
		{"Name": "T-shirt-1", "Price": 100, "Picture": "..."},
		{"Name": "T-shirt-2", "Price": 110, "Picture": "..."},
		{"Name": "T-shirt-3", "Price": 120, "Picture": "..."},
		{"Name": "T-shirt-4", "Price": 130, "Picture": "..."},
		{"Name": "T-shirt-5", "Price": 140, "Picture": "..."},
		{"Name": "T-shirt-6", "Price": 150, "Picture": "..."},
	}

	resp["title"] = "Hot Sales"

	return resp, nil
}
