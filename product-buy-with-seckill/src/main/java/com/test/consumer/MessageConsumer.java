/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月16日 下午3:42:51
* 
*/
package com.test.consumer;

import com.test.mq.MQEventPublisher.AppEvent;

/**
* 
* @When
* @Description
* @Detail
* @author joker 
* @date 创建时间：2018年10月16日 下午3:42:51
*/
public interface MessageConsumer
{
	void process(AppEvent event);
	String getType();
}
