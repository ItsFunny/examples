package com.demo.eureka.server.common;

import com.demo.eureka.server.model.EurekaClient;

import java.util.concurrent.LinkedBlockingQueue;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-02-02 10:37
 */
public class CommonHolder
{
    public static LinkedBlockingQueue<EurekaClient> CLIENTS = new LinkedBlockingQueue<EurekaClient>();

}
