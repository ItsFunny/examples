package com.demo.common.annotation;

import java.lang.annotation.*;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-02-03 17:29
 */
@Target({ElementType.TYPE})
@Retention(RetentionPolicy.RUNTIME)
//@Inherited
@Documented
public @interface EnableMyEurekaClientAnnotation
{
    String value() default "qwe";

}
