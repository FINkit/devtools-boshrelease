#!/usr/bin/env bash
set -euo pipefail

cd bbl-state/bbl-state
source <(bbl print-env)
JUMPBOX_USER="jumpbox"; export JUMPBOX_USER
JUMPBOX_HOST=$(bosh int <(bbl outputs) --path /jumpbox_url | cut -d: -f1); export  JUMPBOX_HOST
# set up jumpbox ssh tunnel
instance_info=$(bosh --json instances -i); export instance_info
jenkins_ip=$(jq -r .Tables[0].Rows[0].ips <<< "$instance_info")
# shellcheck disable=SC2029
ssh -o StrictHostKeyChecking=no -fNnL "8080:$jenkins_ip:8080" -i                    "$JUMPBOX_PRIVATE_KEY" "$JUMPBOX_USER@$JUMPBOX_HOST"
trap "pkill ssh" EXIT
BOSH_NAME=$(bosh int <(bbl outputs) --path /director_name); export BOSH_NAME
JENKINS_PASSWORD=$(credhub get -j -n "/$BOSH_NAME/jenkins/jenkins_admin_password" | jq -r .value)
sleep 10 # Sleeping because this curl lacks the --retry-connrefused
curl -S --retry-max-time 300 --retry-delay 1 --retry 300 http://localhost:8080/login
curl -SLo jenkins-cli.jar http://localhost:8080/jnlpJars/jenkins-cli.jar
chmod +r jenkins-cli.jar
java -jar jenkins-cli.jar -auth administrator:$JENKINS_PASSWORD -s http://localhost:8080 create-job print_env < devtools-boshrelease-ci/ci/tasks/test-jenkins-run-job.xml
java -jar jenkins-cli.jar -auth administrator:$JENKINS_PASSWORD -s http://localhost:8080 build print_env -f -v
