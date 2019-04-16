package main

import (
	"fmt"
	"log"

	"github.com/micro/go-micro"

	pb "github.com/kareemarab/space/user/proto/user"
)

func main() {

	// create db connection
	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// will migrate User struct to database and will check for
	// changes and will automatically migrate them everytime
	// the service is started.

	db.AutoMigrate(&pb.User{})

	repo := &UserRepository{db}
	tokenService := &TokenService{repo}

	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	srv.Init()

	// Register handler
	pb.RegisterUserServiceHandler(srv.Server(), &service{repo, tokenService})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}
