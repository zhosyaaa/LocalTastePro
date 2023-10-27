package main

//
//import (
//	"context"
//	"fmt"
//	"google.golang.org/grpc"
//	"log"
//)
//
//const locationServiceAddress = "127.0.0.1:50051" // Адрес и порт вашего сервиса местоположения
//
//func main() {
//	// Устанавливаем соединение с сервисом местоположения
//	conn, err := grpc.Dial(locationServiceAddress, grpc.WithInsecure())
//	if err != nil {
//		log.Fatalf("Failed to connect to the location service: %v", err)
//	}
//
//	defer conn.Close()
//
//	// Создаем клиент для сервиса местоположения
//	client := pb.NewLocationServiceClient(conn)
//
//	// Создаем запрос
//	location := &pb.Location{
//		Latitude:  123.456, // Замените на координаты местоположения пользователя
//		Longitude: 789.012, // Замените на координаты местоположения пользователя
//	}
//
//	request := &pb.LocationRequest{
//		Location: location,
//	}
//
//	// Вызываем SendLocation
//	response, err := client.SendLocation(context.Background(), request)
//	if err != nil {
//		log.Fatalf("Failed to call SendLocation: %v", err)
//	}
//
//	// Обработка ответа от SendLocation
//	// Ваши дальнейшие действия с ответом от SendLocation
//	fmt.Printf("Location response: %v\n", response)
//}
