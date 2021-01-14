package main

import (
	"go-api-mux/controller"
	"go-api-mux/repository"
	"go-api-mux/router"
	"go-api-mux/service"
	"go-api-mux/utils"
	"os"
)

var (
	postRepository repository.PostRepository = repository.NewFireStoreRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewChiRouter()
)

func main() {
	utils.LoadEnv()
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPosts)
	httpRouter.SERVE((os.Getenv("PORT")))
}
