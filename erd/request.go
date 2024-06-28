package erd

import (
	"fmt"
	"github.com/gosnmp/gosnmp"
	"time"
)

func parsing(pdu gosnmp.SnmpPDU) (string, error) {
	var response string
	switch pdu.Type {
	case gosnmp.OctetString:
		b := pdu.Value.([]byte)
		response = fmt.Sprintf("STRING: %s", string(b))
	default:
		response = fmt.Sprintf("TYPE %d: %d", pdu.Type, gosnmp.ToBigInt(pdu.Value))
	}
	return response, nil
}

func RequestSNMP(ip, password, oid string, port uint16) (string, error) {
	snmp := &gosnmp.GoSNMP{}
	snmp.Target = ip
	snmp.Port = port
	snmp.Community = password
	snmp.Timeout = time.Duration(5 * time.Second)
	snmp.Retries = 3
	snmp.Transport = "udp"

	err := snmp.Connect()
	if err != nil {
		return "", fmt.Errorf("Connect error: %v", err)
	}
	defer snmp.Conn.Close()

	var result string
	err = snmp.Walk(oid, func(pdu gosnmp.SnmpPDU) error {
		response, err := parsing(pdu)
		if err != nil {
			return err
		}
		result += response + "\n"
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("Walk error: %v", err)
	}
	return result, nil
}

func snmpSet(target string, community string, oid string, value int) error {
	params := &gosnmp.GoSNMP{
		Target:    target,
		Port:      161,
		Community: community,
		Version:   gosnmp.Version1,
		Timeout:   time.Duration(2) * time.Second,
	}

	// Подключение к SNMP агенту
	err := params.Connect()
	if err != nil {
		return fmt.Errorf("error connecting to target: %v", err)
	}
	defer params.Conn.Close()

	// Создание PDU для запроса
	pdu := gosnmp.SnmpPDU{
		Name:  oid,
		Type:  gosnmp.Integer,
		Value: value,
	}

	// Выполнение SNMP Set запроса
	_, err = params.Set([]gosnmp.SnmpPDU{pdu})
	if err != nil {
		return fmt.Errorf("error performing SNMP set: %v", err)
	}

	return nil
}
