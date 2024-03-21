package Golang_Redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

var client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB:   0,
})

func TestConnection(t *testing.T) {
	assert.NotNil(t, client)
	//err := client.Close()
	//assert.Nil(t, err)
}

var ctx = context.Background()

func TestPing(t *testing.T) {
	result, err := client.Ping(ctx).Result()
	assert.Nil(t, err)

	assert.Equal(t, "PONG", result)
}

func TestString(t *testing.T) {
	client.SetEx(ctx, "name", "Hanif Al Irsyad", 3*time.Second)
	result, err := client.Get(ctx, "name").Result()
	assert.Nil(t, err)
	assert.Equal(t, "Hanif Al Irsyad", result)

	time.Sleep(5 * time.Second)

	result, err = client.Get(ctx, "name").Result()
	assert.NotNil(t, err)
}

func TestList(t *testing.T) {
	client.RPush(ctx, "names", "Hanif")
	client.RPush(ctx, "names", "Al")
	client.RPush(ctx, "names", "Irsyad")

	assert.Equal(t, "Hanif", client.LPop(ctx, "names").Val())
	assert.Equal(t, "Al", client.LPop(ctx, "names").Val())
	assert.Equal(t, "Irsyad", client.LPop(ctx, "names").Val())

	client.Del(ctx, "names")
}

func TestSet(t *testing.T) {
	client.SAdd(ctx, "students", "Hanif")
	client.SAdd(ctx, "students", "Al")
	client.SAdd(ctx, "students", "Irsyad")

	assert.Equal(t, int64(3), client.SCard(ctx, "students").Val())
	assert.Equal(t, []string{"Irsyad", "Al", "Hanif"}, client.SMembers(ctx, "students").Val())
}

func TestSortedSet(t *testing.T) {
	client.ZAdd(ctx, "scores", redis.Z{Score: 100, Member: "Hanif"})
	client.ZAdd(ctx, "scores", redis.Z{Score: 85, Member: "Al"})
	client.ZAdd(ctx, "scores", redis.Z{Score: 95, Member: "Irsyad"})

	assert.Equal(t, []string{"Al", "Irsyad", "Hanif"}, client.ZRange(ctx, "scores", 0, 2).Val())

	assert.Equal(t, "Hanif", client.ZPopMax(ctx, "scores").Val()[0].Member)
	assert.Equal(t, "Irsyad", client.ZPopMax(ctx, "scores").Val()[0].Member)
	assert.Equal(t, "Al", client.ZPopMax(ctx, "scores").Val()[0].Member)
}

func TestHash(t *testing.T) {
	client.HSet(ctx, "user:1", "id", "1")
	client.HSet(ctx, "user:1", "name", "Hanif")
	client.HSet(ctx, "user:1", "email", "hanif@example.com")

	user := client.HGetAll(ctx, "user:1").Val()

	assert.Equal(t, "1", user["id"])
	assert.Equal(t, "Hanif", user["name"])
	assert.Equal(t, "hanif@example.com", user["email"])

	client.Del(ctx, "user:1")
}

func TestGeoPoint(t *testing.T) {
	client.GeoAdd(ctx, "sellers", &redis.GeoLocation{
		Name:      "Toko A",
		Longitude: 110.408589,
		Latitude:  -7.761040,
	})

	client.GeoAdd(ctx, "sellers", &redis.GeoLocation{
		Name:      "Toko B",
		Longitude: 110.408811,
		Latitude:  -7.759897,
	})

	dist := client.GeoDist(ctx, "sellers", "Toko A", "Toko B", "km").Val()

	assert.Equal(t, 0.1296, dist)

	sellers := client.GeoSearch(ctx, "sellers", &redis.GeoSearchQuery{
		Longitude:  110.408666,
		Latitude:   -7.760741,
		Radius:     5,
		RadiusUnit: "km",
	}).Val()

	fmt.Println(sellers)

	assert.Equal(t, []string{"Toko A", "Toko B"}, sellers)
}

func TestHyperLogLog(t *testing.T) {
	client.PFAdd(ctx, "visitors", "hanif", "al", "irsyad")
	client.PFAdd(ctx, "visitors", "hanif", "al", "qodri")
	client.PFAdd(ctx, "visitors", "hanif", "al", "irsyad")

	total := client.PFCount(ctx, "visitors").Val()
	assert.Equal(t, int64(4), total)
}

func TestPipeline(t *testing.T) {
	_, err := client.Pipelined(ctx, func(pipeliner redis.Pipeliner) error {
		pipeliner.SetEx(ctx, "name", "Hanif Al Irsyad", 5*time.Second)
		pipeliner.SetEx(ctx, "address", "Indonesia", 5*time.Second)
		return nil
	})

	assert.Nil(t, err)

	assert.Equal(t, "Hanif Al Irsyad", client.Get(ctx, "name").Val())
	assert.Equal(t, "Indonesia", client.Get(ctx, "address").Val())
}

func TestTransaction(t *testing.T) {
	_, err := client.TxPipelined(ctx, func(pipeliner redis.Pipeliner) error {
		pipeliner.SetEx(ctx, "name", "Hanif Al Irsyad", 5*time.Second)
		pipeliner.SetEx(ctx, "address", "Indonesia", 5*time.Second)
		return nil
	})

	assert.Nil(t, err)

	assert.Equal(t, "Hanif Al Irsyad", client.Get(ctx, "name").Val())
	assert.Equal(t, "Indonesia", client.Get(ctx, "address").Val())
}

func TestPublishStream(t *testing.T) {
	for i := 0; i < 10; i++ {
		err := client.XAdd(ctx, &redis.XAddArgs{
			Stream: "members",
			Values: map[string]interface{}{"name": "Hanif Al Irsyad", "email": "hanif@example.com"},
		}).Err()
		assert.Nil(t, err)
	}
}

func TestCreateConsumerGroup(t *testing.T) {
	client.XGroupCreate(ctx, "members", "group-1", "0")
	client.XGroupCreateConsumer(ctx, "members", "group-1", "consumer-1")
	client.XGroupCreateConsumer(ctx, "members", "group-1", "consumer-2")
}

func TestGetStream(t *testing.T) {
	result := client.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    "group-1",
		Consumer: "consumer-1",
		Streams:  []string{"members", ">"},
		Count:    2,
		Block:    5 * time.Second,
	}).Val()

	for _, stream := range result {
		for _, message := range stream.Messages {
			fmt.Println(message.ID)
			fmt.Println(message.Values)
		}
	}
}

func TestSubscriberPubSub(t *testing.T) {
	subscriber := client.Subscribe(ctx, "channel-1")
	defer subscriber.Close()
	for i := 0; i < 10; i++ {
		messages, err := subscriber.ReceiveMessage(ctx)
		assert.Nil(t, err)
		fmt.Println(messages.Payload)
	}
}

func TestPublisherPubSub(t *testing.T) {
	for i := 0; i < 10; i++ {
		err := client.Publish(ctx, "channel-1", "Hello World!"+strconv.Itoa(i)).Err()
		assert.Nil(t, err)
	}
}
