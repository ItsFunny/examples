package com.charlie.controller;

import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-01-09 14:32
 */
@RestController
public class TestController
{

    @RequestMapping("/test")
    public String test()
    {
        return "test";
    }

    @RequestMapping("/role")
    public String role()
    {
        return "role";
    }


    @RequestMapping("/test/token")
    public String testToken()
    {
        return "test-token";
    }

//    @GetMapping("/login")
//    public String login()
//    {
//
//        return "login";
//    }
}
