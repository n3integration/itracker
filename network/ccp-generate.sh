#!/bin/bash

function one_line_pem {
    echo "`awk 'NF {sub(/\\n/, ""); printf "%s\\\\\\\n",$0;}' $1`"
}

function json_ccp {
    local PP=$(one_line_pem $6)
    local CP=$(one_line_pem $7)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${MSP}/$2/" \
        -e "s/\${P0PORT}/$3/" \
        -e "s/\${P1PORT}/$4/" \
        -e "s/\${CAPORT}/$5/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        ccp-template.json
}

function yaml_ccp {
    local PP=$(one_line_pem $6)
    local CP=$(one_line_pem $7)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${MSP}/$2/" \
        -e "s/\${P0PORT}/$3/" \
        -e "s/\${P1PORT}/$4/" \
        -e "s/\${CAPORT}/$5/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        ccp-template.yaml | sed -e $'s/\\\\n/\\\n        /g'
}

ORG=warehouse
MSP=Warehouse
P0PORT=7051
P1PORT=8051
CAPORT=7054
PEERPEM=crypto-config/peerOrganizations/warehouse.widgets.com/tlsca/tlsca.warehouse.widgets.com-cert.pem
CAPEM=crypto-config/peerOrganizations/warehouse.widgets.com/ca/ca.warehouse.widgets.com-cert.pem

echo "$(json_ccp $ORG $MSP $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > connection-warehouse.json
echo "$(yaml_ccp $ORG $MSP $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > connection-warehouse.yaml

ORG=iad-factory
MSP=IADFactory
P0PORT=9051
P1PORT=10051
CAPORT=8054
PEERPEM=crypto-config/peerOrganizations/iad-factory.widgets.com/tlsca/tlsca.iad-factory.widgets.com-cert.pem
CAPEM=crypto-config/peerOrganizations/iad-factory.widgets.com/ca/ca.iad-factory.widgets.com-cert.pem

echo "$(json_ccp $ORG $MSP $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > connection-iad-factory.json
echo "$(yaml_ccp $ORG $MSP $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM)" > connection-iad-factory.yaml
