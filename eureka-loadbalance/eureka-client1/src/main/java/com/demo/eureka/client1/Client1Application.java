package com.demo.eureka.client1;

import com.demo.common.annotation.EnableMyEurekaClientAnnotation;
import com.demo.common.annotation.EnableMyFeignClients;
import org.mybatis.spring.boot.autoconfigure.MybatisAutoConfiguration;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.autoconfigure.jdbc.DataSourceAutoConfiguration;
import org.springframework.boot.builder.SpringApplicationBuilder;
import org.springframework.context.annotation.EnableAspectJAutoProxy;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-02-03 10:07
 */
@SpringBootApplication(exclude = {MybatisAutoConfiguration.class, DataSourceAutoConfiguration.class})
@EnableMyEurekaClientAnnotation
@EnableAspectJAutoProxy
@EnableMyFeignClients(basepackages = "com.demo.eureka")
public class Client1Application
{
    public static void main(String[] args)
    {
        new SpringApplicationBuilder(Client1Application.class).run(args);
    }
}
