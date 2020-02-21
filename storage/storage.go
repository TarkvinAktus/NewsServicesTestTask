package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
	"time"

	pb "github.com/TarkvinAktus/NewsServicesTestTask/messages"

	"github.com/go-redis/redis"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/go-nats"
	"gopkg.in/yaml.v2"
)

type configFile struct {
	ListenPort         string `yaml:"listen_port"`
	NatsStorageSubject string `yaml:"nats_storage_subject"`
	RedisAddr          string `yaml:"redis_addr"`
	RedisPass          string `yaml:"redis_pass"`
	RedisDB            int    `yaml:"redis_db"`
}

func getConfig() (configFile, error) {
	var conf configFile

	filename, _ := filepath.Abs("./storage_config.yaml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return conf, err
	}

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Println(err)
		return conf, err
	}

	return conf, nil
}

func getDataFromRedis(conf configFile, key string) string {
	rClient := redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddr,
		Password: conf.RedisPass,
		DB:       conf.RedisDB,
	})

	//Imitation of real data set
	rClient.Set(key, "some news text", 20*time.Millisecond)

	//Get data from redis by title
	redisResp := rClient.Get(key)

	return redisResp.Val()
}

func main() {
	// Get configs
	conf, err := getConfig()
	if err != nil {
		panic(err)
	}

	// Create NATS connection
	natsConnection, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Println(err)
	}

	// Subscribe to subject
	natsConnection.Subscribe(conf.NatsStorageSubject, func(m *nats.Msg) {
		log.Printf("Received a message: %s\n", string(m.Data))

		//Create protobuf struct and send it with data from db
		news := pb.News{
			Name: string(m.Data),
			Data: getDataFromRedis(conf, string(m.Data)),
		}

		data, err := proto.Marshal(&news)
		if err == nil {
			natsConnection.Publish(m.Reply, data)
		} else {
			log.Println("Proto marshal err")
		}
	})
	runtime.Goexit()
}
