/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月16日 上午10:53:44
* 
*/
package com.test.dao;

import java.util.List;

import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;
import org.apache.ibatis.annotations.Select;
import org.apache.ibatis.annotations.Update;

import com.test.model.ProductDTO;


/**
 * 
 * @When
 * @Description
 * @Detail
 * @author joker
 * @date 创建时间：2018年10月16日 上午10:53:44
 */
@Mapper
public interface ProductDao
{
	@Insert("insert into product values (null,#{productName},#{productPrice},#{productStock},#{productLevel})")
	int insert(ProductDTO product);
	
	@Select("select * from product")
	List<ProductDTO>findAll();
	
	@Select("select * from product where product_id =#{productId}")
	ProductDTO findById(long productId);
	
	@Update("update product set product_stock=product_stock-#{buyNum} where product_id=#{productId} and product_stock-#{buyNum}>0")
	int decreaseStock(@Param("productId")long productId,@Param("buyNum")int buyNum);
	
	
}
