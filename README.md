# GolangMachinery
Basic application for sending the protobuf type in rabbit mq using protobuf and machinery package

Steps:[ubuntu 16.06]

install go 1.12.0

install protobuf compiler using tar.gz file 
curl -OL https://github.com/google/protobuf/releases/download/v3.2.0/protoc-3.2.0-linux-x86_64.zip
unzip protoc-3.2.0-linux-x86_64.zip -d protoc3
sudo mv protoc3/bin/* /usr/local/bin/
sudo mv protoc3/include/* /usr/local/include/
export PATH=$PATH:$GOPATH/bin
protc --version



go get github.com/RichardKnop/machinery/v1 for machinery package
go get github.com/golang/protobuf for proto buf package
go get github.com/golang/protobuf/protoc-gen-go for proto file generation package 

commnads to check for the file space 
sudo stat person.xml
sudo stat person.pb.go

command to generate *.pb.go from *.proto 
protoc --go_out=. *.proto

sending the file as argument command
go run main.go person.pb.go

Install rabbit Mq:




Starting the application 

Go to MachineryGolangApp/receiver 
call$ go run main.go
INFO: 2019/04/03 16:37:38 env.go:17 Successfully loaded config from the environment
INFO: 2019/04/03 16:37:38 worker.go:46 Launching a worker with the following settings:
INFO: 2019/04/03 16:37:38 worker.go:47 - Broker: amqp://guest:guest@localhost:5672/
INFO: 2019/04/03 16:37:38 worker.go:49 - DefaultQueue: machinery_tasks
INFO: 2019/04/03 16:37:38 worker.go:53 - ResultBackend: amqp://guest:guest@localhost:5672/
INFO: 2019/04/03 16:37:38 worker.go:55 - AMQP: machinery_exchange
INFO: 2019/04/03 16:37:38 worker.go:56   - Exchange: machinery_exchange
INFO: 2019/04/03 16:37:38 worker.go:57   - ExchangeType: direct
INFO: 2019/04/03 16:37:38 worker.go:58   - BindingKey: machinery_task
INFO: 2019/04/03 16:37:38 worker.go:59   - PrefetchCount: 3
INFO: 2019/04/03 16:37:38 amqp.go:95 [*] Waiting for messages. To exit press CTRL+C


Go to MachineryGolangApp/sender
call$ go run send.go person.pb.go            

INFO: 2019/04/03 18:42:12 env.go:17 Successfully loaded config from the environment
INFO: 2019/04/03 18:42:12 send.go:70  Starting batch: 373a2def-fd63-4550-967d-66eeca16da35 
INFO: 2019/04/03 18:42:12 send.go:76  Single task: 
INFO: 2019/04/03 18:42:13 send.go:87 1 + 1 = []

             

After calling the sender, check it in the receiver

INFO: 2019/04/03 16:38:21 amqp.go:316 Received new message: {"UUID":"task_7ba2031e-0865-4142-ab58-386d69ba569b","Name":"de_serialize","RoutingKey":"machinery_task","ETA":null,"GroupUUID":"","GroupTaskCount":0,"Args":[{"Name":"","Type":"string","Value":"\n\u0008John Doe\u0010\u001b"}],"Headers":{},"Immutable":false,"RetryCount":0,"RetryTimeout":0,"OnSuccess":null,"OnError":null,"ChordCallback":null,"BrokerMessageGroupId":""}
INFO: 2019/04/03 16:38:22 main.go:32  I am a start of task handler for: de_serialize 
INFO: 2019/04/03 16:38:22 tasks.go:10  Deserialization has started 
INFO: 2019/04/03 16:38:22 tasks.go:16  name:"John Doe" age:27  
DEBUG: 2019/04/03 16:38:22 worker.go:248 Processed task task_7ba2031e-0865-4142-ab58-386d69ba569b. Results = []
INFO: 2019/04/03 16:38:22 main.go:36  I am an end of task handler for: de_serialize


