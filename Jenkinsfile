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
                    docker.build("jenkins-test")
                }
            }
        }
        stage('Push to Docker Hub'){
            steps {
                script{
                    docker.withRegistry('https://registry.hub.docker.com', '94e58d8b-6783-4262-b1a0-96250f45f61c') {
                        docker.image("jenkins-test").push()
                    }
                }
            }
        }
    }
}