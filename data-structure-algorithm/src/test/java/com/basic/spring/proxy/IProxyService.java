package com.basic.spring.proxy;


public interface IProxyService
{
    String call(String name);


    IProxyService JDKProxyTestServiceImpl = (name) ->
    {
        System.out.println("call the name:" + name + " and the name should be " + name + "____");
        return name + "_____";
    };

}
