
cwlVersion: v1.0
class: CommandLineTool
hints:
  DockerRequirement:
    dockerPull: ensembl-transform:latest


baseCommand:
  - go
  - run
  - /opt/run.go

inputs:
  GAF_GZ:
    type: File
    inputBinding:
      position: 1

outputs:
  ENSEMBL_JSON:
    type: 
      type: array
      items: File
    outputBinding:
      glob: "*.json"
