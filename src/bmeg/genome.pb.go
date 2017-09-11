// Code generated by protoc-gen-go. DO NOT EDIT.
// source: genome.proto

/*
Package bmeg is a generated protocol buffer package.

It is generated from these files:
	genome.proto

It has these top-level messages:
	Position
	Gene
	Transcript
	GeneSynonym
	GeneDatabase
	GeneFamily
*/
package bmeg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A `Position` is one or more (consecutive***) bases in some `Reference`. A `Position` is
// represented by a `Reference` name, and a base number on that `Reference`
// (0-based).
// ***Modified from GA4GH to include end position.***
type Position struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	// The name of the `Reference` on which the `Position` is located.
	// (e.g. `chr20` or `X`)
	Chromosome string `protobuf:"bytes,3,opt,name=chromosome" json:"chromosome,omitempty"`
	// The 0-based offset from the start of the forward strand for that
	// `Reference`. Genomic positions are non-negative integers less than
	// `Reference` length.
	//
	// For Variant Calls: The start position at which this variant occurs (0-based).
	// This corresponds to the first base of the string of reference bases.
	// Genomic positions are non-negative integers less than reference length.
	// Variants spanning the join of circular genomes are represented as
	// two variants one on each side of the join (position 0).
	Start int64 `protobuf:"varint,4,opt,name=start" json:"start,omitempty"`
	// The end position (exclusive), resulting in [start, end) closed-open
	// interval.
	//
	// For Variantcalls: This is typically calculated by `start + referenceBases.length`.
	End int64 `protobuf:"varint,5,opt,name=end" json:"end,omitempty"`
	// Strand the position is associated with.
	Strand string `protobuf:"bytes,6,opt,name=strand" json:"strand,omitempty"`
}

func (m *Position) Reset()                    { *m = Position{} }
func (m *Position) String() string            { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()               {}
func (*Position) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Position) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Position) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Position) GetChromosome() string {
	if m != nil {
		return m.Chromosome
	}
	return ""
}

func (m *Position) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Position) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *Position) GetStrand() string {
	if m != nil {
		return m.Strand
	}
	return ""
}

// Annotation of a contiguous region of a sequence.
// An example might be a gene or a protein.
// For example, a position X might corresponds to gene/Feature TP53.
type Gene struct {
	Id          string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Symbol      string `protobuf:"bytes,2,opt,name=symbol" json:"symbol,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	SeqId       string `protobuf:"bytes,4,opt,name=seqId" json:"seqId,omitempty"`
	Start       int32  `protobuf:"varint,5,opt,name=start" json:"start,omitempty"`
	End         int32  `protobuf:"varint,6,opt,name=end" json:"end,omitempty"`
	Strand      string `protobuf:"bytes,7,opt,name=strand" json:"strand,omitempty"`
	Accession   string `protobuf:"bytes,8,opt,name=accession" json:"accession,omitempty"`
	Refseq      string `protobuf:"bytes,9,opt,name=refseq" json:"refseq,omitempty"`
	// Name/value attributes of the annotation.  Attribute names follow the GFF3
	// naming convention of reserved names starting with an upper cases
	// character, and user-define names start with lower-case.  Most GFF3
	// pre-defined attributes apply, the exceptions are ID and Parent, which are
	// defined as fields. Additional, the following attributes are added:
	// * Score - the GFF3 score column
	// * Phase - the GFF3 phase column for CDS features.
	Info map[string]string `protobuf:"bytes,11,rep,name=info" json:"info,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Gene) Reset()                    { *m = Gene{} }
func (m *Gene) String() string            { return proto.CompactTextString(m) }
func (*Gene) ProtoMessage()               {}
func (*Gene) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Gene) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Gene) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Gene) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Gene) GetSeqId() string {
	if m != nil {
		return m.SeqId
	}
	return ""
}

func (m *Gene) GetStart() int32 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Gene) GetEnd() int32 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *Gene) GetStrand() string {
	if m != nil {
		return m.Strand
	}
	return ""
}

func (m *Gene) GetAccession() string {
	if m != nil {
		return m.Accession
	}
	return ""
}

func (m *Gene) GetRefseq() string {
	if m != nil {
		return m.Refseq
	}
	return ""
}

func (m *Gene) GetInfo() map[string]string {
	if m != nil {
		return m.Info
	}
	return nil
}

