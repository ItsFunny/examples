package com.demo.common.annotation;

import java.lang.annotation.*;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-02-05 14:43
 */
@Target({ElementType.TYPE})
@Retention(RetentionPolicy.RUNTIME)
//@Inherited
@Documented
public @interface MyFeignClientAnnotation
{
    String name();
}
