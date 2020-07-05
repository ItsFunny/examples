/**
*
* @Description
* @author joker 
* @date 创建时间：2018年9月18日 上午10:11:33
* 
*/
package com.test.tranwithmq.spi;

import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestParam;

/**
* 
* @When
* @Description
* @Detail
* @author joker 
* @date 创建时间：2018年9月18日 上午10:11:33
*/
@FeignClient(name="message")
public interface IMessageServerFeignServiec
{
	@PostMapping(value="/addMessage",produces=MediaType.APPLICATION_JSON_UTF8_VALUE)
	public Integer addMessage(@RequestBody String message);
	
	@GetMapping(value="/updateStatus",produces=MediaType.APPLICATION_JSON_UTF8_VALUE)
	Integer updateMsgStatus(@RequestParam("id")Long id,@RequestParam("status")Integer status);

}
