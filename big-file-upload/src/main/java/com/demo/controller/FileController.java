package com.demo.controller;

import com.charile.bigfile.BigFileUploadProcessInfo;
import com.charile.bigfile.ChunkInfo;
import com.charile.bigfile.DefaultBigFileProcessor;
import com.charile.bigfile.MemoryChunkHandler;
import org.springframework.beans.factory.InitializingBean;
import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.multipart.MultipartFile;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.io.InputStream;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-05-28 10:04
 */
@RestController
public class FileController implements InitializingBean
{
    private DefaultBigFileProcessor bigFileProcessor;

    @PostMapping(value = "/upload/chunk", produces = MediaType.APPLICATION_JSON_VALUE)
    public String uploadMultiPartFile(@RequestParam("chunk") Integer chunk,
                                      @RequestParam("totalChunk") Integer totalChunk,
                                      @RequestParam("chunkSize") Integer chunkSize,
                                      @RequestParam("chunkMd5") String chunkMd5,
                                      @RequestParam("originFileName") String originFileName,
                                      @RequestParam("processId") String processId,
                                      @RequestParam("file") MultipartFile file,
                                      HttpServletRequest request, HttpServletResponse response) throws IOException
    {
        BigFileUploadProcessInfo info = new BigFileUploadProcessInfo();
        info.setChunk(chunk);
        info.setTotalChunk(totalChunk);
        info.setChunkSize(chunkSize);
        info.setChunkMd5(chunkMd5);
        info.setInputStream(file.getInputStream());
        info.setFileOriginName(originFileName);
        info.setProcessId(processId);
        this.bigFileProcessor.process(info);
        response.setStatus(200);
        return "ok";
    }

    @GetMapping(value = "/chunk/percent", produces = MediaType.APPLICATION_JSON_VALUE)
    public void getCent(@RequestParam("processId") String processId, @RequestParam("chunkMd5") String chunkMd5, HttpServletResponse response) throws IOException
    {
        ChunkInfo percentage = bigFileProcessor.getChunkHandler().getPercentage(processId, chunkMd5);
        float res = (float) percentage.getUploadedSize() / percentage.getChunkSize();
        response.getWriter().write(String.valueOf(res));
    }

    @Override
    public void afterPropertiesSet() throws Exception
    {
        this.bigFileProcessor = new DefaultBigFileProcessor(new MemoryChunkHandler());
        this.bigFileProcessor.setStorePath("/Users/joker/Desktop/split");
    }
}
