package com.charlie.authentication;

import org.apache.commons.collections4.MapUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.Authentication;
import org.springframework.security.oauth2.common.OAuth2AccessToken;
import org.springframework.security.oauth2.provider.*;
import org.springframework.security.oauth2.provider.token.AuthorizationServerTokenServices;
import org.springframework.security.web.authentication.SavedRequestAwareAuthenticationSuccessHandler;
import org.springframework.stereotype.Service;
import org.springframework.util.StringUtils;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-01-09 17:09
 */
@Service
public class MySuccessHandler extends SavedRequestAwareAuthenticationSuccessHandler
{
    @Autowired
    private ClientDetailsService clientDetailsService;

    @Autowired
    private AuthorizationServerTokenServices tokenServices;

    @Override
    public void onAuthenticationSuccess(HttpServletRequest request, HttpServletResponse response, Authentication authentication) throws ServletException, IOException
    {
        String authHeader = request.getHeader("authentication");
        String name = authentication.getName();
//        if (StringUtils.isEmpty(authHeader))
//        {
//            throw new RuntimeException("请求头错误");
//        }
//        String[] split = authHeader.split(",");
//        if (split.length != 3)
//        {
//            throw new RuntimeException("请求参数缺失");
//        }
//        String clientId = split[0];
//        String clientSec = split[1];
        String clientId = "clientid";
        String clientSec = "clientSec";

        ClientDetails clientDetail = clientDetailsService.loadClientByClientId(clientId);
        if (clientDetail == null)
        {
            throw new RuntimeException("client 不存在");
        }
        if (!clientDetail.getClientSecret().equals(clientSec))
        {
            throw new RuntimeException("secret 错误");
        }
        TokenRequest tokenRequest = new TokenRequest(MapUtils.EMPTY_SORTED_MAP, clientId, clientDetail.getScope(), "custom");
        OAuth2Request oAuth2Request = tokenRequest.createOAuth2Request(clientDetail);
        OAuth2Authentication auth2Authentication = new OAuth2Authentication(oAuth2Request, authentication);
        OAuth2AccessToken accessToken = tokenServices.getAccessToken(auth2Authentication);
        if (accessToken == null)
        {
            accessToken = tokenServices.createAccessToken(auth2Authentication);
        }
        response.setContentType("application/json;charset=UTF-8");
        response.getWriter().write(accessToken.toString());

    }
}
