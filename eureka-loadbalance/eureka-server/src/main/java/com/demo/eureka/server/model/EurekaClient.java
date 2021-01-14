package com.demo.eureka.server.model;

import com.demo.common.configuration.EurekaClientRequest;
import lombok.Data;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-02-02 10:38
 */
@Data
public class EurekaClient
{
    private String clientName;
    private String clientUrl;

    public void from(EurekaClientRequest request)
    {
        this.clientName = request.getHostName();
        this.clientUrl = request.getClientUrl();
    }

}
