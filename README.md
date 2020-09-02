你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# Transer 

## 简介
golang 翻译器

目前实现了百度和有道的翻译API对接

提供了两个命令行工具:

## Transer

单条翻译，有道 只支持13国语言。

![](Preview/Transer.png)

## QtTranser

Qt的多语言文件批量翻译

![](Preview/QtTranser.png)

## Go环境

golang 1.12及以上

## 编译/运行

在源码目录执行go run并提供参数即可，不提供参数则显示Uasge。 例如：

```
PS E:\Dev\go\Transer> go run .\cmd\QtTranser
  -api string
        baidu | youdao (default "baidu")
  -inputFile string
        the input file need translate
  -outputFile string
        the output file (default "out.tr")
  -secret string
        your secret (default "0d2RvCho9XZNEO5GCGNs")
  -targetLang string
        zh | en | ja | ko | fr | es | pt | it | ru | vi | de | ar | id (default "zh")
  -userID string
        your id (default "20190502000293463")

```

默认userID和secret是我个人申请的有道账号，可能过期，请自行替换为百度/有道账号。

## Build status
[![Build Status](https://travis-ci.org/jaredtao/Transer.svg?branch=master)](https://travis-ci.org/jaredtao/Transer)
## License
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/jaredtao/Transer/blob/master/LICENSE)

参考资料：https://zhuanlan.zhihu.com/p/37690859
