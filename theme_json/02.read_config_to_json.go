package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type AuditAgentConf struct {
	Instance []struct {
		CdbKafkaEndpoint        string `json:"cdb_kafka_endpoint"`
		CdbKafkaName            string `json:"cdb_kafka_name"`
		CdbKafkaPartition       int    `json:"cdb_kafka_partition"`
		CustomerKafkaNames      string `json:"customer_kafka_names"`
		CustomerKafkaPartitions string `json:"customer_kafka_partitions"`
		CustomerKafkaTopics     string `json:"customer_kafka_topics"`
		Entrance                string `json:"entrance"`
		EsHost                  string `json:"es_host"`
		InstanceID              string `json:"instance_id"`
		LogPath                 string `json:"log_path"`
	} `json:"instance"`
}

func DealErr(e error) {
	if e != nil {
		fmt.Println("err: ", e.Error())
		panic(e)
	}
}

func main() {
	filePath := os.Args[1]
	fmt.Printf("The file path is :%s\n", filePath)

	fileData, err := ioutil.ReadFile(filePath)
	DealErr(err)
	fmt.Println(string(fileData))

	fmt.Println("bufio read---------------------------->")
	f, err := os.Open(filePath)
	defer f.Close()
	DealErr(err)
	bio := bufio.NewReader(f)

	bfRead, isPrefix, err := bio.ReadLine()
	DealErr(err)
	fmt.Printf("This mess is  [ %q ] [%v]\n", bfRead, isPrefix)

	res := &AuditAgentConf{}
	//json.Unmarshal([]byte(bfRead), &res)
	json.Unmarshal(fileData, &res)
	fmt.Fprintf(os.Stdout, "struct: %+v\n", res)

	fmt.Println("size = ", len(res.Instance))
	fmt.Println("ckafkanmae: ", res.Instance[0].CdbKafkaName)

}