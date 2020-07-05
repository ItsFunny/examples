package com.demo.eureka.client1.service;

import com.demo.common.annotation.MyFeignClientAnnotation;
import org.springframework.stereotype.Component;
import org.springframework.web.bind.annotation.RequestMapping;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-02-05 21:24
 */
@Component
@MyFeignClientAnnotation(name = "client1")
public interface IFeignTestService
{
    @RequestMapping(value = "/test")
    String test();
}
