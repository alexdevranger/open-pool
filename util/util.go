package util

import (
	"math/big"
	"regexp"
	"strconv"
	"time"

	"github.com/alexdevranger/node-1.8.27/common"
	"github.com/alexdevranger/node-1.8.27/common/math"
)

var Ether = math.BigPow(10, 18)
var Shannon = math.BigPow(10, 9)

var pow256 = math.BigPow(2, 256)
var addressPattern = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
var zeroHash = regexp.MustCompile("^0?x?0+$")

func IsValidHexAddress(s string) bool {
	if IsZeroHash(s) || !addressPattern.MatchString(s) {
		return false
	}
	return true
}

func IsZeroHash(s string) bool {
	return zeroHash.MatchString(s)
}

func MakeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func GetTargetHex(diff int64) string {
	difficulty := big.NewInt(diff)
	diff1 := new(big.Int).Div(pow256, difficulty)
	return string(common.ToHex(diff1.Bytes()))
}

func TargetHexToDiff(targetHex string) *big.Int {
	targetBytes := common.FromHex(targetHex)
	return new(big.Int).Div(pow256, new(big.Int).SetBytes(targetBytes))
}

func ToHex(n int64) string {
	return "0x0" + strconv.FormatInt(n, 16)
}

func FormatReward(reward *big.Int) string {
	return reward.String()
}

func FormatRatReward(reward *big.Rat) string {
	wei := new(big.Rat).SetInt(Ether)
	reward = reward.Quo(reward, wei)
	return reward.FloatString(8)
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func MustParseDuration(s string) time.Duration {
	value, err := time.ParseDuration(s)
	if err != nil {
		panic("util: Can't parse duration `" + s + "`: " + err.Error())
	}
	return value
}

func String2Big(num string) *big.Int {
	n := new(big.Int)
	n.SetString(num, 0)
	return n
}

// for NiceHash...
// fixme: rounding error causes invalid shares



// func DiffToTarget(diff float64) (target *big.Int) {
//     mantissa := 0x0000ffff / diff
//     exp := 1
//     tmp := mantissa
//     for tmp >= 256.0 {
//         tmp /= 256.0
//         exp++
//     }
//     for i := 0; i < exp; i++ {
//         mantissa *= 256.0
//     }
//     target = new(big.Int).Lsh(big.NewInt(int64(mantissa)), uint(26-exp)*8)
//     return
// }

// func DiffFloatToDiffInt(diffFloat float64) (diffInt *big.Int) {
//     target := DiffToTarget(diffFloat)
//     return new(big.Int).Div(pow256, target)
// }

