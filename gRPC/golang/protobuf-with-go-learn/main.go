package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"protobuf-go/src/complexpb"
	"protobuf-go/src/enumpb"
	"protobuf-go/src/first"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	// pm := NewPersonMessage()
	// // _ = writeToFile("person.bin", pm)
	// // pm2 := &first.PersonMessage{}
	// // _ = readFromFile("person.bin", pm2)
	// // fmt.Println(pm2)
	// pmStr := toJSON(pm)
	// fmt.Println(pmStr)

	// pm3 := first.PersonMessage{}

	// _ = fromJSON(pmStr, &pm3)
	// fmt.Println(pm3)
	// em := NewEnumMessage()
	// fmt.Println(enumpb.Gender_name[int32(em.Gender)])
	dm := NewDepartmentMessage()
	fmt.Println(dm)
}

func NewDepartmentMessage() *complexpb.DepartmentMessage {
	dm := &complexpb.DepartmentMessage{
		Id:   5678,
		Name: "开发部",
		Employees: []*complexpb.EmployeeMessage{
			&complexpb.EmployeeMessage{Id: 11, Name: "Dave"},
			&complexpb.EmployeeMessage{Id: 22, Name: "Mike"},
		},
		ParentDepartment: &complexpb.DepartmentMessage{
			Id:   1122,
			Name: "总公司",
		},
	}
	return dm
}

func NewEnumMessage() *enumpb.EnumMessage {
	em := enumpb.EnumMessage{
		Id:     345,
		Gender: enumpb.Gender_FEMALE,
	}
	em.Gender = enumpb.Gender_MALE
	return &em
}

func fromJSON(in string, pb proto.Message) error {
	err := protojson.Unmarshal([]byte(in), pb)
	if err != nil {
		log.Fatalln("读取json时发生错误", err.Error())
	}
	return nil
}

func toJSON(pb proto.Message) string {
	marshaler := protojson.MarshalOptions{
		Indent: "    ",
	}
	str, err := marshaler.Marshal(pb)
	// marshaler := jsonpb.Marshaler{}

	// str, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("转化为JSON时发生错误", err.Error())
	}

	return string(str)
}

func readFromFile(fileName string, pb proto.Message) error {
	dataBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln("读取文件时发送错误", err.Error())
		return err
	}

	err = proto.Unmarshal(dataBytes, pb)
	if err != nil {
		log.Fatalln("转化为struct时发送错误", err.Error())
	}
	return nil
}

func writeToFile(filename string, pb proto.Message) error {
	dataBytes, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("无法序列化到bytes")
		return err
	}

	if err := ioutil.WriteFile(filename, dataBytes, 0644); err != nil {
		log.Fatalln("无法写入到文件", err.Error())
		return err
	}

	log.Println("成功写入到文件")
	return nil

}

func NewPersonMessage() *first.PersonMessage {
	pm := first.PersonMessage{
		Id:         123,
		IsAdult:    true,
		Name:       "Dave",
		LuckNumber: []int32{1, 2, 3, 4, 5},
	}
	pm.Name = "Nick"
	return &pm
}
