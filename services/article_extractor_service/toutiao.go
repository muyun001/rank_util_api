package article_extractor_service

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
	"html"
	"gitlab.fxt.cn/fxt/rank-util/structs/article_extractor"
	"gitlab.fxt.cn/fxt/rank-util/utils"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ArticleExtractorToutiao(req *article_extractor.ArticleExtractorRequest) (*article_extractor.ParseArticleResponse, error) {
	if strings.Contains(req.Body, `ignoreStatic: [/\.tanx\.com\//, /\.alicdn\.com\//, /\.mediav\.com/]`) ||
		strings.Contains(req.Body, `ignoreStatic: [/mediav\.com/, /meimotuan\.com/, /comments8\.com/, /baidu\.com/, /google-analytics\.com/, /dist\/vtt\.js/]`) {
		return nil, errors.New("不是有效的html页面")
	}

	articleExtractorResponse := &article_extractor.ParseArticleResponse{}
	if req.RequestType == 1 {
		sugJson := gjson.Get(req.Body, "data")
		sugJson.ForEach(func(key, value gjson.Result) bool {
			hasVideo := gjson.Get(value.String(), "has_video")
			if hasVideo.Exists() && hasVideo.Raw == "false" {
				articleUrl := gjson.Get(value.String(), "article_url").String()
				intro := gjson.Get(value.String(), "abstract").String()
				name := gjson.Get(value.String(), "media_name").String()
				headImage := gjson.Get(value.String(), "media_avatar_url").String()
				guid := gjson.Get(value.String(), "id").String()
				fmt.Println(guid)
				coverImage := gjson.Get(value.String(), "image_url").String()
				issueTime := strings.Split(gjson.Get(value.String(), "datetime").String(), " ")[0]
				//title := gjson.Get(value.String(), "title").String()
				readCount, err := strconv.Atoi(gjson.Get(value.String(), "read_count").String())
				if err != nil {
					readCount = -2
				}

				article := article_extractor.Article5118{
					Guid: guid,
					//Title:                title,
					Url:                  articleUrl,
					Name:                 name,
					Intro:                intro,
					HeadImage:            headImage,
					IssueTime:            issueTime,
					CoverImage:           coverImage,
					ReadCount:            readCount,
					AddTime:              time.Now().Format("2006-01-02 15:04:05"),
					ResourcePlatformName: req.ResourcePlatformName,
				}
				articleExtractorResponse.Article = append(articleExtractorResponse.Article, article)

			}
			return true
		})

		nextPageUrl := ""
		r := regexp.MustCompile(`&offset=(.*?)&`)
		subMatch := r.FindAllStringSubmatch(req.Url, -1)
		if len(subMatch) > 0 {
			offset, _ := strconv.Atoi(gjson.Get(req.Body, "offset").String())
			nextPageUrl = strings.Replace(req.Url, fmt.Sprintf("offset=%s", subMatch[0][1]), fmt.Sprintf("offset=%d", offset), -1)
		}
		articleExtractorResponse.NextPageUrl = nextPageUrl
		articleExtractorResponse.UniqueKey = utils.DcArticleUniqueKey(articleExtractorResponse.NextPageUrl)

	} else {
		articleContent := ""
		isOriginal := 1
		catalogName := ""
		title := ""

		if strings.Contains(req.Body, "bid: 'toutiao_pc'") { // 第一类型网页
			unescapeStr := html.UnescapeString(req.Body)
			unquoteStr, err := strconv.Unquote(strings.Replace(strconv.Quote(string(unescapeStr)), `\\u`, `\u`, -1))
			if err != nil {
				unquoteStr = unescapeStr
			}

			r := regexp.MustCompile(`(?s)content: '(.*?)'.slice\(6, -6\)`)
			subMatch := r.FindAllStringSubmatch(unquoteStr, -1)
			if len(subMatch) > 0 {
				articleContent = subMatch[0][1]
			}

			r = regexp.MustCompile(`isOriginal: (.*?),`)
			subMatch = r.FindAllStringSubmatch(req.Body, -1)
			if len(subMatch) > 0 {
				original := subMatch[0][1]
				if original == "false" {
					isOriginal = 0
				}
			}

			r = regexp.MustCompile(`chineseTag: '(.*?)',`)
			subMatch = r.FindAllStringSubmatch(req.Body, -1)
			if len(subMatch) > 0 {
				catalogName = subMatch[0][1]
			}

			r = regexp.MustCompile(`title: '&quot;(.*?)&quot;'.slice\(6, -6\),`)
			subMatch = r.FindAllStringSubmatch(req.Body, -1)
			if len(subMatch) > 0 {
				title = subMatch[0][1]
			}

		} else { // 第二类型网页
			dom, err := goquery.NewDocumentFromReader(strings.NewReader(req.Body))
			if err != nil {
				return nil, err
			}

			title = strings.Replace(strings.Replace(dom.Find("header h1").Text(), " ", "", -1), "\n", "", -1)

			dom.Find("article p").Each(func(i int, selection *goquery.Selection) {
				imgUrl := selection.Find("img").AttrOr("src", "")
				if imgUrl != "" {
					articleContent += fmt.Sprintf("<p><img src=\"%s\"></p>", imgUrl)
				}

				paragraph := selection.Text()
				if paragraph != "" {
					articleContent += fmt.Sprintf("<p>%s</p>", paragraph)
				}
			})
		}

		article := article_extractor.Article5118{
			Title:          title,
			CharsCount:     len(articleContent),
			IsKOL:          0,
			IsOriginal:     isOriginal,
			ArticleContent: articleContent,
			CatalogName:    catalogName,
			LikeCount:      0,
			ForwardCount:   0,
			CommentCount:   0,
			EnTitle:        "",
			ChTitle:        "",
		}
		articleExtractorResponse.Article = append(articleExtractorResponse.Article, article)
	}

	articleExtractorResponse.RequestType = req.RequestType

	return articleExtractorResponse, nil
}
