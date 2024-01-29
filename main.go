package main

import (
	"fmt"
	"generate_number/helpers"
)

func main() {
	data := helpers.Participant{Gender: "Ikhwan", Date: "2023-12-21", StartId: 375}

	nip, err := helpers.CreateNIP(data)
	nips, err := helpers.GenerateNIP(data, 10)

	nextNip, err := helpers.CreateNextNIP(nip)
	nextNips, err := helpers.GenerateNextNIP(nip, 10)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Create NIP :", nip)
	fmt.Println("Generate NIP :", nips)

	fmt.Println("Create NextNIP :", nextNip)
	fmt.Println("Generate NextNIP :", nextNips)

}
