package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"

	examplepb "examples/protos"
)

func getFieldOptions() {
	fmt.Println("field options:")
	msg := &examplepb.FieldOption{}
	_, md := descriptor.MessageDescriptorProto(msg)
	stringOpt, _ := proto.GetExtension(md.Field[0].Options, examplepb.E_FieldOptString)
	objOpt, _ := proto.GetExtension(md.Field[0].Options, examplepb.E_FieldOptObj)
	fmt.Println("	obj.foo_string:", objOpt.(*examplepb.ExtObj).FooString)
	fmt.Println("	obj.bar_int", objOpt.(*examplepb.ExtObj).BarInt)
	fmt.Println("	string:", *stringOpt.(*string))
}

func getMessageOptions() {
	fmt.Println("message options:")
	msg := &examplepb.MessageOption{}
	_, md := descriptor.MessageDescriptorProto(msg)
	objOpt, _ := proto.GetExtension(md.Options, examplepb.E_MsgOptObj)
	stringOpt, _ := proto.GetExtension(md.Options, examplepb.E_MsgOptString)
	fmt.Println("	obj.foo_string:", objOpt.(*examplepb.ExtObj).FooString)
	fmt.Println("	obj.bar_int", objOpt.(*examplepb.ExtObj).BarInt)
	fmt.Println("	string:", *stringOpt.(*string))
}

func getServiceOptions() {
	fmt.Println("service options:")
	msg := &examplepb.MessageOption{}
	md, _ := descriptor.MessageDescriptorProto(msg)
	srv := md.Service[1] // ServiceOption
	stringOpt, _ := proto.GetExtension(srv.Options, examplepb.E_SrvOptString)
	fmt.Println("	string:", *stringOpt.(*string))
}
func getMethodOptions() {
	fmt.Println("method options:")
	msg := &examplepb.MessageOption{}
	md, _ := descriptor.MessageDescriptorProto(msg)
	srv := md.Service[1] // ServiceOption
	objOpt, _ := proto.GetExtension(srv.Method[0].Options, examplepb.E_MethodOptObj)
	stringOpt, _ := proto.GetExtension(srv.Method[0].Options, examplepb.E_MethodOptString)
	fmt.Println("	obj.foo_string:", objOpt.(*examplepb.ExtObj).FooString)
	fmt.Println("	obj.bar_int", objOpt.(*examplepb.ExtObj).BarInt)
	fmt.Println("	string:", *stringOpt.(*string))
}

func getFileOptions() {
	fmt.Println("file options:")
	msg := &examplepb.MessageOption{}
	md, _ := descriptor.MessageDescriptorProto(msg)
	stringOpt, _ := proto.GetExtension(md.Options, examplepb.E_FileOptString)
	objOpt, _ := proto.GetExtension(md.Options, examplepb.E_FileOptObj)
	fmt.Println("	obj.foo_string:", objOpt.(*examplepb.ExtObj).FooString)
	fmt.Println("	obj.bar_int", objOpt.(*examplepb.ExtObj).BarInt)
	fmt.Println("	string:", *stringOpt.(*string))
}

func getSetAny() {
	fmt.Println("getSetAny")
	req := &examplepb.SearchRequest{
		Query: "query",
	}
	// 将SearchRequest打包成Any类型
	a, err := ptypes.MarshalAny(req)
	if err != nil {
		log.Println(err)
		return
	}
	// 赋值
	anyMsg := &examplepb.AnyMessage{
		Message: "any message",
		Details: a,
	}

	req = &examplepb.SearchRequest{}
	// 从Any类型中还原proto消息
	err = ptypes.UnmarshalAny(anyMsg.Details, req)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("	any:", req)
}

func getSetOneof() {
	fmt.Println("getSetOneof")
	oneof := &examplepb.OneOfMessage{
		// 同一时间只能设值一个值
		TestOneof: &examplepb.OneOfMessage_M1{
			M1: "this is m1",
		},
	}
	fmt.Println("	m1:", oneof.GetM1())
	fmt.Println("	m2:", oneof.GetM2())
}

func main() {
	getSetAny()
	getSetOneof()
	getFieldOptions()
	getMessageOptions()
	getServiceOptions()
	getFileOptions()
}
