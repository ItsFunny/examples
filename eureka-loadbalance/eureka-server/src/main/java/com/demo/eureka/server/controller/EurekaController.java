package com.demo.eureka.server.controller;

import com.demo.common.configuration.EurekaClientRequest;
import com.demo.eureka.server.common.CommonHolder;
import com.demo.eureka.server.model.EurekaClient;
import com.demo.eureka.server.service.EmptyService;
import com.joker.library.model.HttpClientResult;
import com.joker.library.utils.HttpUtils;
import org.apache.http.client.utils.HttpClientUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.client.RestTemplate;
import org.springframework.web.servlet.ModelAndView;

import javax.servlet.http.HttpServletRequest;
import java.util.Map;

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
    @Autowired
    private EmptyService emptyService;

    @RequestMapping(method = RequestMethod.POST, path = "/myeureka/")
    public String registerClient(@RequestBody(required = false) EurekaClientRequest request)
    {
        EurekaClient client = new EurekaClient();
        client.from(request);
        CommonHolder.CLIENTS.add(client);
        request.getCallback().callBack(request);
        return "ok";
    }

    @RequestMapping(value = "/call/{clientName}")
    public HttpClientResult callResult(@PathVariable String clientName, @RequestParam Map<String, String> params, HttpServletRequest request)
    {
        // FIXME 如果服务响应超时
        String method = request.getMethod();
        HttpClientResult result = null;
        for (EurekaClient client : CommonHolder.CLIENTS)
        {
            if (!client.getClientName().equals(clientName)) continue;
            String clientUrl = client.getClientUrl();
            try
            {
                if (method.equals("get"))
                {

                    result = HttpUtils.doGet(clientUrl, params);
                } else
                {
                    result = HttpUtils.doPost(clientUrl, params);
                }
            } catch (Exception e)
            {
                // 返回失败
            }

        }
        return HttpClientResult.buildSuccess();
    }

    @GetMapping(value = "/test")
    public String test(HttpServletRequest request)
    {
        String url = request.getScheme() + "://" + request.getServerName()
                + ":" + request.getServerPort();
        return url;
    }

    @Autowired
    private RestTemplate restTemplate;

    @GetMapping(value = "/clients")
    public ModelAndView testView()
    {
        ModelAndView modelAndView = new ModelAndView("clients");
        modelAndView.addObject("clients", CommonHolder.CLIENTS);
        return modelAndView;
    }
}
