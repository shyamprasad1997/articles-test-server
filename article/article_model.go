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

type PostArticlesRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedBy   uint64 `json:"created_by"`
}

type PostArticlesResponse struct {
	Status bool `json:"status"`
}
