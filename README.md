# Blockchain기반의 기부플랫폼 - Hyperledger Fabric Network 구현

- Hyperledger Fabric 공부한 내용 및 프로젝트 내용 
https://luz0911.tistory.com/category/Hyperledger%20Fabric


<br></br>

## Hyperledger Fabric Network
: 기존 BYFN파일들을 수정하여 3개의 Organization과 6개의 peer, 1개의 채널로 구성된 Network입니다.   
또한, 기부플랫폼에 적합한 chaincode를 작성하였습니다. 

This repository is about the Hyperledger Fabric Network that includes 3 organizations,6 peers,1 channel and 1 chaincode.



<br></br>

###  ✔  준비 사항 
- Ubuntu환경에서 구동하였습니다. 
1. Install cURL
```
sudo apt install curl
```
2. Install Docker & Docker Compose
- Docker
```
curl -fsSL https://get.docker.com/ | sudo sh
```
- Docker Compose
: 최신 버전 설치해야 함- https://github.com/docker/compose/releases 
```
sudo curl -L https://github.com/docker/compose/releases/download/1.22.0-rc2/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
```
```
sudo chmod +x /usr/local/bin/docker-compose
```

3. Install GO lang 
: homepage참조해서 최신버전 설치 - https://golang.org/dl/
```
sudo wget https://storage.googleapis.com/golang/go1.10.2.linux-amd64.tar.gz
```
4. GO path 설정해주기 
```
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```
<br></br>

### ✔ File Location
- Network 구동 - fabricnetwork/first-network
- chaincode - fabricnetwork/chaincode/realcode/realcode.go


<br>
</br>
