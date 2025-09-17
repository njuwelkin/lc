package main

import (
	"encoding/binary"
	"os"
)

func convert(fname string) {
	fin, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer fin.Close()

	fout, err := os.Create(fname + ".out")
	if err != nil {
		panic(err)
	}
	defer fout.Close()

	for {
		var num byte
		err := binary.Read(fin, binary.LittleEndian, &num)
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}
		err = binary.Write(fout, binary.LittleEndian, ^num)
		if err != nil {
			panic(err)
		}
	}

}

func main() {
	convert("gaussdb_docker.taraa.out")
}
