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

# 还没有想好设计的方式

* 动态sql咋实现？

* 一对多go咋实现,多对多？
    ```
    通过日志可以清楚地看到，SQL执行的结果数有3条，后面输出的用户数是2，
    也就是说本来查询出的3条结果经过MyBatis对collection数据的处理后，变成了两条。
    我们都知道，因为第一个用户拥有两个角色，所以转换为一对多的数据结构后就变成了两条结果。
    那么，MyBatis又是怎么知道要处理成这样的结果呢？

    先来看MyBatis是如何知道要合并ad.min的两条数据的，为什么不把test这条数据也合井进去呢？
    MyBatis在处理结果的时候，会判断结果是否相同，如果是相同的结果，则只会保留第一个结果，
    所以这个问题的关键点就是MyBatis如何判断结果是否相同。

    MyBatis判断结果是否相同时，最简单的情况就是在映射配置中至少有一个id标签，在userMap中配置如下<id property=”id”column=”id”/>
    我们对id（构造方法中为id Arg）的理解一般是，它配置的字段为表的主键（联合主键时可以配置多个id标签），
    因为MyBatis的resultMap只用于配置结果如何映射，并不知道这个表具体如何。
    
    id的唯一作用就是在嵌套的映射配置时判断数据是否相同，当配置id标签时，MyBatis只需要逐条比较所有数据中id标签配置的字段值是否相同即可。
    在配置嵌套结果查询时，配置id标签可以提高处理效率

    大家通过这个简单的例子应该明白id的作用了。需要注意，很可能会出现一种没有配置工d的情况。
    没有配置id时，MyBatis就会把resultMap中配置的所有字段进行比较，如果所有字段的值都相同就合并，只要有一个字段值不同，就不合井。

    在嵌套结果配直id属性时如果查询语句中没有查询id属性配直的列，就会导致id对应的值为null。这种情况下，所有值的id都相同，因此会使嵌套的集合中只有一条数据。所以在配直id71］时，查询语句中必须包含该列
    ```

### http功能
http的参数名称和类型等同数据库表的字段，url的route应该从gin的group而来，这样可以由外部的gin去指定route的路径深度
* get从query而来
* post从body而来
* delete从query而来
* update从body而来  

应该使用shouldbindquery和shouldbindbody来自动绑定参数到结构体，当数据库的函数不是结构体时，在gin模块内生成局部结构体定义


####
java -jar ./antlr-4.7.1-complete.jar -visitor ./g4/PREFIX.g4  -o ./parser  -no-listener -Dlanguage=Go
java -jar ./antlr-4.7.1-complete.jar -visitor ./g4/GOTYPES.g4  -o ./gotypes  -no-listener -Dlanguage=Go -package gotypes

### go reflect 
https://studygolang.com/articles/12348?fr=sidebar