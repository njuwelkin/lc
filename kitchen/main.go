package main

import (
	"flag"
	"fmt"

	"github.com/njuwelkin/lc/kitchen/pkg/clerk"
	"github.com/njuwelkin/lc/kitchen/pkg/cook"
	"github.com/njuwelkin/lc/kitchen/pkg/core"
	"github.com/njuwelkin/lc/kitchen/pkg/courier"
	"github.com/njuwelkin/lc/kitchen/pkg/kitchen"
	"github.com/njuwelkin/lc/kitchen/pkg/parser"
)

func usage() {
	fmt.Println("kitchen inputfile [-conf confpath]")
}

var (
	configPath = flag.String("conf", "./conf.yaml", "path to the config file")
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		usage()
		return
	}
	orderPath := args[0]

	// get context
	ctx, err := core.NewContext(*configPath)
	if err != nil {
		return
	}

	// parse orders
	orders := []core.OrderRequest{}
	err = parser.Parse(orderPath, &orders)
	if err != nil {
		return
	}

	clerk := clerk.NewClerk(ctx)
	cookMgr := cook.NewCookMgr(ctx)
	courierMgr := courier.NewCourierMgr(ctx)
	kitchen := kitchen.NewKitchen(ctx, clerk, cookMgr, courierMgr)

	for _, order := range orders {
		kitchen.PlaceOrder(&order)
	}
}
