package main

import (
	"fmt"
	"log"
	"os"
	datastore "github.com/ipfs/go-datastore"
	dshelp "github.com/ipfs/go-ipfs-ds-help"
)

func main() {	
	argsWithoutProg := os.Args[1:]
	
	if(0 == len(argsWithoutProg)) {
		fmt.Printf("missing file\n")
		fmt.Printf("example ipfs-block-to-cid CIQENVCICS44LLYUDQ5KVN6ALXC6QRHK2X4R6EUFRMBB5OSFO2FUYDQ\n")
		os.Exit(-1)
	}

	key := fmt.Sprintf("/%s",os.Args[1])

	cid, err := dshelp.DsKeyToCid(datastore.NewKey(key))
	
	if err != nil {
		fmt.Printf("Error\n");
		log.Fatal(err)
	} else {
		fmt.Printf("%s\n", cid);
	}
}