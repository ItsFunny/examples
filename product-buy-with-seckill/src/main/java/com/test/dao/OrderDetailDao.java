/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月16日 下午2:56:01
* 
*/
package com.test.dao;

import org.apache.ibatis.annotations.Delete;
import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;

import com.test.model.OrderDetail;

/**
 * 
 * @When
 * @Description
 * @Detail
 * @author joker
 * @date 创建时间：2018年10月16日 下午2:56:01
 */
@Mapper
public interface OrderDetailDao
{
	@Insert("insert into order_detail values (null,#{orderId},#{productId},#{buyNum})")
	int createOrderDetail(OrderDetail orderDetail);

	@Delete("delete from order_detail where order_id=#{orderId}")
	int deleteByOrderId(@Param("orderId") long orderId);

}
