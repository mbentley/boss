// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: worker.proto

/*
	Package moby_buildkit_v1_types is a generated protocol buffer package.

	It is generated from these files:
		worker.proto

	It has these top-level messages:
		WorkerRecord
*/
package moby_buildkit_v1_types

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import pb "github.com/moby/buildkit/solver/pb"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type WorkerRecord struct {
	ID        string            `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Labels    map[string]string `protobuf:"bytes,2,rep,name=Labels" json:"Labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Platforms []pb.Platform     `protobuf:"bytes,3,rep,name=platforms" json:"platforms"`
}

func (m *WorkerRecord) Reset()                    { *m = WorkerRecord{} }
func (m *WorkerRecord) String() string            { return proto.CompactTextString(m) }
func (*WorkerRecord) ProtoMessage()               {}
func (*WorkerRecord) Descriptor() ([]byte, []int) { return fileDescriptorWorker, []int{0} }

func (m *WorkerRecord) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *WorkerRecord) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *WorkerRecord) GetPlatforms() []pb.Platform {
	if m != nil {
		return m.Platforms
	}
	return nil
}

func init() {
	proto.RegisterType((*WorkerRecord)(nil), "moby.buildkit.v1.types.WorkerRecord")
}
func (m *WorkerRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WorkerRecord) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintWorker(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if len(m.Labels) > 0 {
		for k, _ := range m.Labels {
			dAtA[i] = 0x12
			i++
			v := m.Labels[k]
			mapSize := 1 + len(k) + sovWorker(uint64(len(k))) + 1 + len(v) + sovWorker(uint64(len(v)))
			i = encodeVarintWorker(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintWorker(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintWorker(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if len(m.Platforms) > 0 {
		for _, msg := range m.Platforms {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintWorker(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintWorker(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *WorkerRecord) Size() (n int) {
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovWorker(uint64(l))
	}
	if len(m.Labels) > 0 {
		for k, v := range m.Labels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovWorker(uint64(len(k))) + 1 + len(v) + sovWorker(uint64(len(v)))
			n += mapEntrySize + 1 + sovWorker(uint64(mapEntrySize))
		}
	}
	if len(m.Platforms) > 0 {
		for _, e := range m.Platforms {
			l = e.Size()
			n += 1 + l + sovWorker(uint64(l))
		}
	}
	return n
}

func sovWorker(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozWorker(x uint64) (n int) {
	return sovWorker(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WorkerRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorker
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WorkerRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WorkerRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWorker
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWorker
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Labels == nil {
				m.Labels = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowWorker
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowWorker
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthWorker
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowWorker
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthWorker
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipWorker(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthWorker
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Labels[mapkey] = mapvalue
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Platforms", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWorker
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Platforms = append(m.Platforms, pb.Platform{})
			if err := m.Platforms[len(m.Platforms)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWorker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWorker
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipWorker(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWorker
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWorker
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWorker
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthWorker
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowWorker
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipWorker(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthWorker = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWorker   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("worker.proto", fileDescriptorWorker) }

var fileDescriptorWorker = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x86, 0xbf, 0x4d, 0x3e, 0x0b, 0xdd, 0x06, 0x91, 0x45, 0x24, 0xe4, 0x10, 0x8b, 0xa7, 0x1e,
	0x74, 0xb6, 0xea, 0x45, 0x3d, 0x96, 0x0a, 0x16, 0x3c, 0x48, 0x2e, 0x9e, 0xb3, 0xed, 0x36, 0x86,
	0x24, 0xce, 0xb2, 0xd9, 0x44, 0xf2, 0x0f, 0x7b, 0xf4, 0xe2, 0x55, 0x24, 0xbf, 0x44, 0xba, 0x89,
	0x98, 0x83, 0xb7, 0x79, 0x87, 0x67, 0x1e, 0xde, 0xa1, 0xde, 0x1b, 0xea, 0x4c, 0x6a, 0x50, 0x1a,
	0x0d, 0xb2, 0x93, 0x02, 0x45, 0x03, 0xa2, 0x4a, 0xf3, 0x4d, 0x96, 0x1a, 0xa8, 0x2f, 0xc1, 0x34,
	0x4a, 0x96, 0xc1, 0x45, 0x92, 0x9a, 0x97, 0x4a, 0xc0, 0x1a, 0x0b, 0x9e, 0x60, 0x82, 0xdc, 0xe2,
	0xa2, 0xda, 0xda, 0x64, 0x83, 0x9d, 0x3a, 0x4d, 0x70, 0x3e, 0xc0, 0xf7, 0x46, 0xfe, 0x63, 0xe4,
	0x25, 0xe6, 0xb5, 0xd4, 0x5c, 0x09, 0x8e, 0xaa, 0xec, 0xe8, 0xb3, 0x0f, 0x42, 0xbd, 0x67, 0xdb,
	0x22, 0x92, 0x6b, 0xd4, 0x1b, 0x76, 0x48, 0x9d, 0xd5, 0xd2, 0x27, 0x53, 0x32, 0x1b, 0x47, 0xce,
	0x6a, 0xc9, 0x1e, 0xe8, 0xe8, 0x31, 0x16, 0x32, 0x2f, 0x7d, 0x67, 0xea, 0xce, 0x26, 0x57, 0x73,
	0xf8, 0xbb, 0x26, 0x0c, 0x2d, 0xd0, 0x9d, 0xdc, 0xbf, 0x1a, 0xdd, 0x44, 0xfd, 0x3d, 0x9b, 0xd3,
	0xb1, 0xca, 0x63, 0xb3, 0x45, 0x5d, 0x94, 0xbe, 0x6b, 0x65, 0x1e, 0x28, 0x01, 0x4f, 0xfd, 0x72,
	0xf1, 0x7f, 0xf7, 0x79, 0xfa, 0x2f, 0xfa, 0x85, 0x82, 0x5b, 0x3a, 0x19, 0x88, 0xd8, 0x11, 0x75,
	0x33, 0xd9, 0xf4, 0xdd, 0xf6, 0x23, 0x3b, 0xa6, 0x07, 0x75, 0x9c, 0x57, 0xd2, 0x77, 0xec, 0xae,
	0x0b, 0x77, 0xce, 0x0d, 0x59, 0x78, 0xbb, 0x36, 0x24, 0xef, 0x6d, 0x48, 0xbe, 0xda, 0x90, 0x88,
	0x91, 0x7d, 0xf6, 0xfa, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x5c, 0x8f, 0x26, 0x71, 0x01, 0x00,
	0x00,
}
