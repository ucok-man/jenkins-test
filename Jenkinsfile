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
                sh(script: 'docker image build -t jenkins-test .')
            }
        }
        stage('Start App') {
            steps {
                sh(script: 'docker run --name jenkins-test --publish 8081:8081 --detach jenkins-test')
            }
        }
        stage('Run Tests') {
            steps {
                sh(script: 'go test -v ./...')
            }
            post {
                success {
                    echo('Test passed :)')
                }
                failure {
                    echo('Test failed :)')
                }
            }
        }
    }
    post {
        always {
            sh(script: 'docker container stop jenkins-test')
        }
    }
}
