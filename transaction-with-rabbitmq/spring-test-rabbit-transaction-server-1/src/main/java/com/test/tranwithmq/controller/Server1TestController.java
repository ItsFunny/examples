/**
*
* @Description
* @author joker 
* @date 创建时间：2018年9月18日 上午9:37:03
* 
*/
package com.test.tranwithmq.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.alibaba.fastjson.JSONObject;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.google.gson.Gson;
import com.test.common.mq.AppEventPublisher.AppEvent;
import com.test.tranwithmq.dao.MessageDao;
import com.test.tranwithmq.dao.UserDao;
import com.test.tranwithmq.model.MessageModel;
import com.test.tranwithmq.service.MQTransactionService;
import com.test.tranwithmq.service.UserService;

/**
* 
* @When
* @Description
* @Detail
* @author joker 
* @date 创建时间：2018年9月18日 上午9:37:03
*/
@RestController
public class Server1TestController
{

	@Autowired
	private MessageDao messageDao;
	
	@Autowired
	private UserService userService;
	
	@Autowired
	private MQTransactionService mqTransactionService;
	
	@RequestMapping(value="/server1/test")
	public String test1()
	{
//		//1.现在本地消息表中插入数据,
//		//2.然后通知消息服务器新的消息来了
//		//3.通知成功后,进行本地业务
//		//4.本地业务执行成功,通知消息服务器更新状态,发送消息
		
		//传入的是AppEvent的json格式
		//这里重复定义id,自行删除吧,
		String s="ceshi";
		MessageModel messageModel=new MessageModel();
		messageModel.setDetail(s);
		messageModel.setId((long) s.hashCode());
		AppEvent event=new AppEvent();
		event.setData(messageModel);
		event.setType("test");
		//这里自己设置吧,用这种方式是不安全的
		event.setId(System.currentTimeMillis());
		return mqTransactionService.testRabbitMqTransaction(event);
	}
}
