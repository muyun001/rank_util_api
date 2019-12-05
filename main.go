package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.fxt.cn/fxt/rank-util/actions/article_extractor"
	"gitlab.fxt.cn/fxt/rank-util/actions/keyword_include_extractor"
	"gitlab.fxt.cn/fxt/rank-util/actions/domain_include_extractor"
	"gitlab.fxt.cn/fxt/rank-util/actions/rank-extractor"
	"gitlab.fxt.cn/fxt/rank-util/actions/request_builder"
)

func main() {
	router := gin.Default()

	//	构建请求
	router.POST("/request-builder/baidu-pc", request_builder.RequestBuilderBaiduPc)
	router.POST("/request-builder/baidu-mobile", request_builder.RequestBuilderBaiduMobile)
	router.POST("/request-builder/baidu-mini-program", request_builder.RequestBuilderBaiduMiniProgram)
	router.POST("/request-builder/360-pc", request_builder.RequestBuilder360Pc)
	router.POST("/request-builder/sogou-pc", request_builder.RequestBuilderSogouPc)
	router.POST("/request-builder/sm-mobile", request_builder.RequestBuilderSmMobile)
	router.POST("/request-builder/sug-baidu-pc", request_builder.RequestBuilderSugBaiduPc)
	router.POST("/request-builder/sug-baidu-mobile", request_builder.RequestBuilderSugBaiduMobile)
	router.POST("/request-builder/sug-sogou-pc", request_builder.RequestBuilderSugSogouPc)
	router.POST("/request-builder/sug-sogou-mobile", request_builder.RequestBuilderSugSogouMobile)
	router.POST("/request-builder/sug-360-pc", request_builder.RequestBuilderSug360Pc)
	router.POST("/request-builder/sug-360-mobile", request_builder.RequestBuilderSug360Mobile)
	router.POST("/request-builder/sug-sm-mobile", request_builder.RequestBuilderSugSmMobile)
	router.POST("/request-builder/weixin", request_builder.RequestBuilderWeixin)
	router.POST("/request-builder/toutiao", request_builder.RequestBuilderToutiao)

	//	排名解析
	router.POST("/rank-extractor/baidu-pc", rank_extractor.RankExtractorBaiduPc)
	router.POST("/rank-extractor/baidu-mobile", rank_extractor.RankExtractorBaiduMobile)
	router.POST("/rank-extractor/baidu-mini-program", rank_extractor.RankExtractorBaiduMiniProgram)
	router.POST("/rank-extractor/sogou-pc", rank_extractor.RankExtractorSogouPc)
	router.POST("/rank-extractor/360-pc", rank_extractor.RankExtractor360Pc)
	router.POST("/rank-extractor/sm-mobile", rank_extractor.RankExtractorSmMobile)
	router.POST("/rank-extractor/sug-baidu-pc", rank_extractor.RankExtractorSugBaiduPc)
	router.POST("/rank-extractor/sug-baidu-mobile", rank_extractor.RankExtractorSugBaiduMobile)
	router.POST("/rank-extractor/sug-sogou-pc", rank_extractor.RankExtractorSugSogouPc)
	router.POST("/rank-extractor/sug-sogou-mobile", rank_extractor.RankExtractorSugSogouMobile)
	router.POST("/rank-extractor/sug-360-pc", rank_extractor.RankExtractorSug360Pc)
	router.POST("/rank-extractor/sug-360-mobile", rank_extractor.RankExtractorSug360Mobile)
	router.POST("/rank-extractor/sug-sm-mobile", rank_extractor.RankExtractorSugSmMobile)

	//	文章解析
	router.POST("/article-extractor/weixin", article_extractor.ArticleExtractorWeixin)
	router.POST("/article-extractor/toutiao", article_extractor.ArticleExtractorToutiao)

	// keyword收录解析
	router.POST("/keyword-include-extractor/baidu-pc", keyword_include_extractor.KeywordIncludeExtractorBaiduPc)
	router.POST("/keyword-include-extractor/sogou-pc", keyword_include_extractor.KeywordIncludeExtractorSogouPc)
	router.POST("/keyword-include-extractor/360-pc", keyword_include_extractor.KeywordIncludeExtractor360Pc)
	router.POST("/keyword-include-extractor/sm-mobile", keyword_include_extractor.KeywordIncludeExtractorSmMobile)

	// domain收录解析
	router.POST("/domain-include-extractor/baidu-pc", domain_include_extractor.DomainIncludeExtractorBaiduPc)
	router.POST("/domain-include-extractor/sogou-pc", domain_include_extractor.DomainIncludeExtractorSogouPc)
	router.POST("/domain-include-extractor/360-pc", domain_include_extractor.DomainIncludeExtractor360Pc)
	router.POST("/domain-include-extractor/sm-mobile", domain_include_extractor.DomainIncludeExtractorSmMobile)

	router.Run(":8081")
}
