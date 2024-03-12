// Central Services for the Ranch Systems
// A rules engine written in go
//
//  Current rules
//
//  Turn on solar pump if
//     It's been above 40 degrees for an hour
//     The Sun is out
//     No water in basement
//  
//  Turn on main pump
//     It's been above 40 degrees for an hour
//     The Sun is out
//     No water in basement
//     Main Batteries are above 90%



// 192.168.0.106 -p 1883 -t "solar_assistant/total/battery_state_of_charge/state"
package main


import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/stianeikeland/go-rpio/v4"
	"strconv"
	"time"
	"os"
)



// Global Variables, these are set by MQTT Callbacks
var globalMainSOC int = 00

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	globalMainSOC, err := strconv.Atoi(string(msg.Payload()))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(globalMainSOC)
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func main() {
	var broker = "192.168.0.106"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("CentralServices")
	opts.SetUsername("emqx")
	opts.SetPassword("public")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	sub(client)

	// Open the gpio memory, bail out if it fails
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	// While loop to check the rules
	for {
		time.Sleep(10 * time.Second)
	}
	client.Disconnect(250)
}

func publish(client mqtt.Client) {
	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("Message %d", i)
		token := client.Publish("topic/test", 0, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}
}

func sub(client mqtt.Client) {
	topic := "solar_assistant/total/battery_state_of_charge/state"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}
