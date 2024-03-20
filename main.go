package main

import (
	"flag"
	"fmt"
    "math/rand"
    "time"
	"encoding/base64"

	"github.com/sinohope/shamir-pasasword/shamir"
)


var (
	passwordT = flag.Int("t", 2, "File path to asset.db")
	passwordN = flag.Int("n", 3, "File path to asset.db")
)


const charset = "!@#$%^&*()_+abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"


func randomString(length int) string {
	seededRand  := rand.New(rand.NewSource(time.Now().UnixNano()))
    b := make([]byte, length)
    for i := range b {
        b[i] = charset[seededRand.Intn(len(charset))]
    }
    return string(b)
}

func main() {
	flag.Parse()
	password := randomString(16)
	var BKey = []byte(password)

	byteShares, err := shamir.Split(BKey, *passwordN, *passwordT)
	if err != nil {
		fmt.Printf("5 failed: %v", err)
		panic(err)
	}
	for _, byteShare := range byteShares {
		fmt.Println(base64.StdEncoding.EncodeToString(byteShare))
	}
}

