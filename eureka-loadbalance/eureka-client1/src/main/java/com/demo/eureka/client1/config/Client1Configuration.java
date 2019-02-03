package com.demo.eureka.client1.config;

import com.demo.common.annotation.EnableMyEurekaClientAnnotation;
import com.demo.common.configuration.EurekaConfiguration;
import com.demo.eureka.client1.listener.ClientApplicationListener;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.InitializingBean;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.boot.web.servlet.ServletListenerRegistrationBean;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;

import javax.annotation.PostConstruct;
import java.util.EventListener;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-02-03 10:28
 */
@Configuration
@Slf4j
@ComponentScan(basePackages = "com.demo.common")
public class Client1Configuration implements InitializingBean
{
    @Autowired
    private EurekaConfiguration configuration;


    @Override
    public void afterPropertiesSet() throws Exception
    {
        log.info("{}", configuration);
    }


//    @Bean
//    public ServletListenerRegistrationBean<EventListener>listenerRegistrationBean()
//    {
//        ServletListenerRegistrationBean<EventListener>servletListenerRegistrationBean=new ServletListenerRegistrationBean<>();
//        servletListenerRegistrationBean.setListener(new ClientApplicationListener());
//        return servletListenerRegistrationBean;
//    }
}
