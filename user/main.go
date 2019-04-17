package main

import (
	"fmt"
	"log"

	pb "github.com/kareemarab/space/user/proto/auth"
	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/kubernetes"
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
	// srv := k8s.NewService(

	// 	// This name must match the package name given in your protobuf definition
	// 	micro.Name("shippy.auth"),
	// )

	srv := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("space.auth"),
		// micro.Version("latest"),
	)

	srv.Init()

	// Register handler
	pb.RegisterAuthHandler(srv.Server(), &service{repo, tokenService})
	// Create(pb.User{Email: "kareem3043@gmail.com", Password: "lecfej"})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}
