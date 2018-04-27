import java.net.URLEncoder

class SonarApiClient
{
    static def sonarApiUrl = 'https://sonar.dev-build-create.build.finkit.io/api/'

    private static String encode(value) {
        return URLEncoder.encode(value, 'UTF-8')
    }

    static def buildSingleValuedQueryString(String key, String value) {
        return 'key=' + encode(key) + '&value=' + encode(value)
    }
    
    static def buildMultiValuedQueryString(String key, List<String> values) {
        if (values.count == 1) {
            return buildQueryString(key, values.first)
        }
    
        def queryString = 'key='
    
        for (value in values) {
          queryString += 'key=' + encode(key) + '&values=' + encode(value)
        }
    
        return queryString
    }
    
    private static String generateAuthValue() {
        def username = 'admin'
        def password = 'admin'
    
        return "${username}:${password}".getBytes().encodeBase64().toString()
    }
    
    static boolean postQueryString(String url, String queryString) {
        def connection = new URL(url).openConnection() as HttpURLConnection
    
        connection.setRequestProperty('Accept', 'application/json')
        connection.setRequestProperty('Authorization', "Basic ${generateAuthValue()}")
        connection.setRequestMethod('POST')
        connection.doOutput = true
    
        def writer = new OutputStreamWriter(connection.outputStream)
        
        try {
            writer.write(queryString)
            writer.flush()
        
            if (connection.responseCode == HttpURLConnection.HTTP_OK ||
                connection.responseCode == HttpURLConnection.HTTP_NO_CONTENT) {
                println "Request to ${url} with query string ${queryString} succeeded"
            } else {
                println "Request to ${url} with query string ${queryString} failed with status " +
                        "${connection.responseCode} and response ${connection.responseMessage}"
                return false
            }
        } catch (IOException ioe) {
            println "Request failed with error: ${ioe}"
            return false
        } finally {
            writer.close()
        }
    
        return true
    }
}
