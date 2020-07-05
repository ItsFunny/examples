package com.demo.sso.resouce2.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.ResponseBody;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-01-13 15:40
 */
@Controller
public class ResourceController
{
    @ResponseBody
    @GetMapping(value = "/test2")
    public String test2()
    {
        return "ok 2";
    }

}
