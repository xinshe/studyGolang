package main

import (
	"fmt"
	"io/ioutil"
	"studyGolang/protobuf_demo/pb"

	"github.com/golang/protobuf/proto"

	)

// protobuf demo

func main() {
	var cb pb.ContactBook

	p1 := pb.Person{
		Name:   "小王子",
		Gender: pb.GenderType_MALE,
		Number: "7878778",
	}
	fmt.Println(p1)
	cb.Persons = append(cb.Persons, &p1)
	// 序列化
	data, err := proto.Marshal(&p1)
	if err != nil {
		fmt.Printf("marshal failed,err:%v\n", err)
		return
	}
	ioutil.WriteFile("./proto.dat", data, 0644)

	data2, err := ioutil.ReadFile("./proto.dat")
	if err != nil {
		fmt.Printf("read file failed, err:%v\n", err)
		return
	}
	var p2 pb.Person
	proto.Unmarshal(data2, &p2)
	fmt.Println(p2)
}
