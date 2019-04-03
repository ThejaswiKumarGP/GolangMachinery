package main

import (
	examplerTasks "MachineryGolangApp/tasks"
	"context"
	fmt "fmt"
	"time"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/log"
	"github.com/RichardKnop/machinery/v1/tasks"
	proto "github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	opentracing_log "github.com/opentracing/opentracing-go/log"
)

func main() {
	if err := send(); err != nil {
		fmt.Errorf("Error in starting send: %s", err.Error())
	}
}

func send() error {

	server, err := startServer()
	if err != nil {
		return err
	}

	var (
		deserialize tasks.Signature
	)

	p := &Person{
		Name: "John Doe",
		Age:  27,
	}

	out, err := proto.Marshal(p)
	if err != nil {
		log.FATAL.Fatalln("Failed to encode address book:", err)
	}

	var initTasks = func() {
		deserialize = tasks.Signature{
			Name: "de_serialize",
			Args: []tasks.Arg{
				{
					Type:  "string",
					Value: string(out),
				},
			},
		}
	}

	/*
	 * Lets start a span representing this run of the `send` command and
	 * set a batch id as baggage so it can travel all the way into
	 * the worker functions.
	 */
	span, ctx := opentracing.StartSpanFromContext(context.Background(), "send")
	defer span.Finish()

	batchID := uuid.New().String()
	span.SetBaggageItem("batch.id", batchID)
	span.LogFields(opentracing_log.String("batch.id", batchID))

	log.INFO.Println("Starting batch:", batchID)
	/*
	 * First, let's try sending a single task
	 */
	initTasks()

	log.INFO.Println("Single task:")

	asyncResult, err := server.SendTaskWithContext(ctx, &deserialize)
	if err != nil {
		return fmt.Errorf("Could not send task: %s", err.Error())
	}

	results, err := asyncResult.Get(time.Duration(time.Millisecond * 5))
	if err != nil {
		return fmt.Errorf("Getting task result failed with error: %s", err.Error())
	}
	log.INFO.Printf("1 + 1 = %v\n", tasks.HumanReadableResults(results))

	return nil
}

func loadConfig() (*config.Config, error) {
	return config.NewFromEnvironment(true)
}

func startServer() (*machinery.Server, error) {
	cnf, err := loadConfig()
	if err != nil {
		return nil, err
	}

	// Create server instance
	server, err := machinery.NewServer(cnf)
	if err != nil {
		return nil, err
	}

	// Register tasks
	tasks := map[string]interface{}{
		"de_serialize": examplerTasks.Deserialize,
	}

	return server, server.RegisterTasks(tasks)
}
