package com.demo.eureka.client1.plugins;

import org.springframework.beans.BeansException;
import org.springframework.beans.factory.config.BeanPostProcessor;
import org.springframework.core.Ordered;
import org.springframework.stereotype.Component;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-02-03 18:18
 */
@Component
public class ConfigurationPlugin implements BeanPostProcessor, Ordered
{
    @Override
    public Object postProcessBeforeInitialization(Object bean, String beanName) throws BeansException
    {
        return bean;
    }

    @Override
    public Object postProcessAfterInitialization(Object bean, String beanName) throws BeansException
    {
        return bean;
    }


    @Override
    public int getOrder()
    {
        return -1*Integer.MAX_VALUE;
    }
}
