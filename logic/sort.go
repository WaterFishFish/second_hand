package logic

import (
	"second_hand/DAO"
	"sort"
)

func SortByPrice(books *[]DAO.Sell) *[]DAO.Sell {
	sort.Slice(books, func(i, j int) bool {
		return (*books)[i].Price < (*books)[j].Price
	})
	return books
}

func SortBySales(books *[]DAO.Books) *[]DAO.Books {
	sort.Slice(books, func(i, j int) bool {
		return (*books)[i].Count > (*books)[j].Count
	})
	return books
}
