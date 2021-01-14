package com.basic.spring.proxy;

import com.basic.spring.DynamicProxy.DynamicProxyByCglib;
import org.junit.Test;
import org.springframework.cglib.proxy.Enhancer;

import java.util.HashMap;
import java.util.concurrent.ThreadFactory;
import java.util.concurrent.ThreadPoolExecutor;
import java.util.concurrent.locks.ReadWriteLock;
import java.util.concurrent.locks.ReentrantLock;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-01-27 16:00
 */
public class DynamicProxyByCglibTest
{
    // 当使用的是cglib的时候,无法使用lambda表达式,lambda表达式生成的是final类型的,final类型代表无法被继承
    static IProxyService CglibProxyTestServiceImpl = (name) -> name + "____";

    public static class CglibProxyServiceImpl implements IProxyService
    {
        @Override
        public String call(String name)
        {
            return name + "____";
        }

    }

    @Test
    public void testProxyByCglib()
    {
        Enhancer enhancer = new Enhancer();

        enhancer.setSuperclass(CglibProxyServiceImpl.class);

        enhancer.setCallback(DynamicProxyByCglib.CglibProxyTest);
        IProxyService service = (IProxyService) enhancer.create();
        String joker = service.call("joker");
        System.out.println(joker);
    }


}
