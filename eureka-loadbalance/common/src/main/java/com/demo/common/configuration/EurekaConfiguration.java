package com.demo.common.configuration;

import lombok.Data;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-02-02 08:58
 */
@Data
@Component
public class EurekaConfiguration
{
    @Value("${eureka.service.url:qwe}")
    private String eurekaServeUrl;

    @Value("${eureka.service.clientId:0}")
    private String eurekaClientId;

    @Value("${eureka.client.instance.hostname:null}")
    private String hostName;

    @Value("${eureka.client.instance.register:false}")
    private boolean needRegister;   // 是否需要注册到server上

    @Override
    public String toString()
    {
        return "EurekaConfiguration{" +
                "eurekaServeUrl='" + eurekaServeUrl + '\'' +
                ", eurekaClientId='" + eurekaClientId + '\'' +
                ", hostName='" + hostName + '\'' +
                ", needRegister=" + needRegister +
                '}';
    }

}



