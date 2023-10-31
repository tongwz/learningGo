# supervisor启动项目状态检测

功能主要是用来检测项目运行状态是否还存在，如果是非running状态则重启服务，将异常和重启状态发送给企微机器人。



##### 如果自己有项目需要做提醒

基础同步路径是在config/config.toml文件中的

> NoticeUrl 是机器人提醒url
>
> supervisorPrograms里面是要检测的服务，直接改代码，没有加到配置文件里。

