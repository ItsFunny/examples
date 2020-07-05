/**
*
* @Description
* @author joker 
* @date 创建时间：2018年9月18日 上午9:54:50
* 
*/
package com.test.tranwithmq.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import com.alibaba.fastjson.JSONObject;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.google.gson.Gson;
import com.test.common.mq.AppEventPublisher;
import com.test.common.mq.AppEventPublisher.AppEvent;
import com.test.common.mq.MessageContainer;
import com.test.tranwithmq.dao.MessageDao;

/**
* 
* @When
* @Description
* @Detail
* @author joker 
* @date 创建时间：2018年9月18日 上午9:54:50
*/
@RestController
public class MessageController
{
	@Autowired
	private MessageDao messageDao;
	@Autowired
	private MessageContainer container;
	
	@Autowired
	private AppEventPublisher eventPublisher;

	@PostMapping(value="addMessage",produces=MediaType.APPLICATION_JSON_UTF8_VALUE)
	public Integer addMessage(@RequestBody String message)
	{
		JSONObject obj=JSONObject.parseObject(message);
		AppEvent event = obj.toJavaObject(AppEvent.class);
		container.addMessage(event);
		return 1;
//		throw new RuntimeException("手动失败测试");
	}
	@GetMapping(value="updateStatus",produces=MediaType.APPLICATION_JSON_UTF8_VALUE)
	public Integer updateMessageStatus(@RequestParam("id")Long id,@RequestParam("status")Integer status)
	{
		container.updateMessageStatus(status, id);
		//发送消息,如果想保证次序的话用BlockingQueue即可
		AppEvent event = container.getModel(id);
		event.setStatus(1);
		eventPublisher.publishe(event);
		return 1;
	}
	
}
