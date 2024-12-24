# Polyhedra-Simple-Library-Golang

Polyhedra Network 后端开发 Golang 工具库集合。

[English](README.md)

## 简介

这是一个为 Polyhedra Network 后端开发设计的 Golang 工具库集合。目前包含以下组件：

### AlloyDB Connector

用于连接 Google Cloud Platform (GCP) AlloyDB 的 Go 库，集成了 GORM 支持。

#### 特性

- 简单易用的 API
- 环境变量配置
- GORM ORM 集成
- 完善的错误处理
- 灵活的连接池配置
- 详细的文档

#### 安装

```bash
go get github.com/LinJianan2023/Polyhedra-Simple-Library-Golang/alloydb-connector@v1.0.0
```

#### 环境变量

##### 必需变量

| 变量 | 描述 |
|----------|-------------|
| DB_HOST | AlloyDB 实例 URI |
| DB_USER | 数据库用户名 |
| DB_PASS | 数据库密码 |
| DB_NAME | 数据库名称 |
| DB_CERT_PATH | 服务账号密钥文件路径 |

##### 可选变量（连接池配置）

| 变量 | 描述 | 默认值 | 说明 |
|----------|-------------|---------|------|
| DB_MAX_OPEN_CONNS | 最大打开连接数 | 0 | 0 表示无限制 |
| DB_MAX_IDLE_CONNS | 最大空闲连接数 | 2 | Go 标准库默认值 |
| DB_CONN_MAX_LIFETIME | 连接最大生命周期（分钟） | 0 | 0 表示无限制 |
| DB_CONN_MAX_IDLE_TIME | 空闲连接最大生命周期（分钟） | 0 | 0 表示无限制 |

#### 快速开始

##### 1. 基本用法

```go
import (
    "log"
    "github.com/LinJianan2023/Polyhedra-Simple-Library-Golang/alloydb-connector"
)

func main() {
    // 初始化数据库连接
    if err := alloydbconnector.InitDB(); err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }

    // 获取 GORM 实例
    db := alloydbconnector.GetDB()

    // 现在可以使用 db 进行数据库操作
}
```

##### 2. 使用 GORM

```go
// 定义模型
type User struct {
    ID   uint   `gorm:"primarykey"`
    Name string `gorm:"type:varchar(100)"`
}

// 自动迁移
if err := db.AutoMigrate(&User{}); err != nil {
    log.Fatal(err)
}

// 创建记录
user := User{Name: "test user"}
db.Create(&user)

// 查询记录
var users []User
db.Find(&users)
```

#### 连接池配置建议

##### 1. 最大打开连接数 (DB_MAX_OPEN_CONNS)
- 默认值：0（无限制）
- 建议：
  * 小型应用：5-20
  * 中型应用：20-50
  * 大型应用：50-100
  * 根据服务器资源和实际负载调整

##### 2. 最大空闲连接数 (DB_MAX_IDLE_CONNS)
- 默认值：2（Go 标准库默认值）
- 建议：
  * 设置为最大打开连接数的 1/2 到 1/3
  * 避免设置过高以防止资源浪费
  * 避免设置过低以防止频繁创建连接

##### 3. 连接最大生命周期 (DB_CONN_MAX_LIFETIME)
- 默认值：0（无限制）
- 建议：
  * 生产环境：30-120 分钟
  * 考虑数据库和网络稳定性
  * 较短的生命周期有助于资源回收

##### 4. 空闲连接最大生命周期 (DB_CONN_MAX_IDLE_TIME)
- 默认值：0（无限制）
- 建议：
  * 生产环境：5-30 分钟
  * 较短的时间有助于释放不活跃连接
  * 根据应用访问模式调整

#### 最佳实践

1. 环境变量管理
   - 使用 .env 文件或环境变量管理工具
   - 不要硬编码敏感信息
   - 在生产环境中妥善管理密钥文件

2. 连接池配置
   - 在生产环境中配置所有池参数
   - 根据实际负载调整
   - 定期监控池状态

3. 错误处理
   - 始终检查 InitDB() 返回的错误
   - 妥善处理数据库操作错误

4. 安全性
   - 确保服务账号具有适当的权限
   - 安全管理数据库凭证
   - 在生产环境中使用 SSL 连接

#### 常见问题

1. 连接失败
   - 检查环境变量
   - 验证服务账号密钥文件路径
   - 检查网络连接和防火墙设置

2. 性能问题
   - 检查连接池配置
   - 监控连接使用情况
   - 考虑调整最大连接数

3. 权限问题
   - 检查服务账号权限
   - 验证数据库用户权限

## 版本管理

当前版本：v1.0.0

遵循 [语义化版本 2.0.0](https://semver.org/lang/zh-CN/):
- 主版本号：不兼容的 API 变更
- 次版本号：向下兼容的功能性新增
- 修订号：向下兼容的问题修正

### 版本检查
```go
import "github.com/LinJianan2023/Polyhedra-Simple-Library-Golang/alloydb-connector"

version := alloydbconnector.GetVersion()
fmt.Printf("Current version: %s\n", version)
```

## 许可证

MIT License
