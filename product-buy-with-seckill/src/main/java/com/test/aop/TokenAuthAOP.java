/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月21日 下午7:41:54
* 
*/
package com.test.aop;

import javax.servlet.http.HttpServletRequest;

import org.aspectj.lang.JoinPoint;
import org.aspectj.lang.ProceedingJoinPoint;
import org.aspectj.lang.annotation.Around;
import org.aspectj.lang.annotation.Aspect;
import org.aspectj.lang.annotation.Before;
import org.aspectj.lang.annotation.Pointcut;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.redis.core.StringRedisTemplate;
import org.springframework.stereotype.Component;

import com.alibaba.druid.util.StringUtils;
import com.test.constants.ProductConstants;
import com.test.redis.MyRedisService;

/**
 * 
 * @When
 * @Description
 * @Detail
 * @author joker
 * @date 创建时间：2018年10月21日 下午7:41:54
 */
@Aspect
@Component
public class TokenAuthAOP
{
	@Autowired
	private StringRedisTemplate stringRedisTemplate;

	@Pointcut("@annotation(com.test.aop.AuthTokenAnnotation)")
	public void authToken()
	{
	}

	/*
	 * 下单之前校验token的有效性,如果token对应的redis不存在说明是个无效记录 如果存在,则可用原子性++ 判断是否是第一次使用还是已经使用多次
	 */
	@Before("authToken() && @annotation(annotation)")
	public void authTokenByAop(JoinPoint point, AuthTokenAnnotation annotation)
	{
		Object[] args = point.getArgs();
		int index = annotation.requestIndex();
		if (index >= args.length || !(args[index] instanceof HttpServletRequest))
		{
			throw new RuntimeException("参数不匹配");
		}
		HttpServletRequest request = (HttpServletRequest) args[index];
		String token = request.getParameter("token");
		String existToken = stringRedisTemplate.opsForValue().get(ProductConstants.PRODUCT_BUY_TOKEN + token);
		if (StringUtils.isEmpty(existToken))
			throw new RuntimeException("请尝试重新刷新页面访问");
		Long res = stringRedisTemplate.opsForValue().increment("sec_product_kill_" + token, 1);
		if (res != 1)
		{
			stringRedisTemplate.delete("sec_product_kill_" + token);
			throw new RuntimeException("请尝试重新刷新页面访问");
		}
	}
}