type Transcript struct {
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Parent string `protobuf:"bytes,2,opt,name=parent" json:"parent,omitempty"`
	Symbol string `protobuf:"bytes,3,opt,name=symbol" json:"symbol,omitempty"`
	SeqId  string `protobuf:"bytes,4,opt,name=seqId" json:"seqId,omitempty"`
	Start  int32  `protobuf:"varint,5,opt,name=start" json:"start,omitempty"`
	End    int32  `protobuf:"varint,6,opt,name=end" json:"end,omitempty"`
	Strand string `protobuf:"bytes,7,opt,name=strand" json:"strand,omitempty"`
}

func (m *Transcript) Reset()                    { *m = Transcript{} }
func (m *Transcript) String() string            { return proto.CompactTextString(m) }
func (*Transcript) ProtoMessage()               {}
func (*Transcript) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Transcript) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Transcript) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *Transcript) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Transcript) GetSeqId() string {
	if m != nil {
		return m.SeqId
	}
	return ""
}

func (m *Transcript) GetStart() int32 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Transcript) GetEnd() int32 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *Transcript) GetStrand() string {
	if m != nil {
		return m.Strand
	}
	return ""
}

type GeneSynonym struct {
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Gid    string `protobuf:"bytes,2,opt,name=gid" json:"gid,omitempty"`
	Type   string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Symbol string `protobuf:"bytes,4,opt,name=symbol" json:"symbol,omitempty"`
	// Target: Gene
	SynonymFor []string `protobuf:"bytes,5,rep,name=synonymFor" json:"synonymFor,omitempty"`
	// Target: GeneDatabase
	InDatabase []string `protobuf:"bytes,6,rep,name=inDatabase" json:"inDatabase,omitempty"`
}

func (m *GeneSynonym) Reset()                    { *m = GeneSynonym{} }
func (m *GeneSynonym) String() string            { return proto.CompactTextString(m) }
func (*GeneSynonym) ProtoMessage()               {}
func (*GeneSynonym) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GeneSynonym) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GeneSynonym) GetGid() string {
	if m != nil {
		return m.Gid
	}
	return ""
}

func (m *GeneSynonym) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GeneSynonym) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *GeneSynonym) GetSynonymFor() []string {
	if m != nil {
		return m.SynonymFor
	}
	return nil
}

func (m *GeneSynonym) GetInDatabase() []string {
	if m != nil {
		return m.InDatabase
	}
	return nil
}

type GeneDatabase struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Gid  string `protobuf:"bytes,2,opt,name=gid" json:"gid,omitempty"`
	Type string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Name string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
}

func (m *GeneDatabase) Reset()                    { *m = GeneDatabase{} }
func (m *GeneDatabase) String() string            { return proto.CompactTextString(m) }
func (*GeneDatabase) ProtoMessage()               {}
func (*GeneDatabase) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GeneDatabase) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GeneDatabase) GetGid() string {
	if m != nil {
		return m.Gid
	}
	return ""
}

