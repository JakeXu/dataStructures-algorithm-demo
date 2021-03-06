package 抽奖问题

import (
	"testing"
	"time"
	"fmt"
	"math/rand"
)

/**
 * WuMing 
 *2017/9/5 下午3:42
 *
 */

func Test_GetAwardUserName(t *testing.T) {
	var users map[string]int64 = map[string]int64{
		"a": 10,
		"b": 6,
		"c": 3,
		"d": 12,
		"f": 1,
	}

	rand.Seed(time.Now().Unix())
	awardCount := make(map[string]int)
	for i := 0; i <= 1000000; i++ {
		awardName := GetAwardUserName(users)
		if count, ok := awardCount[awardName]; ok {
			awardCount[awardName] = count + 1
		} else {
			awardCount[awardName] = 0
		}
	}
	for n, c := range awardCount {
		fmt.Printf("%v:%v\n",n,c)
	}
}

func Test_getAwardUser_weight(t *testing.T) {
	var users map[string]int64 = map[string]int64{
		"a": 10,
		"b": 6,
		"c": 3,
		"d": 12,
		"f": 1,
	}

	rand.Seed(time.Now().Unix())
	awardCount := make(map[string]int)
	for i := 0; i <= 1000000; i++ {
		awardName := getAwardUser_weight(users)
		if count, ok := awardCount[awardName]; ok {
			awardCount[awardName] = count + 1
		} else {
			awardCount[awardName] = 0
		}
	}
	for n,c := range awardCount {
		fmt.Printf("%v:%v \n",n,c)
	}
}