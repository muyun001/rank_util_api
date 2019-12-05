package article_extractor_service_test

import (
	"io/ioutil"
	"gitlab.fxt.cn/fxt/rank-util/services/article_extractor_service"
	"gitlab.fxt.cn/fxt/rank-util/structs/article_extractor"
	"strings"
	"testing"
)

const dataFileToutiao = "./test_html/toutiao.html"

func TestArticleExtractorToutiao(t *testing.T) {
	contents, err := ioutil.ReadFile(dataFileToutiao)

	if err != nil {
		t.Fatal("读取文件错误")
	}
	html := strings.Replace(string(contents), "\n", "", 1)
	req := &article_extractor.ArticleExtractorRequest{
		Url:                  "https://www.toutiao.com/api/search/content/?aid=24&app_name=web_search&offset=0&format=json&keyword=%E7%A1%95%E5%A3%AB&autoload=true&count=20&en_qc=1&cur_tab=1&from=search_tab&pd=synthesis&timestamp=1564537212795",
		Body:                 html,
		ResourcePlatformName: "toutiao",
		RequestType:          2,
	}
	res, err := article_extractor_service.ArticleExtractorToutiao(req)

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