func (m *GeneDatabase) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GeneDatabase) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GeneFamily struct {
	Id          string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Gid         string `protobuf:"bytes,2,opt,name=gid" json:"gid,omitempty"`
	Type        string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Tag         string `protobuf:"bytes,4,opt,name=tag" json:"tag,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
}

func (m *GeneFamily) Reset()                    { *m = GeneFamily{} }
func (m *GeneFamily) String() string            { return proto.CompactTextString(m) }
func (*GeneFamily) ProtoMessage()               {}
func (*GeneFamily) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GeneFamily) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GeneFamily) GetGid() string {
	if m != nil {
		return m.Gid
	}
	return ""
}

func (m *GeneFamily) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GeneFamily) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *GeneFamily) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*Position)(nil), "bmeg.Position")
	proto.RegisterType((*Gene)(nil), "bmeg.Gene")
	proto.RegisterType((*Transcript)(nil), "bmeg.Transcript")
	proto.RegisterType((*GeneSynonym)(nil), "bmeg.GeneSynonym")
	proto.RegisterType((*GeneDatabase)(nil), "bmeg.GeneDatabase")
	proto.RegisterType((*GeneFamily)(nil), "bmeg.GeneFamily")
}

func init() { proto.RegisterFile("genome.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x4f, 0x8f, 0x9b, 0x30,
	0x10, 0xc5, 0x45, 0x0c, 0x34, 0x0c, 0xa4, 0x52, 0x69, 0x0f, 0x56, 0x4f, 0x88, 0x53, 0xa4, 0x4a,
	0x1c, 0xda, 0x4b, 0xd5, 0x6b, 0xdb, 0x54, 0xb9, 0x55, 0x6a, 0xf7, 0x03, 0x18, 0x98, 0xb0, 0xd6,
	0x62, 0x9b, 0xd8, 0xde, 0x95, 0xf8, 0xc2, 0xfb, 0x39, 0x56, 0x36, 0xb0, 0xf9, 0x73, 0x8a, 0xb4,
	0xc7, 0x37, 0x1a, 0xbd, 0xdf, 0x7b, 0x63, 0x43, 0xd6, 0xa1, 0x54, 0x02, 0xab, 0x41, 0x2b, 0xab,
	0xf2, 0xb0, 0x16, 0xd8, 0x95, 0x2d, 0xac, 0xff, 0x2a, 0xc3, 0x2d, 0x57, 0x32, 0x07, 0x58, 0xf1,
	0x96, 0x06, 0x45, 0xb0, 0x4d, 0xf2, 0x0c, 0x42, 0x3b, 0x0e, 0x48, 0x57, 0x5e, 0xe5, 0x00, 0xcd,
	0xbd, 0x56, 0x42, 0x19, 0x25, 0x90, 0x12, 0x3f, 0xdb, 0x40, 0x64, 0x2c, 0xd3, 0x96, 0x86, 0x45,
	0xb0, 0x25, 0x79, 0x0a, 0x04, 0x65, 0x4b, 0x23, 0x2f, 0xde, 0x43, 0x6c, 0xac, 0x66, 0xb2, 0xa5,
	0xb1, 0xdb, 0x2d, 0x9f, 0x03, 0x08, 0xff, 0xa0, 0xc4, 0x0b, 0x84, 0x5b, 0x1a, 0x45, 0xad, 0xfa,
	0x19, 0xf2, 0x11, 0xd2, 0x16, 0x4d, 0xa3, 0xf9, 0xe0, 0xd2, 0x9c, 0x51, 0xf0, 0xb8, 0x6f, 0x3d,
	0xe5, 0x0c, 0xea, 0x38, 0xd1, 0x02, 0x8d, 0xbd, 0x38, 0x41, 0xdf, 0xf9, 0xdd, 0x0f, 0x90, 0xb0,
	0xa6, 0x41, 0x63, 0x9c, 0xdb, 0x7a, 0x41, 0x6a, 0x3c, 0x18, 0x3c, 0xd2, 0xc4, 0xeb, 0x12, 0x42,
	0x2e, 0x0f, 0x8a, 0xa6, 0x05, 0xd9, 0xa6, 0x5f, 0x3f, 0x55, 0xee, 0x24, 0x95, 0x0b, 0x5a, 0xed,
	0xe5, 0x41, 0xfd, 0x96, 0x56, 0x8f, 0x9f, 0xbf, 0x40, 0xf2, 0x2a, 0x1c, 0xf0, 0x01, 0xc7, 0xb9,
	0xc0, 0x06, 0xa2, 0x27, 0xd6, 0x3f, 0xce, 0x47, 0xfa, 0xb1, 0xfa, 0x1e, 0x94, 0x06, 0xe0, 0xbf,
	0x66, 0x72, 0x6a, 0x71, 0xdd, 0x76, 0x60, 0x1a, 0xa5, 0x9d, 0xdb, 0x9e, 0xda, 0xbf, 0xb5, 0x68,
	0xd9, 0x43, 0xea, 0x32, 0xff, 0x1b, 0xa5, 0x92, 0xa3, 0xb8, 0xa0, 0xa6, 0x40, 0x3a, 0xde, 0xce,
	0xc8, 0xe5, 0x4d, 0xc9, 0x55, 0x80, 0x70, 0x79, 0x63, 0x33, 0x39, 0xec, 0x94, 0xa6, 0x51, 0x41,
	0xa6, 0x19, 0x97, 0xbf, 0x98, 0x65, 0x35, 0x33, 0x48, 0x63, 0x37, 0x2b, 0x7f, 0x42, 0xe6, 0x68,
	0xcb, 0xf4, 0x56, 0x5c, 0x06, 0xa1, 0x64, 0x02, 0x27, 0x58, 0x79, 0x07, 0xe0, 0x4c, 0x76, 0x4c,
	0xf0, 0x7e, 0xbc, 0xd5, 0x22, 0x05, 0x62, 0x59, 0x37, 0xc7, 0xbd, 0xfa, 0x2d, 0xee, 0x4c, 0x49,
	0x1d, 0xfb, 0xaf, 0xfd, 0xed, 0x25, 0x00, 0x00, 0xff, 0xff, 0x73, 0xe5, 0x81, 0xf6, 0xea, 0x02,
	0x00, 0x00,
}
