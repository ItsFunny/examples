package com.demo.sso.common.utils;

import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwt;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;
import io.jsonwebtoken.security.Keys;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import javax.servlet.http.HttpServletRequest;
import java.security.Key;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-01-12 19:36
 */
public class JWTUtil
{
    private static Logger logger = LoggerFactory.getLogger(JWTUtil.class);

    public static Key JWT_KEY = Keys.secretKeyFor(SignatureAlgorithm.HS256);

    public static class JWTParserResult
    {
        private boolean isExpired;

        private boolean isValid;

        public boolean isExpired()
        {
            return isExpired;
        }

        public static JWTParserResult success()
        {
            JWTParserResult result = new JWTParserResult();
            result.setExpired(false);
            result.setValid(true);
            return result;
        }

        public JWTParserResult()
        {
            this.isExpired = true;
            this.isValid = true;
        }

        public void setExpired(boolean expired)
        {
            isExpired = expired;
        }

        public boolean isValid()
        {
            return isValid;
        }

        public void setValid(boolean valid)
        {
            isValid = valid;
        }


        public boolean isSuccess()
        {
            return this.isExpired && this.isValid;
        }
    }

    public JWTUtil()
    {
        JWT_KEY = Keys.secretKeyFor(SignatureAlgorithm.HS256);
    }


    public static String buildToken(Map<String, Object> claims)
    {
        return Jwts.builder()
                .setClaims(claims)
                .setIssuer("sso-server-token")
                .signWith(JWT_KEY)
                .setExpiration(new Date(System.currentTimeMillis() + 1000 * 60 * 5))
                .compact();
    }


    // 在这里看来就显现出了Go的好处啊,多返回值,而Java可能需要返回exception或者model
    public static JWTParserResult parseJWT(String token)
    {

        return new JWTParserResult();
    }

    public static void main(String[] args)
    {
        Key JWT_KEY = Keys.secretKeyFor(SignatureAlgorithm.HS256);
        Map<String, Object> cc = new HashMap<>();
        cc.put("key", "test-t");
        String token = Jwts.builder()
                .signWith(JWT_KEY)
                .setClaims(cc)
                .setIssuer("sso-server-token")
                .setExpiration(new Date(System.currentTimeMillis() + 1000 * 60 * 5))
                .compact();
//        String token = "eyJhbGciOiJIUzI1NiJ9.eyJrZXkiOiJ0ZXN0LXQiLCJleHAiOjE1NDczNjEzNTB9.CDbUExqUnz7s2FWckwB5G6P8EGhT_s92_Sr1vvnWsFA";
        System.out.println(token);
        Claims claims = Jwts.parser().setSigningKey(JWT_KEY).parseClaimsJws(token).getBody();
        System.out.println(claims.getSubject());
        boolean validKey = claims.getIssuer().equals("sso-server-token");
        if (!validKey)
        {
            System.out.println("valid");
        }
        // 校验日期
        boolean before = new Date().after(claims.getExpiration());
        if (!before)
        {
            System.out.println("not expired");
        }

        // 校验详细信息,确认是否是sso-server发布的信息
        String json = claims.get("key", String.class);
        // 公钥解密,示例demo直接判断最后一个是否是't'
        if (json.charAt(json.length() - 1) == 't')
        {
            System.out.println("is ok");
        }

    }
    //eyJhbGciOiJIUzI1NiJ9.eyJrZXkiOiJ0ZXN0LXQiLCJleHAiOjE1NDczNjEzNTB9.CDbUExqUnz7s2FWckwB5G6P8EGhT_s92_Sr1vvnWsFA

}
