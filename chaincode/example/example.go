//코인 기부에 대한 chaincode
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
// 코인 기부
type Donation struct{
	Userid   string `json:"userid"`  // 기부자 id
	Campno string `json:"campno"` //캠페인 NO
	Donatecoin int `json:"donatecoin,string"`  //기부코인
}
func (t *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Chaincode Init!!")
	return shim.Success(nil)
}
func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "Donate" {
		return t.Donate(stub, args)
	}else if function == "query" {
		return t.query(stub, args)	
	}
	return shim.Error("Invalid invoke function name.")
}
//사용자가 기부하는 경우(사용자id, 캠페인NO, 기부코인)
func (t *SmartContract) Donate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
	fmt.Println("start Donate")
	userid := args[0]
	campno := args[1]
	donatecoin, err := strconv.Atoi(args[2])
	
	DonateInfo := &Donation{userid, campno,donatecoin}
	DonateInfoBytes, err := json.Marshal(DonateInfo)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(userid, DonateInfoBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

    fmt.Println("Donate-putState complete")
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
