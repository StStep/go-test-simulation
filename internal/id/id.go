package id

type Uid uint64
type Eid uint64

type UidGen interface {
	Reset()
	Id() Uid
}

type EidGen interface {
	Reset()
	Id() Eid
}

func NewUidGen() UidGen {
	return &uidgen{}
}

func NewEidGen() EidGen {
	return &eidgen{}
}

type uidgen struct {
	lastid Uid
}

func (gen *uidgen) Reset() {
	gen.lastid = 0
}

func (gen *uidgen) Id() Uid {
	gen.lastid = gen.lastid + 1
	return gen.lastid
}

type eidgen struct {
	lastid Eid
}

func (gen *eidgen) Reset() {
	gen.lastid = 0
}

func (gen *eidgen) Id() Eid {
	gen.lastid = gen.lastid + 1
	return gen.lastid
}
