package main

import (
	"fmt"
	"net/http"

	"github.com/EsneiderGrVc/go_server/controller"
	"github.com/EsneiderGrVc/go_server/dbconfig"
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
	dbconfig.NewDbConfig().Config()

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Up and Running...")
	})

	httpRouter.GET("/deliveries", deliverController.GetAllDeliveries)
	httpRouter.GET("/deliveries/{id}", deliverController.GetDelivery)
	httpRouter.POST("/deliveries", deliverController.CreateDelivery)

	httpRouter.POST("/bots", botController.CreateBot)
	httpRouter.GET("/bots", botController.GetBotsOrderByZone)

	httpRouter.SERVE(port)
}
