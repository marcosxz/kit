package kit

import (
	"strconv"
)

// 分转元
func BranchToYuan(branch int) string {
	var fee string
	if branch <= 0 {
		fee = "0"
	} else {
		fee = strconv.Itoa(branch)
		switch l := len(fee); l {
		case 0:
			fee = "0.00"
		case 1:
			fee = "0.0" + fee[:]
		case 2:
			fee = "0." + fee[:]
		default:
			fee = fee[:l-2] + "." + fee[l-2:]
		}
	}
	return fee
}

// 元转分
func YuanToBranch(yuan string) (int64, error) {
	if yuanFloat64, err := strconv.ParseFloat(yuan, 64); err != nil {
		return 0, err
	} else {
		return int64(yuanFloat64 * 100), nil
	}
}
