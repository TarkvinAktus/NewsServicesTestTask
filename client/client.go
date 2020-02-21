package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"time"

	pb "../messages"

	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats"
	"gopkg.in/yaml.v2"
)

type configFile struct {
	ListenPort         string `yaml:"listen_port"`
	NatsStorageSubject string `yaml:"nats_storage_subject"`
	NatsURL            string `yaml:"nats_url"`
}

func getConfig() (configFile, error) {
	var conf configFile

	filename, _ := filepath.Abs("./client_config.yaml")
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

func mainHandler(w http.ResponseWriter, r *http.Request, conf configFile) {

	//Get http param
	key, ok := r.URL.Query()["title"]

	if !ok || len(key[0]) < 1 {
		log.Println("Url Param 'title' is missing")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Url Param 'title' is missing", http.StatusBadRequest)
		return
	}

	// Create NATS connection
	nc, err := nats.Connect(conf.NatsURL)
	if err != nil {
		log.Println("Nats connect err - " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer nc.Close()

	// NATS request
	resp, err := nc.Request(conf.NatsStorageSubject, []byte(key[0]), 1000*time.Millisecond)
	if err == nil && resp != nil {
		protoMsg := pb.News{}
		err := proto.Unmarshal(resp.Data, &protoMsg)
		if err != nil {
			log.Fatalf("Error on unmarshal: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		if protoMsg.Data == "" {
			//404 if no data by this key
			w.WriteHeader(http.StatusNotFound)
		}

		result := "Title - " + protoMsg.Name + " With text - " + protoMsg.Data

		w.Write([]byte(result))
	} else if resp == nil {
		//404 if no response by this key
		w.WriteHeader(http.StatusNotFound)
	} else {
		log.Println("nats request err - " + err.Error())
		//500 if no response from nats
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	// Get configs
	conf, err := getConfig()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mainHandler(w, r, conf)
	})

	http.ListenAndServe(conf.ListenPort, nil)
}
