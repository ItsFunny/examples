package com.demo.sso.common.filter;

import com.demo.sso.common.utils.JWTUtil;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.util.StringUtils;

import javax.servlet.*;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.servlet.http.HttpSession;
import java.io.IOException;
import java.net.URLDecoder;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-01-12 19:24
 */
public class TestFilter implements Filter
{
    private Logger log = LoggerFactory.getLogger(TestFilter.class);

    @Override
    public void init(FilterConfig filterConfig) throws ServletException
    {
        log.debug("[init]");
    }

    @Override
    public void doFilter(ServletRequest servletRequest, ServletResponse servletResponse, FilterChain filterChain) throws IOException, ServletException
    {
        HttpServletRequest req = (HttpServletRequest) servletRequest;

        // 这段代码请忽略,只为了拦截/test而已
        String uri = req.getRequestURI();
        if (!uri.startsWith("/test"))
        {
            filterChain.doFilter(servletRequest, servletResponse);
            return;
        }

        HttpServletResponse response = (HttpServletResponse) servletResponse;
        HttpSession session = req.getSession(true);
        Object decodeToken = session.getAttribute("token");
        if (null == decodeToken)
        {
            // 这里要从body 中获取,而不是从url上获取,但是未涉及前端,因而直接url中获取
            String token = req.getParameter("token");
            if (StringUtils.isEmpty(token) || !JWTUtil.parseJWT(URLDecoder.decode(token, "utf-8")).isSuccess())
            {
                log.error("[token为空或者无效]");
                // 跳转页面
                response.sendRedirect("http://www.sso.com:8888/login" + "?returnUrl=" + req.getRequestURL());
                return;
            }
        }else{
            JWTUtil.JWTParserResult result = JWTUtil.parseJWT((String) decodeToken);
            if (!result.isSuccess())
            {
                // 消息处理,简单的demo就直接跳转了
                response.sendRedirect("http://www.sso.com:8888/login" + "?returnUrl=" + req.getRequestURL());
                return;
            }
        }
        // 因为是多服务stateless ,所以每次请求都得携带token,因而先暂时放在session中
        session.setAttribute("token", decodeToken);
        filterChain.doFilter(servletRequest, servletResponse);
    }

    @Override
    public void destroy()
    {

    }
}
