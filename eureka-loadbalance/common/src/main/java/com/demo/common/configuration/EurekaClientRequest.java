package com.demo.common.configuration;

import com.demo.common.service.RegisterCallback;
import lombok.Data;

import java.io.Serializable;

/**
 * @author joker
 * @When register to eureka
 * @Description eureka client注册时候的model
 * @Detail
 * @date 创建时间：2019-02-02 10:31
 */
@Data
public class EurekaClientRequest implements Serializable
{
    private String hostName;
    private boolean needRegister;   // 是否需要注册到server上
    private String clientUrl;
    private RegisterCallback callback;


    public static EurekaClientRequest build()
    {
        EurekaClientRequest request = new EurekaClientRequest();
        return request;
    }

    public EurekaClientRequest setHostName(String hostName)
    {
        this.hostName = hostName;
        return this;
    }
    public EurekaClientRequest setClientUrl(String clientUrl)
    {
        this.clientUrl=clientUrl;
        return this;
    }
    public EurekaClientRequest setCallback(RegisterCallback callback)
    {
        this.callback=callback;
        return this;
    }
}
