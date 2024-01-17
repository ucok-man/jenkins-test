pipeline {
    agent any
    stages {
        stage('Checkout'){
            steps {
                checkout scm
            }
        }
        stage('Build & Dockerize'){
            steps {
                script{
                    docker.build("first-go-app")
                }
            }
        }
        stage('Push to Docker Hub'){
            steps {
                script{
                    docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-credentials') {
                        docker.image("jumatberkah/first-go-app").push()
                    }
                }
            }
        }
    }
}