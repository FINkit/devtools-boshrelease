#!/usr/bin/env bash
set -euxo pipefail

cd bbl-state/bbl-state
source <(bbl print-env)
BOSH_NAME=$(bosh int <(bbl outputs) --path /director_name); export BOSH_NAME
JENKINS_PASSWORD=$(credhub get -j -n "/$BOSH_NAME/$BOSH_DEPLOYMENT/jenkins_admin_password" | jq -r .value)
sleep 10 # Sleeping because this curl lacks the --retry-connrefused
bosh -d $BOSH_DEPLOYMENT ssh jenkins-master/0 -c "curl -S --retry-max-time 300 --retry-delay 1 --retry 300 http://localhost:8080/login"
curl -SLo /opt/jenkins-cli.jar http://localhost:8080/jnlpJars/jenkins-cli.jar
jenkins="java -jar /opt/jenkins-cli.jar -auth administrator:$JENKINS_PASSWORD -s http://localhost:8080"

cat <<EOF > test_job.xml
<?xml version='1.0' encoding='UTF-8'?>
<project>
    <description>temp job for testing purposes</description>
    <keepDependencies>false</keepDependencies>
    <properties>
    <com.chikli.hudson.plugin.naginator.NaginatorOptOutProperty plugin="naginator@1.17.2">
        <optOut>false</optOut>
    </com.chikli.hudson.plugin.naginator.NaginatorOptOutProperty>
    <com.sonyericsson.rebuild.RebuildSettings plugin="rebuild@1.27">
        <autoRebuild>false</autoRebuild>
        <rebuildDisabled>false</rebuildDisabled>
    </com.sonyericsson.rebuild.RebuildSettings>
    <com.synopsys.arc.jenkinsci.plugins.jobrestrictions.jobs.JobRestrictionProperty plugin="job-restrictions@0.6"/>
    <hudson.plugins.throttleconcurrents.ThrottleJobProperty plugin="throttle-concurrents@2.0.1">
        <categories class="java.util.concurrent.CopyOnWriteArrayList"/>
        <throttleEnabled>false</throttleEnabled>
        <throttleOption>project</throttleOption>
        <limitOneJobWithMatchingParams>false</limitOneJobWithMatchingParams>
        <paramsToUseForLimit></paramsToUseForLimit>
    </hudson.plugins.throttleconcurrents.ThrottleJobProperty>
    </properties>
    <scm class="hudson.scm.NullSCM"/>
    <canRoam>false</canRoam>
    <disabled>false</disabled>
    <blockBuildWhenDownstreamBuilding>false</blockBuildWhenDownstreamBuilding>
    <blockBuildWhenUpstreamBuilding>false</blockBuildWhenUpstreamBuilding>
    <triggers/>
    <concurrentBuild>false</concurrentBuild>
    <builders>
    <hudson.tasks.Shell>
        <command>env</command>
    </hudson.tasks.Shell>
    </builders>
    <publishers/>
    <buildWrappers/>
    <assignedNode>build</assignedNode>
</project>
EOF

$jenkins create-job print_env < test_job.xml
$jenkins build print_env -f -v