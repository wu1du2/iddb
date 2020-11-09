# iddb
**2020 RUC ddb course project**   
***Assignment:***  
**Kunyao Wu: rpc**  
**Hongyao Zhao: mysql installation, query processing**  
**Zhiyu Shui: etcd installation, data allocation**  
**Dalin Wang: SQL parser**  
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
Next, we will introduce the design of iddb. iddb consists of 5 main modules, iparser(parsing input SQL statements), irpc (data transmission and networking), iquerymanager(query analyzation and optimization), imeta(synchronizing metadata) and iexecuter(executing SQL in local mysql databases)

### overview  
123  



### iparser
123123213
21332132

### irpc
123123213
21332132

### iquerymanager
123123213
21332132

### imeta
123123213
21332132

### iexecuter
123123213
21332132

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
the procedure of installing etcd
```

6. install iddb
```
git clone https://github.com/wu1du2/iddb.git  
```


## working log
2020.11.9 discuss the design of iddb

## 设计文档
