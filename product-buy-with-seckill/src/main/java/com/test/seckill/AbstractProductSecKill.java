/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月16日 下午7:28:22
* 
*/
package com.test.seckill;

import java.util.HashMap;
import java.util.Map;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.redis.core.StringRedisTemplate;

import com.alibaba.druid.util.StringUtils;
import com.google.gson.Gson;
import com.test.constants.ProductConstants;
import com.test.dao.OrderDao;
import com.test.dao.ProductDao;
import com.test.model.ProductDTO;
import com.test.seckill.killinstrategy.SecKillInstrantegy;

/**
 * 
 * @When
 * @Description
 * @Detail
 * @author joker
 * @date 创建时间：2018年10月16日 下午7:28:22
 */
public abstract class AbstractProductSecKill implements ProductBuyService
{
	private int type;
	private AbstractProductSecKill nextHandler;

	@Autowired
	protected StringRedisTemplate stringRedisTemplate;
	@Autowired
	private ProductDao productDao;

	// 根据srp原则,这个更适合拆出一个单独的接口
	@Autowired
	protected SecKillInstrantegy instrantegy;

	public Map<String, Object> findFromRedisOrDb(long productId, int buyNum)
	{
		Map<String, Object> params = new HashMap<>();
		String productStr = stringRedisTemplate.opsForValue().get(ProductConstants.SEC_PRODUCT_INFO + productId);
		ProductDTO productDTO = null;
		if (StringUtils.isEmpty(productStr))
		{
			// 这里为了防止redis穿透,可以设置默认值
			productDTO = productDao.findById(productId);
		} else
		{
			Gson gson = new Gson();
			productDTO = gson.fromJson(productStr, ProductDTO.class);
		}
		if (null == productDTO)
			params.put("error", "商品部存在");
		// 这里是有缺陷的,因为如果并发大的情况下,如果这里显示库存充足,但是刚好另外的线程消费完毕导致卖完了
		// 就会导致超卖的问题
		else if (productDTO.getProductStock() - buyNum < 1)
			params.put("error", "库存不足");
		else
			params.put("productDTO", productDTO);
		return params;
	}

	@Override
	public void killProduct(long productId, long userId, int buyNum, Map<String, Object> res)
	{
		if( !(this.instrantegy.filter(productId, userId, res)))
		{
			return;
		}
		if (res.containsKey("error"))
			return;
		// res.putAll(instrantegy.prepare(productId, userId));
		res.putAll(findFromRedisOrDb(productId, buyNum));
		if (res.containsKey("error"))
			return;
		ProductDTO product = (ProductDTO) res.get("productDTO");
		if (null == product)
			res.put("error", "商品信息不存在");
		else
			doKill(product, userId, res);
	}

	private void doKill(ProductDTO product, long userId, Map<String, Object> res)
	{
		if (product.getProductLevel() == type)
		{
			this.doKillByDifLevel(product, userId, res);
		} else if (this.nextHandler != null)
		{
			this.nextHandler.doKill(product, userId, res);
		} else
		{
			// 可以抛出异常
			return;
		}
	}

	public abstract void doKillByDifLevel(ProductDTO product, long userId, Map<String, Object> res);

	public int getType()
	{
		return type;
	}

	public void setType(int type)
	{
		this.type = type;
	}

	public AbstractProductSecKill getNextHandler()
	{
		return nextHandler;
	}

	public void setNextHandler(AbstractProductSecKill nextHandler)
	{
		this.nextHandler = nextHandler;
	}

}
