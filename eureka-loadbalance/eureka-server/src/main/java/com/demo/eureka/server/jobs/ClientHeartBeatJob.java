package com.demo.eureka.server.jobs;

import com.demo.eureka.server.configuration.EurekaServerConfigurationProperties;
import org.springframework.aop.framework.ProxyFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.cglib.proxy.Enhancer;
import org.springframework.context.ApplicationContext;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;
import org.springframework.web.context.ServletContextAware;
import org.springframework.web.context.WebApplicationContext;
import org.springframework.web.context.support.WebApplicationContextUtils;

import javax.servlet.ServletContext;
import java.util.HashMap;

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
        private ServletContext c;
        private WebApplicationContext context;
    @Scheduled(cron = "")
    public void run()
    {
        ProxyFactory factory=new ProxyFactory();
    }
}
