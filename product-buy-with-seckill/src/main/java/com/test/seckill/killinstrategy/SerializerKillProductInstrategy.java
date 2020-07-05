package com.test.seckill.killinstrategy;

import java.math.BigDecimal;
import java.util.HashMap;
import java.util.Map;

import org.springframework.data.redis.core.StringRedisTemplate;
import org.springframework.stereotype.Service;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Primary;

import com.alibaba.druid.util.StringUtils;
import com.google.gson.Gson;
import com.joker.library.service.IdWorkerService;
import com.test.constants.ProductConstants;
import com.test.dao.OrderDao;
import com.test.model.Order;
import com.test.model.ProductDTO;
import com.test.mq.MQEventPublisher;
import com.test.mq.MQEventPublisher.AppEvent;
import com.test.redis.MyRedisService;

import lombok.extern.slf4j.Slf4j;

/**
 * 
 * @When
 * @Description 串行消费
 * @Detail
 * @author joker
 * @date 创建时间：2018年10月22日 上午8:52:13
 */
@Slf4j
public class SerializerKillProductInstrategy extends AbstractSecKillStrategy
{

	@Autowired
	private StringRedisTemplate redisTemplate;
	@Autowired
	private MyRedisService myRedisService;
	@Autowired
	private IdWorkerService idWorkerService;
	@Autowired
	private OrderDao orderDao;
	@Autowired
	private MQEventPublisher publisher;

	@Override
	public boolean filter(Long productId, long userId, Map<String, Object> params)
	{
		return true;
	}

	@Override
	public void doKill(ProductDTO product, long userId, Map<String, Object> params)
	{
		long productId = product.getProductId();
		String key = ProductConstants.SEC_PRODUCT_INFO + product.getProductId();
		long expireTime = System.currentTimeMillis() + 1000 * 60 * 3;
		// 尝试等待3秒
		if (myRedisService.tryLock(ProductConstants.SEC_PRODUCT_LOCK + productId, userId + ":" + expireTime, 3000))
		{
			try
			{
				// 占库存双重检测机制,直接用下面的取代
				// 看的出来,这步其实有重复相同的操作,所以做法是在redis中新增一块,就单独product_id 对应库存的信息
				Gson gson = new Gson();
				product = gson.fromJson(redisTemplate.opsForValue().get(key), ProductDTO.class);
				int stock = product.getProductStock();
				if (stock - 1 < 0)
				{
					params.put("error", "库存不足");
					return;
				}
				product.setProductStock(product.getProductStock() - 1);
				String json = gson.toJson(product);
				redisTemplate.opsForValue().set(key, json);
			} finally
			{
				if (myRedisService.tryReleaseLock(ProductConstants.SEC_PRODUCT_LOCK + productId,
						userId + ":" + expireTime))
				{
					log.info("userId:{}[成功释放锁]", userId);
				} else
				{
					log.error("userId:{},[释放锁失败]", userId);
				}
			}
		} else
		{
			params.put("error", "当前购买人数过多,请刷新再试");
		}
	}

	@Override
	public void rollBack(ProductDTO productDTO, long userId,Map<String, Object>params)
	{
		Gson gson=new Gson();
		//为了防止其他线程改变,需要一直等待获取锁 -- --但是这个方法更适合放到消息队列中做,不然这个线程就一直等待着了
		// redis中的值应该回滚
		String key = ProductConstants.SEC_PRODUCT_INFO + productDTO.getProductId();

		String json = redisTemplate.opsForValue().get(key);
		System.out.println(json);
		System.out.println(gson);
		
	}

}
