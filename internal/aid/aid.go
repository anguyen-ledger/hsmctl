package aid

import (
        "fmt"
        "os"
        "github.com/ghodss/yaml"
        "encoding/json"
        "github.com/itchyny/gojq"
        "encoding/hex"
)

type Compartment struct {
  AID int `json:"attestation_id"`
  CID int `json:"compartment_id"`
}

func check_err(e error) {
  fmt.Println(e)
  os.Exit(1)
}

func Generate(p []string) {
  var result Compartment
  var input map[string]interface{}
  var src_file, dst_json, tmp_json []byte
  var dst_file *os.File
  var err error
  //f,_ := os.ReadFile(os.Args[1])
  src_file,err = os.ReadFile(p[0])
  if err != nil {
    check_err(err)
  }
  dst_json,err = yaml.YAMLToJSON(src_file)
  if err != nil {
    check_err(err)
  }
  if err = json.Unmarshal(dst_json, &input); err != nil {
    check_err(err)
  }
  query,_ := gojq.Parse(".gates | to_entries| .[].value | select(.\"compartment_id\" != null) |. += if has(\"attestation_id\") then null else {\"attestation_id\": .\"compartment_id\"} end | {\"compartment_id\", \"attestation_id\"}")
  iter := query.Run(input)
  for {
        v, ok := iter.Next()
        if !ok {
          break
        }
        tmp_json,err = json.Marshal(v)
          if err != nil {
            check_err(err)
          }
        json.Unmarshal([]byte(tmp_json), &result)
        dst_file, err = os.OpenFile(fmt.Sprintf("/tmp/attest/attest_%d",result.CID), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
        if err != nil {
          check_err(err)
        }
        src := []byte(fmt.Sprintf("0300564c54004f5247%08X", result.AID))
        dst := make([]byte, hex.DecodedLen(len(src)))
        n, err := hex.Decode(dst, src)
        if err != nil {
          check_err(err)
        }
        dst_file.WriteString(fmt.Sprintf("%s", dst[:n]))
        dst_file.Close()
  }
}
