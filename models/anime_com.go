package models

type AnimeCom struct {
	Id           int    `json:"id" form:"id" url:"id"`
	AnimeCover   string `json:"anime_cover" form:"anime_cover" url:"anime_cover"`
	AnimeTitle   string `json:"anime_title" form:"anime_title" url:"anime_title"`
	AnimeContent string `json:"anime_content" form:"anime_content" url:"anime_content"`
	MyComment    string `json:"my_comment" form:"my_comment" url:"my_comment"`
	VisitTime    string `json:"visit_time" form:"visit_time" url:"visit_time"`
	CreateTime   string `json:"create_time" form:"create_time" url:"create_time"`
}
