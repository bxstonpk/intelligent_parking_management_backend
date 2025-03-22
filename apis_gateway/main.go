package main

import (
	"strings"

	"github.com/IBM/sarama"
	"github.com/spf13/viper"
)

var (
	servers string
	topic   string
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func produceMessageToKafka(messages string) {
	consumer, err := sarama.NewConsumerGroup(viper.GetStringSlice("kafka.servers"), viper.GetString("kafka.group"), nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

}

func main() {

}

/* func main() {
	// Slice of server addresses
	servers := []string{"localhost:9092"}
	consuer, err := sarama.NewConsumer(servers, nil)
	if err != nil {
		panic(err)
	}
	defer consuer.Close()

} */
