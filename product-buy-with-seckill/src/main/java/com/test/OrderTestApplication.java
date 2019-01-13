package com.test;

import java.util.HashMap;

import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.builder.SpringApplicationBuilder;
import org.springframework.context.annotation.EnableAspectJAutoProxy;

/**
 * Hello world!
 *
 */
@SpringBootApplication
@EnableAspectJAutoProxy
public class OrderTestApplication 
{
    public static void main( String[] args )
    {
    	new SpringApplicationBuilder(OrderTestApplication.class).run(args);
    }
}
