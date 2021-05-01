package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/gustavohmsilva/Protocol-Buffers-study/13_-_Protobuf_with_go_Exercises_III/tutorial"
	"google.golang.org/protobuf/proto"
)

type data struct {
	file string
	data *tutorial.AddressBook
	sync.RWMutex
}

func main() {
	var addressBook data
	addressBook.file = "address_book.db"

	// Read data from the file (RwMutex just for fun)
	addressBook.Lock()
	addressBook.data = initiateAddressBook("address_book.db")
	addressBook.Unlock()

	callMenu(&addressBook)
}

func callMenu(ab *data) {
	for {
		input := bufio.NewReader(os.Stdin)
		fmt.Print("\nWant to 'write' new person, 'read' new person or 'exit'? ")
		option, err := input.ReadString('\n')
		if err != nil {
			log.Fatalln("Unexpected error")
		}

		switch strings.TrimSpace(option) {
		case "read":
			fmt.Printf("%+v", ab.data)
		case "write":
			fmt.Printf("option write not implemented\n")
			addPerson(ab)
		default:
			fmt.Printf("saving and exiting.\n")
			writeAddressBook(ab)
			os.Exit(0)
		}
	}
}

func initiateAddressBook(filename string) *tutorial.AddressBook {
	addressBookFile, err := ioutil.ReadFile(filename)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatalln("Error reading file", err)

		}
		fmt.Printf("Address Book not found... Creating...\n")
	}

	ab := &tutorial.AddressBook{}
	if err := proto.Unmarshal(addressBookFile, ab); err != nil {
		log.Fatalln("Failed to parse address book")
	}
	return ab
}

func writeAddressBook(ab *data) {
	ab.RLock()
	wireData, err := proto.Marshal(ab.data)
	if err != nil {
		log.Fatalln("Failed to encode address book")
	}
	ab.RUnlock()

	if err := ioutil.WriteFile(ab.file, wireData, 0644); err != nil {
		log.Fatalln("Failed to update address book")
	}
}

func addPerson(ab *data) {
	p := &tutorial.Person{}
	rd := bufio.NewReader(os.Stdin)
	fmt.Print("Enter person ID number: ")
	if _, err := fmt.Fscanf(rd, "%d\n", &p.Id); err != nil {
		log.Fatalln("Error reading ID")
	}

	fmt.Print("Enter name: ")
	name, err := rd.ReadString('\n')
	if err != nil {
		log.Fatalln("Error reading name")
	}
	p.Name = strings.TrimSpace(name)

	fmt.Print("Enter email address (blank for none): ")
	email, err := rd.ReadString('\n')
	if err != nil {
		log.Fatalln("Error reading email address")
	}
	p.Email = strings.TrimSpace(email)

	for {
		fmt.Print("Enter a phone number  (or E to finish): ")
		phone, err := rd.ReadString('\n')
		if err != nil || phone == "" {
			log.Fatalln("Error reading phone number")
		}
		phone = strings.TrimSpace(phone)
		if phone == "E" {
			break
		}
		pn := &tutorial.Person_PhoneNumber{Number: phone}

		fmt.Print("Is this a mobile, home, or work phone? ")
		ptype, err := rd.ReadString('\n')
		if err != nil {
			log.Fatalln("Error reading type of phone")
		}
		ptype = strings.TrimSpace(ptype)

		if tutorial.Person_PhoneType_value[strings.ToUpper(ptype)] != 0 {
			pn.Type = tutorial.Person_PhoneType(
				tutorial.Person_PhoneType_value[strings.ToUpper(ptype)],
			)
		}
		p.Phones = append(p.Phones, pn)
	}
	ab.Lock()
	ab.data.People = append(ab.data.People, p)
	ab.Unlock()
}
