## spring-test-distributed-db:

*FOR*
这个项目是为了测试当数据库分库分表(水平)情况下如何进行数据的查询,分页查询等
*DETAIL*
如果想看流程的话需要去查看我的library这个工程中的sqlextention包
[点击传送门](https://github.com/ItsFunny/Tmall_MicroService/tree/master/library/src/main/java/com/joker/library/sqlextention) 
sql 文件也一并上传了,只需要在本地创建多个库,然后复制,同时修改下配置文件(账户密码,以及一些关键的配置,具体可以看我的那个libray包README)


BUG
---
* `DONE` 当pageSize<10的时候,偏移量会<0,有空的时候再debug更改
    -   原因在于: sql脚本错误(导致2个库中的数据一模一样,id也一样,因而会发生错误)
* `DONE` 当数据分布不均匀的时候,会发生越界 
    -   原因在于: 获取maxId 当判断发现数据为空的时候,并不会插入,因而发生了数组越界


TODO
---
* [x] 常规测试: 数据分布均匀
* [] 非常规测试: 
    -   [x] 部分非常规测试 OK
    -   [x] 极端测试

--- 

TODO-配置
---
* DataSource和SqlSessionFactoryBean 只需要配置数据库相关参数,其余的直接反射注入
* 只需要配置
    -   分库个数(要与DataSource的配置一致)
    -   表前缀名称
    -   
* 自定义注解,被这个注解所使用的dao自动继承接口然后注入到容器中
使用方式:
---
1. 修改config 下的数据库配置,主要是用户名和密码
2. 在本地数据库中创建相关数据库
    -   创建2个库:
        -   test_distribute_db0:
            -   创建3个表,结构相同: 执行test_distribute.sql文件创建3个表
        -   test_distribute_db1:
            -   创建3个表,结构相同: 执行test_distribute.sql文件创建3个表
3. test目录下右键插入`testInsert初始化`->初始化完毕之后->`执行testFindByPage即可`


---

## 遇到的问题

* 在test目录下执行test的时候报错提示:包不存在: 
    -   解决方法: 
        -   原因在于idea编译器的问题 : file->project structure->modules->删除所有的module->再重新导入


---