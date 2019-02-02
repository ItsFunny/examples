package com.demo.eureka.server.jobs;

import com.demo.eureka.server.configuration.EurekaServerConfigurationProperties;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-02-02 10:20
 */
@Component
public class ClientHeartBeatJob implements Runnable
{
    @Autowired
    private EurekaServerConfigurationProperties properties;

    @Scheduled(cron = "")
    public void run()
    {
    }
}
