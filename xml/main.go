package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s := `<?xml version="1.0"?>
<!DOCTYPE dble:schema SYSTEM "schema.dtd">
<dble:schema xmlns:dble="http://dble.cloud/">
    <schema name="schema1" sqlMaxLimit="100">
        <table name="test1" rule="hash-four" dataNode="dn1,dn2"></table>
        <table name="test2" rule="hash-four" dataNode="dn3,dn4"></table>
    </schema>
    <!--
    这是一个注释
    
    zhangqiang2
    
    zhangqiang3
    
        zhangqiang4
    -->
    <dataNode name="dn1" dataHost="mysql-1" database="dbtest1"></dataNode>
    <dataNode name="dn2" dataHost="mysql-1" database="dbtest2"></dataNode>
    <dataNode name="dn3" dataHost="mysql-2" database="dbtest3"></dataNode>
    <dataNode name="dn4" dataHost="mysql-2" database="dbtest4"></dataNode>
    <dataHost name="mysql-1" maxCon="1000" minCon="10" balance="0" switchType="-1" slaveThreshold="100">
        <heartbeat>select user()</heartbeat>
        <writeHost host="hostM1" url="172.20.134.3:3306" user="root" password="Iyhmqit7bUqf7yOvxDRcHxXH7hNi4R6lDA5qCBGKdl2MpHIqZLd0QiSAHBCWqif+pS2eNAIoXxgO5k1ujm/KyA==" id="mysql-dble01" usingDecrypt="1"></writeHost>
    </dataHost>
    <dataHost name="mysql-2" maxCon="1000" minCon="10" balance="0" switchType="-1" slaveThreshold="100">
        <heartbeat>select user()</heartbeat>
        <writeHost host="hostM1" url="172.20.134.3:3307" user="root" password="Iyhmqit7bUqf7yOvxDRcHxXH7hNi4R6lDA5qCBGKdl2MpHIqZLd0QiSAHBCWqif+pS2eNAIoXxgO5k1ujm/KyA==" id="mysql-dble02" usingDecrypt="1"></writeHost>
    </dataHost>
</dble:schema>`
	t := ""
	json.Unmarshal([]byte(s), &t)
	fmt.Println(t)

}
