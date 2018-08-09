// Please, run go generate on root of this repository (github.com/devimteam/proto-utils)
package _go

import (
	"github.com/devimteam/proto-utils/src/main/go/page"
)

type Order struct {
	Column    string
	Direction uint32
}

type Sort struct {
	Orders []*Order
}

type Page struct {
	TotalPages uint32
	TotalSize  uint64
	Page       uint32
	Size       uint64
}

type PageRequest struct {
	Page uint32
	Size uint64
	Sort *Sort
}

func NewSort(orders ...*Order) *Sort {
	return &Sort{Orders: orders}
}

func NewOrder(column string, direction page.Order_Direction) *Order {
	return &Order{
		Column:    column,
		Direction: uint32(direction),
	}
}

func OrderToProto(o *Order) *page.Order {
	if o == nil {
		return nil
	}
	return &page.Order{
		Column:    o.Column,
		Direction: page.Order_Direction(o.Direction),
	}
}

func ProtoToOrder(o *page.Order) *Order {
	if o == nil {
		return nil
	}
	return &Order{
		Column:    o.Column,
		Direction: uint32(o.Direction),
	}
}

func SortToProto(s *Sort) *page.Sort {
	if s == nil {
		return nil
	}
	var orders []*page.Order
	for _, o := range s.Orders {
		orders = append(orders, OrderToProto(o))
	}
	return &page.Sort{
		Orders: orders,
	}
}

func ProtoToSort(s *page.Sort) *Sort {
	if s == nil {
		return nil
	}
	var orders []*Order
	for _, o := range s.Orders {
		orders = append(orders, ProtoToOrder(o))
	}
	return &Sort{
		Orders: orders,
	}
}

func PageToProto(p *Page) *page.Page {
	if p == nil {
		return nil
	}
	return &page.Page{
		TotalPages: p.TotalPages,
		TotalSize:  p.TotalSize,
		Page:       p.Page,
		Size:       p.Size,
	}
}

func ProtoToPage(p *page.Page) *Page {
	if p == nil {
		return nil
	}
	return &Page{
		TotalPages: p.TotalPages,
		TotalSize:  p.TotalSize,
		Page:       p.Page,
		Size:       p.Size,
	}
}

func PageRequestToProto(pr *PageRequest) *page.PageRequest {
	if pr == nil {
		return nil
	}
	return &page.PageRequest{
		Page: pr.Page,
		Size: pr.Size,
		Sort: SortToProto(pr.Sort),
	}
}

func ProtoToPageRequest(pr *page.PageRequest) *PageRequest {
	if pr == nil {
		return nil
	}
	return &PageRequest{
		Page: pr.Page,
		Size: pr.Size,
		Sort: ProtoToSort(pr.Sort),
	}
}
