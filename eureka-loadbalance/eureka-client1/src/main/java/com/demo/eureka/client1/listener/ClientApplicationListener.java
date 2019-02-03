package com.demo.eureka.client1.listener;

import com.demo.common.annotation.EnableMyEurekaClientAnnotation;
import org.springframework.context.ApplicationContext;
import org.springframework.context.ApplicationListener;
import org.springframework.context.event.ContextRefreshedEvent;
import org.springframework.stereotype.Component;

/**
 * @author joker
 * @When
 * @Description 监听容器:
 * 1. 查看是否有使用@EnableMyEurekaClientAnnotation 这个注解
 * 2. 如果有使用,则获取到默认的配置类EurekaConfiguration
 * 3. 对注解类进行校验
 * @Detail
 * @date 创建时间：2019-02-03 22:55
 */
@Component
public class ClientApplicationListener implements ApplicationListener<ContextRefreshedEvent>
{
    @Override
    public void onApplicationEvent(ContextRefreshedEvent contextRefreshedEvent)
    {
        System.out.println(contextRefreshedEvent);
        ApplicationContext applicationContext = contextRefreshedEvent.getApplicationContext();
        String[] beanNamesForAnnotation = applicationContext.getBeanNamesForAnnotation(EnableMyEurekaClientAnnotation.class);
        for (String beanName: beanNamesForAnnotation)
        {
            System.out.println(beanName);
        }
    }
}
