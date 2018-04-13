# HelloGolang REST API  脚手架项目.
## 最近刚刚学习 golang 这门语言的成果。


- 使用 sqlx 处理数据库，使用 echo 做 mvc
- 格式化输出 json ，处理 nil 值。 统一数据返回格式


### exmaple
``` cmd
    curl http://127.0.0.1
```

### response
``` json
{
    "code": 200,
    "data": null,
    "message": "service run",
    "success": true
}
```
### How to build

``` shell

yum install go -y

vi /etc/profile
    export GOPATH=/root/go
    export PATH=$PATH:$GOPATH/bin
    
:wq


source /etc/profile
    
cd /root/go && mkdir src

cd ./src


git clone https://github.com/bestK/hello.git

cd hello


go get

go build

nohup ./hello &
```
