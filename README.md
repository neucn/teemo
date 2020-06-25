<p align="center">
    <img src="https://raw.githubusercontent.com/neucn/teemo/master/.github/img/logo.png" alt="logo" width="200">
</p>

<h1 align="center">Teemo</h1>

<p align="center">
    <img src="https://img.shields.io/github/v/tag/neucn/teemo?label=version&style=flat-square" alt="">
    <img src="https://img.shields.io/github/license/neucn/teemo?style=flat-square" alt="">
</p>


> 东北大学 GPA 变动哨兵

<p align="center">Windows</p>
<p align="center">
    <img src="https://raw.githubusercontent.com/neucn/teemo/master/.github/img/demo@windows.png" alt="windows demo">
</p>
<p align="center">Linux</p>
<p align="center">
    <img src="https://raw.githubusercontent.com/neucn/teemo/master/.github/img/demo@linux.png" alt="linux demo">
</p>

## 系统要求
满足以下之一即可
- Windows 8+
- Linux

暂未测试 OSX

## 下载

进入本仓库 [Release 页面](https://github.com/neucn/teemo/releases/latest) 下载压缩包

- Windows:

    将解压后的`teemo.exe`所在目录的路径加入环境变量 `path` 中
    
- Linux:
  
    建议直接解压到 `/usr/local/bin` 下

## 更新

下载最新版压缩包并解压覆盖原来的程序。

## 使用

本工具为命令行工具，需要在控制台使用。

```
teemo -u 学号 -p 密码
    如 teemo -u 2018xxxx -p abcdefg

teemo -u 学号 -p 密码 -f 监控频率(单位秒)
    如 teemo -u 2018xxxx -p abcdefg -f 60

teemo -u 学号 -p 密码 -v 使用webvpn
    如 teemo -u 2018xxxx -p abcdefg -v

若不指定f，默认60
```

> 监听中请不要关闭程序和控制台
> 
> - Linux 下可以使用类似于`nohup`等工具使程序在后台运行



## 开源协议
Mit License.