/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月22日 上午8:49:59
* 
*/
package com.test.seckill.killinstrategy;

import java.util.Map;

import com.test.model.ProductDTO;

/**
 * 
 * @When
 * @Description
 * @Detail
 * @author joker
 * @date 创建时间：2018年10月22日 上午8:49:59
 */
public interface SecKillInstrantegy
{
	void kill(ProductDTO productDTO, long userId, Map<String, Object> params);
	
	boolean filter(Long  productId,long userId,Map<String, Object>params);

//	Map<String, Object> prepare(long productId, long userId);
}
