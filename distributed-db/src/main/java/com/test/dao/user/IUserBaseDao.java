/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月27日 下午2:04:05
* 
*/
package com.test.dao.user;

import java.util.List;

import org.apache.ibatis.annotations.Param;
import org.apache.ibatis.annotations.Select;

import com.joker.library.sqlextention.ISQLExtentionBaseCRUDDao;
import com.test.dto.UserDTO;
import com.test.model.User;

/**
 * 
 * @When
 * @Description
 * @Detail
 * @author joker
 * @date 创建时间：2018年10月27日 下午2:04:05
 */
public interface IUserBaseDao extends ISQLExtentionBaseCRUDDao<User>
{
	// for temporary ,mbg will generate this automically
	@Select("select * from ${tableName} where user_id between #{min} and #{end}")
	List<User> findByUserIdBetween(@Param("tableName") String tableName, @Param("min") long min,
			@Param("end") long end);

}
