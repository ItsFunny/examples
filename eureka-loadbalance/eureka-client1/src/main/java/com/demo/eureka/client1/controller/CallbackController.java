package com.demo.eureka.client1.controller;

import com.demo.common.configuration.EurekaClientRequest;
import lombok.Data;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.util.HashSet;
import java.util.Set;
import java.util.concurrent.CopyOnWriteArrayList;
import java.util.concurrent.CopyOnWriteArraySet;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-02-05 18:09
 */
@RestController
public class CallbackController
{
    public static CopyOnWriteArrayList<EurekaServer> SERVERS = null;

    public CallbackController()
    {
        SERVERS = new CopyOnWriteArrayList<>();
    }

    @Data
    public static class EurekaServer
    {
        private String serverUrl;
        private String serverName;

        public EurekaServer(String serverUrl, String serverName)
        {
            this.serverUrl = serverUrl;
            this.serverName = serverName;
        }
    }

    // FIXME 返回值修改
    @PostMapping(value = "/eureka-callback")
    public String registerCallback(HttpServletRequest request, HttpServletResponse response)
    {
        String serverUrl = request.getParameter("serverUrl");
        String serverName = request.getParameter("serverName");
        boolean isExist = SERVERS.add(new EurekaServer(serverUrl, serverName));
        if (isExist)
        {
            return "already exist";
        }
        return "ok";
    }

}
