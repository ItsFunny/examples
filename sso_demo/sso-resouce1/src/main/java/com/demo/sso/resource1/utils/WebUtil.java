package com.demo.sso.resource1.utils;

import javax.servlet.http.HttpServletRequest;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-01-12 20:09
 */
public class WebUtil
{

    public static String getReturnUrl(HttpServletRequest request)
    {
        StringBuffer requestURL = request.getRequestURL();

        return requestURL.toString();

    }

}
