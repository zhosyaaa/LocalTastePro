package main

import "RecommendationService/internal/server"

//
//func main() {
//	lis, err := net.Listen("tcp", ":50051")
//	if err != nil {
//		log.Fatalf("Failed to listen: %v", err)
//	}
//
//	serv := grpc.NewServer()
//	pb.RegisterRecommendationServiceServer(serv, &server.RecommendationServiceImpl{})
//
//	log.Println("gRPC server is running on :50051")
//
//	if err := serv.Serve(lis); err != nil {
//		log.Fatalf("Failed to serve: %v", err)
//	}
//}

func main() {
	server.Parse()
}
