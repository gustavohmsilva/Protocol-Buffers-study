package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gustavohmsilva/Protocol-Buffers-study/12_-_Protobuf-with_go/src/complexpb"
	"github.com/gustavohmsilva/Protocol-Buffers-study/12_-_Protobuf-with_go/src/enumpb"
	"github.com/gustavohmsilva/Protocol-Buffers-study/12_-_Protobuf-with_go/src/simplepb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	simpleProtoTests()
	enumProtoTest()
	complexProtoTest()
}

func complexProtoTest() {
	cm := &complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "single dummy",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			{
				Id:   1,
				Name: "single multiple dummy",
			}, {
				Id:   2,
				Name: "single multiple dummy II",
			},
		},
	}

	fmt.Println(&cm)
	// printing name of empty fields as well:
	fmt.Printf("%+v\n", cm)
}

func enumProtoTest() {
	em := &enumpb.EnumMessage{
		Id:           1,
		DayOfTheWeek: enumpb.DayOfTheWeek_SATURDAY,
	}

	fmt.Println(em)
	// printing name of empty fields as well:
	fmt.Printf("%+v\n", em)
}

func simpleProtoTests() {
	sm := doSimple()

	// Write into file the generated simplepb.SimpleMessage
	err := writeToFile("simple.bin", sm)
	if err != nil {
		panic(err)
	}

	// read file into new simplepb.SimpleMessage
	// PS: note that simplepbSimpleMessage feed the req. of proto.Message
	newSm := &simplepb.SimpleMessage{}
	err = readFromFile("simple.bin", newSm)
	if err != nil {
		panic(err)
	}
	fmt.Println(newSm)

	// Convert protomessage to JSON
	smParsedAsText, err := toJSON(newSm)
	if err != nil {
		panic(err)
	}
	fmt.Println(smParsedAsText)

	// Revert it back to a simplepb.SimpleMessage
	baseSm := &simplepb.SimpleMessage{}
	newUnmarshaledSm, err := fromJSON(smParsedAsText, baseSm)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success! ", newUnmarshaledSm)
}

func doSimple() *simplepb.SimpleMessage {
	// Lets create a object of type SimpleMessage
	sm := &simplepb.SimpleMessage{
		Id:         1234,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 2, 3, 4},
	}
	fmt.Println(sm)
	// programatically change value
	sm.Name = "I renamed the message"
	fmt.Println(sm)
	// using getter
	fmt.Println(sm.GetId())
	return sm
}

func writeToFile(fname string, pb proto.Message) error {
	fileBytes, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, fileBytes, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}
	fmt.Println("Data has been written!")
	return nil
}

func readFromFile(fname string, pbMessage proto.Message) error {
	fileBytes, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading simple.bin")
		return err
	}
	if err := proto.Unmarshal(fileBytes, pbMessage); err != nil {
		log.Fatalln(
			"Couldn't read the file into protocol buffer struct",
			err,
		)
		return err
	}
	return nil
}

func toJSON(pb proto.Message) (string, error) {
	marshaler := protojson.MarshalOptions{
		Indent:          "\t",
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}
	rawOut, err := marshaler.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return "", err
	}
	return string(rawOut), nil
}

func fromJSON(in string, pb proto.Message) (proto.Message, error) {
	unmarshaler := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
	if err := unmarshaler.Unmarshal([]byte(in), pb); err != nil {
		log.Fatalln("Can't convert to proto.Message", err)
		return nil, err
	}
	return pb, nil
}
