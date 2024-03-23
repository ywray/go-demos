# go-demos

提高 go 代码的开发效率

该项目包含了标准库、第三方开源库的快速使用, 还有一些个人编写的工具类代码段 ( 如计算两个字符串的编辑距离)

有些代码需要跑起来才能看到效果, 具体测试方法会写在各子路径中的main.go文件中

## 代码结构

### stl 存放标准库相关的demo
* sort_map:  按value将map排序
* errgroup:  在 WaitGroup 的基础上实现子协程错误传递, 同时使用 context 控制协程的生命周期
* format: 对比Sprintf和math.Round
* use_regexp: 字符串是否满足正则, 如不满足, 定制化错误提示
* use_url: 解析url中的地址/参数
* md5
* use_string/base64: 编解码
* use_string/index: 对string和[]rune索引的区别

### third 存放第三方库相关的demo
* use_rollingwriter: 本地写日志文件, 可配置单个log的文件名、文件大小、滚动策略、拆分规则、最大存留数等
* use_redis/redislock: go-redis lib
* use_zerolog: 日志库
* use_uuid: 截取uuid生成纯数字id
* use_yaml: yaml配置文件读取、解析
* use_gin: gin实现简单ping pong server
* fuzzysearch: 字符串模糊匹配

### persional 本人写的一些工具类片段
* read_file: 遍历某路径下所有文件
* ldflags: 配合makefile文件, 将makefile中的变量赋值给go内部
