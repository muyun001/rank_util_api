package article_extractor_service_test

import (
	"io/ioutil"
	"gitlab.fxt.cn/fxt/rank-util/services/article_extractor_service"
	"gitlab.fxt.cn/fxt/rank-util/structs/article_extractor"
	"strings"
	"testing"
)

const dataFileWeixin = "./test_html/weixin.html"

func TestArticleExtractorWeixin(t *testing.T) {
	contents, err := ioutil.ReadFile(dataFileWeixin)

	if err != nil {
		t.Fatal("读取文件错误")
	}
	html := strings.Replace(string(contents), "\n", "", 1)
	req := &article_extractor.ArticleExtractorRequest{
		Body:                 html,
		ResourcePlatformName: "weixin",
		RequestType:          2,
	}
	res, err := article_extractor_service.ArticleExtractorWeixin(req)

	if err != nil {
		t.Fatalf("解析出错: %s", err.Error())
	}

	if req.RequestType == 1 {
		if len(res.NextPageUrl) == 0 {
			t.Fatalf("未能检测到下一页的url")
		}
	} else {

	}
}
