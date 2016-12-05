package main

import (
	"fmt"
	"container/list"
	"strconv"
)

var co Checkout

func main(){

	stock :=  setupCatelog();
	fmt.Print("In stoock to choose " + strconv.Itoa(stock.Len()) + " items:")
	priceRules := PriceRule{}
	co:= NewCheckout(priceRules)
	co.Scan("VOUCHER")
	co.Scan("VOUCHER")
	co.Scan("TSHIRT")
	total :=co.GetTotal()
	fmt.Printf(total)
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


func NewCheckout(priceRules PriceRule) Checkout{

	checkout := Checkout{
		priceRules: priceRules,
	};
	return checkout;
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
	rules []Rule
}

type Checkout struct {
	itens list.List
	total float64
	priceRules PriceRule
}



func (c Checkout) Scan(item string) string{


	co.itens = setupCatelog();



   return ""
}

func (c Checkout) GetTotal() float64{

	for e := co.itens.Front(); e != nil; e = e.Next() {
		product := e.Value.(*Product) // print out the elements
		fmt.Println(product.name)

	}
	return 30;
}

type Product struct{
	name string
	price int
	code string
}
