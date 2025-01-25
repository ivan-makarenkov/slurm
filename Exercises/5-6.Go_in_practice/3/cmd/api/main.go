package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"gocloud.dev/pubsub"

	// Import the pubsub driver packages we want to be able to open.
	_ "gocloud.dev/pubsub/awssnssqs"
	_ "gocloud.dev/pubsub/azuresb"
	_ "gocloud.dev/pubsub/gcppubsub"
	_ "gocloud.dev/pubsub/kafkapubsub"
	_ "gocloud.dev/pubsub/natspubsub"
	_ "gocloud.dev/pubsub/rabbitpubsub"
)

var (
	t             *pubsub.Topic
	rc            *redis.Client
	queueTopic    string
	redisHostport string
)

// RABBIT_SERVER_URL=amqp://guest:guest@localhost:5672/

func init() {
	flag.StringVar(&queueTopic, "t", "rabbit://rates", "rabbit exchange")
	flag.StringVar(&redisHostport, "r", "localhost:6379", "redis host:port")
	flag.Parse()

	u, err := url.Parse(queueTopic)
	if err != nil {
		panic(err)
	}

	if u.Scheme != "rabbit" {
		return
	}

	if err := setupRabbit(u.Host); err != nil {
		panic(err)
	}
}

func PostRateHandler(w http.ResponseWriter, r *http.Request) {
	rate := r.FormValue("rate")

	if _, err := strconv.Atoi(rate); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := topic().Send(context.Background(), &pubsub.Message{
		Body: []byte(rate),
	})
	if err != nil {
		log.Println("[ERROR]", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetTotalHandler(w http.ResponseWriter, r *http.Request) {
	rates, err := storage().LRange(context.TODO(), "result", 0, 10).Result()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(rates) == 0 {
		_, _ = w.Write([]byte(fmt.Sprintf("%.2f", 0.0)))
		return
	}

	var sum int
	for _, rate := range rates {
		v, err := strconv.Atoi(rate)
		if err != nil {
			continue
		}
		sum += v
	}

	result := float64(sum) / float64(len(rates))
	_, _ = w.Write([]byte(fmt.Sprintf("%.2f", result)))
}

func topic() *pubsub.Topic {
	if t != nil {
		return t
	}

	var err error
	t, err = pubsub.OpenTopic(context.Background(), queueTopic)
	if err != nil {
		panic(err)
	}

	return t
}

func storage() *redis.Client {
	if rc != nil {
		return rc
	}

	rc = redis.NewClient(&redis.Options{
		Addr: redisHostport,
	})

	err := rc.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}

	return rc
}

func setupRabbit(qName string) error {
	cn, err := amqp.Dial(os.Getenv("RABBIT_SERVER_URL"))
	if err != nil {
		return err
	}
	defer cn.Close()

	ch, err := cn.Channel()
	if err != nil {
		return fmt.Errorf("channel: %w", err)
	}

	_, err = ch.QueueDeclare(qName, true, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("queue declare: %w", err)
	}

	err = ch.ExchangeDeclare(qName, "topic", false, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("exchange declare: %w", err)
	}

	err = ch.QueueBind(qName, "", qName, false, nil)
	if err != nil {
		return fmt.Errorf("queue bind: %w", err)
	}

	return nil
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/rate", PostRateHandler).Methods(http.MethodPost)
	r.HandleFunc("/total", GetTotalHandler).Methods(http.MethodGet)

	log.Println("start http server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("error on listen and serve: %v", err)
	}
}
