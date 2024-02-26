package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"publisher/sensors"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

var BROKER_URL string
var MQTT_USER string
var MQTTT_PASSWORD string

// function to generate a random sleep time
func randSleep() {
	sleepTime := rand.Intn(5) + 1
	time.Sleep(time.Duration(sleepTime) * time.Second)
}

// function to publish a messagem in a given topic
func pubMessage(topic string, csvPath string, broker_url *string, user *string, password *string) {
	// connecting to a broker
	opts := MQTT.NewClientOptions().AddBroker(fmt.Sprintf("ssl://%s", *broker_url))
	opts.SetClientID("go_publisher")
	opts.SetUsername(*user)
	opts.SetPassword(*password)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// loop to emit the messages
	for {
		// Getting the readings of a given sensor
		solarReading, err := sensors.GenerateReading(csvPath)
		if err == nil {
			randSleep()
			fmt.Printf("\033[33mDado lido: %s  \033[0m\n", solarReading)
			token := client.Publish(topic, 1, false, solarReading)
			token.Wait()
		} else {
			fmt.Printf("\033[31m%s\033[0m\n", err.Error())
			break
		}
	}
	fmt.Println("\033[35mPublisher encerrado! \033[0m")
}
func init() {
	log.Print("Loading .env variables...")

	envPath := "./.env"

	if len(os.Args) > 2 {
		envPath = os.Args[1]
	}

	godotenv.Load(envPath)

	BROKER_URL = os.Getenv("BROKER_URL")
	MQTT_USER = os.Getenv("MQTT_PUB")
	MQTTT_PASSWORD = os.Getenv("MQTT_PUB_PASSWORD")

	if BROKER_URL == "" || MQTT_USER == "" || MQTTT_PASSWORD == "" {
		log.Printf("ENV variables missing.\n")
		os.Exit(1)
	}

	log.Println("ENV variables loaded!")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("\033[31mMissing csv path\033[0m")
		os.Exit(1)
	}
	pubMessage("ponderada3/data", os.Args[len(os.Args)-1], &BROKER_URL, &MQTT_USER, &MQTTT_PASSWORD)
}
