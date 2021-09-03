package UntisAPI

import (
	"math/rand"
	"testing"
	"time"
)

func TestGoUntisGO(t *testing.T)  {
	for i := 0; i < 100; i++ {
		date := time.Date(0,time.Month(0), 0, rand.Int()%23, rand.Int()%59,0,0,time.Local)
		if date != ToGoTime(ToUnitsTime(date)){
			t.Error("converting date to Untis format and back isn't the same", date, ToGoTime(ToUnitsTime(date)))
		}
	}
}

func TestUntisGOUntis(t *testing.T)  {
	for i := 0; i < 100; i++ {
		date := (rand.Int()%23)*100 + rand.Int()%59
		if date != ToUnitsTime(ToGoTime(date)){
			t.Error("converting date to Untis format and back isn't the same", date, ToUnitsTime(ToGoTime(date)))
		}
	}
}