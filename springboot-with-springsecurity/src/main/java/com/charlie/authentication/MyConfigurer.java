package com.charlie.authentication;

import lombok.extern.log4j.Log4j2;
import org.springframework.security.config.annotation.SecurityConfigurerAdapter;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.web.DefaultSecurityFilterChain;
import org.springframework.stereotype.Component;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-01-09 16:19
 */
@Component
@Log4j2
public class MyConfigurer extends SecurityConfigurerAdapter<DefaultSecurityFilterChain, HttpSecurity>
{
    @Override
    public void configure(HttpSecurity builder) throws Exception
    {
        log.info("MyConfigurer config");
        super.configure(builder);
    }


}
