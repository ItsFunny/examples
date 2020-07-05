package com.demo.eureka.client1.listener;

import com.demo.common.annotation.EnableMyEurekaClientAnnotation;
import com.demo.common.annotation.MyFeignClientAnnotation;
import com.demo.common.annotation.MyRequestMethod;
import com.demo.common.annotation.MyRestAnnotation;
import com.demo.common.configuration.EurekaClientRequest;
import com.demo.common.configuration.EurekaClientResponse;
import com.demo.common.configuration.EurekaConfiguration;
import com.demo.eureka.client1.Client1Application;
import com.demo.eureka.client1.controller.CallbackController;
import com.joker.library.model.HttpClientResult;
import com.joker.library.utils.HttpUtils;
import com.joker.library.utils.UrlUtils;
import lombok.Data;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.cglib.proxy.Enhancer;
import org.springframework.cglib.proxy.MethodInterceptor;
import org.springframework.cglib.proxy.MethodProxy;
import org.springframework.context.ApplicationContext;
import org.springframework.context.ApplicationListener;
import org.springframework.context.event.ContextRefreshedEvent;
import org.springframework.http.MediaType;
import org.springframework.stereotype.Component;
import org.springframework.util.StringUtils;
import org.springframework.web.client.RestTemplate;
import org.springframework.web.context.request.RequestContextHolder;
import org.springframework.web.context.request.ServletRequestAttributes;

import javax.servlet.http.HttpServletRequest;
import java.lang.annotation.Annotation;
import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;
import java.util.List;
import java.util.Observable;
import java.util.Observer;
import java.util.Random;

/**
 * @author joker
 * @When
 * @Description 监听容器:
 * 1. 查看是否有使用@EnableMyEurekaClientAnnotation 这个注解
 * 2. 如果有使用,则获取到默认的配置类EurekaConfiguration
 * 3. 对注解类进行校验
 * @Detail
 * @date 创建时间：2019-02-03 22:55
 */
@Component
@Slf4j
public class ClientApplicationListener implements ApplicationListener<ContextRefreshedEvent>, Observer
{


    @Override
    public void onApplicationEvent(ContextRefreshedEvent contextRefreshedEvent)
    {
        ApplicationContext applicationContext = contextRefreshedEvent.getApplicationContext();
        // 获取配置信息
        EurekaConfiguration eurekaConfiguration = applicationContext.getBean(EurekaConfiguration.class);
        if (null == eurekaConfiguration)
        {
            throw new RuntimeException("error configuration,missing eurekaConfiguration");
        }
        if (!eurekaConfiguration.isNeedRegister())
        {
            log.warn("[EurekaConfiguration] wont register to the server");
            return;
        }

        String eurekaServeUrl = eurekaConfiguration.getEurekaServeUrl();
        if (StringUtils.isEmpty(eurekaServeUrl))
        {
            log.error("[EurekaConfiguration] register eureka server url is null");
            return;
        }
        ServletRequestAttributes attributes = (ServletRequestAttributes) RequestContextHolder.getRequestAttributes();
        // FIXME 有错误
//        HttpServletRequest request = attributes.getRequest();
//        EurekaClientRequest eurekaClientRequest = EurekaClientRequest.build()
//                .setHostName(eurekaConfiguration.getHostName())
//                .setClientUrl(UrlUtils.getServerAddWithPort(request));
//        // 通过http向server端发送注册消息
//        EurekaClientResponse response = restTemplate.postForObject(eurekaServeUrl, eurekaClientRequest, EurekaClientResponse.class);
//        if (!response.isSuccess())
//        {
//            log.error("[RegisterToEurekaServer] fail for reason:{}", response.getMsg());
//            throw new RuntimeException(response.getMsg());
//        }
        // 为接口生成动态代理

    }

    // 校验哪些方法添加了@MyFeignCline
    // 1. 遍历这些所有的bean
    // 2. 为其下的所有方法遍历,查看哪些方法有@RequestMapping等注解

    public void searchFeigns(ApplicationContext context)
    {

        String[] names = context.getBeanNamesForAnnotation(MyFeignClientAnnotation.class);
        for (String name : names)
        {
            Object bean = context.getBean(name);
            if (null == bean) continue;
            Method[] methods = bean.getClass().getMethods();
            for (Method method : methods)
            {
                Annotation[] annotations = method.getAnnotations();
                for (Annotation annotation : annotations)
                {
                    if (!(annotation instanceof MyRestAnnotation)) continue;
                    String restMethod = ((MyRestAnnotation) annotation).method();
                    String url = ((MyRestAnnotation) annotation).url();

                }
            }
        }
    }

    @Override
    public void update(Observable o, Object arg)
    {

    }


}
