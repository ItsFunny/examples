package com.demo.eureka.server.controller;

import com.demo.common.configuration.EurekaClientRequest;
import com.demo.eureka.server.common.CommonHolder;
import com.demo.eureka.server.model.EurekaClient;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.servlet.ModelAndView;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-02-02 10:29
 */
@RestController
public class EurekaController
{
    @RequestMapping(method = RequestMethod.POST, path = "/myeureka/")
    public String registerClient(@RequestBody EurekaClientRequest request)
    {
        EurekaClient client = new EurekaClient();
        client.from(request);
        CommonHolder.CLIENTS.add(client);
        request.getCallback().callBack(request);
        return "ok";
    }

    @GetMapping(value = "/test")
    public String test()
    {
        return "ok";
    }

    @GetMapping(value = "/clients")
    public ModelAndView testView()
    {
        ModelAndView modelAndView = new ModelAndView("clients");
        modelAndView.addObject("clients", CommonHolder.CLIENTS);
        return modelAndView;
    }
}
