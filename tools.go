package main

import (
	"protoc-gen-verifier/verifier"

	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func DefaultValueOfKind(field *protogen.Field) interface{} {
	kind := field.Desc.Kind()
	logrus.Debugf("kind=%v", kind)
	couldBeMessage := false
	switch kind {
	case protoreflect.BoolKind:
		return false
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return int32(0)
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return int64(0)
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return uint32(0)
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return uint64(0)
	case protoreflect.FloatKind:
		return float32(0.0)
	case protoreflect.DoubleKind:
		return float64(0.0)
	case protoreflect.StringKind:
		return ""
	case protoreflect.BytesKind:
		return []byte("")
	case protoreflect.EnumKind:
		return protoreflect.EnumNumber(0)
	case protoreflect.MessageKind, protoreflect.GroupKind:
		couldBeMessage = true
	default:
		return nil
	}

	// 针对Message类型需要特殊处理
	if couldBeMessage {
		if field.Desc.IsList() {
			// 该字段为repeated
			return make([]protoreflect.MessageDescriptor, 0)
		} else if field.Desc.IsMap() {
			// 该字段为map
			return make(map[protoreflect.FieldDescriptor]protoreflect.FieldDescriptor, 0)
		} else {
			// TODO 其它情况 现在暂时忽略
			return verifier.Skipper{}
		}
	}

	return nil
}

func continueParsing() bool {
	return *mode == LooseMode
}