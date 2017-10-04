// Please, run go generate on root of this repository (github.com/devimteam/proto-utils)
package _go

import "github.com/devimteam/proto-utils/src/main/go/relay"

type Cursor struct {
	Before *string
	After  *string
	First  *uint32
	Last   *uint32
}

type RelayPage struct {
	HasNextPage     bool
	HasPreviousPage bool
	Start           *string
	End             *string
}

type RelayPageRequest struct {
	Cursor *Cursor
	Sort   *Sort
}

func CursorToProto(cursor *Cursor) *relay.Cursor {
	return &relay.Cursor{
		Before: NilStringToProto(cursor.Before),
		After:  NilStringToProto(cursor.After),
		First:  NilUInt32ToProto(cursor.First),
		Last:   NilUInt32ToProto(cursor.Last),
	}
}

func ProtoToCursor(cursor *relay.Cursor) *Cursor {
	if cursor == nil {
		return nil
	}
	return &Cursor{
		Before: ProtoToNilString(cursor.Before),
		After:  ProtoToNilString(cursor.After),
		First:  ProtoToNilUInt32(cursor.First),
		Last:   ProtoToNilUInt32(cursor.Last),
	}
}

func RelayPageToProto(page *RelayPage) *relay.Page {
	return &relay.Page{
		Start:           NilStringToProto(page.Start),
		End:             NilStringToProto(page.End),
		HasNextPage:     page.HasNextPage,
		HasPreviousPage: page.HasPreviousPage,
	}
}

func ProtoToRelayPage(protoPage *relay.Page) (page *RelayPage) {
	if protoPage == nil {
		return nil
	}
	return &RelayPage{
		Start:           ProtoToNilString(protoPage.Start),
		End:             ProtoToNilString(protoPage.End),
		HasPreviousPage: protoPage.HasPreviousPage,
		HasNextPage:     protoPage.HasNextPage,
	}
}

func RelayPageRequestPtrToProto(pageRequest *RelayPageRequest) *relay.PageRequest {
	if pageRequest == nil {
		return nil
	}
	return &relay.PageRequest{
		Sort:   SortToProto(pageRequest.Sort),
		Cursor: CursorToProto(pageRequest.Cursor),
	}
}

func ProtoToRelayPageRequestPtr(pr *relay.PageRequest) *RelayPageRequest {
	if pr == nil {
		return nil
	}
	return &RelayPageRequest{
		Sort:   ProtoToSort(pr.Sort),
		Cursor: ProtoToCursor(pr.Cursor),
	}
}
