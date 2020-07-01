<p align="center">
    <img src="https://github.com/neucn/teemo/blob/master/docs/logo.png?raw=true" alt="logo" width="200">
</p>

<h1 align="center">Teemo</h1>

<p align="center">
    <img src="https://img.shields.io/github/v/tag/neucn/teemo?label=version&style=flat-square" alt="">
    <img src="https://img.shields.io/github/license/neucn/teemo?style=flat-square" alt="">
</p>


> 东北大学 GPA 变动哨兵

## 系统要求
满足以下之一即可
- Windows 8+
- Linux

暂未测试 OSX

## 下载

进入本仓库 [Release 页面](https://github.com/neucn/teemo/releases/latest) 下载最新压缩包后解压即可。

目前提供了供 amd64 下的 Linux, OSX, Windows 以及 arm64 的 Linux 使用的压缩包。

其他架构或系统请自行构建。

## 更新

下载最新版压缩包并解压覆盖原来的程序。

## 配置

配置分为两部分，`notifiers` 以及 `tasks`，分别定义了通知相关的配置和监听相关的配置

目前支持三种通知方式: 邮件通知，钉钉通知，桌面通知。三者可以同时使用。

更多配置说明请参阅 [Schema](https://github.com/neucn/teemo/blob/master/schema.yaml).

## 使用

解压后修改目录下的 `config.yaml` 或参照 [Schema](https://github.com/neucn/teemo/blob/master/schema.yaml) 自行编写配置

- Windows:

    编写好配置文件后双击打开`teemo.exe`即可
    
    也可以通过控制台使用
    ```shell script
    teemo.exe
    
    # 也可以自己指定配置文件的路径，默认是程序所在目录下的 config.yaml
    teemo.exe -i path/to/config.yaml 
    
    # 如果把目录加入了环境变量 path 中，则可以
    teemo
    ```
    
- Linux:
  
    编写好配置文件后可直接通过命令行使用
    ```shell script
    # 初次使用时赋予执行权限
    chmod +x ./teemo
    
    # 直接使用
    ./teemo
    
    # 指定配置文件的路径
    ./teemo -i path/to/config.yaml 
    ```
  
    建议对`teemo`建立软连接，从而可以全局使用
    
    ```shell script
    ln -s teemo可执行文件的路径 /usr/local/bin/teemo
    
    # 之后就可以直接
    teemo
    ```


> - 监听中请不要关闭程序和控制台
> 
> - Linux 下可以用类似 `nohup teemo &` 的方式使程序在后台运行


## 开源协议
MIT License.