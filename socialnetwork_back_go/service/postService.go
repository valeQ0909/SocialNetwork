package service

import "socialnetwork_back_go/models"

type FmtPost struct {
	models.TablePost
	Author        FmtUser `json:"author,omitempty"`
	FavoriteCount int64   `json:"favorite_count,omitempty"`
	CommentCount  int64   `json:"comment_count,omitempty"`
	IsFavorite    bool    `json:"is_favorite,omitempty"`
}
