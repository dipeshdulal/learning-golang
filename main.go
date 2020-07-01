package main

import (
	"context"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	base64encdec "wesionary.team/dipeshdulal/golang-test/base64"
	"wesionary.team/dipeshdulal/golang-test/fileupload"
	gogrpc "wesionary.team/dipeshdulal/golang-test/gogrpc/quickstart"
	"wesionary.team/dipeshdulal/golang-test/googlebucket"
	"wesionary.team/dipeshdulal/golang-test/sendmail"
	"wesionary.team/dipeshdulal/golang-test/structparser"
	"wesionary.team/dipeshdulal/golang-test/websockets"
)

// GRPC SERVER START

const (
	grpcPort = ":50051"
)

type server struct {
	gogrpc.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *gogrpc.HelloRequest) (*gogrpc.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &gogrpc.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func runGrpcServer() {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	gogrpc.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// GRPC SERVER END

func main() {

	go runGrpcServer()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env file error: ", err.Error())
	}

	router := gin.Default()

	router.GET("/send-email-sync", sendmail.SendMailSyncController)

	router.GET("/send-email-async", sendmail.SendMailAsyncController)

	router.GET("/check-smtp-server", sendmail.SendMailCheckConnection)

	router.POST("/single-file-upload", fileupload.HandleSingleFileUpload)

	router.POST("/base64-file-upload", base64encdec.HandleBase64FileUpload)

	router.GET("/base64-octet-stream", base64encdec.HandleBase64OctetStream)

	router.POST("/multiple-file-upload", fileupload.HandleMultipleFileUpload)

	router.POST("/cloud-storage-bucket", googlebucket.HandleFileUploadToBucket)

	router.GET("/handle-request-parsing", structparser.HandleGetRequestParsing)

	router.GET("/handle-websocket-echo", websockets.HandleWebsocketEchoRequest)

	router.Run(":5000")

}
