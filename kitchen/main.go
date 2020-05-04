package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/njuwelkin/lc/kitchen/pkg/cleaner"
	"github.com/njuwelkin/lc/kitchen/pkg/cook"
	"github.com/njuwelkin/lc/kitchen/pkg/core"
	"github.com/njuwelkin/lc/kitchen/pkg/courier"
	"github.com/njuwelkin/lc/kitchen/pkg/kitchen"
	"github.com/njuwelkin/lc/kitchen/pkg/parser"
)

func usage() {
	fmt.Println("kitchen inputfile [-conf confpath]")
}

func main() {
	configPath := flag.String("conf", "./conf.yaml", "path to the config file")
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
	ctx.IsDebug = true

	// build kitchen
	cookMgr := cook.NewCookMgr(ctx)
	courierMgr := courier.NewCourierMgr(ctx)
	cleaner := cleaner.NewCleaner(ctx)
	kitchen := kitchen.NewKitchen(ctx, cookMgr, courierMgr, cleaner).Run()

	// parse orders
	orders := []core.OrderRequest{}
	err = parser.Parse(orderPath, &orders)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(orders))
	for _, order := range orders {
		kitchen.PlaceOrder(&order)
		time.Sleep(time.Duration(ctx.IngestInterval) * time.Millisecond)
	}

	kitchen.Stop()
	fmt.Println("stopped")
}
