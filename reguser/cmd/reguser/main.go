package main

import (
	"context"
	"goback1/lesson6/reguser/internal/infrastructure/api/handler"
	"goback1/lesson6/reguser/internal/infrastructure/api/routerchi"
	"goback1/lesson6/reguser/internal/infrastructure/db/files/userfilemanager"
	"goback1/lesson6/reguser/internal/infrastructure/server"
	"goback1/lesson6/reguser/internal/usecases/app/repos/userrepo"
	"log"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	// ust := usermemstore.NewUsers()
	ust, err := userfilemanager.NewUsers("./data.json", "mem://userRefreshTopic")
	if err != nil {
		log.Fatal(err)
	}

	us := userrepo.NewUsers(ust)
	hs := handler.NewHandlers(us)
	// h := defmux.NewRouter(hs)
	h := routerchi.NewRouterChi(hs)
	srv := server.NewServer(":8000", h)

	srv.Start(us)
	log.Print("Start")

	<-ctx.Done()

	srv.Stop()
	cancel()
	ust.Close()

	log.Print("Exit")
}
