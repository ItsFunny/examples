package com.test;

import java.util.HashMap;
import java.util.Map;

import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.builder.SpringApplicationBuilder;
import org.springframework.context.annotation.EnableAspectJAutoProxy;
import org.springframework.mail.MailParseException;

/**
 * Hello world!
 */
@SpringBootApplication
@EnableAspectJAutoProxy
public class OrderTestApplication
{
    public static void main(String[] args)
    {

        new SpringApplicationBuilder(OrderTestApplication.class).run(args);
        Byte[] bytes = new Byte[]{};
        String s = "";
        Map<String, String> m = new HashMap<>();
        m.put("key", "value");
        System.out.println(m.containsKey("key"));
        m.remove("key");
        Object o = new Object();
        if (o instanceof String)
        {

        }
    }

}
