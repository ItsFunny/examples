package com.demo.sso.resouce2.config;

import com.demo.sso.common.filter.TestFilter;
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
 * @date 创建时间：2019-01-13 15:35
 */
@Configuration
public class Resource2Configuration
{
    @Bean
    public FilterRegistrationBean<Filter> filterFilterRegistrationBean()
    {
        FilterRegistrationBean<Filter> filterFilterRegistrationBean = new FilterRegistrationBean<>();
        filterFilterRegistrationBean.setFilter(new TestFilter());
        filterFilterRegistrationBean.setUrlPatterns(Arrays.asList("/*"));
        return filterFilterRegistrationBean;
    }

}
