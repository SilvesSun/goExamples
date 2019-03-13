有许多的业务场景需要读取配置文件, 常见的配置文件格式为 `ini`, `YAML`, `JSON`等. 这里选用 [Viper](https://github.com/spf13/viper)

其特点有
- 支持默认值
- 可读取JSON、TOML、YAML、HCL
- 可热加载配置文件
- 可读取环境变量值
- 从远程配置中心读取配置（etcd/consul），并监控变动
- 支持命令行
- 支持从缓存读取配置
- 可直接设置值

配置的读取顺序如下:
- viper.Set() 所设置的值
- 命令行 flag
- 环境变量
- 配置文件
- 配置中心
- 默认值

这里采用 `YAML` 文件作为配置示例, 文件类容为

```yaml
common:
  database:
    name: test
    host: 127.0.0.1
    user: test
    passwd: testpwd
```