package com.demo.config;

import lombok.Builder;
import lombok.Data;
import org.springframework.boot.context.properties.ConfigurationProperties;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-03-17 15:14
 */
@ConfigurationProperties(prefix = "wechat")
@Data
public class WechatConfig
{
    private String payAppId;
    private String payMchId;
    private String payApiKey;
    private String payCallBack;
    private String payNotify;
    private String payApiv3Key;
}
