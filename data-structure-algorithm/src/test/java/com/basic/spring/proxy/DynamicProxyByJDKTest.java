package com.basic.spring.proxy;

import com.basic.spring.DynamicProxy.DynamicProxyByJDK;
import org.junit.Test;

import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Proxy;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-01-27 13:45
 */
public class DynamicProxyByJDKTest
{

    @Test
    public void testJDKProxy()
    {
        DynamicProxyByJDK.MyInvocationHandler handler = new DynamicProxyByJDK.MyInvocationHandler(IProxyService.JDKProxyTestServiceImpl);
        IProxyService jdkProxyInterface = (IProxyService)
                Proxy.newProxyInstance(
                        IProxyService.JDKProxyTestServiceImpl.getClass().getClassLoader(),
                        IProxyService.JDKProxyTestServiceImpl.getClass().getInterfaces(), handler);
        String joker = jdkProxyInterface.call("joker");
        System.out.println(joker);
    }
}
