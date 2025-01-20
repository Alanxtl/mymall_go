package service

import (
	"context"
	"github.com/Alanxtl/mymall_go/app/product/biz/dal/mysql"
	"github.com/Alanxtl/mymall_go/app/product/biz/model"
	product "github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)

	c, err := categoryQuery.GetProductsByCategoryName(req.CategoryName)
	if err != nil {
		return nil, err
	}

	resp = &product.ListProductsResp{}

	for _, vl := range c {
		for _, v := range vl.Products {
			resp.Products = append(resp.Products, &product.Product{
				Id:          uint32(v.ID),
				Picture:     v.Picture,
				Price:       v.Price,
				Description: v.Description,
				Name:        v.Name,
			})
		}
	}

	return resp, nil
}
