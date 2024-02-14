package httpServer

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	natsCustom "nats_example/server/nats"
	"time"
)

func (s *Server) MapHandlers(ctx context.Context, app *fiber.App) error {

	ctx, cancel := context.WithCancel(ctx)

	// -------------------------------------------------------------------------------------

	nts, err := natsCustom.NewNats(s.cfg)
	if err != nil {
		panic(err)
	}

	// -------------------------------------------------------------------------------------

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	// -------------------------------------------------------------------------------------

	var cnt int
	go func(ctx context.Context) {
		time.Sleep(3 * time.Second)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				cnt++
				msg := fmt.Sprintf("MSG %d", cnt)
				fmt.Println("-> Send msg: ", msg)
				err := nts.PublishMessage(s.cfg.Nats.Topic, msg)
				if err != nil {
					fmt.Printf("error in publish message: %v\n", err)
					cancel()
				}

				time.Sleep(time.Second * 6)
			}
		}
	}(ctx)

	go func(ctx context.Context) {
		time.Sleep(2 * time.Second)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				var (
					msg string
					err error
				)

				msg, err = nts.SubscribeAndReceiveMessage(s.cfg.Nats.Topic)
				if err != nil {
					fmt.Printf("error in receive message: %v", err)
					cancel()
				}
				fmt.Println("<- msg receive: ", msg)
				time.Sleep(time.Second * 5)
			}
		}
	}(ctx)

	return nil
}
