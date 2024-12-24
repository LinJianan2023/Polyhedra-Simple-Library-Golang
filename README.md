# Polyhedra-Simple-Library-Golang

A collection of Golang libraries for Polyhedra Network backend development.

[中文文档](README_CN.md)

## Introduction

This is a collection of Golang libraries designed for Polyhedra Network backend development. Currently includes the following components:

### AlloyDB Connector

A Go library for connecting to Google Cloud Platform (GCP) AlloyDB with GORM integration.

#### Features

- Simple and easy-to-use API
- Environment variable configuration
- GORM ORM integration
- Comprehensive error handling
- Flexible connection pool configuration
- Detailed documentation

#### Installation

```bash
go get github.com/LinJianan2023/Polyhedra-Simple-Library-Golang/alloydb-connector@v1.0.0
```

#### Environment Variables

##### Required Variables

| Variable | Description |
|----------|-------------|
| DB_HOST | AlloyDB instance URI |
| DB_USER | Database username |
| DB_PASS | Database password |
| DB_NAME | Database name |
| DB_CERT_PATH | Service account key file path |

##### Optional Variables (Connection Pool)

| Variable | Description | Default | Note |
|----------|-------------|---------|------|
| DB_MAX_OPEN_CONNS | Maximum number of open connections | 0 | 0 means unlimited |
| DB_MAX_IDLE_CONNS | Maximum number of idle connections | 2 | Go standard library default |
| DB_CONN_MAX_LIFETIME | Maximum connection lifetime (minutes) | 0 | 0 means unlimited |
| DB_CONN_MAX_IDLE_TIME | Maximum idle connection lifetime (minutes) | 0 | 0 means unlimited |

#### Quick Start

##### 1. Basic Usage

```go
import (
    "log"
    "github.com/LinJianan2023/Polyhedra-Simple-Library-Golang/alloydb-connector"
)

func main() {
    // Initialize database connection
    if err := alloydbconnector.InitDB(); err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }

    // Get GORM instance
    db := alloydbconnector.GetDB()

    // Now you can use db for database operations
}
```

##### 2. Using GORM

```go
// Define model
type User struct {
    ID   uint   `gorm:"primarykey"`
    Name string `gorm:"type:varchar(100)"`
}

// Auto migrate
if err := db.AutoMigrate(&User{}); err != nil {
    log.Fatal(err)
}

// Create record
user := User{Name: "test user"}
db.Create(&user)

// Query records
var users []User
db.Find(&users)
```

#### Connection Pool Configuration

##### 1. Maximum Open Connections (DB_MAX_OPEN_CONNS)
- Default: 0 (unlimited)
- Recommendations:
  * Small applications: 5-20
  * Medium applications: 20-50
  * Large applications: 50-100
  * Adjust based on server resources and actual load

##### 2. Maximum Idle Connections (DB_MAX_IDLE_CONNS)
- Default: 2 (Go standard library default)
- Recommendations:
  * Set to 1/2 to 1/3 of maximum open connections
  * Avoid setting too high to prevent resource waste
  * Avoid setting too low to prevent frequent connection creation

##### 3. Maximum Connection Lifetime (DB_CONN_MAX_LIFETIME)
- Default: 0 (unlimited)
- Recommendations:
  * Production: 30-120 minutes
  * Consider database and network stability
  * Shorter lifetime helps with resource recycling

##### 4. Maximum Idle Connection Lifetime (DB_CONN_MAX_IDLE_TIME)
- Default: 0 (unlimited)
- Recommendations:
  * Production: 5-30 minutes
  * Shorter time helps release inactive connections
  * Adjust based on application access patterns

#### Best Practices

1. Environment Variable Management
   - Use .env file or environment variable management tools
   - Never hardcode sensitive information
   - Properly manage key files in production

2. Connection Pool Configuration
   - Configure all pool parameters in production
   - Adjust based on actual load
   - Monitor pool status regularly

3. Error Handling
   - Always check InitDB() return error
   - Properly handle database operation errors

4. Security
   - Ensure service account has appropriate permissions
   - Securely manage database credentials
   - Use SSL connection in production

#### Common Issues

1. Connection Failures
   - Check environment variables
   - Verify service account key file path
   - Check network connection and firewall settings

2. Performance Issues
   - Review connection pool configuration
   - Monitor connection usage
   - Consider adjusting maximum connections

3. Permission Issues
   - Check service account permissions
   - Verify database user permissions

## Version Management

Current Version: v1.0.0

Following [Semantic Versioning 2.0.0](https://semver.org/):
- Major version: Incompatible API changes
- Minor version: Backwards-compatible functionality additions
- Patch version: Backwards-compatible bug fixes

### Version Check
```go
import "github.com/LinJianan2023/Polyhedra-Simple-Library-Golang/alloydb-connector"

version := alloydbconnector.GetVersion()
fmt.Printf("Current version: %s\n", version)
```

## License

MIT License
