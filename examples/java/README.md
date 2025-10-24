# Java Examples for Reolink Camera API

This directory contains Java examples for using the Reolink Camera API.

## Prerequisites

- Java 11+
- Maven or Gradle

## Option 1: Using OpenAPI Generator (Recommended)

Generate a Java client from the OpenAPI specification:

```bash
# Using Maven plugin (add to pom.xml)
<plugin>
    <groupId>org.openapitools</groupId>
    <artifactId>openapi-generator-maven-plugin</artifactId>
    <version>7.0.0</version>
    <executions>
        <execution>
            <goals>
                <goal>generate</goal>
            </goals>
            <configuration>
                <inputSpec>${project.basedir}/../../docs/reolink-camera-api-openapi.yaml</inputSpec>
                <generatorName>java</generatorName>
                <library>okhttp-gson</library>
                <configOptions>
                    <sourceFolder>src/gen/java/main</sourceFolder>
                </configOptions>
            </configuration>
        </execution>
    </executions>
</plugin>

# Or using CLI
openapi-generator-cli generate \
  -i ../../docs/reolink-camera-api-openapi.yaml \
  -g java \
  -o ./generated-client \
  --library okhttp-gson
```

### Using the Generated Client

```java
import com.reolink.api.*;
import com.reolink.api.auth.*;
import com.reolink.api.models.*;
import com.reolink.api.api.DefaultApi;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://192.168.1.100/cgi-bin");

        DefaultApi apiInstance = new DefaultApi(defaultClient);
        
        // Login request
        List<Object> loginRequest = Arrays.asList(
            new HashMap<String, Object>() {{
                put("cmd", "Login");
                put("param", new HashMap<String, Object>() {{
                    put("User", new HashMap<String, String>() {{
                        put("userName", "admin");
                        put("password", "your_password");
                    }});
                }});
            }}
        );
        
        try {
            List<Object> result = apiInstance.loginPost(loginRequest);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling DefaultApi#loginPost");
            e.printStackTrace();
        }
    }
}
```

## Option 2: Manual Implementation

### Maven Dependencies

```xml
<dependencies>
    <dependency>
        <groupId>com.squareup.okhttp3</groupId>
        <artifactId>okhttp</artifactId>
        <version>4.11.0</version>
    </dependency>
    <dependency>
        <groupId>com.google.code.gson</groupId>
        <artifactId>gson</artifactId>
        <version>2.10.1</version>
    </dependency>
</dependencies>
```

### Basic Usage

