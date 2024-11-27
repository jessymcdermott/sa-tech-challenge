package io.honeydemo.meminator.imagepicker.controller;

import com.fasterxml.jackson.databind.ObjectMapper;

import jakarta.annotation.PostConstruct;

import java.io.File;
import java.io.IOException;
import java.util.List;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class ImageController {

    private List<String> imageList;

    @PostConstruct
    public void init() {
        ObjectMapper objectMapper = new ObjectMapper();
        try {
            File file = new File("images.json");
            if (!file.exists()) {
                throw new RuntimeException("images.json can not be found");
            }
            ImageData imageData = objectMapper.readValue(file, ImageData.class);
            this.imageList = imageData.getImages();
        } catch (IOException e) {
            throw new RuntimeException("Failed to parse images.json", e);
        }
    }

    private static final String bucketName = System.getenv("BUCKET_NAME");
    private static final String imageUrlPrefix = "https://" + bucketName + ".s3.amazonaws.com/";

    @GetMapping("/imageUrl")
    public ImageResult imageUrl() {
        // choose a random image from the list
        String chosenImage = imageList.get((int) (Math.random() * imageList.size()));
        // INSTRUMENTATION: add a useful attribute
        String imageUrl = imageUrlPrefix + chosenImage;
        return new ImageResult(imageUrl);
    }

    public static class ImageResult {
        private String imageUrl;

        public ImageResult(String imageUrl) {
            this.imageUrl = imageUrl;
        }

        public String getImageUrl() {
            return this.imageUrl;
        }

        public void setImageUrl(String imageUrl) {
            this.imageUrl = imageUrl;
        }
    }
}


