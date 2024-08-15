package piscine

type food struct {
	goodlack int
}

func FoodDeliveryTime(order string) int {
	more := food{}
	if order == "burger" {
		more.goodlack = 15
	} else if order == "chips" {
		more.goodlack = 10
	} else if order == "nuggets" {
		more.goodlack = 12
	} else {
		return 404
	}
	return more.goodlack
}
