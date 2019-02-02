package com.demo.common.service;

import com.demo.common.configuration.EurekaClientRequest;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-02-02 10:49
 */
public interface RegisterCallback
{

    void callBack(EurekaClientRequest request);


    RegisterCallback DEFAULT_CALLBACK = (request) ->
    {
        // 最好是发送请求给原先的服务器,common中需要也添加controller,让client也自启动注册或者啥的
        System.out.println(request.getHostName() + "注册成功");
    };
}
