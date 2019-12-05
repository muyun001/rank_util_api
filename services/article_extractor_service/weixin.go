package article_extractor_service

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"gitlab.fxt.cn/fxt/rank-util/structs/article_extractor"
	"gitlab.fxt.cn/fxt/rank-util/utils"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ArticleExtractorWeixin(req *article_extractor.ArticleExtractorRequest) (*article_extractor.ParseArticleResponse, error) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(req.Body))
	if err != nil {
		return nil, err
	}

	if strings.Contains(req.Body, "您的访问出错了") && strings.Contains(req.Body, "请输入图中的验证码") {
		return nil, errors.New("出现验证码问题")
	}

	articleExtractorResponse := &article_extractor.ParseArticleResponse{}
	if req.RequestType == 1 {
		dom.Find("ul.news-list div.txt-box").Each(func(i int, selection *goquery.Selection) {
			articleUrl := selection.Find("h3 a").AttrOr("data-share", "")
			name := selection.Find("div.s-p").Text()
			name = strings.Replace(strings.Replace(name[0:strings.Index(name, "document")], " ", "", -1), "\n", "", -1)
			intro := selection.Find("p.txt-info").Text()
			headImage := selection.Find("a.account").AttrOr("data-headimage", "")
			guid := selection.Find("a.account").AttrOr("i", "")

			issueTime := ""
			subMatch := regexp.MustCompile(`timeConvert\('(.*?)'\)`).FindStringSubmatch(selection.Find("span.s2").Text())
			if len(subMatch) == 2 {
				timeInt, _ := strconv.Atoi(subMatch[1])
				issueTime = time.Unix(int64(timeInt), 0).Format("2006-01-02")
				if issueTime == "1970-01-01" {
					issueTime = ""
				}
			}

			article := article_extractor.Article5118{
				Guid:                 guid,
				Url:                  articleUrl,
				Name:                 name,
				Intro:                intro,
				HeadImage:            headImage,
				IssueTime:            issueTime,
				ReadCount:            -2,
				AddTime:              time.Now().Format("2006-01-02 15:04:05"),
				ResourcePlatformName: req.ResourcePlatformName,
			}
			articleExtractorResponse.Article = append(articleExtractorResponse.Article, article)
		})

		dom.Find("ul.news-list div.img-box").Each(func(i int, selection *goquery.Selection) {
			coverImage := ""
			img := selection.Find("a img").AttrOr("src", "")
			subMatch := regexp.MustCompile(`&url=(.*)`).FindStringSubmatch(img)
			if len(subMatch) == 2 {
				coverImage = subMatch[1]
			}
			articleExtractorResponse.Article[i].CoverImage = coverImage
		})

		nextPageUrl := dom.Find("#sogou_next").AttrOr("href", "")
		if nextPageUrl != "" {
			articleExtractorResponse.NextPageUrl = "https://weixin.sogou.com/weixin" + nextPageUrl
			articleExtractorResponse.UniqueKey = utils.DcArticleUniqueKey(articleExtractorResponse.NextPageUrl)
		}

	} else {
		dom.Find("div#img-content").Each(func(i int, selection *goquery.Selection) {
			title := selection.Find("h2.rich_media_title").Text()
			title = strings.Replace(strings.Replace(title, " ", "", -1), "\n", "", -1)
			content := selection.Find("div.rich_media_content").Find(" p").Text()

			article := article_extractor.Article5118{
				Title:          title,
				CharsCount:     len(content),
				IsKOL:          0,
				IsOriginal:     1,
				ArticleContent: content,
				CatalogName:    "",
				LikeCount:      0,
				ForwardCount:   0,
				CommentCount:   0,
				EnTitle:        "",
				ChTitle:        "",
			}
			articleExtractorResponse.Article = append(articleExtractorResponse.Article, article)
		})
	}

	articleExtractorResponse.RequestType = req.RequestType

	return articleExtractorResponse, nil
}
