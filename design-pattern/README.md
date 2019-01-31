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