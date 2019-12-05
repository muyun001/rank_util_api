package rank_extractor_service

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/panwenbin/ghttpclient"
	"gitlab.fxt.cn/fxt/rank-util/structs/rank_extractor"
	"gitlab.fxt.cn/fxt/rank-util/utils"
	"net/url"
	"regexp"
	"strings"
	"time"
)

func RankExtractorBaiduPc(req *rank_extractor.RankExtractorRequest) (*rank_extractor.RankExtractorResponse, error) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(req.Body))
	if err != nil {
		return nil, err
	}

	if false == strings.Contains(req.Body, "html") ||
		strings.Contains(req.Body, "location.href.replace") {
		return nil, errors.New("不是有效的html页面")
	}

	if strings.Contains(req.Body, `id="wrap"`) ||
		strings.Contains(req.Body, "页面不存在_百度搜索") ||
		false == strings.Contains(req.Body, "<title>") ||
		false == strings.Contains(req.Body, `id="container"`) ||
		false == strings.Contains(req.Body, `id="content_left"`) {
		return &rank_extractor.RankExtractorResponse{}, nil
	}

	ckDomain := utils.TopDomain(req.CheckMatch)
	ckSiteName := req.SiteName
	rank := req.StartRank
	var ranks []int
	chanLen := 0
	type RankDomain struct {
		Rank   int
		Domain string
	}
	rankDomainChan := make(chan *RankDomain, 10)

	dom.Find("div.c-container").Each(func(i int, selection *goquery.Selection) {
		rank += 1
		domainOrSiteName := ""
		indexedUrl := ""
		indexedSiteName := ""

		cssList := []string{
			"a.c-showurl > span",
			"a.c-showurl",
			"span.c-showurl",
			"div.g",
			"span.g",
		}
		for _, css := range cssList {
			domainOrSiteName = strings.Replace(selection.Find(css).Text(), " ", "", -1)
			if domainOrSiteName == "" {
				continue
			}
		}

		hasDots := strings.Contains(domainOrSiteName, "...")
		hasSlash := strings.Contains(strings.Replace(domainOrSiteName, "://", "", 1), "/")
		notComplete := !hasSlash && hasDots
		if domainOrSiteName == "" || notComplete {
			address := selection.Find("h3").Find("a").AttrOr("href", "")
			chanLen++
			go func(rank int) {
				realHost := findRealAddress(address)
				rankDomainChan <- &RankDomain{Rank: rank, Domain: realHost}
			}(rank)
			return
		}

		if strings.Contains(domainOrSiteName, ".") {
			indexedUrl = domainOrSiteName
			var indexedDomain string
			uri, err := url.ParseRequestURI(utils.FormatUrl(indexedUrl))
			if err == nil {
				indexedDomain = utils.TopDomain(uri.Host)
			}
			if strings.Compare(indexedDomain, ckDomain) == 0 {
				ranks = append(ranks, rank)
			}
		} else {
			indexedSiteName = domainOrSiteName
			if strings.Compare(indexedSiteName, ckSiteName) == 0 {
				ranks = append(ranks, rank)
			}
		}
	})

	for i := 0; i < chanLen; i++ {
		rankDomain := <-rankDomainChan
		if rankDomain.Domain == "" {
			continue
		}
		showDomain := utils.TopDomain(rankDomain.Domain)

		if strings.Compare(showDomain, ckDomain) == 0 {
			ranks = append(ranks, rankDomain.Rank)
		}
	}

	rankExtractorResponse := &rank_extractor.RankExtractorResponse{}
	rankExtractorResponse.Ranks = ranks

	return rankExtractorResponse, nil
}

func findRealAddress(sourceUrl string) string {
	uri, err := url.ParseRequestURI(sourceUrl)
	if err != nil {
		return ""
	}

	client := ghttpclient.NewClient().Timeout(time.Second * 5).Url(sourceUrl).Get()
	response, err := client.Response()
	if err != nil {
		return ""
	}

	if response.Request.URL.Host == uri.Host { // baidu
		body, err := client.ReadBodyClose()
		if err != nil {
			return ""
		}

		re := regexp.MustCompile(`window.location.replace\("(.*?)"\)`)
		subMatch := re.FindStringSubmatch(string(body))
		if len(subMatch) == 2 {
			return subMatch[1]
		}
	} else {
		_ = response.Body.Close()
	}

	return response.Request.URL.Host
}
