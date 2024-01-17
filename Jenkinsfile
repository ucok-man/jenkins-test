/* groovylint-disable CompileStatic */
pipeline {
    agent any

    stages {
        stage('Verify Branch') {
            steps {
                echo("$GIT_BRANCH")
            }
        }
        stage('Docker Build') {
            steps {
                wrap([$class: 'BuildUser']) {
                    echo('BUILD USER: ')
                    echo(env.BUILD_USER)
                    echo(env.BUILD_USER_ID)
                    echo(env.BUILD_USER_GROUP)
                }
                sh(script: 'docker image build -t jenkins-test .')
            }
        }
    }
}
