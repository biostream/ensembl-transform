package main

import (
  "io"
  "os"
  "fmt"
  "log"
  "flag"
  "bmeg"
  "compress/gzip"
  "github.com/blachlylab/gff3"
  "github.com/golang/protobuf/jsonpb"
)

func main() {
  flag.Parse()
  
  json := jsonpb.Marshaler{}
  
  gz_handle, _ := os.Open(flag.Arg(0))
  handle, _ := gzip.NewReader(gz_handle)
  read := gff3.NewReader(handle)
  for feat, err := read.Read(); err != io.EOF; feat, err = read.Read() {
    switch feat.TypeField {
    case "CDS":
      //fmt.Printf("%s %s\n", feat.AttributesField["ID"], feat.TypeField)
    case "five_prime_UTR":
    
    case "three_prime_UTR":
      //test
    
    case "biological_region":
      
    case "mRNA":
    case "tRNA":
    case "vaultRNA_primary_transcript":
    case "lnc_RNA":
    
    case "gene":      
      g := bmeg.Gene{
        Id:feat.AttributesField["ID"],
        Symbol:feat.AttributesField["Name"],
        Description:feat.AttributesField["description"],
        Strand:string(feat.StrandField),
        Start:int32(feat.StartField),
	      End:int32(feat.EndField),
        SeqId:feat.SeqidField,
      }
      s, _ := json.MarshalToString(&g)
      fmt.Printf("%s\n", s )
    case "transcript":
      t := bmeg.Transcript{
        Id:feat.AttributesField["ID"],
        Parent:feat.AttributesField["Parent"],
        Symbol:feat.AttributesField["Name"],
        Strand:string(feat.StrandField),
        Start:int32(feat.StartField),
	      End:int32(feat.EndField),
        SeqId:feat.SeqidField,
      }
      s, _ := json.MarshalToString(&t)
      fmt.Printf("%s\n", s )

    case "exon":
      //fmt.Printf("%s %s\n", feat.AttributesField["Parent"], feat.TypeField)
    
    case "miRNA":
    case "snoRNA":
    case "snRNA":
    case "rRNA":
    case "pseudogene":
    case "scRNA":
    case "chromosome":  
    case "pseudogenic_transcript":    
    case "ncRNA":
    case "ncRNA_gene":
    case "V_gene_segment":
    case "J_gene_segment":
    case "C_gene_segment":
    case "D_gene_segment":

    default:
      log.Printf("Unknown %s\n", feat.TypeField )
    }
  }
  gz_handle.Close()
  
}