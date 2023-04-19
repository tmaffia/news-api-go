package newsapi

type ArticleResponse struct {
	code         string
	status       string
	message      string
	totalResults int32
	articles     []Article
}
