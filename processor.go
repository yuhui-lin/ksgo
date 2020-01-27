package ksgo

// type Processor interface {
// 	process(key string, val interface{})
// }

type Processor func(ctx ProcessorContext, val interface{})

type ProcessorContext interface {
	Topic() string
	Partition() int32
	Key() interface{}
	Offset() int64
	Forward(key string, val interface{})
	GetStateStore(name string) StateStore
}
