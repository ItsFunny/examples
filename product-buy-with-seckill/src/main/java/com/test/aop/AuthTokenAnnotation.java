/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月21日 下午7:42:49
* 
*/
package com.test.aop;

import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;

/**
 * 
 * @When
 * @Description
 * @Detail
 * @author joker
 * @date 创建时间：2018年10月21日 下午7:42:49
 */
@Target(
{ ElementType.METHOD })
@Retention(RetentionPolicy.RUNTIME)
public @interface AuthTokenAnnotation
{
	int requestIndex() default 0;
}
