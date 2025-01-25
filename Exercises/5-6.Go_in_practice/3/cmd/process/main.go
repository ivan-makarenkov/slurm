package main

import (
	"context"
	"flag"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	rc  *redis.Client
	sub *pubsub.Subscription
)

func main() {
	queueTopic := flag.String("t", "rabbit://rates", "rabbit exchange")
	redisHostport := flag.String("r", "localhost:6379", "redis host:port")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM)
	defer stop()

	log.Println("start consume", *queueTopic)

	for {
		if ctx.Err() != nil {
			break
		}

		s, err := subscription(*queueTopic)
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second)
			continue
		}
		msg, err := s.Receive(ctx)
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second)
			continue
		}

		err = storage(*redisHostport).LPush(ctx, "result", string(msg.Body)).Err()
		if err != nil {
			log.Println(err)
		}

		if rand.Float64() < .05 { // периодическая очистка бд
			_ = storage(*redisHostport).LTrim(ctx, "result", 0, 9)
		}

		msg.Ack()
	}
}

func subscription(topic string) (*pubsub.Subscription, error) {
	if sub != nil {
		return sub, nil
	}
	var err error
	sub, err = pubsub.OpenSubscription(context.Background(), topic)
	if err != nil {
		return nil, err
	}
	return sub, nil
}

func storage(redisurl string) *redis.Client {
	if rc != nil {
		return rc
	}

	rc = redis.NewClient(&redis.Options{
		Addr: redisurl,
	})

	err := rc.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}

	return rc
}
