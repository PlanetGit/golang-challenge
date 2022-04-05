package main

import (
	"fmt"
	"testing"
)

func Teststruct(x interface{}) {
	// type switch
	switch x.(type) {
	case Repository:
		fmt.Println("Success response")
	default:
		fmt.Println("Invalid repository")
	}
}

// 3. tests for service layer
func TestService(t *testing.T) {
	test_repository := getRepository()

	Teststruct(test_repository)

	if len(test_repository.Projects.Nodes) != 10 {
		t.Errorf("Invalid Nodes")
	}

}
