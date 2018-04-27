// Script must be called from the script directory to pick up the imported Class.
import SonarApiClient

def sonarApiSettingsUrl = SonarApiClient.sonarApiUrl + 'settings/set'
// Could connect to 'settings/get' after to test without exercising functionality.

def keySingleValuePairs = [
    new Tuple2('sonar.core.serverBaseURL', 'https://sonar.dev-build-create.build.finkit.io'),
    new Tuple2('sonar.auth.github.enabled', 'true')
]

for (pair in keySingleValuePairs) {
    def queryString = SonarApiClient.buildSingleValuedQueryString(pair.first, pair.second)
    def querySucceeded = SonarApiClient.postQueryString(sonarApiSettingsUrl, queryString)

    if (!querySucceeded) {
        System.exit(1)
    }
}

System.exit(0)
