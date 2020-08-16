package com.demo.log;

import lombok.extern.log4j.Log4j2;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.builder.SpringApplicationBuilder;

/**
 * Hello world!
 */
@SpringBootApplication
@Log4j2
public class App
{
    public static void main(String[] args)
    {
        log.debug("11");
        new SpringApplicationBuilder(App.class).run(args);
    }
}
