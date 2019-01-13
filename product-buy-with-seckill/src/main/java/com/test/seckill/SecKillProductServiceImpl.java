/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月16日 下午3:22:43
* 
*/
package com.test.seckill;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

import javax.annotation.PostConstruct;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.redis.core.RedisTemplate;
import com.alibaba.druid.util.StringUtils;
import com.google.gson.Gson;
import com.test.cache.OrderProductTestCache;
import com.test.constants.ProductConstants;
import com.test.model.ProductDTO;

/**
 * 
 * @When
 * @Description 通过redis 锁的形式串行的购买
 * @Detail
 * @author joker
 * @date 创建时间：2018年10月16日 下午3:22:43
 */
public class SecKillProductServiceImpl extends AbstractProductSecKill
{

	


	@Override
	public void doKillByDifLevel(ProductDTO product, long userId, Map<String, Object> res)
	{
		this.instrantegy.kill(product, userId, res);
	}

}
