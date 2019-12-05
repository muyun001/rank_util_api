package domain_include_extractor_service_test

import (
	"io/ioutil"
	"gitlab.fxt.cn/fxt/rank-util/services/domain_include_extractor_service"
	"gitlab.fxt.cn/fxt/rank-util/structs/include_extractor"
	"strings"
	"testing"
)

const dataFileSmMobile = "./test_html/sm_mobile.html"

func TestIncludeExtractorSmMobile(t *testing.T) {
	contents, err := ioutil.ReadFile(dataFileSmMobile)
	if err != nil {
		t.Fatal("读取文件错误")
	}

	html := strings.Replace(string(contents), "\n", "", 1)
	req := &include_extractor.IncludeExtractorRequest{
		Body:   html,
	}

	res, err := domain_include_extractor_service.DomainIncludeExtractorSmMobile(req)
	if err != nil {
		t.Fatal("解析失败")
	}

	expectIncludeNum := 7
	if res.IncludeNum != expectIncludeNum {
		t.Fatalf("期待收录数%d,实际解出%d", expectIncludeNum, res.IncludeNum)
	}
}
