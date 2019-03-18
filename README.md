## 项目
这个项目提供了飒漫画的漫画相关服务

## 项目说明
- bin为可执行文件及部署包
- config为服务及log配置
- doc文件夹为db脚本等
- src为源码，内modles为定义，process为主要逻辑处理，server为http服务搭建
- glide.yaml为glide管理的相关依赖，如到新环境可直接用glide下载全部依赖
- script为各种运行、功能脚本

### 脚本使用说明

#### 使用提示
- 设置一下run2.sh --- server_name
- 设置一下base_setting.sh --- short_server_name,server_dir,develop_addr,test_addr, user_name
- 下面常用功能中远程操作部分需要相应用户可以免密码登录--https://my.oschina.net/jean/blog/290461

#### 常用功能
- auto_deploy.sh在编译的机上使用，
- 开发、测试机直接重启服务可用上层server下的run2.sh b&&./run2.sh r env(dev或test)--b参数会移除旧目录并解压缩包，so会是压缩包内的旧配置，仅重启的话只用./run2.sh r env(dev或test)就行
- 重启或用script目录下的run.sh restart env(dev或test)）--更改配置后可用这个，不会还原配置）
- 编译打包到开发、测试机并重启服务：./auto_deploy.sh
- 本地编译运行并tail -f log：./auto_deploy.sh l    （后面可加命令，如 |grep db，下6项同）
- 编译打包到开发机并重启服务并tail -f log：./auto_deploy.sh d
- 编译打包到测试机并重启服务并tail -f log：./auto_deploy.sh t
- 本地log：./log l
- 开发机log：./log d
- 测试机log：./log t
- 开发、测试机log：./run2.sh l
- vim打开本地log：./log ol
- vim打开开发机log：./log od
- vim打开测试机log：./log ot
- 开发、测试机vim打开log：./run2.sh ol
