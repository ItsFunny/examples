# 设计模式
---

###   `FINISH`   单例模式
-   `FINISH` 传统的单例模式无法避免因为`反射`和`反序列化`产生新的对象,因而需要额外的措施:
    -   避免反射生成新的对象:
    ```
     private Singleton()
     {  
         if (null!=INSTANCE_LAZY)
         {
             throw new RuntimeException("不允许反射生成对象");
         }
     }
        
    ```
    -   避免因反序列化生成的对象:
        
    ```
      private Object readResolve() throws ObjectStreamException
      {
          return INSTANCE_LAZY;
      }  
    ```
    
-   `FINISH` 懒汉模式
    -   缺点: 线程不安全
    
    ```
    private static Singleton INSTANCE_LAZY = null;
    
    public static Singleton LazyGetInstance()
    {
        if (null == INSTANCE_LAZY)
        {
            INSTANCE_LAZY = new Singleton();
        }
        return INSTANCE_LAZY;
    }
    ```
-   `FINISH` 恶汉模式 
    -   缺点: 初次使用既加载
    
    ```
    private static Singleton INSTANCE_HUNGARY = new Singleton();
    
    public static Singleton HungerGetInstance()
    {
        return INSTANCE_HUNGARY;
    }
    ```
    
-   `FINISH` 双重锁模式   
    -   代码冗余:
    
    ```
    private static volatile Singleton INSTANCE_DOUBLE_LOCK = null;
    
    public static Singleton DoubleLockGetInstance()
    {
        if (null == INSTANCE_DOUBLE_LOCK)
        {
            synchronized (Singleton.class)
            {
                if (null == INSTANCE_DOUBLE_LOCK)
                {
                    INSTANCE_DOUBLE_LOCK = new Singleton();
                }
            }
        }
        return INSTANCE_DOUBLE_LOCK;
    }
    
    ```
    
-   `FINISH` 枚举单例

    ```
    public enum SingletonEnum
    {
        INSTANCE_ENUM;

        public void function()
        {

        }
    }
        
     ```    
-   `FINISH` 静态内部类的形式
    -   与懒汉,饿汉模式的区别在于,静态内部类默认是线程安全的,不需要加锁(因为类加载是线程安全的)
    ```
    private static class InnerSingleton
    {
        private static final Singleton INSTANCE_INNER = new Singleton();
    }

    public static Singleton GetByInnerClass()
    {
        return InnerSingleton.INSTANCE_INNER;
    }

    ```
 
 ###    `FINISH`    模板模式
 
* 模板模式有三个角色:
* 1.service接口,用于提供具体的方法
* 2.抽象类,抽象类会复写接口中的方法,将相同的部分代码在抽象类中实现,这层对外部而言是不可见的
* 3.具体实现类,继承抽象类,实现具体的细节
* 如下所示IObjectService用于生成各种对象
* AbstractObjectService对于一个参数都会有前置校验,这是相同的部分,因而可以抽出来放到抽象类中
* 至于具体怎么找,找什么交给具体的实现类来写 

```
interface IObjectService
{
    void findByName(String name);
}
abstract class AbstractObjectService implements IObjectService
{
    protected abstract String doFindObject(String user);
    public void findByName(String obj)
    {
        if (null == obj || "".equals(obj))
        {
            return;
        }
    }
}
class UserServiceImpl extends AbstractObjectService
{
    protected String doFindObject(String user)
    {
        return "find user:" + user;
    }
}

class AnimalServiceImpl extends AbstractObjectService
{
    protected String doFindObject(String animal)
    {
        return "find animal:" + animal;
    }
}
```

### `FINISH`    责任链模式

* 理解中的责任链模式的角色有:
    -   接口: 用于分发具体外抛给外部的功能
    -   抽象handler: 复写接口中的方法,并且这个抽象类是具有接口变量的引用充当nextHandler
    -   具体的handler: 不同对象的handler具有不同的执行策略,执行之前会判断对象是否该交由这个handler处理
    -   被执行的对象: 顾名思义,是被handler所执行的对象,同抽象handler会有一个唯一的标识标识这个应该由谁处理
    
    ```
    
    interface IUserService
    {
        void login(UserBO userBO);
    }
    
    class UserBO
    {
        byte level;
        String userName;
        String passWord;
    
        public UserBO(byte level, String userName, String passWord)
        {
            this.level = level;
            this.userName = userName;
            this.passWord = passWord;
        }
    
    
    }
    
    
    abstract class AbstractUserServiceHandler implements IUserService
    {
        protected byte level;    // 用于判断不同的实现类处理不同的角色
        protected IUserService nextHandler;
    
        protected abstract void doLogin(String userName, String password);
    
        protected AbstractUserServiceHandler(byte level)
        {
            this.level = level;
        }
    
        public void login(UserBO userBO)
        {
            if (userBO.level == this.level)
            {
                this.doLogin(userBO.userName, userBO.passWord);
            } else if (null != this.nextHandler)
            {
                this.nextHandler.login(userBO);
            } else
            {
                throw new RuntimeException("no concrete handler to handle the request");
            }
        }
    }
    
    class NormalUserHandler extends AbstractUserServiceHandler
    {
    
    
        protected NormalUserHandler(byte level)
        {
            super(level);
            this.level = level;
        }
    
        protected void doLogin(String userName, String password)
        {
            System.out.println("normal user login:" + userName);
        }
    }
    
    class VIPUserHandler extends AbstractUserServiceHandler
    {
    
        protected VIPUserHandler(byte level)
        {
            super(level);
            this.level = level;
        }
    
        protected void doLogin(String userName, String password)
        {
            System.out.println("vipUserLogin:" + userName);
        }
    }
    
    ```
    
