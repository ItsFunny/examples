package com.demo.common.configuration;

import lombok.Data;
import org.springframework.beans.factory.annotation.Value;

/**
 * @author joker
 * @When register to eureka
 * @Description eureka client注册时候的model
 * @Detail
 * @date 创建时间：2019-02-02 10:31
 */
@Data
public class EurekaClientRequest
{


    private String hostName;

    private boolean needRegister;   // 是否需要注册到server上


    private String clientUrl;




}
