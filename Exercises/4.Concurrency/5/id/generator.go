package id

import (
	"fmt"
	"math/rand"
	"strconv"
)

func GenerateID() string {
	return fmt.Sprintf("user:%s", strconv.Itoa(rand.Intn(2023)))
}
