package models

type PostRequest struct {
	Text     string `json:"text"`
	PostType string `json:"post_type"`
}

type PostResponse struct {
	Text     string `json:"text"`
	PostType string `json:"post_type"`
}

type Post struct {
	Text     string `json:"text"`
	PostType string `json:"post_type"`
}
