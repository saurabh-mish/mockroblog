package routes

var routes = []route{
	newRoute("GET", "/hello", Hello),
	newRoute("GET", "/api/v1/users", AllUsers),
	newRoute("GET", "/api/v1/posts", AllPosts),
}
