MySQL
---

TODO
---
* [Github上更新的较这里勤](https://github.com/ItsFunny/data-structure-algorithm/tree/master/src/main/source)


---

# 索引

* 什么是索引:索引是指数据库管理系统中的一个排序的数据结构,索引的两大类型有b tree,hash,lsm树
    -   hash索引: 基于hash的特性,检索效率非常的高,一次定位,不需要像b tree一样多次io
        -   缺点: 
            -   仅仅只能满足"=","IN",和"<=>"查询,`不能使用范围查询`,`因为hash当数据变更之后并不能保证变更后的hash与变更前的一致`
            -   `Hash索引无法用来排序`,原因一样,hash一旦数据变更之后就不能保证前后一致
            -   `Hash索引不能利用部分查询`,当通过组合索引的时候,hash索引会将索引和相加,而不是单独计算,组合索引也就无效了
            -   `无法避免全局扫描`,因为可能不同的值具有相同的hash索引
            -   `大量hash碰撞后,并不能保证效率比b tree 高`
    -   b tree索引(是一种k,v结构的树): b tree 是一种多路平衡二叉树,k,v在同一个节点上
        -   具有如下特征:对于`m阶的b tree`
            -   根节点的孩子数目为[ceil(m/2),m],关键字的数目为[ceil(m/2)-1,m-1] 
            -   **非叶节点的根节点至少有2个孩子节点**
            ![](https://img-blog.csdnimg.cn/20190212005404642.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0NvZGVyX0pva2Vy,size_16,color_FFFFFF,t_70)
            btree中叶子节点和根节点以及分支节点都是同一种类型(class)
    -   b+ tree索引(也是一种k,v结构的树): b+ tree 也是一种多路平衡二叉树,但是与b tree的不同点在于,非叶子节点不存放值v,只存放键k****
        -   具有如下的特征:对于`m阶的b+ tree`
            -   根节点的孩子数目为:[ceil(m/2),m],关键字的数目为[ceil(m/2)-1,m]
            -   根节点和分支节点只用来保存关键字(索引),数据地址都存放在叶子节点上
            ![](https://img-blog.csdnimg.cn/20190212005018445.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0NvZGVyX0pva2Vy,size_16,color_FFFFFF,t_70)     
            叶子节点中的5,8,9分别就是数据库中的记录,因此**b+tree中叶子节点与根节点和分支节点不是同一种类型(class)**
    -   lsm(Log Structured Merge Tree)树索引: 
        -   就是将对数据的修改尽量保存在内存中,当达到一定量的时候批量写入磁盘,读取的时候合并磁盘的和内存中的最近修改记录
    -   为什么b+tree 更适合作为索引:`因为btree只是提高了磁盘io的性能,但是并没有解决元素遍历效率低下的问题`
        -   在mysql中每次读取数据都是以页为单位,而每页的数据由操作系统而定,4k或者8k或者16k,但是**磁盘io是昂贵的操作,因而操作系统会预读,既多读相连的几页数据**
        -   b+tree因为数据都在叶子节点,且叶子节点都是有序的,这也就导致了范围查询的时候b+tree比btree要快太多了,也因为如此使得数据页中的数据量尽可能的多,从而减少了io次数
    -   b+tree 索引可以分为聚集索引和辅助索引,`聚集索引的b+tree存放的是行记录数据,而辅助索引存放的是,当通过辅助索引的时候先通过辅助索引找到聚集索引键,然后聚集索引键在聚集索引中找到数据`    
    -   索引设置的疑问点:
        -   为什么`索引要尽可能的小`:因为io次数与b+tree的高度有关,原因在于`每个磁盘页的大小是一定的`,数据项占的空间越小->从而每页占据的数据量也就越多->树的高度也就会越低,这同时也是为什么将数据值都放在叶子节点的缘故
        -   `索引的最左匹配原则`: b+树的key项是复合索引结构,是按照从左到右建立搜索树的,如(a,b,c)索引,优先比较a,然后b,最后c
* 索引的分类:`一般来说除了聚集索引之外就是非聚集索引`
    -   `聚集索引`:表中各行的物理顺序与物理磁盘键值的逻辑(索引)顺序一致,`表中只能包含一个聚集索引,主键列默认为聚集索引`
        ![](https://img-blog.csdnimg.cn/20190212195238171.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0NvZGVyX0pva2Vy,size_16,color_FFFFFF,t_70)
        -   如果`主键被定义了`: 则这个主键就是聚集索引
        -   若主键未定义,则这个表的`唯一非空索引作为聚集索引`
        -   若没有主键也没有唯一非空索引,则会自动生成一个`隐藏的主键,会自动递增`
    -   `非聚集索引,也可以称为辅助索引`:非聚集索引是指,叶子节点上存放的不是完整的记录,而是指向某个数据块的值,`如有username作为非聚集索引,则叶子节点还指向了聚集索引中username为xxx的一个数据块,同时还包含了主键`            
        ![](https://img-blog.csdnimg.cn/20190212195306627.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0NvZGVyX0pva2Vy,size_16,color_FFFFFF,t_70)
        -   普通索引(name): 仅用户加速查找
        -   唯一索引: 
            -   唯一索引: unique(uuid)    `加速查找且约束:唯一`
        -   联合索引(组合索引,覆盖索引):
            -   联合主键索引:(id,name)
            -   联合唯一索引:(uuid,name)
            -   联合普通索引:(name,asd)
        -   全文索引: 用于搜索一篇很长的文章的时候
    -   关于索引的注意点:
        -   二次查询问题:从非聚集索引定义可知,`叶子节点有指向索引值的指针`,因而如果查询列中还有其他的属性,则会发生二次查询(通过查询出来的主键再查询数据),也就是说会访问数据表
            -   如: select name,score from user where name="qwe"的时候,会发生二次查询,因为只有name是个索引
            -   避免的方法就是通过`组合索引`: 为score也创建索引,这样就是覆盖索引了,不会二次查询
    -   关于建立聚集索引还是非聚集索引的建议:
    
|动作 | 是否推荐使用聚簇索引 |是否推荐使用非聚簇索引 |
| ------ | ------ | ------ |
|列经常被分组排序 | 是 | 是 |
|返回某范围内的数据|是|是|
|一个或极少不同值| 否|否|
|小数目的不同值| 是|否|
|大数目的不同值|否|是|
|频繁更新的列|否|是|
|外键列|是|是|
|主键列|是|是|
|频繁修改索引列|否|是
**判断是否要用聚集索引其实只需要判断这个属性列的值是否频繁改变,如果频繁改变不推荐,因为需要树中移动和呼唤**

-   关于索引的建议:
    -   `重复值少的建立索引`
    -   `Where字句条件频繁的建立索引`
    -   `在有原生函数:distint,min,max,order by,group by操作的列和join连接列建立索引`
    -   `不要用select *,或者尽量少用查询的列不仅仅包含索引`:这样会导致二次查询,从而会访问数据表,因而有可能会造成全表扫描
        
        
        
# SQL语句编写的建议

* 不要使用 select * 
    -   `数据太多`:* 会返回所有的业务字段
    -   `性能可能低下`:因为一个业务表不可能为所有的字段都做索引,而非聚集索引的叶子节点存放的是`指向当前索引值的指针`,而如果想再得到其他的数据就需要二次查询,因而性能会降低
    
* 不要返回大批量的数据,让业务做筛选
* 擅用,慎用索引   
* 复合索引的最左匹配原则
        
# MySQL的存储引擎
* InnoDB(默认引擎)
    -   支持事务
    -   最小锁为`行级锁`,支持外键
    -   聚集索引:叶子节点存放数据,非聚集索引:叶子节点不存放数据,存放指向数据的指针
    (图在上面)

* MyISAM
    -   不支持事务,但是每次操作都是原子性的
    -   最小锁为`表级锁`
    -   每个表都会在磁盘中存储三个文件:`.frm:存储表的定义`,`.MYD:存储数据`,`.MYI:存储索引`;既myisam数据文件,索引文件,表文件分开存储
    -   聚集索引:类似于InnoDB的非聚集索引,叶子节点存放指向数据的指针,但是这个数据`是全部数据`,非聚集索引与聚集索引一样,但是非聚集索引不用保证一致性
    * 聚集索引结构如下:![](https://img-blog.csdnimg.cn/20190213202356422.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0NvZGVyX0pva2Vy,size_16,color_FFFFFF,t_70)
    - 非聚集索引结构:![](https://img-blog.csdnimg.cn/20190213202431447.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0NvZGVyX0pva2Vy,size_16,color_FFFFFF,t_70)
- MyISAM和InnoDB,左侧为innodb的索引查找过程:![](https://img-blog.csdnimg.cn/20190213203034518.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0NvZGVyX0pva2Vy,size_16,color_FFFFFF,t_70)
* Memory
    
* 存储引擎总结:![](https://img-blog.csdnimg.cn/20190213154835156.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0NvZGVyX0pva2Vy,size_16,color_FFFFFF,t_70)

# MySQL的事务:    

#### 锁
* MySQL中的锁可以分为两类:
    -   `共享锁`: 既可以认为是读锁,读锁共享
        -   select name from user lock in share mode
    -   `排它锁`: 既可以认为是写锁
        -   `行锁`: recordLock 锁定之后就不能再操作了 
        -   `表锁`: tableLock 对整个表锁定
        -   `间隙锁`: gapLock 既对复合记录的所有行数加锁,这种锁也就彻底避免了幻读

#### 事务之前需要了解2pl(Two-Phase Locking)

* 2pl: Two-Phase Locking ;既2阶段加锁,与Java类似,对于lock需要`加锁和解锁`,具体如下图所示:
![](https://img-blog.csdnimg.cn/20190213232452824.jpg?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0NvZGVyX0pva2Vy,size_16,color_FFFFFF,t_70)
`2PL就是将加锁/解锁分为两个完全不相交的阶段。加锁阶段：只加锁，不放锁。解锁阶段：只放锁，不加锁。`
#### 事务之前需要先了解:MVCC
* 什么是MVCC: MVCC既多版本控制,读不加锁,与MVCC相对的就是基于锁的并发控制
    -   在MVCC中,读可以分为两类:`快照读(snapshot read)`和`当前读(current read)`
        -   快照读(`一致性读`): 只会查询当前版本号<=当前版本号的数据行(`既读取到的数据要么已经commit了的,要么就是当前事务的数据`)
            -   定义: `指的是读取到的是可见版本,可能是历史版本`
            -   默认的select 操作是快照读,不加锁
                -   当insert的时候: `保存当前事务版本号的行作为行的行创建版本号`:既当前事务的版本号作为这条记录的版本号
                -   update: 会产生两条数据;然后会执行2个操作: `将当前事务的版本号作为新的记录的版本号`,`将当前版本号作为老数据的要更新的版本号`
                -   delete: 同理`也是用当前版本号的行作为删除记录`
        -   当前读: 
            -   定义: `读取到的是最新的记录,最新的,最新的并返回,会加上锁`
            - select * from table where ? lock in share mode;
            - select * from table where ? for update;
            - insert into table values (…);
            - update table set ? where ?;
            - delete from table where ?;   
            - **注意这些使用到的都是当前读,这里可能会有疑惑,明明还有写操作,这里的读是指定位这条数据记录行的时候是当前读**,除了in share mode使用的是共享锁,其他的都是默认加的排它锁,流程大致如下图所示:
            -   ![](https://img-blog.csdnimg.cn/20190213220121775.jpg?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0NvZGVyX0pva2Vy,size_16,color_FFFFFF,t_70)可以发现是先where定位返回然后锁住的,
            并且是一条一条交互的
        -   `总结`: **注意,这里的读不是仅仅的指select,这里的读是每个操作都要用到的,无论是select,update,delete,insert,都要用到的,这里的读特指定位到数据的之后的操作**
-   什么是事务:事务是指一段连续的操作,`要么全部成功,要么全部失败`
-   事务的特性:ACID
    -   A: Atomic: 原子性,要么成功,要么失败
    -   C: Cosistent: 一致性,既事务开始前和结束后数据只有预期的变化
    -   I: Isolation: 隔离性,指的不是不同的事务不会产生影响
    -   Durable: 持久化: 既持久化到数据库中
-   事务的并发:
    -   `脏读`: 指的是读取到其他事务`失败前的数据,注意其他事务回滚导致读取到的数据其实是失败的数据`
    -   `不可重复读`: 读取到其他事务提交`修改之后的数据,注意是修改,与原来读的数据不一致,再注意一点:是同一事务第二次读取,第二次,第二次`例如：事务T1读取某一数据，事务T2读取并修改了该数据，T1为了对读取值进行检验而再次读取该数据，便得到了不同的结果。

    -   `幻读`: 指的是读取到其他事务提交的`新增的数据,注意是添加,添加,与原来读取到的数目不一致`
-   事务的隔离级别:
    -   `未提交读`: 事务中的修改，即使没有提交，对其他会话也是可见的,`啥并发问题都不能解决,只能解决更新丢失`
    -   `提交读`: 保证了一个事务如果没有完全成功（未执行commit），事务中的操作对其他会话是不可见的 。,`因而能够解决不可重复读`
        -   原理: 快照度不考虑(因为没必要,不会加锁),当使用`当前读的时候,会对读到的记录加锁(行锁),如下面的那幅图`,但是`第二次读的时候,第一次和第二次这个间隔不能保证`
        其实就想volatile 的int变量一样,+1是原子性的,但是++不是原子性的一样,`因而还是会存在幻读的情况`
    -   `重复读`: 一个事务中多次执行统一读SQL,返回结果一样
        -   原理: 快照读同样忽略,当前读的时候,`读取到的记录加锁(行锁),以及当where条件是一个范围的时候,会加间隙锁`,如:delete from t1 where id = 10 就会在符合记录之间加锁,如果是范围的话(where id >10 and id<20 )也一样会加间隙锁(10与20之间),如下图所示:
        `既通过间隙锁+写锁的形式消除幻读`
        ![](https://img-blog.csdnimg.cn/20190213225332403.jpg?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0NvZGVyX0pva2Vy,size_16,color_FFFFFF,t_70)
    -   `串行读`: 隔离级别最高的,在RC,RR隔离级别中,select默认都是快照读,而在这个隔离级别,默认是当前读
    
    -   总结: 在MySQL/InnoDB中所谓的读不加锁,其实是针对隔离级别而言的,Serialize隔离级别所有的读操作都是需要加锁的
# 实例分析:

#### 表:有属性列: id,name 
#### sql语句为select * from t1 where id = 10;

**RC级别**
1. `id为主键,且rc隔离级别`: 走聚簇索引定位之后之后`加行锁`
2. `id为unique,name为主键,rc级别`:先走非聚簇索引id=10之后,得到name主键,然后走聚簇索引定位到name匹配的记录,所以`会加2个锁,一个是非聚簇索引上的锁,一个则是聚簇索引上的记录`
3. `id非唯一索引,name为主键,rc级别`: 与上述类似,但是存在多条记录的可能,因此会对所有服务的加行锁(注意是行锁)
4. `id无索引,级别为rc`:会上`表锁`,但是会有优化,既mysql会对不匹配的记录解锁(违背了2pl)
<br>**RR级别**
5. `id为主键`: 与1一样,聚簇中加行锁
6. `id唯一索引`:与2一样,对满足的记录加行锁 
7. `id非唯一索引,name为注解,rr级别`: 与3类似,但是加的锁不同,加的是`间隙锁`,对范围加锁
8. `id非索引`: 与4类似,加表锁,`同时会加间隙锁`,效率非常低下,但是myslq也会优化,对于不满足记录的条件会自动解锁(同样违背2pl)


# 疑问点:

* `为什么要最左匹配原则`: 假设有表复合索引(name,cid)
    -   首先:什么是最左匹配原则,`最左匹配原则是指where条件中,必须存在复合索引中的第一个索引`,至于是不是第一个条件,其实不是必然,但是最好还是放在第一个,因为有的sql查询器可能并不会优化(使之成为第一个条件)
    -   提出这个疑问之前我们需要明白:**通过索引查询之所以很快,是因为会通过搜索树维护索引,使得索引是有序的,是有序的,是有序的**
    -   对于索引而言,`单个索引结构上都是有序的`,因而能快速的查找到
    -   对于复合索引而言,则在内存中索引的结构会如下:(大概通过如此建立)
    -   ![](https://img-blog.csdnimg.cn/20190212214601348.jpg),**从左到右建立维护索引树**,既先通过name排序,name:abcdef排好序之后,再对cid进行排序,`也就是说右边的总是基于左边的进行排序,只有最左边的完全的肯定有序的`,这里就可以回答为什么要最左匹配了,我们从图中可以发现,只有当左侧等值(c)的时候,右侧才是有序的,所以这就是最左匹配的缘故
        -   因而在这种复合表索引的情况下,有如下的sql写法
        -   where顺序是name,cid :先匹配name,再匹配cid,这时候是完美走索引的,name是有序的,而cid在name的基础下也是有序的,
        -   where只有cid: 在这种情况下,当explain的时候发现也会走索引`但是这种索引很慢,是遍历索引来查找的`

* `何时会使用全表扫描`:
    -   首先: 什么是全表扫描:既对数据库中所有的记录一条一条匹配
    -   `sql编写不规范`:
        -   `无索引的查询`
        -   `使用null作为判断条件` select name from user where `age=null`
        -   `左模糊查询` select age from user where name `like %qwe%` 或者是`like %qwe`
            -   建议使用 `like qwe%` 右模糊查询
        -   `使用or作为条件`:select name from user where age=1 or age =2
            -   建议使用 union 
        -   `使用in时` select name from user where age in (1,2,3)
            -   如果范围是连续的话,使用`between`
            -   如果是嵌套查询的话,可以使用 `exist`代替:但是适用于条件少的情况: select name from user where id in (select id from user where age =3)------>select name from user where id exist(select id from user where name=3) 
            -   同理not in时使用not exist即可
        -   `使用!=或者<>时`
            -   推荐使用<,<=,>,>=,between等
        
        -   `当=号左边有计算的时候`:select name from user where age+1>3
         
        -   `当使用参数作为查询条件的时候`:select name from user where age=@age 
            -   原因: 不确定性,
    -   `数据量太少的时候,sql查询器认为还不如全表查询块`
    
* `如何避免幻读,为什么隔离级别为RC(提交读)无法避免幻读`:
    -   首先: 幻读是指同一个事务中读取到的数据量不一致,既可以认为是`范围中插入了数据`
    -   如何避免呢,其实在mysql中幻读已经`永远不会发生了`,
        -   MVCC机制:默认的select会使用快照读(既一次性读),只能读取到事务版本号<=当前事务版本号的记录
        -   而在当前读中(insert,delete,update)则是通过间隙锁来实现的(既对这个范围加锁)
        -   还有一个很关键的因素在于:`2pl`,既两阶段加锁,加锁和解锁互相隔离,这也就使得间隙锁是事务结束止之后再解锁
    -   RC为什么无法避免呢,究其原因在于:加的锁的姿势不对,应该加的是间隙锁而不是行锁    
        