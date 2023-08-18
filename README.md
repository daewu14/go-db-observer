# Database Observer - GO

### Installation
- Get package
  ```shell
  go get github.com/daewu14/go-db-observer
  ```

### Example
- Structure
  ```
    PROJECT_ROOT
    │
    ├── observer
    │   └── user.go
    └── main.go
    ```
- main.go
  ```go
    func main() {
        dbo.Observe(dbo.Config{
            Host:     "localhost",
            User:     "root",
            Password: "db-password",
            Flavor:   "mysql",
            Port:     "3306",
        }, []dbo.Observer{
            // register struct implemented dbo.Observer
            observer.User{},
        })
    }
  ```
- user.go
  ```go
    package observer

    import (
        "bytes"
        "context"
        "encoding/json"
        "github.com/daewu14/go-db-observer/dbo"
    )
    
    type User struct{}
    
    func (m User) Table() string {
        return "users" // your table name
    }
    
    func (m User) OnInserted(ctx context.Context, cnc context.CancelFunc, affected dbo.AffectedTable) {
        println("on inserted", affected.Table)
        for _, column := range affected.Columns {
            println("column name is", column.Name, "new value is", dumpToString(column.NewValue))
        }
    }
    
    func (m User) OnUpdated(ctx context.Context, cnc context.CancelFunc, affected dbo.AffectedTable) {
        println("on updated", affected.Table)
        println("on id", dumpToString(affected.OnFirstColumn.Val))
        for _, column := range affected.Columns {
            println("column name is", column.Name)
            println("old value is", dumpToString(column.OldValue))
            println("new value is", dumpToString(column.NewValue))
        }
    }
    
    func (m User) OnDeleted(ctx context.Context, cnc context.CancelFunc, affected dbo.AffectedTable) {
        println("on deleted", affected.Table)
        for _, column := range affected.Columns {
            println("column name is", column.Name, "with value", dumpToString(column.OldValue))
        }
    }
    
    // just helper to convert data into string data type
    func dumpToString(v interface{}) string {
        str, ok := v.(string)
        if !ok {
            buff := &bytes.Buffer{}
            json.NewEncoder(buff).Encode(v)
            return buff.String()
        }
        return str
    }
  ```