package com.demo.common.configuration;

import lombok.Data;
import org.springframework.beans.factory.annotation.Value;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-02-02 08:58
 */
@Data
public class EurekaConfiguration
{
    @Value("${eureka.service.url}")
    private String eurekaServeUrl;

    @Value("${eureka.service.clientId}")
    private String eurekaClientId;

    @Value("${eureka.client.instance.hostname}")
    private String hostName;

    @Value("${eureka.client.instance.register}")
    private boolean needRegister;   // 是否需要注册到server上

}


