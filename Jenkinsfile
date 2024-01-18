/* groovylint-disable CompileStatic */
pipeline {
    agent any

    def appImage

    stages {
        stage('Verify Branch') {
            steps {
                echo("$GIT_BRANCH")
            }
        }
        stage('Docker Build') {
            steps {
                appImage = docker.build('jumatberkah/jenkins-test')
            }
        }
        stage('Start App') {
            steps {
                sh(script: 'docker run --name jenkins-test --publish 8081:8081 --detach jumatberkah/jenkins-test')
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
        stage('Docker Push') {
            docker.withRegistry('', 'docker-auth') {
                appImage.push()
            }
        }
    }
    post {
        always {
            sh(script: 'docker container stop jenkins-test')
        }
    }
}
