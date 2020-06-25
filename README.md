<p align="center">
    <img src="https://github.com/neucn/teemo/blob/master/docs/logo.png?raw=true" alt="logo" width="200">
</p>

<h1 align="center">Teemo</h1>

<p align="center">
    <img src="https://img.shields.io/github/v/tag/neucn/teemo?label=version&style=flat-square" alt="">
    <img src="https://img.shields.io/github/license/neucn/teemo?style=flat-square" alt="">
</p>


> 东北大学 GPA 变动哨兵

<p align="center">Windows</p>
<p align="center">
    <img src="https://github.com/neucn/teemo/blob/master/docs/demo@windows.png?raw=true" alt="windows demo">
</p>
<p align="center">Linux</p>
<p align="center">
    <img src="https://github.com/neucn/teemo/blob/master/docs/demo@linux.png?raw=true" alt="linux demo">
</p>

## 系统要求
满足以下之一即可
- Windows 8+
- Linux

暂未测试 OSX

## 下载

进入本仓库 [Release 页面](https://github.com/neucn/teemo/releases/latest) 下载最新压缩包后解压即可

## 更新

下载最新版压缩包并解压覆盖原来的程序。

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
    teemo
  
    # 指定配置文件的路径
    teemo -i path/to/config.yaml 
    ```
  
    建议对`teemo`建立软连接，从而可以全局使用
    
    ```shell script
    ln -s teemo可执行文件的路径 /usr/local/bin/teemo
    ```


> - 监听中请不要关闭程序和控制台
> 
> - Linux 下可以使用类似于`nohup`等工具使程序在后台运行


## 开源协议
Mit License.