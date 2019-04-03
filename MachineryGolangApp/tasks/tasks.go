package tasks

import (
	"github.com/RichardKnop/machinery/v1/log"
	"github.com/golang/protobuf/proto"
)

//Deserialize the data
func Deserialize(m string) error {
	log.INFO.Println("Deserialization has started")
	newElliot := &Person{}
	err := proto.Unmarshal([]byte(m), newElliot)
	if err != nil {
		log.ERROR.Println(" unmarshalling error: ", err)
	}
	log.INFO.Println(newElliot)
	return nil
}