### `FINISH` 策略模式

* 自我理解中的策略模式具有的2个角色:
    -   策略接口: 提供外部使用的方法
    -   执行者: 不同的策略有不同的执行者
    
```

interface FileStrategy
{
    boolean upload(File file);
}

public class StrategyPattern
{
    public static FileStrategy FTPFileStrategy = (file) ->
    {
        System.out.println("这是文件策略中的ftp策略,文件会上传到远程的ftp服务器");
        return true;
    };
    public static FileStrategy LocalFileStrategy = (file) ->
    {
        System.out.println("本地策略:文件上传到本地");
        return true;
    };
}

```

### `FINISH`   工厂模式

 

-   静态工厂模式(简单工厂模式)
    * 自我理解中的静态工厂模式: `与生活中的相同,工厂是个大杂烩,啥都生成,而简单模式也类似,生产的都是一些相似,但又不同的对象`
        - 抽象接口: 用于定义被生产者相同的特征
        - 工厂者:  用于生成不同的对象,是个大杂烩,都经过他生产
    ```
     public static IHumanService CreateHuman(String type)
        {
            if ("student".equals(type))
            {
                return new Student("joker");
            } else if ("teacher".equals(type))
            {
                return new Teacher("clown");
            } else
            {
                return null;
            }
        }

    ```
-   普通工厂模式
    * 自我理解中的普通工厂模式: `把工厂的职责细化了,这个工厂就专门生产这个,另外一个工厂专门生产其他的,各司其职`,**但是被生产的对象又是有相似的地方的**
    -   被生产对象的公共抽象接口: 用于定义被生产对象的共性
    -   工厂方法接口: 工厂之间的共性
    -   不同的具体的工厂: 一个工厂生产具体的个体
    ![](https://img-blog.csdnimg.cn/20190202075733754.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0NvZGVyX0pva2Vy,size_16,color_FFFFFF,t_70)
    ```
    interface IHumanFactoryService
    {
        IHumanService createHuman();
    }
    
    class StudentFactory implements IHumanFactoryService
    {
    
        @Override
        public IHumanService createHuman()
        {
            return new Student("joker");
        }
    }
    
    class TeacherFactory implements IHumanFactoryService
    {
    
        @Override
        public IHumanService createHuman()
        {
            return new Teacher("clown");
        }
    }
    ```
-   抽象工厂模式
    * 自我理解: 抽象工厂模式是基于**产品族的**,注意是一个族,也就是说创建的对象还是有关联的,如键鼠套装,键盘与鼠标是关联的,但是生产键鼠的有赛睿,有雷蛇,也有其他,而这些就是
    工厂的概念,键鼠则是抽象概念; 附一张之前的图:
    ![](https://img-blog.csdn.net/20180610104815911?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0NvZGVyX0pva2Vy/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)
    -   被生产对象的抽象接口
    -   被生产对象的抽象接口的具体的实现类
    -   抽象工厂: 用于定义统一的共性
    -   抽象工厂具体实现类: 实现自个的个性
    ```
    interface  IHuman
    {
    
    }
    interface IStudent extends IHuman
    {
    
    }
    interface ITeacher extends IHuman
    {
    
    }
    
    class ZheJiangStudent implements IStudent
    {
        String name;
    
        public ZheJiangStudent(String name)
        {
            this.name = name;
        }
    }
    
    class BeiJingStudent implements IStudent
    {
        String name;
    
        public BeiJingStudent(String name)
        {
            this.name = name;
        }
    }
    
    class ZheJiangTeacher implements ITeacher
    {
        String name;
    
        public ZheJiangTeacher(String name)
        {
            this.name = name;
        }
    }
    
    class BeiJingTeacher implements ITeacher
    {
        String name;
    
        public BeiJingTeacher(String name)
        {
            this.name = name;
        }
    }
    
    abstract class AbstractHumanServiceFactory
    {
        public abstract IStudent createStudent();
    
        public abstract ITeacher createTeacher();
    }
    
    class ZheJiangHumanFactory extends AbstractHumanServiceFactory
    {
    
        @Override
        public IStudent createStudent()
        {
            return new ZheJiangStudent("joker");
        }
    
        @Override
        public ITeacher createTeacher()
        {
            return new ZheJiangTeacher("joker");
        }
    }
    
    class BeijingHumanFactory extends AbstractHumanServiceFactory
    {
    
        @Override
        public IStudent createStudent()
        {
            return new BeiJingStudent("clown");
        }
        @Override
        public ITeacher createTeacher()
        {
            return new BeiJingTeacher("clown");
        }
    
    }
    ```