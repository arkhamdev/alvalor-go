// Code generated by capnpc-go. DO NOT EDIT.

package codec

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type Transaction struct{ capnp.Struct }

// Transaction_TypeID is the unique identifier for the type Transaction.
const Transaction_TypeID = 0xec9fd906d129035f

func NewTransaction(s *capnp.Segment) (Transaction, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	return Transaction{st}, err
}

func NewRootTransaction(s *capnp.Segment) (Transaction, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	return Transaction{st}, err
}

func ReadRootTransaction(msg *capnp.Message) (Transaction, error) {
	root, err := msg.RootPtr()
	return Transaction{root.Struct()}, err
}

func (s Transaction) String() string {
	str, _ := text.Marshal(0xec9fd906d129035f, s.Struct)
	return str
}

func (s Transaction) Transfers() (Transfer_List, error) {
	p, err := s.Struct.Ptr(0)
	return Transfer_List{List: p.List()}, err
}

func (s Transaction) HasTransfers() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Transaction) SetTransfers(v Transfer_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewTransfers sets the transfers field to a newly
// allocated Transfer_List, preferring placement in s's segment.
func (s Transaction) NewTransfers(n int32) (Transfer_List, error) {
	l, err := NewTransfer_List(s.Struct.Segment(), n)
	if err != nil {
		return Transfer_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Transaction) Fees() (Fee_List, error) {
	p, err := s.Struct.Ptr(1)
	return Fee_List{List: p.List()}, err
}

func (s Transaction) HasFees() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Transaction) SetFees(v Fee_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewFees sets the fees field to a newly
// allocated Fee_List, preferring placement in s's segment.
func (s Transaction) NewFees(n int32) (Fee_List, error) {
	l, err := NewFee_List(s.Struct.Segment(), n)
	if err != nil {
		return Fee_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s Transaction) Data() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return []byte(p.Data()), err
}

func (s Transaction) HasData() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Transaction) SetData(v []byte) error {
	return s.Struct.SetData(2, v)
}

func (s Transaction) Signatures() (capnp.DataList, error) {
	p, err := s.Struct.Ptr(3)
	return capnp.DataList{List: p.List()}, err
}

func (s Transaction) HasSignatures() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Transaction) SetSignatures(v capnp.DataList) error {
	return s.Struct.SetPtr(3, v.List.ToPtr())
}

// NewSignatures sets the signatures field to a newly
// allocated capnp.DataList, preferring placement in s's segment.
func (s Transaction) NewSignatures(n int32) (capnp.DataList, error) {
	l, err := capnp.NewDataList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.DataList{}, err
	}
	err = s.Struct.SetPtr(3, l.List.ToPtr())
	return l, err
}

// Transaction_List is a list of Transaction.
type Transaction_List struct{ capnp.List }

// NewTransaction creates a new list of Transaction.
func NewTransaction_List(s *capnp.Segment, sz int32) (Transaction_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4}, sz)
	return Transaction_List{l}, err
}

func (s Transaction_List) At(i int) Transaction { return Transaction{s.List.Struct(i)} }

func (s Transaction_List) Set(i int, v Transaction) error { return s.List.SetStruct(i, v.Struct) }

func (s Transaction_List) String() string {
	str, _ := text.MarshalList(0xec9fd906d129035f, s.List)
	return str
}

// Transaction_Promise is a wrapper for a Transaction promised by a client call.
type Transaction_Promise struct{ *capnp.Pipeline }

func (p Transaction_Promise) Struct() (Transaction, error) {
	s, err := p.Pipeline.Struct()
	return Transaction{s}, err
}

const schema_b5f3d18a6c743283 = "x\xdaL\xcd\xb1K\xfbP\x14\xc5\xf1s\xdeM\xf2\x9b" +
	"\xda_\x9e\xcd\x1a\x9c\x1c\x14\x14\x15\x1c\xec\xe4\xa8\xe0\xd0" +
	"'N.\xfa\x88\xa9\x04$-\xc9s(\xb8\xb9\xa8 " +
	"Vpw\x0e\xb8::;9t\xf4/\x10\xc1Iw" +
	"#QD\xb7\xcb\xe7\xcb\xe5,\xcepM-\xf9\x97\x01" +
	"`6\xfd\xa0\xde\x95\xd9I\xf0t\xf3\x0a=\xc5\xfad" +
	"\xd9\x1d\x9eO\xde\xef\xe0{\xff\x80\xce\xbc\xbcuV\xa5" +
	"\xb9V\xe4\x19\x0f\xb5+l^\xda\xc4\xa9l\x90/$" +
	"v\x98\x0f\xbb\xdb_4\x9d\xb8l\x90\xf7H\x13\x8a\x07" +
	"x\x04\xb4\xdd\x02\xcc\x9e\xd0\x1c+j2b\x83\xa39" +
	"\xc08\xa1\x19+j\xa5\"*@_4x*4\xd7" +
	"\x8aZ$\xa2\x00\xfaj\x070c\xa1\xb9U\xfc\x1e\xee" +
	"\xa7\x05X\xb2\x0d\xf6\x84\x0c\xeb*>[\xefn\xdc\x7f" +
	"\x00l\xf0\x7f?M\xff\xd4Q\xf8\x12Wq\xf5\xf8S" +
	"\xf7\xad\xb3lA\xb1\x05\xd6ev\x90[wT@~" +
	"_\x9a\xd6\x06?\x03\x00\x00\xff\xff\x9a\x95A\x08"

func init() {
	schemas.Register(schema_b5f3d18a6c743283,
		0xec9fd906d129035f)
}
