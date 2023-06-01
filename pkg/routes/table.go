package routes

//import "mockroblog/pkg/controllers"

var routes = []route{
	newRoute("GET", "/api/v1/hello", Hello),

	newRoute("GET", "/api/v1/users", GetAllUsers),
	newRoute("POST", "/api/v1/user", CreateUser),
	newRoute("GET", "/api/v1/user/([0-9]+)", RetrieveUser),
	// newRoute("PATCH", "/api/v1/user/:id", ModifyUser),
	// newRoute("DELETE", "/api/v1/user/:id", DeleteUser),

	newRoute("GET", "/api/v1/posts", GetAllPosts),
	// newRoute("POST", "/api/v1/post", CreatePost),
	// newRoute("GET", "/api/v1/post/:id", RetrievePost),
	// newRoute("PATCH", "/api/v1/post/:id", ModifyPost),
	// newRoute("DELETE", "/api/v1/post/:id", DeletePost),
}
