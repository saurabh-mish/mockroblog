package routes

//import "mockroblog/pkg/controllers"

var routes = []route{
	newRoute("GET", "/api/v1/hello", Hello),

	// account microservice
	newRoute("GET", "/api/v1/users", GetAllUsers),
	newRoute("POST", "/api/v1/user", CreateUser),
	newRoute("GET", "/api/v1/user/([0-9]+)", RetrieveUser),

	// post microservice
	newRoute("POST",   "/api/v1/post", CreatePost),
	newRoute("DELETE", "/api/v1/post/id=([0-9]+)", DeletePost),
	newRoute("GET",    "/api/v1/post/id=([0-9]+)", RetrievePost),
	newRoute("GET", "/api/v1/posts/number=([0-9]+)", RetrieveRecentPosts),
	newRoute("GET", "/api/v1/posts/number=([0-9]+)&community=([^/]+)", RetrieveRecentPostsFromCommunity),

	// vote microservice
	newRoute("POST", "/api/v1/vote/post/id=([0-9]+)&action=([^/]+)", CastVoteForPost),
	newRoute("GET", "/api/v1/votes/post/id=([0-9]+)", GetVotesForPost),
	newRoute("GET", "/api/v1/votes/number=([0-9]+)", GetTopScoringPosts),
}
