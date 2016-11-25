package main

import (
	"fmt"
	"container/list"
	"strconv"
)

func main(){

	stock :=  setupCatelog();
	fmt.Print("In stoock to choose " + strconv.Itoa(stock.Len()) + " items:")
}


func setupCatelog() list.List{

	alist := list.List{}

	for e := 0 ; e <10 ; e++{
		product := &Product{"T-Shirt",20,"TSHIRT"}
		alist.PushBack(product)
	}
	for e := 10 ; e <20 ; e++{
		product := &Product{"Voucher",5,"VOUCHER"}
		alist.PushBack(product)
	}
	for e := 20 ; e <30 ; e++{
		product := &Product{"Mug",8,"MUG"}
		alist.PushBack(product)
	}
	return alist;
}


func NewCheckout(priceRules PriceRule) list.List{


}

type Rule struct {
	format string
	isDirectDiscount bool //if not use direct discount in $$ will be used percentage value.
	percentage float64
	isBundle bool
	bundleEqual int
	bundleNotEqual int
	numberProductsToValidate int
}

type PriceRule struct{

}

type Checkout struct {
	itens list.List
	total float64
}


type Product struct{
	name string
	price int
	code string
}
