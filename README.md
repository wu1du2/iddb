# iddb
**2020 RUC ddb course project**
## Overview
The purpose of this project is to implement a simple Relational Distributed Database(RDDB), iddb. iddb supports handling some simple SQL queries, making query optimization and executing SQL sequences in different sites.

iddb is implemented by Golang. This project follows Golang standard of coding.
1. The project iddb is packed up as an isolated directory. iddb consists 3 part, bin, pkg, and src. All codes, including grpc dependency and our packages, are placed in src.
2. All modules of iddb are assembled as Golang Packages. We can use module by using "import PKG_NAME"
3. Make sure the go environment variable “GOPATH” is [xxx/iddb]. Using "go env" to check out.

## Environment
Hardware:  
OS:  
golang version:  
go environment:  
mysql version:  
etcd version:  
network:  

## Modules 
Next, we will introduce the design of iddb. iddb consists of 5 main modules, iparser(analyzing input SQL sentence), irpc (data transmission and networking), iquerymanager(query optimization), imeta(synchronizing metadata) and iexecuter(execute SQL in local mysql)

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
1. install golang newest version (up to 2020/11/7)

2. set go environment variables for all 3 computers

3. install mysql5.7

4. install etcd

5. install iddb

