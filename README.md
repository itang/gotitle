# gotitle

获取指定网页的标题


## Install

```
$ go get -u github.com/itang/gotitle/cmd/gotitle
```

## Examples

```
$ gotitle http://www.sina.com.cn http://www.baidu.com
>http get 'http://www.sina.com.cn'...
	->>title: 新浪首页
>http get 'http://www.baidu.com'...
	->>title: 百度一下，你就知道
```
