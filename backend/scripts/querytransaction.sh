#!/bin/bash

WORKDIR="$GOPATH/src/github.com/hyperledger/fabric-samples/test-network/"

CURDIR="$GOPATH/src/contractdeploy"
cd $WORKDIR || exit

export FABRIC_CFG_PATH=$PWD/../config/
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051
export PATH=${PWD}/../bin:$PATH
export ORDERER_CA=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

#peer chaincode query -C $1 -n $2 -c '20060102150405'
peer chaincode query -C $1 -n $2 -c '{"Args":["querytransaction","20060102150405"]}'
# '{"Args":["queryAllCars"]}'
