1. MacOS 下安装mysqlclient 报  `ld: library not found for -lssl`

mac下找不到openssl动态链接库, 解决方案如下
env LDFLAGS="-I/usr/local/opt/openssl/include -L/usr/local/opt/openssl/lib" pip3 --no-cache install mysqlclient