```java
import okhttp3.*;
import com.google.gson.*;
import java.io.IOException;
import java.util.*;

public class ReolinkCamera {
    private final String host;
    private final String username;
    private final String password;
    private String token;
    private final String baseURL;
    private final OkHttpClient client;
    private final Gson gson;
    
    public static final MediaType JSON = MediaType.get("application/json; charset=utf-8");
    
    public ReolinkCamera(String host, String username, String password) {
        this.host = host;
        this.username = username;
        this.password = password;
        this.baseURL = String.format("http://%s/cgi-bin/api.cgi", host);
        this.client = new OkHttpClient();
        this.gson = new Gson();
    }
    
    private List<Map<String, Object>> sendRequest(List<Map<String, Object>> requests) 
            throws IOException {
        String json = gson.toJson(requests);
        
        RequestBody body = RequestBody.create(json, JSON);
        Request request = new Request.Builder()
                .url(baseURL)
                .post(body)
                .build();
        
        try (Response response = client.newCall(request).execute()) {
            if (!response.isSuccessful()) {
                throw new IOException("Unexpected code " + response);
            }
            
            String responseBody = response.body().string();
            return gson.fromJson(responseBody, 
                new com.google.gson.reflect.TypeToken<List<Map<String, Object>>>(){}.getType());
        }
    }
    
    public String login() throws IOException {
        Map<String, Object> userParam = new HashMap<>();
        userParam.put("userName", username);
        userParam.put("password", password);
        
        Map<String, Object> param = new HashMap<>();
        param.put("User", userParam);
        
        Map<String, Object> request = new HashMap<>();
        request.put("cmd", "Login");
        request.put("param", param);
        
        List<Map<String, Object>> requests = Arrays.asList(request);
        List<Map<String, Object>> responses = sendRequest(requests);
        
        Map<String, Object> response = responses.get(0);
        int code = ((Double) response.get("code")).intValue();
        
        if (code != 0) {
            throw new IOException("Login failed with code: " + code);
        }
        
        Map<String, Object> value = (Map<String, Object>) response.get("value");
        Map<String, Object> tokenObj = (Map<String, Object>) value.get("Token");
        this.token = (String) tokenObj.get("name");
        
        return this.token;
    }
    
    public void logout() throws IOException {
        Map<String, Object> request = new HashMap<>();
        request.put("cmd", "Logout");
        request.put("token", token);
        
        sendRequest(Arrays.asList(request));
    }
    
    public Map<String, Object> getDeviceInfo() throws IOException {
        Map<String, Object> request = new HashMap<>();
        request.put("cmd", "GetDevInfo");
        request.put("token", token);
        
        List<Map<String, Object>> responses = sendRequest(Arrays.asList(request));
        Map<String, Object> response = responses.get(0);
        
        int code = ((Double) response.get("code")).intValue();
        if (code != 0) {
            throw new IOException("Request failed with code: " + code);
        }
        
        return (Map<String, Object>) response.get("value");
    }
    
    public Map<String, Object> getHDDInfo() throws IOException {
        Map<String, Object> request = new HashMap<>();
        request.put("cmd", "GetHddInfo");
        request.put("token", token);
        
        List<Map<String, Object>> responses = sendRequest(Arrays.asList(request));
        Map<String, Object> response = responses.get(0);
        
        int code = ((Double) response.get("code")).intValue();
        if (code != 0) {
            throw new IOException("Request failed with code: " + code);
        }
        
        return (Map<String, Object>) response.get("value");
    }
    
    public void ptzControl(String operation, int speed) throws IOException {
        Map<String, Object> param = new HashMap<>();
        param.put("channel", 0);
        param.put("op", operation);
        param.put("speed", speed);
        
        Map<String, Object> request = new HashMap<>();
        request.put("cmd", "PtzCtrl");
        request.put("param", param);
        request.put("token", token);
        
        List<Map<String, Object>> responses = sendRequest(Arrays.asList(request));
        Map<String, Object> response = responses.get(0);
        
        int code = ((Double) response.get("code")).intValue();
        if (code != 0) {
            throw new IOException("PTZ control failed with code: " + code);
        }
    }
    
    public byte[] getSnapshot(int channel) throws IOException {
        String url = String.format("http://%s/cgi-bin/api.cgi?cmd=Snap&channel=%d&token=%s",
                host, channel, token);
        
        Request request = new Request.Builder()
                .url(url)
                .build();
        
        try (Response response = client.newCall(request).execute()) {
            if (!response.isSuccessful()) {
                throw new IOException("Unexpected code " + response);
            }
            
            return response.body().bytes();
        }
    }
    
    public static void main(String[] args) {
        ReolinkCamera camera = new ReolinkCamera("192.168.1.100", "admin", "your_password");
        
        try {
            // Login
            String token = camera.login();
            System.out.println("Logged in with token: " + token);
            
            // Get device info
            Map<String, Object> info = camera.getDeviceInfo();
            System.out.println("Device Info: " + info);
            
            // Get HDD info
            Map<String, Object> hdd = camera.getHDDInfo();
            System.out.println("HDD Info: " + hdd);
            
            // Get snapshot
            byte[] snapshot = camera.getSnapshot(0);
            java.nio.file.Files.write(
                java.nio.file.Paths.get("snapshot.jpg"), 
                snapshot
            );
            System.out.println("Snapshot saved to snapshot.jpg");
            
            // Logout
            camera.logout();
            System.out.println("Logged out");
            
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
```

## Advanced Examples

### Spring Boot Integration

```java
import org.springframework.stereotype.Service;
import org.springframework.beans.factory.annotation.Value;

@Service
public class ReolinkCameraService {
    
    @Value("${reolink.host}")
    private String host;
    
    @Value("${reolink.username}")
    private String username;
    
    @Value("${reolink.password}")
    private String password;
    
    private ReolinkCamera camera;
    
    @PostConstruct
    public void init() {
        this.camera = new ReolinkCamera(host, username, password);
    }
    
    public String captureSnapshot() throws IOException {
        camera.login();
        byte[] snapshot = camera.getSnapshot(0);
        camera.logout();
        
        // Convert to base64 or save to file
        return Base64.getEncoder().encodeToString(snapshot);
    }
}
```

### Async Operations (CompletableFuture)

```java
import java.util.concurrent.CompletableFuture;

public class AsyncReolinkCamera extends ReolinkCamera {
    
    public AsyncReolinkCamera(String host, String username, String password) {
        super(host, username, password);
    }
    
    public CompletableFuture<String> loginAsync() {
        return CompletableFuture.supplyAsync(() -> {
            try {
                return login();
            } catch (IOException e) {
                throw new RuntimeException(e);
            }
        });
    }
    
    public CompletableFuture<Map<String, Object>> getDeviceInfoAsync() {
        return CompletableFuture.supplyAsync(() -> {
            try {
                return getDeviceInfo();
            } catch (IOException e) {
                throw new RuntimeException(e);
            }
        });
    }
}
```

## See Also

- `BasicExample.java` - Simple example with basic operations
- `AdvancedExample.java` - Advanced features (PTZ, AI, streaming)
- `SpringBootExample.java` - Spring Boot integration

