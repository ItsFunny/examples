package com.demo.sso.resource1;

import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.builder.SpringApplicationBuilder;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：201   9-01-12 12:08
 */

@SpringBootApplication
public class Resource1Application
{

    public static void main(String[] args)
    {
        new SpringApplicationBuilder(Resource1Application.class).run(args);
    }
}
