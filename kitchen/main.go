package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/njuwelkin/lc/kitchen/pkg/cleaner"
	"github.com/njuwelkin/lc/kitchen/pkg/cook"
	"github.com/njuwelkin/lc/kitchen/pkg/core"
	"github.com/njuwelkin/lc/kitchen/pkg/courier"
	"github.com/njuwelkin/lc/kitchen/pkg/kitchen"
)

func usage() {
	fmt.Println("kitchen inputfile [-conf confpath]")
}

func main() {
	// get args
	configPath := flag.String("conf", "./conf.yaml", "path to the config file")
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		usage()
		return
	}
	orderPath := args[0]

	// get config and global context
	ctx, err := core.NewContext(*configPath)
	if err != nil {
		return
	}

	// build kitchen
	cookMgr := cook.NewCookMgr(ctx)
	courierMgr := courier.NewCourierMgr(ctx)
	cleaner := cleaner.NewCleaner(ctx)
	kitchen := kitchen.NewKitchen(ctx, cookMgr, courierMgr, cleaner).Run()

	// inject orders
	orders := []core.OrderRequest{}
	err = parse(orderPath, &orders)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, order := range orders {
		kitchen.PlaceOrder(&order)
		time.Sleep(time.Duration(ctx.IngestInterval) * time.Millisecond)
	}

	// stop
	kitchen.Stop()
}

func load(path string) ([]byte, error) {
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadFile(absPath)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func parse(path string, val interface{}) error {
	content, err := load(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(content, val)
}
