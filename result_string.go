// Code generated by "stringer -type Result result.go"; DO NOT EDIT

package main

import "fmt"

const _Result_name = "InvalidEvenWinLose"

var _Result_index = [...]uint8{0, 7, 11, 14, 18}

func (i Result) String() string {
	if i < 0 || i >= Result(len(_Result_index)-1) {
		return fmt.Sprintf("Result(%d)", i)
	}
	return _Result_name[_Result_index[i]:_Result_index[i+1]]
}
