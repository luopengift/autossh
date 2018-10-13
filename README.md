# autossh

## Build

```(shell)
go get -u github.com/luopengift/autossh
go get -u github.com/luopengift/readline
# 需要替换readline库
rm -rf $GOPATH/src/github.com/chzyer/readline
ln -s $GOPATH/src/github.com/luopengift/readline $GOPATH/src/github.com/chzyer/readline
make
```
