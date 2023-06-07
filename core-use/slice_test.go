package core_use

import (
	"fmt"
	"testing"
)

func TestOrder(t *testing.T) {
	workHead := []string{
		"DBInstanceId(实例ID)",
		"DBInstanceDescription(实例描述)",
		"*DBInstanceType(实例类型)",
		"*DBInstanceClassType(实例规格族)",
		"*Category(系列)",
		"*DBInstanceCPU(CPU（核）)",
		"*DBInstanceMemory(内存)",
		"*MaxIOPS(IOPS)",
		"*DBInstanceStorage(存储（GB）)",
		"*RegionId(地域)",
		"*ZoneId(可用区)",
		"*Engine(数据库类型)",
		"*EngineVersion(数据库版本)",
		"InstanceNetworkType(网络类型)",
		"PayType(付费类型)",
		"CreationTime(创建时间)",
		"ExpireTime(到期时间)",
		"ZoneId(所在可用区)",
		"VpcId(虚拟网络ID)",
		"VSwitchId(虚拟交换机ID)",
		"Tags(标签)",
		"MaxConnections(连接数)",
		"IOPSUsage(IOPS利用率（%）)",
		"CPUUsage(CPU使用率（%）)",
		"SessionUsage(连接数利用率（%）)",
		"DiskUsage(磁盘空间利用率（%）)",
		"PrivatePort(内网端口)",
		"PrivateString(内网连接地址)",
		"PublicPort(外网端口)",
		"PublicString(外网连接地址)",
		"*实例规格",
		"*架构",
	}
	for _, v := range workHead {
		fmt.Printf("%s ", v)
	}
}
