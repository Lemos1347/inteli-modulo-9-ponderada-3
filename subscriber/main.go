package main

import (
	"fmt"
	"log"
	"os"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

var BROKER_URL string
var MQTT_USER string
var MQTTT_PASSWORD string

var messagePubHandler MQTT.MessageHandler = func(_ MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Recebido dado solar: %s do tópico: %s as %s \n", msg.Payload(), msg.Topic(), time.Now().Format(time.RFC3339))
}

func runSub(topic string, callback MQTT.MessageHandler, broker_url *string, user *string, password *string) {
	opts := MQTT.NewClientOptions().AddBroker(fmt.Sprintf("ssl://%s", *broker_url))
	opts.SetClientID("go_subscriber")
	opts.SetUsername(*user)
	opts.SetPassword(*password)
	opts.SetDefaultPublishHandler(callback)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe(topic, 1, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}

	fmt.Println("Subscriber está rodando. Pressione CTRL+C para sair.")
	select {}
}

func init() {
	log.Print("Loading .env variables...")

	envPath := "./.env"

	if len(os.Args) > 1 {
		envPath = os.Args[1]
	}

	godotenv.Load(envPath)

	BROKER_URL = os.Getenv("BROKER_URL")
	MQTT_USER = os.Getenv("MQTT_SUB")
	MQTTT_PASSWORD = os.Getenv("MQTT_SUB_PASSWORD")

	if BROKER_URL == "" || MQTT_USER == "" || MQTTT_PASSWORD == "" {
		log.Printf("ENV variables missing.\n")
		os.Exit(1)
	}

	log.Println("ENV variables loaded!")
}

func main() {
	log.Printf("ENV vars:\n%s\n%s\n%s", BROKER_URL, MQTT_USER, MQTTT_PASSWORD)
	runSub("ponderada3/data", messagePubHandler, &BROKER_URL, &MQTT_USER, &MQTTT_PASSWORD)
}
