/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月16日 下午1:49:53
* 
*/
package com.test.dao;

import org.apache.ibatis.annotations.Delete;
import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;
import org.apache.ibatis.annotations.Update;

import com.test.model.Order;

import lombok.Delegate;

/**
 * 
 * @When
 * @Description
 * @Detail
 * @author joker
 * @date 创建时间：2018年10月16日 下午1:49:53
 */
@Mapper
public interface OrderDao
{
	int createOrder(Order order);

	@Update("update order_master set order_status =#{status} where order_id = #{orderId} ")
	int updateOrder(@Param("orderId") long orderId, @Param("status") int status);

	@Delete("delete from order_master where order_id=#{orderId} and order_status=#{status}")
	int deleteByIdAndStatus(@Param("orderId") long orderId, @Param("status") int status);

}
