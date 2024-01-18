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
                script {
                    docker.build('jumatberkah/jenkins-test')
                }
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
            steps {
                script {
                    docker.withRegistry('', 'docker-auth') {
                        def appimg =  docker.build('jumatberkah/jenkins-test')
                        appimg.push('--password-stdin')
                    }
                }
            }
        }
    }
    post {
        always {
            sh(script: 'docker container rm -f -v jenkins-test')
        }
    }
}
