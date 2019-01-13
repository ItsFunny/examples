package com.demo.sso.server.filter;

import com.demo.sso.common.utils.JWTUtil;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Component;
import org.springframework.util.StringUtils;

import javax.servlet.*;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.servlet.http.HttpSession;
import java.io.IOException;
import java.net.URLEncoder;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-01-13 11:22
 */
@Slf4j
public class LoginFilter implements Filter
{
    @Override
    public void init(FilterConfig filterConfig) throws ServletException
    {
        log.debug("[init]begin");
    }

    @Override
    public void doFilter(ServletRequest servletRequest, ServletResponse servletResponse, FilterChain filterChain) throws IOException, ServletException
    {

        HttpServletRequest request = (HttpServletRequest) servletRequest;
        HttpServletResponse response = (HttpServletResponse) servletResponse;
        String requestURI = request.getRequestURI();

        HttpSession session = request.getSession(true);
        System.out.println(session.getId());

        if (!requestURI.startsWith("/login") && !requestURI.startsWith("/doLogin"))
        {
            filterChain.doFilter(servletRequest, servletResponse);
            return;
        }

        Object token = session.getAttribute("token");
        if (null == token)
        {
            filterChain.doFilter(servletRequest, servletResponse);
            return;
        }

        // 从redis或者其他中获取登录的token,这里为了演示直接从内存中获取
        // token 是不可能为空的
        JWTUtil.JWTParserResult result = JWTUtil.parseJWT((String) token);
        if (result.isSuccess())
        {
            String returnUrl = request.getParameter("returnUrl");
            if (StringUtils.isEmpty(returnUrl))
            {
                returnUrl = "www.resource1.com:8000/test";
            }
            // 这里实际中是需要对? & 进行判断的
            response.sendRedirect(returnUrl + "?token=" + URLEncoder.encode((String) token, "utf-8"));
            return;
        }


        filterChain.doFilter(servletRequest, servletResponse);
    }

    @Override
    public void destroy()
    {

    }
}
