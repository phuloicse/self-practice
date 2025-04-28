package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	// pb "github.com/your-module-path/proto"
	pb "serviceA/proto"
	"google.golang.org/grpc"
)

var podName = getPodName()

func getPodName() string {
	name, err := os.Hostname()
	if err != nil {
		return "unknown-pod"
	}
	return name
}

func connectToB(userID string) error {
	conn, err := grpc.Dial("service-b:50051", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		return fmt.Errorf("failed to connect to B: %w", err)
	}
	defer conn.Close()

	client := pb.NewServiceBClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := client.HandleRequest(ctx, &pb.Request{
		UserId:    userID,
		SenderPod: podName,
	})
	if err != nil {
		return fmt.Errorf("error calling B: %w", err)
	}

	log.Printf("Response from B: %s\n", resp.Message)
	return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user-id")
	if userID == "" {
		http.Error(w, "Missing user-id", http.StatusBadRequest)
		return
	}

	if err := connectToB(userID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Sent request to B with user-id: %s\n", userID)
}

func main() {
	http.HandleFunc("/connect-b", handler)
	log.Println("🚀 Service A HTTP server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
