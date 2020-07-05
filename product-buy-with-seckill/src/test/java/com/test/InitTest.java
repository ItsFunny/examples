/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月16日 上午10:58:24
* 
*/
package com.test;

import java.math.BigDecimal;
import java.util.Random;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.context.annotation.Bean;
import org.springframework.data.redis.core.StringRedisTemplate;
import org.springframework.test.context.junit4.SpringJUnit4ClassRunner;

import com.test.dao.ProductDao;
import com.test.model.ProductDTO;

/**
 * 
 * @When
 * @Description
 * @Detail
 * @author joker
 * @date 创建时间：2018年10月16日 上午10:58:24
 */
@SpringBootTest(classes = OrderTestApplication.class)
@RunWith(SpringJUnit4ClassRunner.class)
public class InitTest
{
	@Autowired
	private StringRedisTemplate redisTemplate;
	@Autowired
	private ProductDao productDao;

	@Test
	public void init()
	{
		Random random = new Random();
		for (int i = 0; i < 100; i++)
		{
			ProductDTO product = new ProductDTO();
			product.setProductName("product_" + i);
			product.setProductPrice(new BigDecimal(new Random(47).nextDouble()));
			product.setProductLevel(random.nextBoolean() ? 0 : 1);
			productDao.insert(product);
		}
	}

	@Test
	public void testMybatis()
	{
		int decreaseStock = productDao.decreaseStock(1L, 1);
		System.out.println(decreaseStock);
	}

	@Test
	public void testRedis()
	{
	}

}
