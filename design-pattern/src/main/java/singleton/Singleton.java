package singleton;

import java.io.ObjectOutputStream;
import java.io.ObjectStreamException;
import java.io.Serializable;

/**
 * @author joker
 * @When
 * @Description 单例模式
 * @Detail
 * @date 创建时间：2019-01-28 15:57
 */
/*
    单例模式的实现有多种
 */
public class Singleton implements Serializable
{

    // 避免因为反射而生成新的对象
    private Singleton()
    {
        // 随便挑选的一个实例,不单指lazy
        if (null != INSTANCE_LAZY)
        {
            throw new RuntimeException("不允许反射生成对象");
        }
    }

    // 避免因为反序列化生成新的对象

    private Object readResolve() throws ObjectStreamException
    {
        return INSTANCE_LAZY;
    }

    // 1.懒汉模式: 既真正调用的时候才会去加载,not thread safe
    // 缺点在于第一次加载的时候会很缓慢
    private static Singleton INSTANCE_LAZY = null;

    public static Singleton LazyGetInstance()
    {
        if (null == INSTANCE_LAZY)
        {
            INSTANCE_LAZY = new Singleton();
        }
        return INSTANCE_LAZY;
    }


    // 2.饿汉模式: 既提前加载数据,not thread safe
    private static Singleton INSTANCE_HUNGARY = new Singleton();

    public static Singleton HungerGetInstance()
    {
        return INSTANCE_HUNGARY;
    }


    // 3.静态内部类的形式

    private static class InnerSingleton
    {
        private static final Singleton INSTANCE_INNER = new Singleton();
    }

    public static Singleton GetByInnerClass()
    {
        return InnerSingleton.INSTANCE_INNER;
    }


    // 4.双重锁机制:配合volatile防止重排序实现
    // 缺点在于还是无法避免反序列化之后获取到的对象是同一个对象
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

    // 5.还有一种则是枚举单例的形式

    public enum SingletonEnum
    {
        INSTANCE_ENUM;

        public void function()
        {

        }
    }

    // 如何避免反序列化造成的对象不同呢
    // 答案是通过实现SerializeAble接口之后


}
