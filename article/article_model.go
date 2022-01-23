package article

type Article struct {
	ArticleId   uint64 `json:"article_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type GetArticlesResponse struct {
	Status   bool      `json:"status"`
	Articles []Article `json:"articles"`
}
