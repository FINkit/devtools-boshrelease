def getAuthValue() {
    def username = 'admin'
    def password = 'admin'

    return "${username}:${password}".getBytes().encodeBase64().toString()
}

def sonarApiUrl = 'https://sonar.dev-build-create.build.finkit.io/api/'
def sonarApiPluginsUrl = sonarApiUrl + 'plugins/install'
def queryString = "key=authgithub"

def connection = new URL(sonarApiPluginsUrl).openConnection() as HttpURLConnection

connection.setRequestProperty('Accept', 'application/json')
connection.setRequestProperty('Authorization', "Basic ${getAuthValue()}")
connection.setRequestMethod('POST')
connection.doOutput = true

def writer = new OutputStreamWriter(connection.outputStream)
def pluginInstalled = false

try {
    writer.write(queryString)
    writer.flush()

    if (connection.responseCode == HttpURLConnection.HTTP_OK ||
        connection.responseCode == HttpURLConnection.HTTP_NO_CONTENT) {
        pluginInstalled = true
        println "Plugin install succeeded!"
    } else {
        println "Plugin install failed with status ${connection.responseCode} and response ${connection.responseMessage}"
    }
} catch (IOException ioe) {
    println "Request failed with error: ${ioe}"
}
finally {
    writer.close()
}

if (pluginInstalled) {
    System.exit(0)
} else {
    System.exit(1)
}

// Probably need to restart now
