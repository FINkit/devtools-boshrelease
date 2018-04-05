#!/usr/bin/env bash
set -euxo pipefail

cd bbl-state/bbl-state
source <(bbl print-env)
BOSH_NAME=$(bosh int <(bbl outputs) --path /director_name); export BOSH_NAME
JENKINS_PASSWORD=$(credhub get -j -n "/$BOSH_NAME/$BOSH_DEPLOYMENT/jenkins_admin_password" | jq -r .value)
sleep 10 # Sleeping because this curl lacks the --retry-connrefused
bosh -d $BOSH_DEPLOYMENT ssh jenkins-master/0 -c "curl -S --retry-max-time 300 --retry-delay 1 --retry 300 http://localhost:8080/login"
cd ../../devtools-bosh-release-ci/tasks/
bosh -d $BOSH_DEPLOYMENT scp test-jenkins-run-job.xml jenkins-master/0:/tmp/test-jenkins-run-job.xml
bosh -d $BOSH_DEPLOYMENT ssh jenkins-master/0 -c "curl -SLo /opt/jenkins-cli.jar http://localhost:8080/jnlpJars/jenkins-cli.jar"
bosh -d $BOSH_DEPLOYMENT ssh jenkins-master/0 -c "java -jar /opt/jenkins-cli.jar -auth administrator:$JENKINS_PASSWORD -s http://localhost:8080 create-job print_env < /tmp/test-jenkins-run-job.xml"
bosh -d $BOSH_DEPLOYMENT ssh jenkins-master/0 -c "java -jar /opt/jenkins-cli.jar -auth administrator:$JENKINS_PASSWORD -s http://localhost:8080 build print_env -f -v"
