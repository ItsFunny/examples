/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月21日 下午7:20:23
* 
*/
package com.test.controller;

import java.math.BigDecimal;
import java.util.HashMap;
import java.util.Map;
import java.util.Random;
import java.util.concurrent.TimeUnit;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.apache.commons.lang3.StringUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.redis.core.StringRedisTemplate;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.servlet.ModelAndView;
import org.springframework.web.servlet.ModelAndViewDefiningException;

import com.joker.library.service.IdWorkerService;
import com.test.aop.AuthTokenAnnotation;
import com.test.cache.OrderProductTestCache;
import com.test.constants.ProductConstants;
import com.test.dao.OrderDao;
import com.test.model.Order;
import com.test.model.ProductDTO;
import com.test.redis.MyRedisService;
import com.test.seckill.ProductBuyService;

import lombok.extern.slf4j.Slf4j;

/**
 * 
 * @When
 * @Description 秒杀情况下的logic
 * @Detail
 * @author joker
 * @date 创建时间：2018年10月21日 下午7:20:23
 */
@Slf4j
@Controller
public class SecKillController
{
	@Autowired
	private IdWorkerService idWorkerService;

	@Autowired
	private StringRedisTemplate redisTemplate;

	@Autowired
	private ProductBuyService productBuyService;

	Random random = new Random(47);

	@RequestMapping(value = "/sec/kill")
	public ModelAndView showSecKillProduct()
	{
		System.out.println("请求进入");
		ProductDTO product = null;
		boolean isSec = false;
		while (!isSec)
		{
			int index = random.nextInt(OrderProductTestCache.PRODUCTS.size());
			product = OrderProductTestCache.PRODUCTS.get(index);
			if (product.getProductLevel() == 1)
				isSec = true;
		}
		product = new ProductDTO();
		product.setProductId(1);
		product.setBuyNum(1);
		product.setProductLevel(1);
		product.setProductName("product_1");
		product.setProductPrice(new BigDecimal(0.724d));
		product.setProductStock(1000);
		ModelAndView modelAndView = new ModelAndView("sec_buy_product");
		// 生成token
		long token = idWorkerService.nextId();
		redisTemplate.opsForValue().set(ProductConstants.PRODUCT_BUY_TOKEN + token, 0 + "", 60000,
				TimeUnit.MILLISECONDS);
		modelAndView.addObject("token", token);
		// 随机挑选某个商品去买
		modelAndView.addObject("buyStock", 1);
		modelAndView.addObject("product", product);
		return modelAndView;
	}

	@AuthTokenAnnotation(requestIndex = 0)
	@RequestMapping(value = "/sec/doKill")
	public ModelAndView doKillSecPro(HttpServletRequest request, HttpServletResponse response)
	{
		Map<String, Object> params = new HashMap<String, Object>();
		long productId = Long.parseLong(request.getParameter("productId"));
		ModelAndView modelAndView = null;
		int userId=random.nextInt();
		int buyNum=Integer.parseInt(request.getParameter("buyNumber"));
		
		productBuyService.killProduct(productId, userId,buyNum, params);
		if(params.containsKey("error")) {
			modelAndView=new ModelAndView("error",params);
		}else {
			modelAndView=new ModelAndView("pay",params);
		}
		return modelAndView;
	}

	public boolean checkToken(String token)
	{
		String key = ProductConstants.PRODUCT_BUY_TOKEN + token;
		String existToken = redisTemplate.opsForValue().get(key);
		if (StringUtils.isEmpty(existToken))
		{
			return false;
		}
		// 利用原子性判断是否使用
		Long res = redisTemplate.opsForValue().increment(key, 1);
		if (res == 1)
		{
			// 未使用,可以正常使用
			return true;
		}
		// 否则,已经使用过了,删除
		redisTemplate.delete(key);
		return false;
	}
}
