# BUPTLogin
![](https://img.shields.io/badge/version-0.0.1-green.svg)
![](https://img.shields.io/github/stars/ingbyr/BUPTNetLoginByGo.svg)
![](https://img.shields.io/github/forks/ingbyr/BUPTNetLoginByGo.svg)
![](https://img.shields.io/github/issues/ingbyr/BUPTNetLoginByGo.svg)

北邮校园网网关登陆工具，适配新网关 ngw.bupt.edu.cn


## 安装
[下载页面](https://github.com/ingbyr/BUPTNetLoginByGo/releases)

## 使用方法
运行 `bnl.exe`

```shell
Usage of bnl.exe:
  -l string
        线路选择，可用参数 xyw（校园网）、lt（联通）、yd（移动）、dx（电信）
  -lo
        注销北邮校园网网关
  -p string
        校园网账户密码
  -u string
        校园网账户名称
  -v    版本信息
```

登陆联通网络举例：
```shell
bnl.exe -l lt -u 用户名 -p 密码
```

注销网络举例
```shell
bnl.exe -lo
```

> 更多请前往 [个人博客](https://www.ingbyr.com)