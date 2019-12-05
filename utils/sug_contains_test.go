package utils_test

import (
	"gitlab.fxt.cn/fxt/rank-util/utils"
	"testing"
)

func TestIsSugContains(t *testing.T) {
	checkWord := "选新东方"
	sugStr := "z选A新东方!"
	if utils.IsSugContains(checkWord, sugStr) == false {
		t.Error("expect true, got false")
	}
	checkWord2 := "选新东方a"
	sugStr2 := "z选A新东方!"
	if utils.IsSugContains(checkWord2, sugStr2) == true {
		t.Error("expect false, got true")
	}
}
