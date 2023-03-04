# onebot-plus-plugin-template
onebot-plus-plugin 模板工程
## 用法
1. 将本仓库克隆到本地
2. 修改go.mod中的module名
3. 运行`go mod tidy`更新依赖
4. 修改main.go中插件的逻辑
## 构建插件
`sh build.sh $plugin_name`
- 不同平台的插件将会生成在plugins目录下