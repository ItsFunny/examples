/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月23日 上午8:05:55
* 
*/
package com.test.seckill.killinstrategy;

import java.math.BigDecimal;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import javax.annotation.PostConstruct;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.redis.core.StringRedisTemplate;

import com.alibaba.druid.util.StringUtils;
import com.fasterxml.jackson.annotation.JacksonInject.Value;
import com.google.gson.Gson;
import com.test.constants.ProductConstants;
import com.test.dao.OrderDao;
import com.test.dao.ProductDao;
import com.test.model.Order;
import com.test.model.ProductDTO;
import com.test.mq.MQEventPublisher;
import com.test.mq.MQEventPublisher.AppEvent;

import lombok.extern.slf4j.Slf4j;

/**
 * 
 * 通过redis的list实现秒杀,具体如下: 具体如下:先通过sku_id 在redis中预设置库存和价格等信息,可以简单的通过map
 * 
 * @When
 * @Description 基于redis的list实现排队机制,就是lpush 然后rpop实现
 * @Detail
 * @author joker
 * @date 创建时间：2018年10月23日 上午8:05:55
 */
@Slf4j
public class ListSerializerKillProduct extends AbstractSecKillStrategy
{
	@Autowired
	private StringRedisTemplate stringRedisTemplate;


	@Override
	public boolean filter(Long productId, long userId, Map<String, Object> params)
	{
		String stockStr = stringRedisTemplate.opsForValue().get(ProductConstants.SEC_PRODUCT_STOCK + productId);
		if (StringUtils.isEmpty(stockStr))
		{
			params.put("error", "商品信息不存在");
			return false;
		} else
		{
			params.put("limit", stockStr);
			Long size = stringRedisTemplate.opsForList().size(ProductConstants.USER_SEC_QUEUE + productId);
			if (size > Long.parseLong(stockStr))
			{
				params.put("error", "库存不足");
				return false;
			}
			else
				return true;
		}

	}

	@Override
	public void doKill(ProductDTO productDTO, long userId, Map<String, Object> params)
	{
		// 正常情况下redis中的key还应该有一个sku属性,这里就只简单的id,不分规格之类的
		// 创建订单
		// Order order = new Order();
		// order.setOrderPayment(productDTO.getProductPrice().multiply(new
		// BigDecimal(productDTO.getBuyNum())));
		// order.setOrderStatus(0);
		// order.setUserId(userId);
		// // 可以选择直接插入或者是发布到消息队列中
		// int validCount = orderDao.createOrder(order);
		// if (validCount == 0)
		// params.put("error", "数据插入出错,请刷新重试,或者过会再试");
		Gson gson = new Gson();
		Map<String, Object> map = new HashMap<String, Object>();
		map.put("productDTO", productDTO);
		map.put("userId", userId);
		String value = gson.toJson(map);
		// 先直接用map代替
		Long number = stringRedisTemplate.opsForList()
				.leftPush(ProductConstants.USER_SEC_QUEUE + productDTO.getProductId(), value);
		long limit = Long.parseLong( params.get("limit").toString());
		if (number > limit) {
			params.put("error", "sold out");
			params.put("index", number);
		}
	}

	@Override
	public void rollBack(ProductDTO productDTO, long userId,Map<String, Object>params)
	{
		log.error("[rollBack]触发了回滚操作,product:{},userId:{},额外的参数:{}",productDTO,userId,params);
	}

}
