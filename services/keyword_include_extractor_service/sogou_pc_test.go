package keyword_include_extractor_service_test

import (
	"io/ioutil"
	"gitlab.fxt.cn/fxt/rank-util/services/keyword_include_extractor_service"
	"gitlab.fxt.cn/fxt/rank-util/structs/include_extractor"
	"strings"
	"testing"
)

const dataFileSogouPc = "./test_html/sogou_pc.html"

func TestIncludeExtractorSogouPc(t *testing.T) {
	contents, err := ioutil.ReadFile(dataFileSogouPc)
	if err != nil {
		t.Fatal("读取文件错误")
	}

	html := strings.Replace(string(contents), "\n", "", 1)
	req := &include_extractor.IncludeExtractorRequest{
		Body:   html,
	}

	res, err := keyword_include_extractor_service.KeywordIncludeExtractorSogouPc(req)
	if err != nil {
		t.Fatal("解析失败")
	}

	expectIsIncluded := true
	if res.IsIncluded != expectIsIncluded {
		t.Fatalf("期待收录情况%v,实际解出%v", expectIsIncluded, res.IsIncluded)
	}
}
