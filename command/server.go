package command

import (
	"attachment/common/logger"
	"attachment/controller"
	"attachment/db"
	"attachment/file_uploader"
	"attachment/pb"
	"attachment/service"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	var serverPort string
	defaultServerPort := viper.GetString("SERVER_PORT")
	serverCmd.PersistentFlags().StringVar(&serverPort, "port", defaultServerPort, "Server port")
	viper.BindPFlag("SERVER_PORT", serverCmd.PersistentFlags().Lookup("port"))

	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run server",
	Run: func(cmd *cobra.Command, args []string) {
		lis, err := net.Listen("tcp", ":8070")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		defer lis.Close()

		grpcServer := grpc.NewServer()

		//gRPCClient:=grpc.Dial()

		db := db.InitDB()
		FileInterface := file_uploader.NewFileUploaderFactory()
		svc := service.NewAttachmentService(db, FileInterface)

		server := controller.AttachmentServer{
			Svc: svc,
		}

		//service.NewAttachmentService(entClient, uploader fileUploader.FileUploaderInterface)
		fmt.Println("server is listening on ", "8070")

		pb.RegisterAttachmentServiceServer(grpcServer, server)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalln("Failed to serve:", err)
		}

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		logger.LogInfo("Shutting down server...")

		//shutdown gin

		time.Sleep(time.Millisecond * 100)
		fmt.Println("bye bye")
	},
}
