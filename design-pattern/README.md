# 设计模式
---

*   `FINISH`   单例模式
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
      
           