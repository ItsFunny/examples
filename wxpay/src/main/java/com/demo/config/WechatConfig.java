package com.demo.config;

import com.ijpay.wxpay.WxPayApiConfig;
import lombok.Builder;
import lombok.Data;
import lombok.Getter;
import lombok.Setter;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.boot.context.properties.EnableConfigurationProperties;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-03-17 15:14
 */
@ConfigurationProperties(prefix = "wechat")
@Getter
@Setter
@Builder
public class WechatConfig
{
    private String payAppId;
    private String payMchId;
    private String payApiKey;
    private String payCallBack;
    private String payNotify;
    private String payApiv3Key;

}
