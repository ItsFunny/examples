package com.demo.log.controller;

import lombok.extern.log4j.Log4j;
import lombok.extern.log4j.Log4j2;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.UUID;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-08-15 22:09
 */
@RestController
@Log4j2
public class TestController
{
    @RequestMapping("/info")
    public void info()
    {
        log.info("info"+ UUID.randomUUID());
    }

    @RequestMapping("/e")
    public void error()
    {
        log.error("error"+ UUID.randomUUID());
    }
    @RequestMapping("/warn")
    public void warn()
    {
        log.warn("warn"+ UUID.randomUUID());
    }
    @RequestMapping("/debug")
    public void debug()
    {
        log.debug("debug"+ UUID.randomUUID());
    }

}
