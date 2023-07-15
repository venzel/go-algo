package src

import (
	"fmt"

	"github.com/google/uuid"
)

func GenerateUUID() {
	id := uuid.New().String()
	fmt.Println(id)
}
