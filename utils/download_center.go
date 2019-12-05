package utils

import (
	"crypto/md5"
	"fmt"
	"time"
)

func DcUniqueKey(requestUrl string, capture bool, searchCycle int) string {
	dateStr := time.Now().Format("2006-01-02")
	sourceStr := fmt.Sprintf("%s%t%s%d", requestUrl, capture, dateStr, searchCycle)
	return fmt.Sprintf("%x", md5.Sum([]byte(sourceStr)))
}

func DcArticleUniqueKey(requestUrl string) string {
	dateStr := time.Now().Format("2006-01-02")
	sourceStr := fmt.Sprintf("%s%s", requestUrl, dateStr)
	return fmt.Sprintf("%x", md5.Sum([]byte(sourceStr)))
}
