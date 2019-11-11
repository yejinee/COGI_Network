//캠페인 등록에 대한 chaincode
package main
import (
	"fmt"
	"strconv"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"bytes"
)
type SmartContract struct {
}
//캠페인 등록
type Campaign struct {
	Campno   string `json:"campno"`  // 캠페인 NO
	Campname string `json:"campname"` //캠페인 명
	Orgname string `json:"orgname"` // 기부단체 명
	Target int `json:"target,string"`  //목표액
}
func (t *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Chaincode Init!!")
	return shim.Success(nil)
}
func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "newCampaign" {
		return t.newCampaign(stub, args)
	}else if function == "query" {
		return t.query(stub, args)	
	}
	return shim.Error("Invalid invoke function name.")
}
//캠페인 새로 등록하는 경우(캠페인no, 캠페인 이름, 기부단체명, 목표액)
func (t *SmartContract) newCampaign(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
	fmt.Println("start newCampaign")
	campno := args[0]
	campname := args[1]
	orgname := args[2]
	target, err := strconv.Atoi(args[3])
	
	CampInfo := &Campaign{campno, campname,orgname,target}
	CampInfoBytes, err := json.Marshal(CampInfo)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(campno, CampInfoBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

    fmt.Println("NewCampaign-putState complete")
	return shim.Success(nil)
}

// query callback representing the query of a chaincode
func (t *SmartContract) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("call query method")
	queryString := "{\"selector\":{}}"
	fmt.Println("queryString" + queryString)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {
	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return nil, err
	}

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())
	return buffer.Bytes(), nil
}

func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	// buffer is a JSON array containing QueryResults
	
	fmt.Println("call constructQueryResponseFromIterator")
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}

		fmt.Println("what")
		fmt.Println(string(queryResponse.Value))
		fmt.Println("what")

		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return &buffer, nil
}

func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

