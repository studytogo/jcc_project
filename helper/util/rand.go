package util

import (
	"fmt"
	"math/rand"
	"time"
)

func CreatePurchaseOrder() string {
	middle := time.Now().Format("20060102-150405.000")
	content1 := middle[0:15]
	content2 := middle[16:19]
	middle = content1 + content2
	rnd := CreateCaptcha()
	newRDh := "RK-" + middle + "-" + rnd
	return newRDh
}

func CreatePurchaseReturnOrder() string {
	middle := time.Now().Format("20060102-150405.000")
	content1 := middle[0:15]
	content2 := middle[16:19]
	middle = content1 + content2
	rnd := CreateCaptcha()
	newRDh := "Ck-" + middle + "-" + rnd
	return newRDh
}

func CreateGiveOrder() string {
	middle := time.Now().Format("20060102-150405.000")
	content1 := middle[0:15]
	content2 := middle[16:19]
	middle = content1 + content2
	rnd := CreateCaptcha()
	newRDh := "HZ-" + middle + "-" + rnd
	return newRDh
}

func CreateCaptcha() string {
	return fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
}