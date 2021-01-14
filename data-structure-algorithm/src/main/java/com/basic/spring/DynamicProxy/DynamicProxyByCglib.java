package com.basic.spring.DynamicProxy;

import org.springframework.cglib.proxy.MethodInterceptor;
import org.springframework.cglib.proxy.MethodProxy;

import java.lang.reflect.Method;

/**
 * @author joker
 * @When
 * @Description 通过cglib实现动态代理
 * @Detail
 * @date 创建时间：2019-01-27 15:03
 */

/*
    cglib实现代理的原理是:
    通过asm字节码操作的框架,在内存中动态生成代理类的子类,并复写非final方法,比jdk具有更高的效率
    cglib的主要角色:
    net.sf.cglib.proxy.Enhancer:主要的增强类  是一个类生成器,用于通过字节码技术生成子类
    net.sf.cglib.proxy.MethodInterceptor: 主要的拦截接口,我们自主实现这个接口
    net.sf.cglib.proxy.MethodProxy – JDK的java.lang.reflect.Method类的代理类，可以方便的实现对源对象方法的调用,如使用：
    Object o = methodProxy.invokeSuper(proxy, args);//就算第一个参数是被代理对象，也不会出现死循环的问题。

    步骤:
        1. 实现MethodInterceptor接口,实现自定义的CglibProxy
     实现类已经自定义完毕,我们只需要调用即可
     调用步骤:
        1. 创建Enhancer对象
        2. 设置superClass (这个步骤表明这个子类是哪个的父类)
        3. 设置回调函数setCallBack(MethodInterceptor是CallBack的子接口)
        4. 通过enhancer对象生成即可:create

     注意点:
        1. 父类不能是由final所修饰
        2. 父类必须提供默认的构造函数(既无参构造函数),且需要定义为visiable,不可以是private的对象
        3. 无法使用lambda表达式,因为是final类型的
 */
public class DynamicProxyByCglib
{
    public static MethodInterceptor CglibProxyTest = (target, method, args, methodProxy) ->
    {
        System.out.println("before call ");
        String superName = methodProxy.getSuperName();
        System.out.println(superName);
        Object result = methodProxy.invokeSuper(target, args);
        System.out.println("after call");
        if (result instanceof String)
        {
            return "-----" + result + "+++++++++";
        }
        return result;
    };
}
