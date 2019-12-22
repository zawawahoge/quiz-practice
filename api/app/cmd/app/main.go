package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/zawawahoge/quiz-practice/api/app/infra/impl/adminserviceimpl"
	"github.com/zawawahoge/quiz-practice/api/app/infra/impl/loginserviceimpl"
	database "github.com/zawawahoge/quiz-practice/api/app/infra/repository"
	"github.com/zawawahoge/quiz-practice/api/app/proto/v1/service"
)

const (
	port = ":9090"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	envMysql := os.Getenv("MYSQL_CONNECTION")
	if envMysql == "" {
		panic("MYSQL_CONNECTION is empty")
	}
	db, err := gorm.Open("mysql", envMysql)
	if err != nil {
		panic(err)
	}

	dbAccountRepository := database.NewAccountRepository(db)

	s := grpc.NewServer()
	service.RegisterLoginServiceServer(s, loginserviceimpl.New(dbAccountRepository))
	service.RegisterAdminServiceServer(s, adminserviceimpl.New(dbAccountRepository))

	fmt.Println("starting server...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
