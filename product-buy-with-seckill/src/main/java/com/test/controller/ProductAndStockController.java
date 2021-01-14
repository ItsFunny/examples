/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月16日 上午10:47:14
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
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.data.redis.core.StringRedisTemplate;
import org.springframework.data.redis.core.ValueOperations;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestAttribute;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.servlet.ModelAndView;

import com.joker.library.service.IdWorkerService;
import com.test.cache.OrderProductTestCache;
import com.test.dao.OrderDao;
import com.test.dao.OrderDetailDao;
import com.test.dao.ProductDao;
import com.test.model.Order;
import com.test.model.OrderDetail;
import com.test.model.ProductDTO;
import com.test.mq.MQEventPublisher;
import com.test.mq.MQEventPublisher.AppEvent;
import com.test.seckill.NormalProductServiceImpl;
import com.test.seckill.ProductBuyService;

/**
 * 
 * @When
 * @Description
 * @Detail
 * @author joker
 * @date 创建时间：2018年10月16日 上午10:47:14
 */
@Controller
public class ProductAndStockController
{
	@Autowired
	private StringRedisTemplate redisTemplate;
	@Autowired
	private IdWorkerService idWorkerService;
	@Autowired
	private ProductDao productDao;
	@Autowired
	private OrderDao orderDao;

	@Autowired
	private ProductBuyService productBuyService;

	private Random random = new Random();

	@RequestMapping(value = "/buy")
	public ModelAndView buy()
	{
		ProductDTO product = null;
		boolean isNorma=false;
		while (!isNorma)
		{
			int index = random.nextInt(OrderProductTestCache.PRODUCTS.size());
			product = OrderProductTestCache.PRODUCTS.get(index);
			if(product.getProductLevel()==0) isNorma=true;
		}
		ModelAndView modelAndView = new ModelAndView("buy_product");
		// 生成token
		long token = idWorkerService.nextId();
		redisTemplate.opsForValue().set("product_buy_token_" + token, 0 + "", 60000, TimeUnit.MILLISECONDS);
		modelAndView.addObject("token", token);
		// 随机挑选某个商品去买
		modelAndView.addObject("buyStock", random.nextInt(10));
		modelAndView.addObject("product", product);
		return modelAndView;
	}

	// 展示购买的商品信息
	// 预扣库存
	@RequestMapping(value = "/doBuy")
	public ModelAndView doBuy(HttpServletRequest request, HttpServletResponse response)
	{
		Map<String, Object> params = new HashMap<String, Object>();
		ModelAndView modelAndView = null;
		// token校验应该放在这,因为需要预扣库存
		String token = request.getParameter("token");
		if (!checkToken(token))
		{
			modelAndView = new ModelAndView("error");
			modelAndView.addObject("error", "无效请求,请重新购买");
			return modelAndView;
		}
		// buyNum 和productId更应该是list类型,而不是单体,因为存在购物车购买的情况
		// 同样,只是为了演示的话,直接单个数据
		int buyNum = Integer.parseInt(request.getParameter("buyNumber"));
		long productId = Long.parseLong(request.getParameter("productId"));
		// 因为只是个demo,所以userId直接伪造的方式
		long userId = Math.abs(random.nextLong());
		ProductDTO product = productDao.findById(productId);
		product.setBuyNum(buyNum);
		productBuyService.killProduct(productId, userId,buyNum, params);
		if (params.containsKey("error"))
		{
			modelAndView = new ModelAndView("error", params);
		} else
		{
			// 跳转到支付页面
			modelAndView = new ModelAndView("pay", params);
		}
		return modelAndView;
	}

	// 校验商品信息
	// @RequestMapping(value = "pay")
	// public ModelAndView payMoney(HttpServletRequest request, HttpServletResponse
	// response)
	// {
	//
	// long orderId=Long.parseLong(request.getParameter("orderId"));
	// //更新order状态
	// ModelAndView modelAndView = null;
	// //生成订单,跳转支付页面
	// return modelAndView;
	// }
	@RequestMapping(value = "/doPay")
	@ResponseBody
	public String doPay(HttpServletRequest request, HttpServletResponse response)
	{
		// 调用接口
		long orderId = Long.parseLong(request.getParameter("orderId"));
		int validCount = orderDao.updateOrder(orderId, 1);
		if (validCount < 1)
		{
			return "fail";
		}
		return "ok";
	}

	public boolean checkToken(String token)
	{
		String key = "product_buy_token_" + token;
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
