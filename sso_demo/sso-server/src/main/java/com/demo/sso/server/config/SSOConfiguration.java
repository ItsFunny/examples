package com.demo.sso.server.config;

import com.demo.sso.server.filter.LoginFilter;
import org.springframework.boot.web.servlet.FilterRegistrationBean;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import javax.servlet.Filter;
import java.util.Arrays;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-01-13 13:17
 */
@Configuration
public class SSOConfiguration
{
    @Bean
    public FilterRegistrationBean<Filter> filterFilterRegistrationBean()
    {
        FilterRegistrationBean<Filter> filterFilterRegistrationBean = new FilterRegistrationBean<>();
        filterFilterRegistrationBean.setFilter(new LoginFilter());
        filterFilterRegistrationBean.setUrlPatterns(Arrays.asList("/*"));
        filterFilterRegistrationBean.setOrder(0);
        return filterFilterRegistrationBean;
    }

}
