package com.demo.eureka.server.listeners;

import com.demo.eureka.server.service.EmptyService;
import org.springframework.context.ApplicationContext;
import org.springframework.context.ApplicationListener;
import org.springframework.context.event.ContextRefreshedEvent;
import org.springframework.stereotype.Component;

import java.util.EventListener;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-02-05 16:16
 */

@Component
public class ContextRefreshListener implements ApplicationListener<ContextRefreshedEvent>
{

    @Override
    public void onApplicationEvent(ContextRefreshedEvent contextRefreshedEvent)
    {
        ApplicationContext applicationContext = contextRefreshedEvent.getApplicationContext();
        EmptyService bean = applicationContext.getBean(EmptyService.class);
        System.out.println(bean);
    }
}
