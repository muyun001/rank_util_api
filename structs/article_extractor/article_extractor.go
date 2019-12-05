package article_extractor

type ArticleExtractorRequest struct {
	Url                  string `json:"url"`
	Body                 string `json:"body" bson:"body"`
	RequestType          int    `json:"request_type"`
	ResourcePlatformName string `json:"resourcePlatformName"`
}

type Article5118 struct {
	ID                   int    `gorm:"primary_key" json:"id"`
	KeywordId            int    `gorm:"type:int;" json:"keyword_id"`
	Guid                 string `gorm:"type:varchar(128);unique_index:guid" json:"guid"`
	Title                string `gorm:"type:varchar(128)" json:"title"`
	IssueTime            string `gorm:"type:varchar(64)" json:"issue_time"`
	CharsCount           int    `gorm:"type:int" json:"chars_count"`
	ReadCount            int    `gorm:"type:int" json:"read_count"`
	LikeCount            int    `gorm:"type:int" json:"like_count"`
	IsKOL                int    `gorm:"type:int" json:"is_kol"`
	IsOriginal           int    `gorm:"type:int" json:"is_original"`
	CoverImage           string `gorm:"type:varchar(255)" json:"cover_image"`
	AddTime              string `gorm:"type:varchar(64)" json:"add_time"`
	Name                 string `gorm:"type:varchar(128)" json:"name"`
	HeadImage            string `gorm:"type:varchar(255)" json:"head_image"`
	ResourcePlatformName string `gorm:"type:varchar(128)" json:"resource_platform_name"`
	CatalogName          string `gorm:"type:varchar(64)" json:"catalog_name"`
	Intro                string `gorm:"type:text" json:"intro"`
	ArticleContent       string `gorm:"type:longtext" json:"article_content"`
	Url                  string `gorm:"type:varchar(64)" json:"url"`
	ForwardCount         int    `gorm:"type:int" json:"forward_count"`
	CommentCount         int    `gorm:"type:int" json:"comment_count"`
	EnTitle              string `gorm:"type:varchar(8)" json:"en_title"`
	ChTitle              string `gorm:"type:varchar(8)" json:"ch_title"`
}

type ParseArticleResponse struct {
	Article     []Article5118 `json:"article"`
	UniqueKey   string        `json:"unique_key"`
	NextPageUrl string        `json:"next_page_url"`
	RequestType int           `json:"request_type"`
}
