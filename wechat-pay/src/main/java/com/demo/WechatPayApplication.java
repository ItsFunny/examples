package com.demo;

import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.builder.SpringApplicationBuilder;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-03-17 15:09
 */
@SpringBootApplication
public class WechatPayApplication
{
    public static void main(String[] args)
    {
        new SpringApplicationBuilder(WechatPayApplication.class).run(args);
    }
}
