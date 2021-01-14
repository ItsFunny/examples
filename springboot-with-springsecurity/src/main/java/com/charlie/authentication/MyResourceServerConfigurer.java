package com.charlie.authentication;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.configurers.CorsConfigurer;
import org.springframework.security.oauth2.config.annotation.web.configuration.EnableResourceServer;
import org.springframework.security.oauth2.config.annotation.web.configuration.ResourceServerConfigurerAdapter;
import org.springframework.security.web.authentication.AuthenticationFailureHandler;
import org.springframework.security.web.authentication.AuthenticationSuccessHandler;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-01-09 15:32
 */
@Configuration
@EnableResourceServer
public class MyResourceServerConfigurer extends ResourceServerConfigurerAdapter
{
    @Autowired
    private AuthenticationSuccessHandler successHandler;
    @Autowired
    private AuthenticationFailureHandler failureHandler;

    @Autowired
    private MyConfigurer myConfigurer;

    @Override
    public void configure(HttpSecurity http) throws Exception
    {
        http
                .formLogin()
                .loginProcessingUrl("/login")
                .successHandler(successHandler)
                .failureHandler(failureHandler);

        http
                .apply(myConfigurer)
                .and()
                .authorizeRequests()
                .antMatchers("/test/token")
                .permitAll()
                .antMatchers(
                        "/register",
                        "/social/**",
                        "/**/*.js",
                        "/**/*.css",
                        "/**/*.jpg",
                        "/**/*.png",
                        "/**/*.woff2",
                        "/code/image")
                .permitAll()//以上的请求都不需要认证
                .anyRequest()
                .authenticated()
                .and()
                .csrf().disable();

    }
}
