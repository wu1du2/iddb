# iddb
**2020 RUC ddb course project**   
***Assignment:***  
**Kunyao Wu: irpc,iqueryoptimizer**  
**Hongyao Zhao: mysql installation, iqueryoptimizer**  
**Zhiyu Shui: imeta,etcd installation,iqueryoptimizer**  
**Dalin Wang: iparser, iqueryanalyzer**  
## Overview  
The purpose of this project is to implement a simple Relational Distributed Database(RDDB), iddb. iddb supports handling some simple SQL queries, making query optimization and executing SQL statements in different sites.

iddb is implemented by Golang. This project follows Golang standard of coding.
1. The project iddb is packed up as an isolated directory. iddb consists 3 part, bin, pkg, and src. All codes, including grpc dependency and our packages, are placed in src.
2. All modules of iddb are assembled as Golang Packages. We can use module by using "import PKG_NAME"
3. Make sure the go environment variable “GOPATH” is [xxx/iddb]. Using "go env" to check out.

## Environment
1. Hardware:  
cpu: 8  Intel Core Processor (Skylake)  
memory:  15885 MB  
disk: 500 GB  

2. OS:  
CentOS Linux release 7.4.1708 (Core)  
(kernel version:3.10.0-693.5.2.el7.x86_64)  

3. golang version:  
go1.15.4  

4. mysql version:  
mysql-5.7.32  

5. etcd version:  
etcd-v3.3.25  

6. network:  
cluster:  
node1: 10.77.70.161  
node2: 10.77.70.162  
node3: 10.77.70.171  
  
protocol buffer: libprotoc 3.13.0  
grpc: 1.34.0-dev  

## Design 
Next, we will introduce the design of iddb. iddb consists of 6 main modules, iparser(parsing input SQL statements), irpc (data transmission and remote executor call), iqueryanalyzer(query analyzation) and iqueryoptimizer(optimization), imeta(synchronizing metadata by etcd) and iexecuter(executing SQL in local mysql databases)

### overview  

![iddbarchitecture](https://github.com/wu1du2/iddb/raw/master/img/architecture.jpeg)


### iparser
123123213
21332132

### iqueryanalyzer
123123213
21332132

### iqueryoptimizer
123123213
21332132

### imeta
123123213
21332132

### iexecuter
123123213
21332132

### irpc
interface:
1. irpc/call_client: RunCallClient(caddress string, txnid int64) int64  

2. irpc/call_server: RunCallServer()  

3. irpc/tran_client: RunTranClient(taddress string, table Table) int64  

4. irpc/tran_server: RunTranServer()  

USAGE: read src/examples/irpc


## installation & configuration
note that all operations need to complete in all 3 computers in cluster.

1. install tools and dependency  
```
sudo yum install git
sudo yum install wget  
```

2. install golang newest version (up to 2020/11/7)
```
wget https://studygolang.com/dl/golang/go1.15.4.linux-amd64.tar.gz  
tar zxf go1.15.4.linux-amd64.tar.gz  
sudo mv go /usr/local/  
```

3. set go environment variables for all 3 computers  
```
export GOROOT=/usr/local/go  
export GOBIN=$GOROOT/bin  
export PATH=$PATH:$GOBIN  
export GOPATH=~/iddb
```

4. install mysql5.7
```
the procedure of installing mysql5.7
```

5. install etcd
```
environment

(1)packages:
$git clone https://github.com.cnpmjs.org/coreos/go-semver.git
$git clone https://github.com.cnpmjs.org/coreos/go-systemd.git
$git clone https://github.com.cnpmjs.org/gogo/protobuf.git
$git clone https://github.com.cnpmjs.org/uber-go/zap.git
$git clone https://github.com.cnpmjs.org/uber-go/multierr.git
$git clone https://github.com.cnpmjs.org/uber-go/atomic.git

(2)etcd:
$git clone https://github.com.cnpmjs.org/etcd-io/etcd.git

for version 3.3.25
$git checkout 7d12776
$./build

(3)goreman (for cluster)
$git clone https://github.com.cnpmjs.org/mattn/goreman.git
$make

if fialed =>
$go env -w GOPROXY=https://goproxy.cn,direct
$go env -w GOPRIVATE=.gitlab.com,.gite.com
$go env -w GOSUMDB="sum.goolang.google.cn"

(4)configfile: etcd/Procfile_C
# Procfile_C

#161
etcd1: bin/etcd --name etcd161 --listen-client-urls http://10.77.70.161:2379,http://127.0.0.1:2379 --advertise-client-urls http://10.77.70.161:2379 --listen-peer-urls http://10.77.70.161:2380 --initial-advertise-peer-urls http://10.77.70.161:2380 --initial-cluster-token etcd-cluster-1 --initial-cluster 'etcd161=http://10.77.70.161:2380,etcd162=http://10.77.70.162:2380,etcd171=http://10.77.70.171:2380' --initial-cluster-state new --enable-pprof

#162
etcd2: bin/etcd --name etcd162 --listen-client-urls http://10.77.70.162:2379,http://127.0.0.1:2379 --advertise-client-urls http://10.77.70.162:2379 --listen-peer-urls http://10.77.70.162:2380 --initial-advertise-peer-urls http://10.77.70.162:2380 --initial-cluster-token etcd-cluster-1 --initial-cluster 'etcd161=http://10.77.70.161:2380,etcd162=http://10.77.70.162:2380,etcd171=http://10.77.70.171:2380' --initial-cluster-state new --enable-pprof

#171
etcd: bin/etcd --name etcd171 --listen-client-urls http://10.77.70.171:2379,http://127.0.0.1:2379 --advertise-client-urls http://127.0.0.1:2379 --listen-peer-urls http://10.77.70.171:2380 --initial-advertise-peer-urls http://10.77.70.171:2380 --initial-cluster-token etcd-cluster-1 --initial-cluster 'etcd161=http://10.77.70.161:2380,etcd162=http://10.77.70.162:2380,etcd171=http://10.77.70.171:2380' --initial-cluster-state new --enable-pprof

(5)run etcd cluster 
$cd /home/centos/iddb/src/github.com/coreos/etcd
/home/centos/iddb/src/github.com/mattn/goreman/goreman -f Procfile_C start

$nohup /home/centos/iddb/src/github.com/mattn/goreman/goreman -f Procfile_C start &

```

6. install iddb
```
git clone https://github.com/wu1du2/iddb.git  
```


## working log
2020.11.9 discussed the design of iddb
2020.11.11 mid-term report

