Spring
---
* 因为最近要重回业务的怀抱,因而重新再看一遍Spring的内容刻不容缓

TODO
---
* [Github上更新的较这里勤](https://github.com/ItsFunny/data-structure-algorithm/tree/master/src/main/source)
* 对于一些Spring关键点源码分析
* 对Spring设计上的一些分析


---

## BeanFactory 和ApplicationContext
`额外的延伸点`:
* 关于父子容器的概念:通过HierarchicalBeanFactory接口,可以实现**子容器可以访问父容器,但是父容器不能访问子容器**,controller层是子容器,视图层可以访问业务层和持久层的bean,但是相反则不行,`既controller可以访问dao,service,而dao,service无法访问controller`

#### BeanFactory 
* 内部是一个hashMap实现的缓存器
* `bean的生命周期`
	-	


#### ApplicationContext

#### WebApplicationContext
* 作用: `wip`专门用于web,`允许从web的根目录文件中装载配置文件从而初始化`
* 如何使用:

* 区别:
	-	ApplicationContext是BeanFactory的一个子类(`wip`)
	-	BeanFactory加载bean的时候,是只有get 的时候才会触发加载,而ApplicationContext则是`初始化上下文的时候就加载了所有的bean信息`

## 关于BeanFactory和ApplicationContext的部分总结

---
## Spring配置(专门用于讲解配置):
##### 配置方式:
* 通过xml配置
* 通过注解@Componnet配置
* 通过JavaConfig配置
* 省略groovy的方式
##### `注意`:
* 在非SpringBoot下,如果使用的是JavaConfig的形式,则web.xml中的<context-param>参数需要指定为
```
<context-param>
	<param-name>contextClass</param-name>
	<param-value>org.springframework.web.context.support.AnnotationConfigWebApplicationContext
</context-param>

然后对于原先的contextConfigLocation则是指向具体的类
<context-param>
	<param-name>contextConfigLocation</param-name>
	<param-value>com.test.config.Configuraiton</param-value>
</context-param>
```
当然你不指定也是没关系的,只要通过 `<component-scan basepackages="com.demo"> `这样讲配置类扫描进去也是可以的,但是`注意要与SpringMVC的配置扫描分开哦,不然会加载两次,血的教训`


##### WebApplicationContext的配置:
-	`在web.xml中配置:`,至于作用在上面已经申明了
	
	```		
	<listener> 
				<listener-class>org.springfarmework.web.context.ContextLoadListener</listener-class>
	</listener>
	这个listener可以获取到<context-param>中的名为contextConfigLocation的值,
	至于这个值是可以采用统配符的: classpath:*/spring-*.xml
	如:
	<context-param>
		<param-name>contextConfigLocation</param-name>
		<param-value>/WEB-INF/spring-mvc.xml</param-value>
	</context-param>		

##### 日志功能的配置:
注意:
* 日志的配置必须在Spring的配置之前
* 日志的配置可以使用listener或者是servlet
	-	如果用listener的话,需要将这个listener`放置在contextConfigListener之前`
	-	如果用servlet的话,`load-on-startup要设置为1`,优先级别最高
* 核心参数就是log4jConfigLocation
```
<context-param>
	<param-name>log4jConfigLocation</param-name>
	<param-value>/WEB-INF/log4j.properties</param-value>
</context-param>
```
* listener方式:
```
<listener>
	<listener-class>Log4jConfigListener</listener-class>
</listener>
```
* servlet方式:
```
<servlet>
	<servlet-name>log4jConfigSerlvet</servlet-name>
	<serlvet-class>org.springframework.web.util.Log4jConfigServlet</servet-class>
	<load-on-startup>1</load-on-startup>
</servlet>

```

设计方面
---
* Spring 的设计基本都是一个主接口,然后衍生出一大堆的子接口,或者是抽象实现类,然后在抽象实现类中再定义一些额外的方法,从而使得丰富多样