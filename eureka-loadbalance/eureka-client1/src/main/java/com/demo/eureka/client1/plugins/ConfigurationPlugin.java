package com.demo.eureka.client1.plugins;

import com.demo.common.annotation.*;
import com.demo.eureka.client1.controller.CallbackController;
import com.demo.eureka.client1.listener.ClientApplicationListener;
import com.demo.eureka.client1.service.IFeignTestService;
import com.joker.library.model.HttpClientResult;
import com.joker.library.utils.HttpUtils;
import org.springframework.beans.BeansException;
import org.springframework.beans.factory.config.BeanPostProcessor;
import org.springframework.core.Ordered;
import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;
import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Proxy;
import java.util.Observable;
import java.util.Observer;
import java.util.Random;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-02-03 18:18
 */
@Component
public class ConfigurationPlugin extends Observable implements BeanPostProcessor, Ordered
{
    private boolean isSoa;

    @Override
    public Object postProcessBeforeInitialization(Object bean, String beanName) throws BeansException
    {
        if (bean instanceof IFeignTestService)
        {
            throw new RuntimeException("okkk");
        }
        System.out.println(beanName);
        if (!isSoa)
        {
            // 判断是否有这个注解
            EnableMyEurekaClientAnnotation annotation = bean.getClass().getAnnotation(EnableMyEurekaClientAnnotation.class);
            if (null == annotation) return bean;
            isSoa = true;
            EnableMyFeignClients feignClients = bean.getClass().getAnnotation(EnableMyFeignClients.class);
            if (null!=feignClients)
            {
                dynamicProxy(feignClients.basepackages());
            }
            return bean;
        }

        // 判断@EnableMyFeignClients 注解是否存在,存在的话则需要通过basepackages扫描 那些有@MyFeignClient的service
        // 使它能够被Spring所管理

        EnableMyFeignClients feignClientAnnotation = bean.getClass().getAnnotation(EnableMyFeignClients.class);
        if (null == feignClientAnnotation) return bean;

        // TODO 还需要判断是否已经存在这个bean了,不要二次动态代理,(既多个地方写了@EnableMyFeignClients,造成了交叉动态代理)
        dynamicProxy(feignClientAnnotation.basepackages());


        return bean;
    }

    // TODO 扫描并且加载到Spring容器中
    private void dynamicProxy(String scanBasePackages)
    {
// 说明是一个feign接口
        // 则还需要判断是否有实现类,如果有实现类的话通过cglib代理
//        bean = Proxy.newProxyInstance(bean.getClass().getClassLoader(), bean.getClass().getInterfaces(), MyJDKSOAInvocationHandler);
        // 扫描特定注解的接口,然后为其生成动态代理
//        setChanged();
//        notifyObservers();
    }

    @Override
    public Object postProcessAfterInitialization(Object bean, String beanName) throws BeansException
    {

        return bean;
    }


    @Override
    public int getOrder()
    {
        return -1 * Integer.MAX_VALUE;
    }

    @PostConstruct
    public void afterPropertiesSet()
    {
        this.addObserver(new ClientApplicationListener());
    }

    // FIXME 接口拆分,这个方法太大了
    public static final InvocationHandler MyJDKSOAInvocationHandler = (proxy, method, args) ->
    {
        MyFeignClientAnnotation myFeignClientAnnotation = proxy.getClass().getAnnotation(MyFeignClientAnnotation.class);
        if (null == myFeignClientAnnotation)
        {
//            throw new RuntimeException("error configuration");
            return null;
        }
        String microName = myFeignClientAnnotation.name();
        MyRestAnnotation myRestAnnotation = method.getAnnotation(MyRestAnnotation.class);
        if (null == myRestAnnotation)
        {
//            throw new RuntimeException("error configuration");
            return null;
        }
        String restMethod = myRestAnnotation.method();
        String microUri = myRestAnnotation.url();
        // 第一种方案是,名字服务,既server 保存了名字对应的地址信息,client先发送名字请求,server返回地址,然后client调用
        // 第二种方案是,client统一将请求和名字一起发送给server,server请求之后返回结果给客户端
        // 这里采取第二种方案
        Random random = new Random();
        int index = random.nextInt(CallbackController.SERVERS.size());
        CallbackController.EurekaServer eurekaServer = CallbackController.SERVERS.get(index);

        String url = eurekaServer.getServerUrl() + "/call" + "/clientName=" + microName + "?microUri=" + microUri;
        HttpClientResult result = null;
        // FIXME 异常处理
//        if (restMethod.equals(MyRequestMethod.GET))
//        {
//            // get
//            result = HttpUtils.doGet(url);
//        } else
//        {
//            // post
//            result = HttpUtils.doPost(url);
//        }
        System.out.println(url);
        return "qwe";
    };


}
