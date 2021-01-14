package com.demo.sso.resource1.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.servlet.ModelAndView;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-01-12 19:22
 */
@Controller
public class ResourceController
{
    @ResponseBody
    @GetMapping(value = "/test")
    public String test()
    {
        return "ok";
    }

    @ResponseBody
    @GetMapping(value = "/url")
    public String url(HttpServletRequest request, HttpServletResponse response)
    {

        String requestURI = request.getRequestURI();
        StringBuffer requestURL = request.getRequestURL();
        System.out.println(requestURI);
        System.out.println(requestURL);
        return "ok";


    }

}
