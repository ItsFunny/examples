package com.demo.sso.server;

import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.builder.SpringApplicationBuilder;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-01-13 11:20
 */
@SpringBootApplication
public class SSOServerApplication
{

    public static void main(String[] args)
    {
        new SpringApplicationBuilder(SSOServerApplication.class).run(args);
    }
}
