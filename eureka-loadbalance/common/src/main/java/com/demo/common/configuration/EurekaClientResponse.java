package com.demo.common.configuration;

import lombok.Data;

import java.io.Serializable;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-02-05 14:03
 */
@Data
public class EurekaClientResponse implements Serializable
{
    public static final int SUCCESS = 1;
    public static final int FAIL = -1;

    private int status;
    private String msg;

    public boolean isSuccess()
    {
        return SUCCESS == this.status;
    }
}
