package com.demo.controller;

import com.charile.utils.QRCodeUtil;
import com.demo.config.WechatConfig;
import com.ijpay.core.kit.QrCodeKit;
import com.ijpay.core.kit.WxPayKit;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail 1. 生成二维码,productId为商品ID,此时不预创建订单
 * 2. 在callback中,获取productId和tType,查询需要支付的信息等,其中,为了后续的积分不足现今补齐的需求
 * 需要采用策略模式或者是责任链模式,此时需要预创建订单
 * 3.
 * @Attention:
 * @Date 创建时间：2020-03-17 15:12
 */
@RestController
@EnableConfigurationProperties(value = {
        WechatConfig.class
})
public class WechatController
{
    @Autowired
    private WechatConfig wechatConfig;

    //
    @GetMapping(path = "/qrCode")
    public void generateCode(HttpServletRequest request, HttpServletResponse response) throws Exception
    {
        String url = WxPayKit.bizPayUrl(wechatConfig.getPayApiKey(), wechatConfig.getPayAppId(), wechatConfig.getPayMchId(), "10000000");
        QRCodeUtil.encode(url, response.getOutputStream());
    }
}
