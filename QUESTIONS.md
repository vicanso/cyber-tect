# 问题汇总

## Mini Redis

用户session等信息会使用redis保存，由于项目主要用于应用监控，登录用户数较少，为了减少对其它服务的依赖，因此使用[miniredis](https://github.com/alicebob/miniredis)替换真正的redis服务，miniredis是基于内存的，因此如果程序重启则sesion信息会清除。

## 超级用户

第一个注册的用户自动创建为超级用户。