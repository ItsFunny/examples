package com.demo.sso.server.controller;

import com.demo.sso.common.utils.JWTUtil;
import org.springframework.http.MediaType;
import org.springframework.stereotype.Controller;
import org.springframework.util.StringUtils;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.servlet.ModelAndView;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.UnsupportedEncodingException;
import java.net.URLEncoder;
import java.util.HashMap;
import java.util.Map;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-01-13 11:21
 */
@Controller
public class SSOController
{
    @GetMapping(value = "/login")
    public ModelAndView login(HttpServletRequest request, HttpServletResponse response)
    {
        String returnUrl = request.getParameter("returnUrl");
        if (StringUtils.isEmpty(returnUrl))
        {
            returnUrl = "www.resource1.com/test";
        }
        ModelAndView modelAndView = new ModelAndView("login");
        modelAndView.addObject("returnUrl", returnUrl);
        return modelAndView;
    }

    @PostMapping(value = "/doLogin")
    public ModelAndView doLogin(HttpServletRequest request, HttpServletResponse response)
    {

        String returnUrl = request.getParameter("returnUrl");
        Map<String, Object> claims = new HashMap<>();
        claims.put("key", "test-t");
        String token = null;
        try
        {
            token = JWTUtil.buildToken(claims);
            request.getSession(true).setAttribute("token", URLEncoder.encode(token, "utf-8"));
        } catch (UnsupportedEncodingException e)
        {
            e.printStackTrace();
        }
        ModelAndView modelAndView = new ModelAndView("redirect:" + returnUrl);
        modelAndView.addObject("token", token);

        return modelAndView;
    }

}
