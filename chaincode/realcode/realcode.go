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
// 코인 기부
type Donation struct{
	Userid   string `json:"userid"`  // 기부자 id
	Campno string `json:"campno"` //캠페인 NO
	Donatecoin int `json:"donatecoin,string"`  //기부코인
}
//코인 충전
type Coin struct{
	Userid   string `json:"userid"`  // 기부자 id
	Putcoin int `json:"putcoin,string"`  //충전 액수
}
// 기부물품 구매
type Purchase struct{
	Campno string `json:"campno"` //캠페인 NO
	Buycoin int `json:"buycoin,string"`  //기부코인
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
	}else if function == "Donate" {
		return t.Donate(stub, args)
	}else if function == "MakeCoin" {
		return t.MakeCoin(stub, args)
	}else if function == "Buyproduct" {
		return t.Buyproduct(stub, args)
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
//코인충전 (사용자 id, 충전액수)
func (t *SmartContract) MakeCoin(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
	fmt.Println("start MakeCoin")
	userid := args[0]
	putcoin, err := strconv.Atoi(args[1])
	
	CoinInfo := &Coin{userid, putcoin}
	CoinInfoBytes, err := json.Marshal(CoinInfo)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(userid, CoinInfoBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

    fmt.Println("MakeCoin-putState complete")
	return shim.Success(nil)
}
//기부물품 구매 (캠페인NO, 기부금 사용액수)
func (t *SmartContract) Buyproduct(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
	fmt.Println("start Buyproduct")
	campno := args[0]
	buycoin, err := strconv.Atoi(args[1])
	
	PurchaseInfo := &Purchase{campno, buycoin}
	PurchaseInfoBytes, err := json.Marshal(PurchaseInfo)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(campno, PurchaseInfoBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

    fmt.Println("BuyProduct-putState complete")
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


