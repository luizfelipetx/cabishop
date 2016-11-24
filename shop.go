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
		product := &Product{"Cabify T-Shirt",20,"TSHIRT"}
		alist.PushBack(product)
	}
	for e := 10 ; e <20 ; e++{
		product := &Product{"Cabify Voucher",5,"VOUCHER"}
		alist.PushBack(product)
	}
	for e := 20 ; e <30 ; e++{
		product := &Product{"Cabify Mug",8,"MUG"}
		alist.PushBack(product)
	}
	return alist;
}


func NewCheckout(){



}





type Product struct{
	name string
	price int
	code string
}
