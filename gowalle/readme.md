### 功能：
基于sqlx和gin的强大功能，简直了  
并没有设计成mybatis的样子，而是做成了代码生成，修改mybatis的配置文件可以动态修改sql语句和功能，而不需要代码重新编译。暂时我不需要这样的功能。如果需要可以参考gomybatis的功能。

#### 定义函数类型
* 定义结构体类型，快速组合已知的数据库表结构体
* 定义go函数类型 func name(argname argtype,...) rettype 
* 参数支持对象和参数列表，对象使用 : 来访问变量， 参数列表使用 ?
 ：访问方式支持嵌套。? 的次序和个数应和函数的参数序列匹配，不支持乱序
#### 查询
 * select 多个根据 返回值是否是slice来该判定：get select 
 * 支持定义新的struct对象，并且支持嵌套数据的查询 根据结构体中的 db tag的字段进行查找 例如 
    ```
    var employees []struct {
      Emp1  *Employee `db:"emp1"`
      Emp2  *Employee `db:"emp2"`
      *Boss `db:"boss"`
    }
    err := db.Select(&employees,
        `SELECT emp.name "emp1.name", emp.id "emp1.id", emp.boss_id "emp1.boss_id",
        emp.name "emp2.name", emp.id "emp2.id", emp.boss_id "emp2.boss_id",
        boss.id "boss.id", boss.name "boss.name" FROM employees AS emp
        JOIN employees AS boss ON emp.boss_id = boss.id
      `)
    ```
* insert 单个时，返回值为新插入的autoincrement的id
 
#### 执行:
   * NamedExec 支持插入，更新，删除
   * 插入支持slice批量操作，使用语句和单个没有任何区别，在声明函数时，根据输入参数的类型来判断
   * ：支持嵌套
        ```
        q3 := `INSERT INTO placeperson (first_name, last_name, email, place_id)
        VALUES (:first_name, :last_name, :email, :place.id)`
        _, err = db.NamedExec(q3, pp)
        ```
#### 动态sql
还没有想好设计的方式

### http功能
http的参数名称和类型等同数据库表的字段，url的route应该从gin的group而来，这样可以由外部的gin去指定route的路径深度
* get从query而来
* post从body而来
* delete从query而来
* update从body而来  

应该使用shouldbindquery和shouldbindbody来自动绑定参数到结构体，当数据库的函数不是结构体时，在gin模块内生成局部结构体定义
