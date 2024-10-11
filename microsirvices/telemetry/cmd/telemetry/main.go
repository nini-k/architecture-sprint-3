package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	telemetry_srv "github.com/nini-k/architecture-sprint-3/microsirvices/telemetry/internal/app/grpc/telemetry"
	pb "github.com/nini-k/architecture-sprint-3/microsirvices/telemetry/internal/pb/telemetry"
	telemetry_svc "github.com/nini-k/architecture-sprint-3/microsirvices/telemetry/internal/services/telemetry"
	telemetry_storage "github.com/nini-k/architecture-sprint-3/microsirvices/telemetry/internal/storage/mongo/telemetry"
	"google.golang.org/grpc"
)

var (
	grpcServerEndpoint        = os.Getenv("GRPC_ENDPOINT")
	grpcGatewayServerEndpoint = os.Getenv("GRPC_GATEWAY_ENDPOINT")
	mongoURL                  = os.Getenv("MONGO_URL")
)

func main() {
	ctx := context.Background()

	telemetryStorage := telemetry_storage.New()
	telemetryService := telemetry_svc.New(telemetryStorage)
	telemetryServiceServer := telemetry_srv.New(telemetryService)

	go func() {
		listener, err := net.Listen("tcp", grpcServerEndpoint)
		if err != nil {
			log.Fatalf("listen: %v", err)
		}
		server := grpc.NewServer()

		pb.RegisterTelemetryServiceServer(server, telemetryServiceServer)

		log.Printf("grpc server listening at %v", grpcServerEndpoint)
		if err = server.Serve(listener); err != nil {
			log.Fatalf("serve grpc server: %v", err)
		}
	}()

	gwmux := runtime.NewServeMux()

	if err := pb.RegisterTelemetryServiceHandlerServer(ctx, gwmux, telemetryServiceServer); err != nil {
		log.Fatal("register device manager service handler from endpoint")
	}

	mux := http.NewServeMux()

	mux.Handle("/", gwmux)

	mux.HandleFunc("/docs/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./swagger-ui/telemetry/telemetry.swagger.json")
	})

	mux.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir("./swagger-ui"))))

	log.Printf("grpc gateway server listening at %v", grpcGatewayServerEndpoint)
	if err := http.ListenAndServe(grpcGatewayServerEndpoint, mux); err != nil {
		log.Fatalf("listen and serve grpc gateway: %v", err)
	}
}
