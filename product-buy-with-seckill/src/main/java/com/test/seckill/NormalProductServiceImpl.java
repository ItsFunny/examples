/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月16日 下午7:46:24
* 
*/
package com.test.seckill;

import java.math.BigDecimal;
import java.util.HashMap;
import java.util.Map;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Primary;
import org.springframework.stereotype.Service;

import com.test.dao.OrderDao;
import com.test.dao.OrderDetailDao;
import com.test.dao.ProductDao;
import com.test.model.Order;
import com.test.model.OrderDetail;
import com.test.model.ProductDTO;
import com.test.mq.MQEventPublisher;
import com.test.mq.MQEventPublisher.AppEvent;

/**
 * 
 * @When
 * @Description
 * @Detail
 * @author joker
 * @date 创建时间：2018年10月16日 下午7:46:24
 */
// @Service(value="normalKil")
public class NormalProductServiceImpl extends AbstractProductSecKill
{
	@Autowired
	private ProductDao productDao;
	@Autowired
	private MQEventPublisher publisher;
	@Autowired
	private OrderDao orderDao;
	@Autowired
	private OrderDetailDao orderDetailDao;

	@Override
	public void doKillByDifLevel(ProductDTO product, long userId, Map<String, Object> res)
	{
		//再以乐观锁的方式检查一遍
		long productId=product.getProductId();
		int buyNum=product.getBuyNum();
		int validCount = productDao.decreaseStock(productId, buyNum);
		if (validCount == 1)
		{
			// 这里的步骤可以放在消息队列中
			// 并且以下这块操作需要用事务包裹,演示demo,就不添加service了
			// 生成订单
			Order order = new Order();
			order.setUserId(userId);
			BigDecimal bigDecimal = new BigDecimal(product.getProductPrice() + "");
			bigDecimal = bigDecimal.multiply(new BigDecimal(buyNum));
			order.setOrderPayment(bigDecimal);
			order.setOrderStatus(0);// 0 新创建,未付款

			orderDao.createOrder(order);
			// 生成订单详情
			OrderDetail orderDetail = new OrderDetail();
			orderDetail.setOrderId(order.getOrderId());
			orderDetail.setProductId(productId);
			orderDetail.setBuyNum(buyNum);
			orderDetailDao.createOrderDetail(orderDetail);
			// 往延时队列中发送消息
			AppEvent event = new AppEvent(order);
			event.setType("order.nonconsumer");
			publisher.publish(event);
			res.put("orderId", order.getOrderId());
		} else
		{
			// 失败,说明有多个买,并且卖完了了
			res.put("error", "卖完了");
		}

	}


}
