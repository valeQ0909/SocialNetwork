package controller

import "socialnetwork_back_go/service"

type PostsListResponse struct {
	Response
	PostsList []service.FmtPost `json:"post_list"`
}
