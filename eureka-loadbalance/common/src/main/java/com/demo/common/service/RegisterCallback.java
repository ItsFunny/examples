package com.demo.common.service;

import com.demo.common.configuration.EurekaClientRequest;
import com.joker.library.utils.HttpUtils;
import com.joker.library.utils.UrlUtils;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.client.RestTemplate;
import org.springframework.web.context.request.RequestContextHolder;
import org.springframework.web.context.request.ServletRequestAttributes;

import java.util.HashMap;
import java.util.Map;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-02-02 10:49
 */
public interface RegisterCallback
{

    void callBack(EurekaClientRequest request);


    RegisterCallback DEFAULT_CALLBACK = (request) ->
    {
        System.out.println(request.getHostName() + "注册成功");
        RestTemplate restTemplate = new RestTemplate();
        Map<String, Object> params = new HashMap<>();
        ServletRequestAttributes attributes = (ServletRequestAttributes) RequestContextHolder.getRequestAttributes();
        javax.servlet.http.HttpServletRequest httpServletRequest = attributes.getRequest();
        params.put("serverUrl", UrlUtils.getServerAddWithPort(httpServletRequest));
        params.put("serverName", "eureka-server1");
        String result = restTemplate.postForObject(request.getClientUrl() + "/eureka-callback", params, String.class);
        System.out.println(result);
    };
}
