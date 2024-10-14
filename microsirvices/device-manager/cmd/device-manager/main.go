package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/pgxpool"
	device_manager_srv "github.com/nini-k/architecture-sprint-3/microsirvices/device-manager/internal/app/grpc/device-manager"
	pb "github.com/nini-k/architecture-sprint-3/microsirvices/device-manager/internal/pb/device-manager"
	device_manager_svc "github.com/nini-k/architecture-sprint-3/microsirvices/device-manager/internal/services/device-manager"
	device_storage "github.com/nini-k/architecture-sprint-3/microsirvices/device-manager/internal/storage/postgresql/device"
	"google.golang.org/grpc"
)

var (
	grpcServerEndpoint        = os.Getenv("GRPC_ENDPOINT")
	grpcGatewayServerEndpoint = os.Getenv("GRPC_GATEWAY_ENDPOINT")
	postgresURL               = os.Getenv("POSTGRES_URL")
)

func main() {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, postgresURL)
	if err != nil {
		log.Fatalf("create postrgres pool: %v", err)
	}
	defer pool.Close()

	deviceStorage := device_storage.New(pool)
	deviceManagerService := device_manager_svc.New(deviceStorage)
	deviceManagerServiceServer := device_manager_srv.New(deviceManagerService)

	go func() {
		listener, err := net.Listen("tcp", grpcServerEndpoint)
		if err != nil {
			log.Fatalf("listen: %v", err)
		}
		server := grpc.NewServer()

		pb.RegisterDeviceManagerServiceServer(server, deviceManagerServiceServer)

		log.Printf("grpc server listening at %v", grpcServerEndpoint)
		if err = server.Serve(listener); err != nil {
			log.Fatalf("serve grpc server: %v", err)
		}
	}()

	gwmux := runtime.NewServeMux()

	if err := pb.RegisterDeviceManagerServiceHandlerServer(ctx, gwmux, deviceManagerServiceServer); err != nil {
		log.Fatal("register device manager service handler from endpoint")
	}

	mux := http.NewServeMux()

	mux.Handle("/", gwmux)

	mux.HandleFunc("/docs/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./swagger-ui/device-manager/device-manager.swagger.json")
	})

	mux.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir("./swagger-ui"))))

	log.Printf("grpc gateway server listening at %v", grpcGatewayServerEndpoint)
	if err := http.ListenAndServe(grpcGatewayServerEndpoint, mux); err != nil {
		log.Fatalf("listen and serve grpc gateway: %v", err)
	}
}
