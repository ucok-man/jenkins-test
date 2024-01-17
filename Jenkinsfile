pipeline {
    agent any
    stages {
        stage('Checkout') {
            steps {
                checkout scmGit(branches: [[name: '*/main']], extensions: [], userRemoteConfigs: [[url: 'https://github.com/ucok-man/jenkins-test']])            }
        }
        stage('Build & Dockerize') {
            steps {
                script{
                    docker.build("jenkins-test")
                }
            }
        }
        stage('push image to dockerhub') {
            steps {
                script{
                   bat 'docker login -u jumatberkah -p ucok27101998'
                   bat 'docker tag jenkins-test jumatberkah/jenkins-test'
                   bat 'docker push jumatberkah/jenkins-test'
                }
            }
        }
    }
}