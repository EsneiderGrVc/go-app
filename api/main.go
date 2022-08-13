package main

import (
	"fmt"
	"net/http"

	"github.com/EsneiderGrVc/go_server/controller"
	"github.com/EsneiderGrVc/go_server/docs"
	_ "github.com/EsneiderGrVc/go_server/docs"
	router "github.com/EsneiderGrVc/go_server/http"
	"github.com/EsneiderGrVc/go_server/repository"
	"github.com/EsneiderGrVc/go_server/services"
	_ "github.com/lib/pq"
)

var (
	deliveryRepository repository.DeliveryRepository = repository.NewDeliveryRepository()
	deliveryService    services.DeliveryService      = services.NewDeliveryService(deliveryRepository)
	deliverController  controller.DeliveryController = controller.NewDeliveryController(deliveryService)

	botRepository repository.BotRepository = repository.NewBotRepository()
	botService    services.BotService      = services.NewBotService(botRepository)
	botController controller.BotController = controller.NewBotController(botService)

	httpRouter router.Router = router.NewMuxRouter()
)

func main() {
	const port string = ":80"
	// dbconfig.NewDbConfig().Config()

	docs.SwaggerInfo.Title = "KiwiBot API"
	docs.SwaggerInfo.Description = "Create deliveries and assign them to bots."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Up and Running...")
	})

	httpRouter.GET("/deliveries", deliverController.GetAllDeliveries)
	httpRouter.GET("/deliveries/{id}", deliverController.GetDelivery)
	httpRouter.POST("/deliveries", deliverController.CreateDelivery)

	httpRouter.POST("/bots", botController.CreateBot)
	httpRouter.GET("/bots", botController.GetBotsOrderByZone)

	httpRouter.SWAGGER("/swagger/")

	httpRouter.SERVE(port)
}
