package com.demo.sso.resource1.config;

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
 * @date 创建时间：2019-01-12 20:41
 */
@Configuration
public class Resource1Configuration
{

    @Bean
    public FilterRegistrationBean<Filter> filterFilterRegistrationBean()
    {

        FilterRegistrationBean<Filter> bean = new FilterRegistrationBean<>();
        bean.setFilter(new TestFilter());
        bean.setUrlPatterns(Arrays.asList("/*"));
        return bean;
    }

}
