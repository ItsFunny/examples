package com.demo;

import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.builder.SpringApplicationBuilder;

/**
 * Hello world!
 */
@SpringBootApplication
public class BigFileApplication
{
    public static void main(String[] args)
    {
        new SpringApplicationBuilder(BigFileApplication.class).run(args);
    }
}
