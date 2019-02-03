package com.demo.eureka.server.configuration;

import freemarker.cache.StringTemplateLoader;
import freemarker.cache.TemplateLoader;
import freemarker.cache.WebappTemplateLoader;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;
import org.springframework.core.io.support.PathMatchingResourcePatternResolver;
import org.springframework.web.context.ServletContextAware;
import org.springframework.web.servlet.ViewResolver;
import org.springframework.web.servlet.config.annotation.EnableWebMvc;
import org.springframework.web.servlet.view.freemarker.FreeMarkerConfig;
import org.springframework.web.servlet.view.freemarker.FreeMarkerConfigurer;
import org.springframework.web.servlet.view.freemarker.FreeMarkerViewResolver;

import javax.servlet.ServletContext;
import java.util.Locale;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-02-02 09:15
 */
@Configuration
@EnableWebMvc
@ComponentScan(basePackages = "com.demo.eureka.server")
public class SpringConfiguration implements ServletContextAware
{
    private ServletContext context;

    // 视图解析器
    @Bean
    public ViewResolver viewResolver()
    {
        FreeMarkerViewResolver viewResolver = new FreeMarkerViewResolver();
        viewResolver.setSuffix(".html");
        viewResolver.setPrefix("/WEB-INF/templates/");
        viewResolver.setCache(true);
        viewResolver.setContentType("text/html;charset=UTF-8");
        viewResolver.setRequestContextAttribute("requestContext");
        viewResolver.setOrder(0);
        return viewResolver;
    }

    @Bean
    public FreeMarkerConfigurer configurer()
    {
        FreeMarkerConfigurer configurer = new FreeMarkerConfigurer();

        freemarker.template.Configuration configuration = new freemarker.template.Configuration(freemarker.template.Configuration.DEFAULT_INCOMPATIBLE_IMPROVEMENTS);
        WebappTemplateLoader webappTemplateLoader = new WebappTemplateLoader(this.context);
        configuration.setTagSyntax(freemarker.template.Configuration.AUTO_DETECT_TAG_SYNTAX);
        configuration.setDefaultEncoding("UTF-8");
        configuration.setOutputEncoding("UTF-8");
        configuration.setLocale(Locale.SIMPLIFIED_CHINESE);
        configuration.setDateFormat("yyyy-MM-dd");
        configuration.setTimeFormat("HH:mm:ss");
        configuration.setDateTimeFormat("yyyy-MM-dd HH:mm:ss");
        configuration.setTemplateLoader(webappTemplateLoader);

        configurer.setConfiguration(configuration);
        return configurer;
    }

    @Override
    public void setServletContext(ServletContext servletContext)
    {
        this.context = servletContext;
    }
}
