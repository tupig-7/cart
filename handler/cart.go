package handler

import (
	"context"
	"github.com/tupig-7/cart/common"
	"github.com/tupig-7/cart/domain/model"
	"github.com/tupig-7/cart/domain/service"
	cart "github.com/tupig-7/cart/proto"
)

type Cart struct{
	CartDataService service.ICartDataService
}

//添加购物
func (h *Cart) AddCart(ctx context.Context, request *cart.CartInfo, response *cart.ResponseAdd) (err error) {
	cart := &model.Cart{}
	common.SwapTo(request, cart)
	response.CartId, err = h.CartDataService.AddCart(cart)
	return err
}
func (h *Cart) ClearCart(ctx context.Context, request *cart.Clean, response *cart.Response) error{
	if err := h.CartDataService.CleanCart(request.UserId); err != nil {
		return err
	}
	response.Msg = "购物车清空成功"
	return nil
}
func (h *Cart) Incr(ctx context.Context, request *cart.Item, response *cart.Response) error {
	if err := h.CartDataService.IncrNum(request.Id, request.ChangeNum); err != nil {
		return err
	}
	response.Msg = "添加购物车成功"
	return nil
}
func (h *Cart) Decr(ctx context.Context, request *cart.Item, response *cart.Response) error {
	if err := h.CartDataService.DecrNum(request.Id, request.ChangeNum); err != nil {
		return err
	}
	response.Msg = "购物车减少成功"
	return nil
}
func (h *Cart) DeleteItemByID(ctx context.Context, request *cart.CartID, response *cart.Response) error {
	if err := h.CartDataService.DeleteCart(request.Id); err != nil {
		return err
	}
	response.Msg = "购物车删除成功"
	return nil
}
func (h *Cart) GetAll(ctx context.Context, request *cart.CartFindAll, response *cart.CartAll) error {
	cartAll, err := h.CartDataService.FindAllCart(request.UserId)
	if err != nil {
		return err
	}

	for _, v := range cartAll {
		cart := &cart.CartInfo{}
		if err := common.SwapTo(v, cart); err != nil {
			return err
		}
		response.CartInfo = append(response.CartInfo, cart)
	}
	return nil
}

