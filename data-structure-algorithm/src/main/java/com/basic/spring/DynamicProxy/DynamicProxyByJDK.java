package com.basic.spring.DynamicProxy;

import lombok.Data;

import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;

/**
 * @author joker
 * @When
 * @Description JDK动态代理的实现
 * @Detail
 * @date 创建时间：2019-01-27 13:25
 */
/*
    JDK动态代理的实现原理是指:常见的代理模式,既然代理类实现了接口的同时,也持有原先实现类的对象,因而
    能实现对实现类的织入


    JDK动态代理的核心是InvocationHandler这个接口,注意,因为JDK代理是面向接口的
    所以这是个接口而非具体的实现类
    1. 自定义一个自己的handler,实现InvocationHandler接口
    2. 重写invoke方法,invoke方法中的参数意义分别为:
        2.1 Object target:被代理的对象,所以自定义的handler中要定义一个对象来接收
        2.2 Method:方法名称,既会对哪些方法进行动态代理
        2.3 Object[] args:指的是调用时候的参数
    3. 添加自己的切面逻辑
    4. 然后通过第二个参数的method.invoke(proxy,args)实现对原来方法的调用,返回的是一个object对象,直接返回即可
    也可对返回的result对象进行进一步处理,就是aop中的after
    5. 最后一步,则是在调用处声明的时候: 通过Proxy.newProxyInstance来实现的,第一个参数是类加载器,第二个参数是
    这个类实现的接口,第三个参数则是自定的handler


    注意点:
        1. 这个proxy对象必须是实现了某个接口的对象
        2. 在实现invoke方法的时候,method.invoke调用的是自定义handler中的target对象,而不是方法参数中的proxy对象
 */
public class DynamicProxyByJDK
{
    public static final InvocationHandler DEFAULT_INTERFACE_HANDLER = (proxy, method, args) ->
    {

        return null;
    };

    @Data
    public static class MyInvocationHandler implements InvocationHandler
    {
        public MyInvocationHandler(Object target)
        {
            this.target = target;
        }

        private Object target;

        @Override
        public Object invoke(Object proxy, Method method, Object[] args) throws Throwable
        {
            System.out.println("before ");
            Object result = method.invoke(target, args);
            System.out.println("after");
            if (result instanceof String)
            {
                return "-----" + result + "+++++";
            }
            return result;
        }
    }
}
